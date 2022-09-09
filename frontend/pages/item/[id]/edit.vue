<script setup lang="ts">
  definePageMeta({
    layout: "home",
  });

  const route = useRoute();
  const api = useUserApi();
  const toast = useNotifier();

  const itemId = computed<string>(() => route.params.id as string);

  const { data: item } = useAsyncData(async () => {
    const { data, error } = await api.items.get(itemId.value);
    if (error) {
      toast.error("Failed to load item");
      navigateTo("/home");
      return;
    }
    return data;
  });

  type FormField = {
    type: "text" | "textarea" | "select" | "date";
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
    <div class="space-y-4">
      <div class="overflow-hidden card bg-base-100 shadow-xl sm:rounded-lg">
        <div class="px-4 py-5 sm:px-6">
          <h3 class="text-lg font-medium leading-6">Item Details</h3>
        </div>
        <div class="border-t border-gray-300 sm:p-0">
          <div v-for="field in mainFields" :key="field.ref" class="sm:divide-y sm:divide-gray-300 grid grid-cols-1">
            <div class="pt-2 pb-4 sm:px-6 border-b border-gray-300">
              <FormTextArea v-if="field.type === 'textarea'" v-model="item[field.ref]" :label="field.label" inline />
              <FormTextField v-else-if="field.type === 'text'" v-model="item[field.ref]" :label="field.label" inline />
              <FormDatePicker v-else-if="field.type === 'date'" v-model="item[field.ref]" :label="field.label" inline />
            </div>
          </div>
        </div>
      </div>

      <div class="overflow-visible card bg-base-100 shadow-xl sm:rounded-lg">
        <div class="px-4 py-5 sm:px-6">
          <h3 class="text-lg font-medium leading-6">Purchase Details</h3>
        </div>
        <div class="border-t border-gray-300 sm:p-0">
          <div v-for="field in purchaseFields" :key="field.ref" class="sm:divide-y sm:divide-gray-300 grid grid-cols-1">
            <div class="pt-2 pb-4 sm:px-6 border-b border-gray-300">
              <FormTextArea v-if="field.type === 'textarea'" v-model="item[field.ref]" :label="field.label" inline />
              <FormTextField v-else-if="field.type === 'text'" v-model="item[field.ref]" :label="field.label" inline />
              <FormDatePicker v-else-if="field.type === 'date'" v-model="item[field.ref]" :label="field.label" inline />
            </div>
          </div>
        </div>
      </div>

      <div class="overflow-visible card bg-base-100 shadow-xl sm:rounded-lg">
        <div class="px-4 py-5 sm:px-6">
          <h3 class="text-lg font-medium leading-6">Sold Details</h3>
        </div>
        <div class="border-t border-gray-300 sm:p-0">
          <div v-for="field in soldFields" :key="field.ref" class="sm:divide-y sm:divide-gray-300 grid grid-cols-1">
            <div class="pt-2 pb-4 sm:px-6 border-b border-gray-300">
              <FormTextArea v-if="field.type === 'textarea'" v-model="item[field.ref]" :label="field.label" inline />
              <FormTextField v-else-if="field.type === 'text'" v-model="item[field.ref]" :label="field.label" inline />
              <FormDatePicker v-else-if="field.type === 'date'" v-model="item[field.ref]" :label="field.label" inline />
            </div>
          </div>
        </div>
      </div>
    </div>
  </BaseContainer>
</template>
