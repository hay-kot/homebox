<template>
  <div class="border-b border-base-200 pb-3">
    <h3 class="text-xl font-medium leading-4 text-base-content">
      <slot />
    </h3>
    <p v-if="$slots.description" class="mt-2 max-w-4xl text-sm text-gray-500">
      <slot name="description" />
    </p>
  </div>
</template>
