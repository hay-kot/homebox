<script setup lang="ts">
  import { ItemSummary } from "~~/lib/api/types/data-contracts";
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";

  definePageMeta({
    middleware: ["auth"],
  });

  useHead({
    title: "Homebox | Home",
  });

  const api = useUserApi();

  const query = ref("");
  const loading = useMinLoader(2000);
  const results = ref<ItemSummary[]>([]);

  async function search() {
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

  onMounted(() => {
    search();
  });

  const locationsStore = useLocationStore();
  const locations = computed(() => locationsStore.locations);

  const labelStore = useLabelStore();
  const labels = computed(() => labelStore.labels);

  const advanced = ref(false);
  const selectedLocations = ref([]);
  const selectedLabels = ref([]);
  const includeArchived = ref(false);

  watchEffect(() => {
    if (!advanced.value) {
      selectedLocations.value = [];
      selectedLabels.value = [];
    }
  });

  watchDebounced(query, search, { debounce: 250, maxWait: 1000 });
  watchDebounced(selectedLocations, search, { debounce: 250, maxWait: 1000 });
  watchDebounced(selectedLabels, search, { debounce: 250, maxWait: 1000 });
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
