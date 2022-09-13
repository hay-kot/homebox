<script setup lang="ts">
  import { ItemUpdate } from "~~/lib/api/types/data-contracts";
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";

  definePageMeta({
    layout: "home",
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
  ];

  const purchaseFields: FormField[] = [
    {
      type: "text",
      label: "Purchased From",
      ref: "purchaseFrom",
    },
    {
      type: "text",
      label: "Purchased Price",
      ref: "purchasePrice",
    },
    {
      type: "date",
      label: "Purchased At",
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
</script>

<template>
  <BaseContainer v-if="item" class="pb-8">
    <section class="px-3">
      <div class="space-y-4">
        <div class="overflow-hidden card bg-base-100 shadow-xl sm:rounded-lg">
          <BaseSectionHeader v-if="item" class="p-5">
            <Icon name="mdi-package-variant" class="-mt-1 mr-2 text-gray-600" />
            <span class="text-gray-600">
              {{ item.name }}
            </span>
            <p class="text-sm text-gray-600 font-bold pb-0 mb-0">Quantity {{ item.quantity }}</p>
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
            <FormSelect v-model="item.location" label="Location" :items="locations ?? []" select-first />
            <FormMultiselect v-model="item.labels" label="Labels" :items="labels ?? []" />
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
