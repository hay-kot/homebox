<template>
  <div class="card bg-base-100 shadow-xl sm:rounded-lg">
    <div v-if="$slots.title" class="px-4 py-5 sm:px-6">
      <h3 class="text-lg font-medium leading-6">
        <slot name="title"></slot>
      </h3>
      <p v-if="$slots.subtitle" class="mt-1 max-w-2xl text-sm text-gray-500">
        <slot name="subtitle"></slot>
      </p>
      <template v-if="$slots['title-actions']">
        <slot name="title-actions"></slot>
      </template>
    </div>
    <slot />
  </div>
</template>
