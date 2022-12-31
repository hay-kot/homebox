<template>
  <NuxtLink
    class="group card bg-neutral text-neutral-content hover:bg-primary transition-colors duration-300"
    :to="`/item/${item.id}`"
  >
    <div class="card-body p-4 gap-3">
      <h2 class="card-title">
        {{ item.name }}
        <Icon v-if="item.archived" class="ml-auto" name="mdi-archive-outline" />
        <Icon v-else class="ml-auto" name="mdi-package-variant" />
      </h2>
      <div v-if="colOne" class="flex gap-x-2 items-center justify-between">
        <NuxtLink
          v-if="item.location"
          class="badge badge-primary group-hover:badge-ghost"
          :to="`/location/${item.location.id}`"
        >
          <Icon name="heroicons-map-pin" class="mr-2 swap-on"></Icon>
          {{ item.location.name }}
        </NuxtLink>
        <div class="flex gap-2 ml-auto items-end">
          <div v-if="item.purchasePrice" class="tooltip" data-tip="Purchase Price">
            <span class="badge badge-ghost">
              <Currency :amount="item.purchasePrice" />
            </span>
          </div>
          <div v-if="item.insured" class="tooltip" data-tip="Insured">
            <Icon class="h-5 w-5 text-base-200" name="mdi-shield-check" />
          </div>
          <div v-if="item.quantity > 1" class="tooltip" data-tip="Quantity">
            <span class="badge badge-ghost">
              {{ item.quantity }}
            </span>
          </div>
        </div>
      </div>
      <span
        class="w-[100%] group-hover:bg-neutral-content h-[3px] transition-colors duration-300 rounded-box bg-primary"
      ></span>
      <div class="flex gap-2 flex-wrap justify-end">
        <LabelChip v-for="label in top3" :key="label.id" :label="label" size="sm" />
      </div>
    </div>
  </NuxtLink>
</template>

<script setup lang="ts">
  import { ItemOut, ItemSummary } from "~~/lib/api/types/data-contracts";

  const colOne = computed(() => {
    return props.item.location || props.item.purchasePrice;
  });

  const top3 = computed(() => {
    return props.item.labels.slice(0, 3) || [];
  });

  const props = defineProps({
    item: {
      type: Object as () => ItemOut | ItemSummary,
      required: true,
    },
  });
</script>
