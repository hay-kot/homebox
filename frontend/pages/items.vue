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

  const api = useUserApi();
  const loading = useMinLoader(2000);
  const items = ref<ItemSummary[]>([]);
  const total = ref(0);

  const page = useRouteQuery("page", 1);
  const perPage = useRouteQuery("perPage", 24);
  const query = useRouteQuery("q", "");
  const advanced = useRouteQuery("advanced", false);
  const includeArchived = useRouteQuery("archived", false);

  const hasNext = computed(() => {
    return page.value * perPage.value < total.value;
  });

  const totalPages = computed(() => {
    return Math.ceil(total.value / perPage.value);
  });

  function next() {
    page.value = Math.min(Math.ceil(total.value / perPage.value), page.value + 1);
  }

  const hasPrev = computed(() => {
    return page.value > 1;
  });

  function prev() {
    page.value = Math.max(1, page.value - 1);
  }

  async function resetPageSearch() {
    page.value = 1;
    items.value = [];
    await search();
  }

  async function search() {
    if (searchLocked.value) {
      return;
    }

    loading.value = true;

    const locations = selectedLocations.value.map(l => l.id);
    const labels = selectedLabels.value.map(l => l.id);

    const { data, error } = await api.items.getAll({
      q: query.value || "",
      locations,
      labels,
      includeArchived: includeArchived.value,
      page: page.value,
      pageSize: perPage.value,
    });
    if (error) {
      page.value--;
      loading.value = false;
      return;
    }

    if (!data.items || data.items.length === 0) {
      page.value--;
    }

    total.value = data.total;
    items.value = data.items;

    loading.value = false;
  }

  const route = useRoute();
  const router = useRouter();

  const queryParamsInitialized = ref(false);

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
  });

  const locationsStore = useLocationStore();
  const locations = computed(() => locationsStore.allLocations);

  const labelStore = useLabelStore();
  const labels = computed(() => labelStore.labels);

  const selectedLocations = ref<LocationOutCount[]>([]);
  const selectedLabels = ref<LabelSummary[]>([]);

  watchPostEffect(() => {
    if (!queryParamsInitialized.value) {
      return;
    }

    const labelIds = selectedLabels.value.map(l => l.id);
    router.push({
      query: {
        ...router.currentRoute.value.query,
        lab: labelIds,
      },
    });
  });

  watchPostEffect(() => {
    if (!queryParamsInitialized.value) {
      return;
    }

    const locIds = selectedLocations.value.map(l => l.id);
    router.push({
      query: {
        ...router.currentRoute.value.query,
        loc: locIds,
      },
    });
  });

  watchEffect(() => {
    if (!advanced.value) {
      selectedLocations.value = [];
      selectedLabels.value = [];
    }
  });

  watchDebounced([selectedLocations, selectedLabels, query, page, perPage], search, { debounce: 250, maxWait: 1000 });
  watch(includeArchived, search);
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
      <BaseSectionHeader> Items </BaseSectionHeader>
      <span class="text-base font-medium"> {{ total }} Results </span>
      <div ref="cardgrid" class="grid mt-4 grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        <TransitionGroup appear name="list">
          <ItemCard v-for="item in items" :key="item.id" :item="item" />
        </TransitionGroup>
        <div class="hidden first:inline text-xl">No Items Found</div>
      </div>
      <div v-if="items.length > 0" class="flex">
        <div class="btn-group mx-auto mt-10">
          <button :disabled="!hasPrev" class="btn" @click="prev">«</button>
          <button class="btn">Page {{ page }} of {{ totalPages }}</button>
          <button :disabled="!hasNext" class="btn" @click="next">»</button>
        </div>
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
