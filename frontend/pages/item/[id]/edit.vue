<script setup lang="ts">
  import { ItemAttachment, ItemUpdate } from "~~/lib/api/types/data-contracts";
  import { AttachmentTypes } from "~~/lib/api/types/non-generated";
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";
  import { capitalize } from "~~/lib/strings";
  import Autocomplete from "~~/components/Form/Autocomplete.vue";

  definePageMeta({
    middleware: ["auth"],
  });

  const route = useRoute();
  const api = useUserApi();
  const toast = useNotifier();
  const preferences = useViewPreferences();

  const itemId = computed<string>(() => route.params.id as string);

  const locationStore = useLocationStore();
  const locations = computed(() => locationStore.locations);

  const labelStore = useLabelStore();
  const labels = computed(() => labelStore.labels);

  const { data: item, refresh } = useAsyncData(async () => {
    const { data, error } = await api.items.get(itemId.value);
    if (error) {
      toast.error("Failed to load item");
      navigateTo("/home");
      return;
    }

    if (locations) {
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

  onMounted(() => {
    refresh();
  });

  async function saveItem() {
    const payload: ItemUpdate = {
      ...item.value,
      locationId: item.value.location?.id,
      labelIds: item.value.labels.map(l => l.id),
      parentId: parent.value ? parent.value.id : null,
    };

    const { error } = await api.items.update(itemId.value, payload);

    if (error) {
      toast.error("Failed to save item");
      return;
    }

    toast.success("Item saved");
    navigateTo("/item/" + itemId.value);
  }

  type FormField = {
    type: "text" | "textarea" | "select" | "date" | "label" | "location" | "number" | "checkbox";
    label: string;
    ref: string;
  };

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
      ref: "warrantyExpires",
    },
    {
      type: "textarea",
      label: "Warranty Notes",
      ref: "warrantyDetails",
    },
  ];

  const soldFields = [
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

  function uploadImage(e: InputEvent) {
    const files = (e.target as HTMLInputElement).files;
    if (!files) {
      return;
    }

    uploadAttachment([files.item(0)], AttachmentTypes.Attachment);
  }

  const dropPhoto = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Photo);
  const dropAttachment = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Attachment);
  const dropWarranty = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Warranty);
  const dropManual = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Manual);
  const dropReceipt = (files: File[] | null) => uploadAttachment(files, AttachmentTypes.Receipt);

  async function uploadAttachment(files: File[] | null, type: AttachmentTypes) {
    if (!files && files.length === 0) {
      return;
    }

    const { data, error } = await api.items.addAttachment(itemId.value, files[0], files[0].name, type);

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

    const { error } = await api.items.deleteAttachment(itemId.value, attachmentId);

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
  });

  const attachmentOpts = Object.entries(AttachmentTypes).map(([key, value]) => ({
    text: capitalize(key),
    value,
  }));

  function openAttachmentEditDialog(attachment: ItemAttachment) {
    editState.id = attachment.id;
    editState.title = attachment.document.title;
    editState.type = attachment.type;
    editState.modal = true;

    editState.obj = attachmentOpts.find(o => o.value === attachment.type);
  }

  async function updateAttachment() {
    editState.loading = true;
    const { error, data } = await api.items.updateAttachment(itemId.value, editState.id, {
      title: editState.title,
      type: editState.type,
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

  // Custom Fields
  // const fieldTypes = [
  //   {
  //     name: "Text",
  //     value: "text",
  //   },
  //   {
  //     name: "Number",
  //     value: "number",
  //   },
  //   {
  //     name: "Boolean",
  //     value: "boolean",
  //   },
  // ];

  function addField() {
    item.value.fields.push({
      id: null,
      name: "Field Name",
      type: "text",
      textValue: "",
      numberValue: 0,
      booleanValue: false,
      timeValue: null,
    });
  }

  const { query, results } = useItemSearch(api, { immediate: false });
  const parent = ref();
</script>

<template>
  <BaseContainer v-if="item" class="pb-8">
    <BaseModal v-model="editState.modal">
      <template #title> Attachment Edit </template>

      <FormTextField v-model="editState.title" label="Attachment Title" />
      {{ editState.type }}
      <FormSelect
        v-model:value="editState.type"
        label="Attachment Type"
        value-key="value"
        name="text"
        :items="attachmentOpts"
      />
      <div class="modal-action">
        <BaseButton :loading="editState.loading" @click="updateAttachment"> Update </BaseButton>
      </div>
    </BaseModal>

    <section class="px-3">
      <div class="space-y-4">
        <div class="card bg-base-100 shadow-xl sm:rounded-lg overflow-visible">
          <BaseSectionHeader v-if="item" class="p-5">
            <Icon name="mdi-package-variant" class="-mt-1 mr-2 text-base-content" />
            <span class="text-base-content">
              {{ item.name }}
            </span>
            <p class="text-sm text-base-content font-bold pb-0 mb-0">Quantity {{ item.quantity }}</p>
            <template #after>
              <div class="modal-action mt-3">
                <div class="mr-auto tooltip" data-tip="Hide the cruft! ">
                  <label class="label cursor-pointer mr-auto">
                    <input v-model="preferences.editorSimpleView" type="checkbox" class="toggle toggle-primary" />
                    <span class="label-text ml-4"> Simple View </span>
                  </label>
                </div>
                <BaseButton size="sm" @click="saveItem">
                  <template #icon>
                    <Icon name="mdi-content-save-outline" />
                  </template>
                  Save
                </BaseButton>
              </div>
            </template>
          </BaseSectionHeader>
          <div class="px-5 mb-6 grid md:grid-cols-2 gap-4">
            <FormSelect
              v-if="item"
              v-model="item.location"
              label="Location"
              :items="locations ?? []"
              compare-key="id"
            />
            <FormMultiselect v-model="item.labels" label="Labels" :items="labels ?? []" />

            <Autocomplete
              v-if="!preferences.editorSimpleView"
              v-model="parent"
              v-model:search="query"
              :items="results"
              item-text="name"
              label="Parent Item"
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
        </div>

        <BaseCard>
          <template #title> Custom Fields </template>
          <div class="px-5 divide-y divide-gray-300 space-y-4">
            <div
              v-for="(field, idx) in item.fields"
              :key="`field-${idx}`"
              class="grid grid-cols-2 md:grid-cols-4 gap-2"
            >
              <!-- <FormSelect v-model:value="field.type" label="Field Type" :items="fieldTypes" value-key="value" /> -->
              <FormTextField v-model="field.name" label="Name" />
              <div class="flex items-end col-span-3">
                <FormTextField v-model="field.textValue" label="Value" />
                <div class="tooltip" data-tip="Delete">
                  <button class="btn btn-sm btn-square mb-2 ml-2" @click="item.fields.splice(idx, 1)">
                    <Icon name="mdi-delete" />
                  </button>
                </div>
              </div>
            </div>
          </div>
          <div class="px-5 pb-4 mt-4 flex justify-end">
            <BaseButton size="sm" @click="addField"> Add </BaseButton>
          </div>
        </BaseCard>

        <div
          v-if="!preferences.editorSimpleView"
          ref="attDropZone"
          class="overflow-visible card bg-base-100 shadow-xl sm:rounded-lg"
        >
          <div class="px-4 py-5 sm:px-6">
            <h3 class="text-lg font-medium leading-6">Attachments</h3>
            <p class="text-xs">Changes to attachments will be saved immediately</p>
          </div>
          <div class="border-t border-gray-300 p-4">
            <div v-if="attDropZoneActive" class="grid grid-cols-4 gap-4">
              <DropZone @drop="dropPhoto"> Photo </DropZone>
              <DropZone @drop="dropWarranty"> Warranty </DropZone>
              <DropZone @drop="dropManual"> Manual </DropZone>
              <DropZone @drop="dropAttachment"> Attachment </DropZone>
              <DropZone @drop="dropReceipt"> Receipt </DropZone>
            </div>
            <button
              v-else
              class="h-24 w-full border-2 border-primary border-dashed grid place-content-center"
              @click="clickUpload"
            >
              <input ref="refAttachmentInput" hidden type="file" @change="uploadImage" />
              <p>Drag and drop files here or click to select files</p>
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
                      <Icon name="mdi-delete" />
                    </button>
                  </div>
                  <div class="tooltip" data-tip="Edit">
                    <button class="btn btn-sm btn-square" @click="openAttachmentEditDialog(attachment)">
                      <Icon name="mdi-pencil" />
                    </button>
                  </div>
                </div>
              </li>
            </ul>
          </div>
        </div>

        <div v-if="!preferences.editorSimpleView" class="overflow-visible card bg-base-100 shadow-xl sm:rounded-lg">
          <div class="px-4 py-5 sm:px-6">
            <h3 class="text-lg font-medium leading-6">Purchase Details</h3>
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

        <div v-if="!preferences.editorSimpleView" class="overflow-visible card bg-base-100 shadow-xl sm:rounded-lg">
          <div class="px-4 py-5 sm:px-6">
            <h3 class="text-lg font-medium leading-6">Warranty Details</h3>
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

        <div v-if="!preferences.editorSimpleView" class="overflow-visible card bg-base-100 shadow-xl sm:rounded-lg">
          <div class="px-4 py-5 sm:px-6">
            <h3 class="text-lg font-medium leading-6">Sold Details</h3>
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
  </BaseContainer>
</template>
