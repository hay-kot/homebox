<script lang="ts" setup>
  defineProps({
    cmp: {
      type: String,
      default: "div",
    },
  });
</script>

<template>
  <component :is="cmp" class="container max-w-6xl mx-auto px-3">
    <slot />
  </component>
</template>
