<script setup lang="ts">
  export type sizes = 'sm' | 'md' | 'lg';

  import { Label } from '~~/lib/api/classes/labels';
  defineProps({
    label: {
      type: Object as () => Label,
      required: true,
    },
    size: {
      type: String as () => sizes,
      default: 'md',
    },
  });

  const badge = ref(null);
  const isHover = useElementHover(badge);
  const { focused } = useFocus(badge);

  const isActive = computed(() => isHover.value || focused.value);
</script>

<template>
  <NuxtLink
    class="badge"
    :class="{
      'p-3': size !== 'sm',
      'p-2 badge-sm': size === 'sm',
    }"
    ref="badge"
    :to="`/label/${label.id}`"
  >
    <label class="swap swap-rotate" :class="isActive ? 'swap-active' : ''">
      <Icon name="heroicons-arrow-right" class="mr-2 swap-on"></Icon>
      <Icon name="heroicons-tag" class="mr-2 swap-off"></Icon>
    </label>
    {{ label.name }}
  </NuxtLink>
</template>
