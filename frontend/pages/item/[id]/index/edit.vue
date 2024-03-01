<script setup lang="ts">
  import { ItemAttachment, ItemField, ItemOut, ItemUpdate } from "~~/lib/api/types/data-contracts";
  import { AttachmentTypes } from "~~/lib/api/types/non-generated";
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";
  import { capitalize } from "~~/lib/strings";
  import Autocomplete from "~~/components/Form/Autocomplete.vue";
  import MdiDelete from "~icons/mdi/delete";
  import MdiPencil from "~icons/mdi/pencil";
  import MdiContentSaveOutline from "~icons/mdi/content-save-outline";

  definePageMeta({
    middleware: ["auth"],
  });

  const route = useRoute();
  const api = useUserApi();
  const toast = useNotifier();
  const preferences = useViewPreferences();

  const itemId = computed<string>(() => route.params.id as string);

  const locationStore = useLocationStore();
  const locations = computed(() => locationStore.allLocations);

  const labelStore = useLabelStore();
  const labels = computed(() => labelStore.labels);

  const {
    data: nullableItem,
    refresh,
    pending: requestPending,
  } = useAsyncData(async () => {
    const { data, error } = await api.items.get(itemId.value);
    if (error) {
      toast.error("Failed to load item");
      navigateTo("/home");
      return;
    }

    if (locations && data.location?.id) {
      // @ts-expect-error - we know the locations is valid
      const location = locations.value.find(l => l.id === data.location.id);
      if (location) {
        data.location = location;
      }
    }

    if (data.parent) {
      parent.value = data.parent;
    }

    return data;
  });

  const item = computed<ItemOut>(() => nullableItem.value as ItemOut);

  onMounted(() => {
    refresh();
  });

  async function saveItem() {
    if (!item.value.location?.id) {
      toast.error("Failed to save item: no location selected");
      return;
    }

    const payload: ItemUpdate = {
      ...item.value,
      locationId: item.value.location?.id,
      labelIds: item.value.labels.map(l => l.id),
      parentId: parent.value ? parent.value.id : null,
      assetId: item.value.assetId,
    };

    const { error } = await api.items.update(itemId.value, payload);

    if (error) {
      toast.error("Failed to save item");
      return;
    }

    toast.success("Item saved");
    navigateTo("/item/" + itemId.value);
  }
  type NoUndefinedField<T> = { [P in keyof T]-?: NoUndefinedField<NonNullable<T[P]>> };

  type StringKeys<T> = { [k in keyof T]: T[k] extends string ? k : never }[keyof T];
  type OnlyString<T> = { [k in StringKeys<T>]: string };

  type NumberKeys<T> = { [k in keyof T]: T[k] extends number ? k : never }[keyof T];
  type OnlyNumber<T> = { [k in NumberKeys<T>]: number };

  type TextFormField = {
    type: "text" | "textarea";
    label: string;
    // key of ItemOut where the value is a string
    ref: keyof OnlyString<NoUndefinedField<ItemOut>>;
  };

  type NumberFormField = {
    type: "number";
    label: string;
    ref: keyof OnlyNumber<NoUndefinedField<ItemOut>> | keyof OnlyString<NoUndefinedField<ItemOut>>;
  };

  // https://stackoverflow.com/questions/50851263/how-do-i-require-a-keyof-to-be-for-a-property-of-a-specific-type
  // I don't know why typescript can't just be normal
  type BooleanKeys<T> = { [k in keyof T]: T[k] extends boolean ? k : never }[keyof T];
  type OnlyBoolean<T> = { [k in BooleanKeys<T>]: boolean };

  interface BoolFormField {
    type: "checkbox";
    label: string;
    ref: keyof OnlyBoolean<NoUndefinedField<ItemOut>>;
  }

  type DateKeys<T> = { [k in keyof T]: T[k] extends Date | string ? k : never }[keyof T];
  type OnlyDate<T> = { [k in DateKeys<T>]: Date | string };

  type DateFormField = {
    type: "date";
    label: string;
    ref: keyof OnlyDate<NoUndefinedField<ItemOut>>;
  };

  type FormField = TextFormField | BoolFormField | DateFormField | NumberFormField;

  const mainFields: FormField[] = [
    {
      type: "text",
      label: "Name",
      ref: "name",
    },
    {
      type: "number",
      label: "Quantity",
      ref: "quantity",
    },
    {
      type: "textarea",
      label: "Description",
      ref: "description",
    },
    {
      type: "text",
      label: "Serial Number",
      ref: "serialNumber",
    },
    {
      type: "text",
      label: "Model Number",
      ref: "modelNumber",
    },
    {
      type: "text",
      label: "Manufacturer",
      ref: "manufacturer",
    },
    {
      type: "textarea",
      label: "Notes",
      ref: "notes",
    },
    {
      type: "checkbox",
      label: "Insured",
      ref: "insured",
    },
    {
      type: "checkbox",
      label: "Archived",
      ref: "archived",
    },
    {
      type: "text",
      label: "Asset ID",
      ref: "assetId",
    },
  ];

  const purchaseFields: FormField[] = [
    {
      type: "text",
      label: "Purchased From",
      ref: "purchaseFrom",
    },
    {
      type: "text",
      label: "Purchase Price",
      ref: "purchasePrice",
    },
    {
      type: "date",
      label: "Purchase Date",
      // @ts-expect-error - we know this is a date
      ref: "purchaseTime",
    },
  ];

  const warrantyFields: FormField[] = [
    {
      type: "checkbox",
      label: "Lifetime Warranty",
      ref: "lifetimeWarranty",
    },
    {
      type: "date",
      label: "Warranty Expires",
      // @ts-expect-error - we know this is a date
      ref: "warrantyExpires",
    },
    {
      type: "textarea",
      label: "Warranty Notes",
      ref: "warrantyDetails",
    },
  ];

  const soldFields: FormField[] = [
    {
      type: "text",
      label: "Sold To",
      ref: "soldTo",
    },
    {
      type: "text",
      label: "Sold Price",
      ref: "soldPrice",
    },
    {
      type: "date",
      label: "Sold At",
      // @ts-expect-error - we know this is a date
      ref: "soldTime",
    },
  ];

  // - Attachments
  const attDropZone = ref<HTMLDivElement>();
  const { isOverDropZone: attDropZoneActive } = useDropZone(attDropZone);

  const refAttachmentInput = ref<HTMLInputElement>();

  function clickUpload() {
    if (!refAttachmentInput.value) {
      return;
    }
    refAttachmentInput.value.click();
  }

  function uploadImage(e: Event) {
    const files = (e.target as HTMLInputElement).files;
    if (!files || !files.item(0)) {
      return;
    }

    const first = files.item(0);
    if (!first) {
      return;
    }

    uploadAttachment([first], null);
  }

  const dropPhoto = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Photo);
  const dropAttachment = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Attachment);
  const dropWarranty = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Warranty);
  const dropManual = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Manual);
  const dropReceipt = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Receipt);

  async function uploadAttachment(files: File[] | null, type: AttachmentTypes | null) {
    if (!files || files.length === 0) {
      return;
    }

    const { data, error } = await api.items.attachments.add(itemId.value, files[0], files[0].name, type);

    if (error) {
      toast.error("Failed to upload attachment");
      return;
    }

    toast.success("Attachment uploaded");

    item.value.attachments = data.attachments;
  }

  const confirm = useConfirm();

  async function deleteAttachment(attachmentId: string) {
    const confirmed = await confirm.open("Are you sure you want to delete this attachment?");

    if (confirmed.isCanceled) {
      return;
    }

    const { error } = await api.items.attachments.delete(itemId.value, attachmentId);

    if (error) {
      toast.error("Failed to delete attachment");
      return;
    }

    toast.success("Attachment deleted");
    item.value.attachments = item.value.attachments.filter(a => a.id !== attachmentId);
  }

  const editState = reactive({
    modal: false,
    loading: false,

    // Values
    obj: {},
    id: "",
    title: "",
    type: "",
    primary: false,
  });

  const attachmentOpts = Object.entries(AttachmentTypes).map(([key, value]) => ({
    text: capitalize(key),
    value,
  }));

  function openAttachmentEditDialog(attachment: ItemAttachment) {
    editState.id = attachment.id;
    editState.title = attachment.document.title;
    editState.type = attachment.type;
    editState.primary = attachment.primary;
    editState.modal = true;

    editState.obj = attachmentOpts.find(o => o.value === attachment.type) || attachmentOpts[0];
  }

  async function updateAttachment() {
    editState.loading = true;
    const { error, data } = await api.items.attachments.update(itemId.value, editState.id, {
      title: editState.title,
      type: editState.type,
      primary: editState.primary,
    });

    if (error) {
      toast.error("Failed to update attachment");
      return;
    }

    item.value.attachments = data.attachments;

    editState.loading = false;
    editState.modal = false;

    editState.id = "";
    editState.title = "";
    editState.type = "";

    toast.success("Attachment updated");
  }

  function addField() {
    item.value.fields.push({
      id: null,
      name: "Field Name",
      type: "text",
      textValue: "",
      numberValue: 0,
      booleanValue: false,
      timeValue: null,
    } as unknown as ItemField);
  }

  const { query, results } = useItemSearch(api, { immediate: false });
  const parent = ref();

  async function deleteItem() {
    const confirmed = await confirm.open("Are you sure you want to delete this item?");

    if (!confirmed.data) {
      return;
    }

    const { error } = await api.items.delete(itemId.value);
    if (error) {
      toast.error("Failed to delete item");
      return;
    }
    toast.success("Item deleted");
    navigateTo("/home");
  }

  async function keyboardSave(e: KeyboardEvent) {
    // Cmd + S
    if (e.metaKey && e.key === "s") {
      e.preventDefault();
      await saveItem();
    }

    // Ctrl + S
    if (e.ctrlKey && e.key === "s") {
      e.preventDefault();
      await saveItem();
    }
  }

  onMounted(() => {
    window.addEventListener("keydown", keyboardSave);
  });

  onUnmounted(() => {
    window.removeEventListener("keydown", keyboardSave);
  });
</script>

<template>
  <div v-if="item" class="pb-8">
    <BaseModal v-model="editState.modal">
      <template #title> {{ $t("item.edit.editState.title") }} </template>

      <FormTextField v-model="editState.title" :label="$t('item.edit.editState.attachment.title')" />
      <FormSelect
        v-model:value="editState.type"
        :label="$t('item.edit.editState.attachment.type')"
        value-key="value"
        name="text"
        :items="attachmentOpts"
      />
      <div v-if="editState.type == 'photo'" class="flex gap-2 mt-3">
        <input v-model="editState.primary" type="checkbox" class="checkbox" />
        <p class="text-sm">
          <span class="font-semibold">Primary Photo</span>
          {{ $t("item.edit.editState.tips") }}
        </p>
      </div>
      <div class="modal-action">
        <BaseButton :loading="editState.loading" @click="updateAttachment"> {{ $t("item.edit.editState.button") }} </BaseButton>
      </div>
    </BaseModal>

    <section class="relative">
      <div class="my-4 justify-end flex gap-2 items-center sticky z-10 top-1">
        <div class="mr-auto tooltip tooltip-right" data-tip="Show Advanced View Options">
          <label class="label cursor-pointer mr-auto">
            <input v-model="preferences.editorAdvancedView" type="checkbox" class="toggle toggle-primary" />
            <span class="label-text ml-4"> {{ $t("item.edit.advanced") }} </span>
          </label>
        </div>
        <BaseButton size="sm" @click="saveItem">
          <template #icon>
            <MdiContentSaveOutline />
          </template>
          {{ $t("item.edit.save_button") }}
        </BaseButton>
        <BaseButton class="btn btn-sm btn-error" @click="deleteItem()">
          <MdiDelete class="mr-2" />
          {{ $t("item.edit.delete_button") }}
        </BaseButton>
      </div>
      <div v-if="!requestPending" class="space-y-6">
        <BaseCard class="overflow-visible">
          <template #title> {{ $t("item.edit.details.title") }} </template>
          <template #title-actions>
            <div class="flex flex-wrap justify-between items-center mt-2 gap-4"></div>
          </template>
          <div class="px-5 pt-2 border-t mb-6 grid md:grid-cols-2 gap-4">
            <LocationSelector v-model="item.location" />
            <FormMultiselect v-model="item.labels" :label="$t('item.edit.details.labels')" :items="labels ?? []" />
            <Autocomplete
              v-if="preferences.editorAdvancedView"
              v-model="parent"
              v-model:search="query"
              :items="results"
              item-text="name"
              :label="$t('item.edit.details.parent')"
              no-results-text="Type to search..."
            />
          </div>

          <div class="border-t border-gray-300 sm:p-0">
            <div v-for="field in mainFields" :key="field.ref" class="sm:divide-y sm:divide-gray-300 grid grid-cols-1">
              <div class="pt-2 px-4 pb-4 sm:px-6 border-b border-gray-300">
                <FormTextArea v-if="field.type === 'textarea'" v-model="item[field.ref]" :label="field.label" inline />
                <FormTextField
                  v-else-if="field.type === 'text'"
                  v-model="item[field.ref]"
                  :label="field.label"
                  inline
                />
                <FormTextField
                  v-else-if="field.type === 'number'"
                  v-model.number="item[field.ref]"
                  type="number"
                  :label="field.label"
                  inline
                />
                <FormDatePicker
                  v-else-if="field.type === 'date'"
                  v-model="item[field.ref]"
                  :label="field.label"
                  inline
                />
                <FormCheckbox
                  v-else-if="field.type === 'checkbox'"
                  v-model="item[field.ref]"
                  :label="field.label"
                  inline
                />
              </div>
            </div>
          </div>
        </BaseCard>

        <BaseCard>
          <template #title> {{ $t("item.edit.custom.title") }} </template>
          <div class="px-5 border-t divide-y divide-gray-300 space-y-4">
            <div
              v-for="(field, idx) in item.fields"
              :key="`field-${idx}`"
              class="grid grid-cols-2 md:grid-cols-4 gap-2"
            >
              <!-- <FormSelect v-model:value="field.type" label="Field Type" :items="fieldTypes" value-key="value" /> -->
              <FormTextField v-model="field.name" :label="$t('item.edit.custom.name')" />
              <div class="flex items-end col-span-3">
                <FormTextField v-model="field.textValue" :label="$t('item.edit.custom.value')" />
                <div class="tooltip" data-tip="Delete">
                  <button class="btn btn-sm btn-square mb-2 ml-2" @click="item.fields.splice(idx, 1)">
                    <MdiDelete />
                  </button>
                </div>
              </div>
            </div>
          </div>
          <div class="px-5 pb-4 mt-4 flex justify-end">
            <BaseButton size="sm" @click="addField"> {{ $t("item.edit.custom.button") }} </BaseButton>
          </div>
        </BaseCard>

        <div
          v-if="preferences.editorAdvancedView"
          ref="attDropZone"
          class="overflow-visible card bg-base-100 shadow-xl sm:rounded-lg"
        >
          <div class="px-4 py-5 sm:px-6">
            <h3 class="text-lg font-medium leading-6">{{ $t("item.edit.attachments.title") }}</h3>
            <p class="text-xs">{{ $t("item.edit.attachments.tips") }}</p>
          </div>
          <div class="border-t border-gray-300 p-4">
            <div v-if="attDropZoneActive" class="grid grid-cols-4 gap-4">
              <DropZone @drop="dropPhoto"> {{ $t("item.edit.attachments.dropZone.photo") }} </DropZone>
              <DropZone @drop="dropWarranty"> {{ $t("item.edit.attachments.dropZone.warranty") }} </DropZone>
              <DropZone @drop="dropManual"> {{ $t("item.edit.attachments.dropZone.manual") }} </DropZone>
              <DropZone @drop="dropAttachment"> {{ $t("item.edit.attachments.dropZone.attachment") }} </DropZone>
              <DropZone @drop="dropReceipt"> {{ $t("item.edit.attachments.dropZone.receipt") }} </DropZone>
            </div>
            <button
              v-else
              class="h-24 w-full border-2 border-primary border-dashed grid place-content-center"
              @click="clickUpload"
            >
              <input ref="refAttachmentInput" hidden type="file" @change="uploadImage" />
              <p>{{ $t("item.edit.attachments.dropZone.tips") }}</p>
            </button>
          </div>

          <div class="border-t border-gray-300 p-4">
            <ul role="list" class="divide-y divide-gray-400 rounded-md border border-gray-400">
              <li
                v-for="attachment in item.attachments"
                :key="attachment.id"
                class="grid grid-cols-6 justify-between py-3 pl-3 pr-4 text-sm"
              >
                <p class="my-auto col-span-4">
                  {{ attachment.document.title }}
                </p>
                <p class="my-auto">
                  {{ capitalize(attachment.type) }}
                </p>
                <div class="flex gap-2 justify-end">
                  <div class="tooltip" data-tip="Delete">
                    <button class="btn btn-sm btn-square" @click="deleteAttachment(attachment.id)">
                      <MdiDelete />
                    </button>
                  </div>
                  <div class="tooltip" data-tip="Edit">
                    <button class="btn btn-sm btn-square" @click="openAttachmentEditDialog(attachment)">
                      <MdiPencil />
                    </button>
                  </div>
                </div>
              </li>
            </ul>
          </div>
        </div>

        <div v-if="preferences.editorAdvancedView" class="overflow-visible card bg-base-100 shadow-xl sm:rounded-lg">
          <div class="px-4 py-5 sm:px-6">
            <h3 class="text-lg font-medium leading-6">{{ $t("item.edit.purchase.title") }}</h3>
          </div>
          <div class="border-t border-gray-300 sm:p-0">
            <div
              v-for="field in purchaseFields"
              :key="field.ref"
              class="sm:divide-y sm:divide-gray-300 grid grid-cols-1"
            >
              <div class="pt-2 px-4 pb-4 sm:px-6 border-b border-gray-300">
                <FormTextArea v-if="field.type === 'textarea'" v-model="item[field.ref]" :label="field.label" inline />
                <FormTextField
                  v-else-if="field.type === 'text'"
                  v-model="item[field.ref]"
                  :label="field.label"
                  inline
                />
                <FormTextField
                  v-else-if="field.type === 'number'"
                  v-model.number="item[field.ref]"
                  type="number"
                  :label="field.label"
                  inline
                />
                <FormDatePicker
                  v-else-if="field.type === 'date'"
                  v-model="item[field.ref]"
                  :label="field.label"
                  inline
                />
                <FormCheckbox
                  v-else-if="field.type === 'checkbox'"
                  v-model="item[field.ref]"
                  :label="field.label"
                  inline
                />
              </div>
            </div>
          </div>
        </div>

        <div v-if="preferences.editorAdvancedView" class="overflow-visible card bg-base-100 shadow-xl sm:rounded-lg">
          <div class="px-4 py-5 sm:px-6">
            <h3 class="text-lg font-medium leading-6">{{ $t("item.edit.warranty.title") }}</h3>
          </div>
          <div class="border-t border-gray-300 sm:p-0">
            <div
              v-for="field in warrantyFields"
              :key="field.ref"
              class="sm:divide-y sm:divide-gray-300 grid grid-cols-1"
            >
              <div class="pt-2 px-4 pb-4 sm:px-6 border-b border-gray-300">
                <FormTextArea v-if="field.type === 'textarea'" v-model="item[field.ref]" :label="field.label" inline />
                <FormTextField
                  v-else-if="field.type === 'text'"
                  v-model="item[field.ref]"
                  :label="field.label"
                  inline
                />
                <FormTextField
                  v-else-if="field.type === 'number'"
                  v-model.number="item[field.ref]"
                  type="number"
                  :label="field.label"
                  inline
                />
                <FormDatePicker
                  v-else-if="field.type === 'date'"
                  v-model="item[field.ref]"
                  :label="field.label"
                  inline
                />
                <FormCheckbox
                  v-else-if="field.type === 'checkbox'"
                  v-model="item[field.ref]"
                  :label="field.label"
                  inline
                />
              </div>
            </div>
          </div>
        </div>

        <div v-if="preferences.editorAdvancedView" class="overflow-visible card bg-base-100 shadow-xl sm:rounded-lg">
          <div class="px-4 py-5 sm:px-6">
            <h3 class="text-lg font-medium leading-6">{{ $t("item.edit.sold.title") }}</h3>
          </div>
          <div class="border-t border-gray-300 sm:p-0">
            <div v-for="field in soldFields" :key="field.ref" class="sm:divide-y sm:divide-gray-300 grid grid-cols-1">
              <div class="pt-2 pb-4 px-4 sm:px-6 border-b border-gray-300">
                <FormTextArea v-if="field.type === 'textarea'" v-model="item[field.ref]" :label="field.label" inline />
                <FormTextField
                  v-else-if="field.type === 'text'"
                  v-model="item[field.ref]"
                  :label="field.label"
                  inline
                />
                <FormTextField
                  v-else-if="field.type === 'number'"
                  v-model.number="item[field.ref]"
                  type="number"
                  :label="field.label"
                  inline
                />
                <FormDatePicker
                  v-else-if="field.type === 'date'"
                  v-model="item[field.ref]"
                  :label="field.label"
                  inline
                />
                <FormCheckbox
                  v-else-if="field.type === 'checkbox'"
                  v-model="item[field.ref]"
                  :label="field.label"
                  inline
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
