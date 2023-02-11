<template>
  <div class="grid grid-cols-1 md:grid-cols-4 gap-10 py-6">
    <div class="col-span-3">
      <h4 class="mb-1 text-lg font-semibold">
        <slot name="title"></slot>
      </h4>
      <p class="text-sm">
        <slot></slot>
      </p>
    </div>
    <BaseButton class="btn-primary mt-auto" @click="$emit('action')">
      <slot name="button">
        <slot name="title"></slot>
      </slot>
    </BaseButton>
  </div>
</template>
