<template>
  <button
    :disabled="disabled || loading"
    class="btn"
    :class="{
      loading: loading,
    }"
  >
    <slot />
  </button>
</template>

<script setup lang="ts">
  defineProps({
    loading: {
      type: Boolean,
      default: false,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
  });
</script>
