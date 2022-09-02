<script setup lang="ts">
  import { Label } from '~~/lib/api/classes/labels';
  defineProps({
    label: {
      type: Object as () => Label,
      required: true,
    },
  });

  const badge = ref(null);
  const isHover = useElementHover(badge);
  const { focused } = useFocus(badge);

  const isActive = computed(() => isHover.value || focused.value);
</script>

<template>
  <NuxtLink ref="badge" :to="`/label/${label.id}`">
    <span class="badge badge-lg p-4">
      <label class="swap swap-rotate" :class="isActive ? 'swap-active' : ''">
        <Icon name="heroicons-arrow-right" class="mr-2 swap-on"></Icon>
        <Icon name="heroicons-tag" class="mr-2 swap-off"></Icon>
      </label>
      {{ label.name }}
    </span>
  </NuxtLink>
</template>
