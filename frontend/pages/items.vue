<script setup lang="ts">
  import { watchPostEffect } from "vue";
  import { ItemSummary, LabelSummary, LocationOutCount } from "~~/lib/api/types/data-contracts";
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";

  definePageMeta({
    middleware: ["auth"],
  });

  useHead({
    title: "Homebox | Home",
  });

  const searchLocked = ref(false);
  const queryParamsInitialized = ref(false);
  const initialSearch = ref(true);

  const api = useUserApi();
  const loading = useMinLoader(2000);
  const items = ref<ItemSummary[]>([]);
  const total = ref(0);

  const page1 = useRouteQuery("page", 1);

  const page = computed({
    get: () => page1.value,
    set: value => {
      page1.value = value;
    },
  });

  const pageSize = useRouteQuery("pageSize", 21);
  const query = useRouteQuery("q", "");
  const advanced = useRouteQuery("advanced", false);
  const includeArchived = useRouteQuery("archived", false);

  const totalPages = computed(() => Math.ceil(total.value / pageSize.value));
  const hasNext = computed(() => page.value * pageSize.value < total.value);
  const hasPrev = computed(() => page.value > 1);

  function prev() {
    page.value = Math.max(1, page.value - 1);
  }

  function next() {
    page.value = Math.min(Math.ceil(total.value / pageSize.value), page.value + 1);
  }

  async function resetPageSearch() {
    if (searchLocked.value) {
      return;
    }

    if (!initialSearch.value) {
      page.value = 1;
    }

    items.value = [];
    await search();
  }

  async function search() {
    if (searchLocked.value) {
      return;
    }

    loading.value = true;

    const { data, error } = await api.items.getAll({
      q: query.value || "",
      locations: locIDs.value,
      labels: labIDs.value,
      includeArchived: includeArchived.value,
      page: page.value,
      pageSize: pageSize.value,
    });

    if (error) {
      page.value = Math.max(1, page.value - 1);
      loading.value = false;
      return;
    }

    if (!data.items || data.items.length === 0) {
      page.value = Math.max(1, page.value - 1);
      loading.value = false;
      return;
    }

    total.value = data.total;
    items.value = data.items;

    loading.value = false;
    initialSearch.value = false;
  }

  const route = useRoute();
  const router = useRouter();

  onMounted(async () => {
    loading.value = true;
    // Wait until locations and labels are loaded
    let maxRetry = 10;
    while (!labels.value || !locations.value) {
      await new Promise(resolve => setTimeout(resolve, 100));
      if (maxRetry-- < 0) {
        break;
      }
    }
    searchLocked.value = true;
    const qLoc = route.query.loc as string[];
    if (qLoc) {
      selectedLocations.value = locations.value.filter(l => qLoc.includes(l.id));
    }

    const qLab = route.query.lab as string[];
    if (qLab) {
      selectedLabels.value = labels.value.filter(l => qLab.includes(l.id));
    }

    queryParamsInitialized.value = true;
    searchLocked.value = false;

    // trigger search if no changes
    if (!qLab && !qLoc) {
      search();
    }

    loading.value = false;
    window.scroll({
      top: 0,
      left: 0,
      behavior: "smooth",
    });
  });

  const locationsStore = useLocationStore();
  const locations = computed(() => locationsStore.allLocations);

  const labelStore = useLabelStore();
  const labels = computed(() => labelStore.labels);

  const selectedLocations = ref<LocationOutCount[]>([]);
  const selectedLabels = ref<LabelSummary[]>([]);

  const locIDs = computed(() => selectedLocations.value.map(l => l.id));
  const labIDs = computed(() => selectedLabels.value.map(l => l.id));

  watchPostEffect(() => {
    if (!queryParamsInitialized.value) {
      return;
    }

    router.push({
      query: {
        ...router.currentRoute.value.query,
        lab: labIDs.value,
      },
    });
  });

  watchPostEffect(() => {
    if (!queryParamsInitialized.value) {
      return;
    }

    router.push({
      query: {
        ...router.currentRoute.value.query,
        loc: locIDs.value,
      },
    });
  });

  watchEffect(() => {
    if (!advanced.value) {
      selectedLocations.value = [];
      selectedLabels.value = [];
    }
  });

  // resetPageHash computes a JSON string that is used to detect if the search
  // parameters have changed. If they have changed, the page is reset to 1.
  const resetPageHash = computed(() => {
    const map = {
      q: query.value,
      includeArchived: includeArchived.value,
      locations: locIDs.value,
      labels: labIDs.value,
    };

    return JSON.stringify(map);
  });

  watchDebounced(resetPageHash, resetPageSearch, { debounce: 250, maxWait: 1000 });

  watchDebounced([page, pageSize], search, { debounce: 250, maxWait: 1000 });
</script>

<template>
  <BaseContainer class="mb-16">
    <FormTextField v-model="query" placeholder="Search" />
    <div class="flex mt-1">
      <label class="ml-auto label cursor-pointer">
        <input v-model="advanced" type="checkbox" class="toggle toggle-primary" />
        <span class="label-text text-base-content ml-2"> Filters </span>
      </label>
    </div>
    <BaseCard v-if="advanced" class="my-1 overflow-visible">
      <template #title> Filters </template>
      <template #subtitle>
        Location and label filters use the 'OR' operation. If more than one is selected only one will be required for a
        match
      </template>
      <div class="px-4 pb-4">
        <FormMultiselect v-model="selectedLabels" label="Labels" :items="labels ?? []" />
        <FormMultiselect v-model="selectedLocations" label="Locations" :items="locations ?? []" />
        <div class="flex pb-2 pt-5">
          <label class="label cursor-pointer mr-auto">
            <input v-model="includeArchived" type="checkbox" class="toggle toggle-primary" />
            <span class="label-text ml-4"> Include Archived Items </span>
          </label>
          <Spacer />
        </div>
      </div>
    </BaseCard>
    <section class="mt-10">
      <BaseSectionHeader ref="itemsTitle"> Items </BaseSectionHeader>
      <p class="text-base font-medium flex items-center">
        {{ total }} Results
        <span class="text-base ml-auto"> Page {{ page }} of {{ totalPages }}</span>
      </p>

      <div ref="cardgrid" class="grid mt-4 grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        <ItemCard v-for="item in items" :key="item.id" :item="item" />

        <div class="hidden first:inline text-xl">No Items Found</div>
      </div>
      <div v-if="items.length > 0 && (hasNext || hasPrev)" class="mt-10 flex gap-2 flex-col items-center">
        <div class="flex">
          <div class="btn-group">
            <button :disabled="!hasPrev" class="btn text-no-transform" @click="prev">
              <Icon class="mr-1 h-6 w-6" name="mdi-chevron-left" />
              Prev
            </button>
            <button v-if="hasPrev" class="btn text-no-transform" @click="page = 1">First</button>
            <button v-if="hasNext" class="btn text-no-transform" @click="page = totalPages">Last</button>
            <button :disabled="!hasNext" class="btn text-no-transform" @click="next">
              Next
              <Icon class="ml-1 h-6 w-6" name="mdi-chevron-right" />
            </button>
          </div>
        </div>
        <p class="text-sm font-bold">Page {{ page }} of {{ totalPages }}</p>
      </div>
    </section>
  </BaseContainer>
</template>

<style lang="css">
  .list-move,
  .list-enter-active,
  .list-leave-active {
    transition: all 0.25s ease;
  }

  .list-enter-from,
  .list-leave-to {
    opacity: 0;
    transform: translateY(30px);
  }

  .list-leave-active {
    position: absolute;
  }
</style>
