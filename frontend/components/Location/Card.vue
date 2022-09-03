<template>
  <NuxtLink
    ref="card"
    :to="`/location/${location.id}`"
    class="card bg-primary text-primary-content transition duration-300"
  >
    <div class="card-body p-4">
      <h2 class="flex items-center gap-2">
        <label class="swap swap-rotate" :class="isActive ? 'swap-active' : ''">
          <Icon name="heroicons-arrow-right" class="swap-on" />
          <Icon name="heroicons-map-pin" class="swap-off" />
        </label>
        {{ location.name }}
        <span class="badge badge-secondary badge-lg ml-auto text-secondary-content">
          {{ location.itemCount }}</span
        >
      </h2>
    </div>
  </NuxtLink>
</template>

<script lang="ts" setup>
  import { Location } from '~~/lib/api/classes/locations';

  defineProps({
    location: {
      type: Object as () => Location,
      required: true,
    },
  });

  const card = ref(null);
  const isHover = useElementHover(card);
  const { focused } = useFocus(card);

  const isActive = computed(() => isHover.value || focused.value);
</script>
