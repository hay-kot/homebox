<script lang="ts" setup>
  defineProps({
    is: {
      type: String,
      default: 'div',
    },
  });
</script>

<template>
  <component :is="is" class="container max-w-6xl mx-auto px-4">
    <slot />
  </component>
</template>
