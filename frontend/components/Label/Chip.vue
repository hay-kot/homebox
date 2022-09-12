<script setup lang="ts">
  import { LabelOut, LabelSummary } from "~~/lib/api/types/data-contracts";

  export type sizes = "sm" | "md" | "lg";
  defineProps({
    label: {
      type: Object as () => LabelOut | LabelSummary,
      required: true,
    },
    size: {
      type: String as () => sizes,
      default: "md",
    },
  });

  const badge = ref(null);
  const isHover = useElementHover(badge);
  const { focused } = useFocus(badge);

  const isActive = computed(() => isHover.value || focused.value);
</script>

<template>
  <NuxtLink
    ref="badge"
    class="badge"
    :class="{
      'p-3': size !== 'sm',
      'p-2 badge-sm': size === 'sm',
    }"
    :to="`/label/${label.id}`"
  >
    <label class="swap swap-rotate" :class="isActive ? 'swap-active' : ''">
      <Icon name="heroicons-arrow-right" class="mr-2 swap-on"></Icon>
      <Icon name="heroicons-tag" class="mr-2 swap-off"></Icon>
    </label>
    {{ label.name }}
  </NuxtLink>
</template>
