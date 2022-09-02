<template>
  <button
    :disabled="disabled || loading"
    class="btn"
    ref="submitBtn"
    :class="{
      loading: loading,
    }"
  >
    <label v-if="$slots.icon" class="swap swap-rotate mr-2" :class="{ 'swap-active': isHover }">
      <slot name="icon" />
    </label>
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

  const submitBtn = ref(null);
  const isHover = useElementHover(submitBtn);
</script>
