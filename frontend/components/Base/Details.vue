<template>
  <div class="overflow-hidden card bg-base-100 shadow-xl sm:rounded-lg">
    <div class="px-4 py-5 sm:px-6">
      <h3 class="text-lg font-medium leading-6">
        <slot name="title"></slot>
      </h3>
      <p v-if="$slots.subtitle" class="mt-1 max-w-2xl text-sm text-gray-500">
        <slot name="subtitle"></slot>
      </p>
    </div>
    <div class="border-t border-gray-300 px-4 py-5 sm:p-0">
      <dl class="sm:divide-y sm:divide-gray-300">
        <div v-for="(dValue, dKey) in details" :key="dKey" class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
          <dt class="text-sm font-medium text-gray-500">
            {{ dKey }}
          </dt>
          <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
            <slot :name="rmSpace(dKey)" v-bind="{ key: dKey, value: dValue }">
              {{ dValue }}
            </slot>
          </dd>
        </div>
      </dl>
    </div>
  </div>
</template>

<script setup lang="ts">
  type StringLike = string | number | boolean;

  function rmSpace(str: string) {
    return str.replace(" ", "");
  }

  defineProps({
    details: {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      type: Object as () => Record<string, StringLike | any>,
      required: true,
    },
  });
</script>

<style scoped></style>
