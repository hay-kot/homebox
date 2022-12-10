<template>
  <NuxtLink
    v-if="to"
    v-bind="attributes"
    ref="submitBtn"
    class="btn"
    :class="{
      loading: loading,
      'btn-sm': size === 'sm',
      'btn-lg': size === 'lg',
    }"
    :style="upper ? '' : 'text-transform: none'"
  >
    <label v-if="$slots.icon" class="swap swap-rotate mr-2" :class="{ 'swap-active': isHover }">
      <slot name="icon" />
    </label>
    <slot />
  </NuxtLink>
  <button
    v-else
    v-bind="attributes"
    ref="submitBtn"
    class="btn"
    :class="{
      loading: loading,
      'btn-sm': size === 'sm',
      'btn-lg': size === 'lg',
    }"
    :style="upper ? '' : 'text-transform: none'"
  >
    <label v-if="$slots.icon" class="swap swap-rotate mr-2" :class="{ 'swap-active': isHover }">
      <slot name="icon" />
    </label>
    <slot />
  </button>
</template>

<script setup lang="ts">
  type Sizes = "sm" | "md" | "lg";

  const props = defineProps({
    upper: {
      type: Boolean,
      default: false,
    },
    loading: {
      type: Boolean,
      default: false,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    size: {
      type: String as () => Sizes,
      default: "md",
    },
    to: {
      type: String as () => string | null,
      default: null,
    },
  });

  const attributes = computed(() => {
    if (props.to) {
      return {
        to: props.to,
      };
    }
    return {
      disabled: props.disabled || props.loading,
    };
  });

  const submitBtn = ref(null);
  const isHover = useElementHover(submitBtn);
</script>
