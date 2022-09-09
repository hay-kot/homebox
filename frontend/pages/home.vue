<script setup lang="ts">
  definePageMeta({
    layout: 'home',
  });
  useHead({
    title: 'Homebox | Home',
  });

  const api = useUserApi();

  const { data: locations } = useAsyncData('locations', async () => {
    const { data } = await api.locations.getAll();
    return data.items;
  });

  const { data: labels } = useAsyncData('labels', async () => {
    const { data } = await api.labels.getAll();
    return data.items;
  });

  const { data: items } = useAsyncData('items', async () => {
    const { data } = await api.items.getAll();
    return data.items;
  });

  const totalItems = computed(() => items.value?.length || 0);
  const totalLocations = computed(() => locations.value?.length || 0);
  const totalLabels = computed(() => labels.value?.length || 0);

  const stats = [
    {
      label: 'Locations',
      value: totalLocations,
    },
    {
      label: 'Items',
      value: totalItems,
    },
    {
      label: 'Labels',
      value: totalLabels,
    },
  ];

  const importDialog = ref(false);
  const importCsv = ref(null);
  const importLoading = ref(false);
  const importRef = ref<HTMLInputElement>(null);
  whenever(
    () => !importDialog.value,
    () => {
      importCsv.value = null;
    }
  );

  function setFile(e: Event & { target: HTMLInputElement }) {
    importCsv.value = e.target.files[0];
    console.log('importCsv.value', importCsv.value);
  }

  const toast = useNotifier();

  function openDialog() {
    importDialog.value = true;
  }

  function uploadCsv() {
    importRef.value.click();
  }

  async function submitCsvFile() {
    importLoading.value = true;

    const { error } = await api.items.import(importCsv.value);

    if (error) {
      toast.error('Import failed. Please try again later.');
    }

    // Reset
    importDialog.value = false;
    importLoading.value = false;
    importCsv.value = null;
    importRef.value.value = null;
  }
</script>

<template>
  <BaseContainer class="space-y-16 pb-16">
    <BaseModal v-model="importDialog">
      <template #title> Import CSV File </template>
      <p>
        Import a CSV file containing your items, labels, and locations. See documentation for more information on the
        required format.
      </p>

      <form @submit.prevent="submitCsvFile">
        <div class="flex flex-col gap-2 py-6">
          <input ref="importRef" type="file" class="hidden" accept=".csv" @change="setFile" />
          <BaseButton type="button" @click="uploadCsv">
            <Icon class="h-5 w-5 mr-2" name="mdi-upload" />
            Upload
          </BaseButton>
          <p class="text-center pt-4 -mb-5">
            {{ importCsv?.name }}
          </p>
        </div>

        <div class="modal-action">
          <BaseButton type="submit" :disabled="!importCsv"> Submit </BaseButton>
        </div>
      </form>
    </BaseModal>

    <section aria-labelledby="profile-overview-title" class="mt-8">
      <div class="overflow-hidden rounded-lg bg-white shadow">
        <h2 class="sr-only" id="profile-overview-title">Profile Overview</h2>
        <div class="bg-white p-6">
          <div class="sm:flex sm:items-center sm:justify-between">
            <div class="sm:flex sm:space-x-5">
              <div class="mt-4 text-center sm:mt-0 sm:pt-1 sm:text-left">
                <p class="text-sm font-medium text-gray-600">Welcome back,</p>
                <p class="text-xl font-bold text-gray-900 sm:text-2xl">Hayden Kotelman</p>
                <p class="text-sm font-medium text-gray-600">User</p>
              </div>
            </div>
            <div class="mt-5 flex justify-center sm:mt-0">
              <a
                href="#"
                class="flex items-center justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50"
                >View profile</a
              >
            </div>
          </div>
        </div>
        <div
          class="grid grid-cols-1 divide-y divide-gray-200 border-t border-gray-200 bg-gray-50 sm:grid-cols-3 sm:divide-y-0 sm:divide-x"
        >
          <div v-for="stat in stats" :key="stat.label" class="px-6 py-5 text-center text-sm font-medium">
            <span class="text-gray-900">{{ stat.value.value }}</span>
            {{ ' ' }}
            <span class="text-gray-600">{{ stat.label }}</span>
          </div>
        </div>
      </div>
    </section>

    <section>
      <BaseSectionHeader class="mb-5"> Storage Locations </BaseSectionHeader>
      <div class="grid grid-cols-1 sm:grid-cols-2 card md:grid-cols-3 gap-4">
        <LocationCard v-for="location in locations" :location="location" />
      </div>
    </section>

    <section>
      <BaseSectionHeader class="mb-5">
        Items
        <template #description>
          <div class="tooltip" data-tip="Import CSV File">
            <button @click="openDialog" class="btn btn-primary btn-sm">
              <Icon name="mdi-database" class="mr-2"></Icon>
              Import
            </button>
          </div>
        </template>
      </BaseSectionHeader>
      <div class="grid sm:grid-cols-2 gap-4">
        <ItemCard v-for="item in items" :item="item" />
      </div>
    </section>

    <section>
      <BaseSectionHeader class="mb-5"> Labels </BaseSectionHeader>
      <div class="flex gap-2 flex-wrap">
        <LabelChip v-for="label in labels" size="lg" :label="label" />
      </div>
    </section>
  </BaseContainer>
</template>
