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
  const results = ref<ItemSummary[]>([]);

  const query = useRouteQuery("q", "");
  const advanced = useRouteQuery("advanced", false);
  const includeArchived = useRouteQuery("archived", false);

  async function search() {
    if (searchLocked.value) {
      return;
    }

    loading.value = true;

    const locations = selectedLocations.value.map(l => l.id);
    const labels = selectedLabels.value.map(l => l.id);

    const { data, error } = await api.items.getAll({
      q: query.value,
      locations,
      labels,
      includeArchived: includeArchived.value,
    });
    if (error) {
      loading.value = false;
      return;
    }

    results.value = data.items;
    loading.value = false;
  }

  const route = useRoute();
  const router = useRouter();

  const queryParamsInitialized = ref(false);

  onMounted(async () => {
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
  });

  const locationsStore = useLocationStore();
  const locations = computed(() => locationsStore.locations);

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

  watchDebounced([selectedLocations, selectedLabels, query], search, { debounce: 250, maxWait: 1000 });
  watch(includeArchived, search);
</script>

<template>
  <BaseContainer class="mb-16">
    <FormTextField v-model="query" placeholder="Search" />
    <div class="flex mt-1">
      <label class="ml-auto label cursor-pointer">
        <input v-model="advanced" type="checkbox" class="toggle toggle-primary" />
        <span class="label-text text-neutral-content ml-2"> Filters </span>
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
      <BaseSectionHeader class="mb-5"> Items </BaseSectionHeader>
      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <TransitionGroup name="list">
          <ItemCard v-for="item in results" :key="item.id" :item="item" />
        </TransitionGroup>
        <div class="hidden first:inline text-xl">No Items Found</div>
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
