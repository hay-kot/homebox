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

  // const purchasePriceOverTime = purchasePriceOverTimeChart(api);

  // const inventoryByLocation = inventoryByLocationChart(api);

  // const refDonutEl = ref<HTMLDivElement>();

  // const donutElWidth = computed(() => {
  //   return refDonutEl.value?.clientWidth || 0;
  // });
</script>

<template>
  <div>
    <BaseContainer class="flex flex-col gap-12 pb-16">
      <!-- <section class="grid grid-cols-6 gap-6">
        <article class="col-span-4">
          <Subtitle> Inventory Value Over Time </Subtitle>
          <BaseCard>
            <div class="p-10 h-[300px]">
              <ClientOnly>
                <ChartLine :chart-data="purchasePriceOverTime" />
              </ClientOnly>
            </div>
          </BaseCard>
        </article>
        <article class="col-span-2">
          <Subtitle>
            Inventory By
            <span class="btn-group">
              <button class="btn btn-xs btn-active text-no-transform">Locations</button>
              <button class="btn btn-xs text-no-transform">Labels</button>
            </span>
          </Subtitle>
          <BaseCard class="h-[300px]">
            <div ref="refDonutEl" class="grid place-content-center h-full">
              <ClientOnly>
                <ChartDonut
                  chart-id="donut"
                  :width="donutElWidth - 50"
                  :height="265"
                  :chart-data="inventoryByLocation"
                />
              </ClientOnly>
            </div>
          </BaseCard>
        </article>
      </section> -->

      <section>
        <Subtitle> Quick Statistics </Subtitle>
        <div class="grid grid-cols-2 gap-2 md:grid-cols-4 md:gap-6">
          <StatCard v-for="(stat, i) in stats" :key="i" :title="stat.label" :value="stat.value" :type="stat.type" />
        </div>
      </section>

      <section>
        <Subtitle> Recently Added </Subtitle>

        <BaseCard v-if="breakpoints.lg">
          <Table :headers="itemTable.headers" :data="itemTable.items">
            <template #cell-name="{ item }">
              <NuxtLink :to="`/item/${item.id}`">
                {{ item.name }}
              </NuxtLink>
            </template>
            <template #cell-warranty="{ item }">
              <Icon v-if="item.warranty" name="mdi-check" class="text-green-500 h-5 w-5" />
              <Icon v-else name="mdi-close" class="text-red-500 h-5 w-5" />
            </template>
            <template #cell-purchasePrice="{ item }">
              <Currency :amount="item.purchasePrice" />
            </template>
            <template #cell-location_name="{ item }">
              <NuxtLink class="badge badge-sm badge-primary p-3" :to="`/location/${item.location.id}`">
                {{ item.location?.name }}
              </NuxtLink>
            </template>
          </Table>
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
