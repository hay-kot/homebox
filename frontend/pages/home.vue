<script setup lang="ts">
  import { useAuthStore } from "~~/stores/auth";
  import { useItemStore } from "~~/stores/items";
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";

  definePageMeta({
    middleware: ["auth"],
  });

  useHead({
    title: "Homebox | Home",
  });

  const api = useUserApi();

  const auth = useAuthStore();

  const itemsStore = useItemStore();
  const items = computed(() => itemsStore.items);

  const locationStore = useLocationStore();
  const locations = computed(() => locationStore.locations);

  const labelsStore = useLabelStore();
  const labels = computed(() => labelsStore.labels);

  const totalItems = computed(() => items.value?.length || 0);
  const totalLocations = computed(() => locations.value?.length || 0);
  const totalLabels = computed(() => labels.value?.length || 0);

  const stats = [
    {
      label: "Locations",
      value: totalLocations,
    },
    {
      label: "Items",
      value: totalItems,
    },
    {
      label: "Labels",
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
  }

  const toast = useNotifier();

  function openDialog() {
    importDialog.value = true;
  }

  function uploadCsv() {
    importRef.value.click();
  }

  const eventBus = useEventBus();

  async function submitCsvFile() {
    importLoading.value = true;

    const { error } = await api.items.import(importCsv.value);

    if (error) {
      toast.error("Import failed. Please try again later.");
    }

    // Reset
    importDialog.value = false;
    importLoading.value = false;
    importCsv.value = null;
    importRef.value.value = null;

    eventBus.emit(EventTypes.ClearStores);
  }
</script>

<template>
  <div>
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
    <BaseContainer class="flex flex-col gap-16 pb-16">
      <section>
        <BaseCard>
          <template #title> Welcome Back, {{ auth.self ? auth.self.name : "Username" }} </template>
          <template #subtitle> {{ auth.self.isSuperuser ? "Admin" : "User" }} </template>
          <template #title-actions>
            <div class="flex justify-end gap-2">
              <div class="tooltip" data-tip="Import CSV File">
                <button class="btn btn-primary btn-sm" @click="openDialog">
                  <Icon name="mdi-database" class="mr-2"></Icon>
                  Import
                </button>
              </div>
              <BaseButton type="button" size="sm" to="/profile">
                <Icon class="h-5 w-5 mr-2" name="mdi-person" />
                Profile
              </BaseButton>
            </div>
          </template>

          <div
            class="grid grid-cols-1 divide-y divide-base-300 border-t border-base-300 sm:grid-cols-3 sm:divide-y-0 sm:divide-x"
          >
            <div v-for="stat in stats" :key="stat.label" class="px-6 py-5 text-center text-sm font-medium">
              <span class="text-base-900 font-bold">{{ stat.value.value }}</span>
              {{ " " }}
              <span class="text-base-600">{{ stat.label }}</span>
            </div>
          </div>
        </BaseCard>
      </section>

      <section>
        <BaseSectionHeader class="mb-5"> Storage Locations </BaseSectionHeader>
        <div class="grid grid-cols-1 sm:grid-cols-2 card md:grid-cols-3 gap-4">
          <LocationCard v-for="location in locations" :key="location.id" :location="location" />
        </div>
      </section>

      <section>
        <BaseSectionHeader class="mb-5"> Labels </BaseSectionHeader>
        <div class="flex gap-2 flex-wrap">
          <LabelChip v-for="label in labels" :key="label.id" size="lg" :label="label" />
        </div>
      </section>
    </BaseContainer>
  </div>
</template>
