<script setup lang="ts">
  import { statCardData } from "./statistics";
  import { itemsTable } from "./table";
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";

  definePageMeta({
    middleware: ["auth"],
  });
  useHead({
    title: "Homebox | Home",
  });

  const api = useUserApi();
  const breakpoints = useBreakpoints();

  const locationStore = useLocationStore();
  const locations = computed(() => locationStore.parentLocations);

  const labelsStore = useLabelStore();
  const labels = computed(() => labelsStore.labels);

  const itemTable = itemsTable(api);
  const stats = statCardData(api);
</script>

<template>
  <div>
    <BaseContainer class="flex flex-col gap-12 pb-16">
      <section>
        <Subtitle> Quick Statistics </Subtitle>
        <div class="grid grid-cols-2 gap-2 md:grid-cols-4 md:gap-6">
          <StatCard v-for="(stat, i) in stats" :key="i" :title="stat.label" :value="stat.value" :type="stat.type" />
        </div>
      </section>

      <section>
        <Subtitle> Recently Added </Subtitle>

        <BaseCard v-if="breakpoints.lg">
          <ItemViewTable :items="itemTable.items" />
        </BaseCard>
        <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <ItemCard v-for="item in itemTable.items" :key="item.id" :item="item" />
        </div>
      </section>

      <section>
        <Subtitle> Storage Locations </Subtitle>
        <div class="grid grid-cols-1 sm:grid-cols-2 card md:grid-cols-3 gap-4">
          <LocationCard v-for="location in locations" :key="location.id" :location="location" />
        </div>
      </section>

      <section>
        <Subtitle> Labels </Subtitle>
        <div class="flex gap-4 flex-wrap">
          <LabelChip v-for="label in labels" :key="label.id" size="lg" :label="label" class="shadow-md" />
        </div>
      </section>
    </BaseContainer>
  </div>
</template>
