<template>
  <NuxtLink
    class="group card bg-neutral text-neutral-content hover:bg-primary transition-colors duration-300"
    :to="`/item/${item.id}`"
  >
    <div class="card-body py-4 px-6">
      <h2 class="card-title">
        <Icon name="mdi-package-variant" />
        {{ item.name }}
      </h2>
      <p>{{ item.description }}</p>
      <div class="flex gap-2 flex-wrap justify-end">
        <LabelChip v-for="label in item.labels" :label="label" class="badge-primary group-hover:badge-secondary" />
      </div>
    </div>
  </NuxtLink>
</template>

<script setup lang="ts">
  import { Item } from '~~/lib/api/classes/items';

  defineProps({
    item: {
      type: Object as () => Item,
      required: true,
    },
  });
</script>

<style scoped></style>
