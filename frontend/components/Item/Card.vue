<template>
  <NuxtLink class="group card rounded-md border border-gray-300" :to="`/item/${item.id}`">
    <div class="relative h-[200px]">
      <img v-if="imageUrl" class="h-[200px] w-full object-cover rounded-t shadow-sm border-gray-300" :src="imageUrl" />
      <div class="absolute bottom-1 left-1">
        <NuxtLink
          v-if="item.location"
          class="text-sm hover:link badge shadow-md rounded-md"
          :to="`/location/${item.location.id}`"
        >
          {{ item.location.name }}
        </NuxtLink>
      </div>
    </div>
    <div class="rounded-b p-4 pt-2 flex-grow col-span-4 flex flex-col gap-y-1 bg-base-100">
      <h2 class="text-lg font-bold two-line">{{ item.name }}</h2>
      <div class="divider my-0"></div>
      <div class="flex justify-between gap-2">
        <div v-if="item.insured" class="tooltip z-10" data-tip="Insured">
          <MdiShieldCheck class="h-5 w-5 text-primary" />
        </div>
        <div class="tooltip" data-tip="Quantity">
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
  import type { ItemOut, ItemSummary } from "~~/lib/api/types/data-contracts";
  import MdiShieldCheck from "~icons/mdi/shield-check";

  const api = useUserApi();

  const imageUrl = computed(() => {
    if (!props.item.imageId) {
      return "/no-image.jpg";
    }

    return api.authURL(`/items/${props.item.id}/attachments/${props.item.imageId}`);
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
