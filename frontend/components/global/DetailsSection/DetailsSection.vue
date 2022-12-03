<template>
  <div class="border-t border-gray-300 px-4 py-5 sm:p-0">
    <dl class="sm:divide-y sm:divide-gray-300">
      <div v-for="(detail, i) in details" :key="i" class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
        <dt class="text-sm font-medium text-base-content">
          {{ detail.name }}
        </dt>
        <dd class="mt-1 text-sm text-base-content sm:col-span-2 sm:mt-0">
          <slot :name="detail.slot || detail.name" v-bind="{ detail }">
            <DateTime v-if="detail.type == 'date'" :date="detail.text" />
            <Currency v-else-if="detail.type == 'currency'" :amount="detail.text" />
            <template v-else-if="detail.type === 'link'">
              <div class="tooltip tooltip-primary tooltip-right" :data-tip="detail.href">
                <a class="btn btn-primary btn-xs" :href="detail.href" target="_blank">
                  <Icon name="mdi-open-in-new" class="mr-2 swap-on"></Icon>
                  {{ detail.text }}
                </a>
              </div>
            </template>
            <template v-else-if="detail.type === 'markdown'">
              <ClientOnly>
                <Markdown :source="detail.text" />
              </ClientOnly>
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
  import type { AnyDetail, Detail } from "./types";

  defineProps({
    details: {
      type: Object as () => (Detail | AnyDetail)[],
      required: true,
    },
  });
</script>
