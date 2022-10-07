<template>
  <div class="border-t border-gray-300 px-4 py-5 sm:p-0">
    <dl class="sm:divide-y sm:divide-gray-300">
      <div v-for="(detail, i) in details" :key="i" class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
        <dt class="text-sm font-medium text-base-content">
          {{ detail.name }}
        </dt>
        <dd class="mt-1 text-sm text-base-content sm:col-span-2 sm:mt-0">
          <slot :name="detail.slot || detail.name" v-bind="{ detail }">
            <template v-if="detail.type == 'date'">
              <DateTime :date="detail.text" />
            </template>
            <template v-else>
              {{ detail.text }}
            </template>
          </slot>
        </dd>
      </div>
    </dl>
  </div>
</template>

<script setup lang="ts">
  import type { DateDetail, Detail } from "./types";

  defineProps({
    details: {
      type: Object as () => (Detail | DateDetail)[],
      required: true,
    },
  });
</script>
