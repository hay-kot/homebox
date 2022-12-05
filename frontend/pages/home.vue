<script setup lang="ts">
  import { useAuthStore } from "~~/stores/auth";
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

  const locationStore = useLocationStore();
  const locations = computed(() => locationStore.parentLocations);

  const labelsStore = useLabelStore();
  const labels = computed(() => labelsStore.labels);

  const { data: statistics } = useAsyncData(async () => {
    const { data } = await api.stats.group();
    return data;
  });

  const stats = computed(() => {
    return [
      {
        label: "Locations",
        value: statistics.value?.totalLocations || 0,
      },
      {
        label: "Items",
        value: statistics.value?.totalItems || 0,
      },
      {
        label: "Labels",
        value: statistics.value?.totalLabels || 0,
      },
    ];
  });

  const importDialog = ref(false);
  const importCsv = ref(null);
  const importLoading = ref(false);
  const importRef = ref<HTMLInputElement>();
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

  const { data: timeseries } = useAsyncData(async () => {
    const { data } = await api.stats.totalPriceOverTime();
    return data;
  });

  const primary = useCssVar("--p");
  const secondary = useCssVar("--s");
  const accent = useCssVar("--a");
  const neutral = useCssVar("--n");
  const base = useCssVar("--b");

  const chartData = computed(() => {
    let start = timeseries.value?.valueAtStart;

    return {
      labels: timeseries?.value.entries.map(t => new Date(t.date).toDateString()) || [],
      datasets: [
        {
          label: "Purchase Price",
          data:
            timeseries.value?.entries.map(t => {
              start += t.value;
              return start;
            }) || [],
          backgroundColor: primary.value,
          borderColor: primary.value,
        },
      ],
    };
  });

  const { data: donutSeries } = useAsyncData(async () => {
    const { data } = await api.stats.locations();
    return data;
  });

  const donutData = computed(() => {
    return {
      labels: donutSeries.value?.map(l => l.name) || [],
      datasets: [
        {
          label: "Value",
          data: donutSeries.value?.map(l => l.total) || [],
          backgroundColor: [primary.value, secondary.value, neutral.value],
          borderColor: [primary.value, secondary.value, neutral.value],
          hoverOffset: 4,
        },
      ],
    };
  });

  const refDonutEl = ref<HTMLDivElement>(null);

  const donutElWidth = computed(() => {
    return refDonutEl.value?.clientWidth || 0;
  });
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
          <input ref="importRef" type="file" class="hidden" accept=".csv,.tsv" @change="setFile" />

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
          <!-- <template #subtitle> {{ auth.self.isSuperuser ? "Admin" : "User" }} </template> -->
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
              <span class="text-base-900 font-bold">{{ stat.value }}</span>
              {{ " " }}
              <span class="text-base-600">{{ stat.label }}</span>
            </div>
          </div>
        </BaseCard>
      </section>

      <section v-if="timeseries" class="grid grid-cols-6 gap-6">
        <BaseCard class="col-span-4">
          <template #title>Total Asset Value {{ fmtCurrency(timeseries.valueAtEnd) }}</template>
          <div class="p-6 pt-0">
            <ClientOnly>
              <ChartLine chart-id="asd" :height="200" :chart-data="chartData" />
            </ClientOnly>
          </div>
        </BaseCard>
        <BaseCard class="col-span-2">
          <template #title> Asset By Location {{ fmtCurrency(timeseries.valueAtEnd) }}</template>
          <div ref="refDonutEl" class="grid place-content-center h-full">
            <ClientOnly>
              <ChartDonut chart-id="donut" :width="donutElWidth - 50" :height="300" :chart-data="donutData" />
            </ClientOnly>
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
