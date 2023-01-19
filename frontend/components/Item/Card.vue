<template>
  <NuxtLink class="group card rounded-md" :to="`/item/${item.id}`">
    <div class="rounded-t flex flex-col justify-center bg-neutral text-neutral-content p-5">
      <h2 class="text-base mb-2 last:mb-0 font-bold two-line">{{ item.name }}</h2>
      <NuxtLink
        v-if="item.location"
        class="inline-flex text-sm items-center hover:link"
        :to="`/location/${item.location.id}`"
      >
        <Icon name="heroicons-map-pin" class="mr-1 h-4 w-4"></Icon>
        <span>
          {{ item.location.name }}
        </span>
      </NuxtLink>
    </div>
    <div class="rounded-b p-4 pt-2 flex-grow col-span-4 flex flex-col gap-y-2 bg-base-100">
      <div class="flex justify-between gap-2">
        <div class="mr-auto tooltip tooltip-tip" data-tip="Purchase Price">
          <span class="badge badge-sm badge-ghost h-5">
            <Currency :amount="item.purchasePrice" />
          </span>
        </div>
        <div v-if="item.insured" class="tooltip z-10" data-tip="Insured">
          <Icon class="h-5 w-5 text-primary" name="mdi-shield-check" />
        </div>
        <div v-if="item.quantity > 1" class="tooltip" data-tip="Quantity">
          <span class="badge h-5 w-5 badge-primary badge-sm text-xs">
            {{ item.quantity }}
          </span>
        </div>
      </div>
      <Markdown class="mb-2 text-clip three-line" :source="item.description" />

      <div class="flex gap-2 flex-wrap -mr-1 mt-auto justify-end">
        <LabelChip v-for="label in top3" :key="label.id" :label="label" size="sm" />
      </div>
    </div>
  </NuxtLink>
</template>

<script setup lang="ts">
  import { ItemOut, ItemSummary } from "~~/lib/api/types/data-contracts";

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

<style lang="css">
  .three-line {
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 3;
    line-clamp: 3;
    -webkit-box-orient: vertical;
  }

  .two-line {
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    line-clamp: 2;
    -webkit-box-orient: vertical;
  }
</style>
