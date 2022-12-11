<template>
  <NuxtLink
    ref="card"
    :to="`/location/${location.id}`"
    class="card bg-white text-neutral rounded-md transition duration-300 shadow-md"
  >
    <div
      class="card-body"
      :class="{
        'p-4': !dense,
        'py-2 px-3': dense,
      }"
    >
      <h2 class="flex items-center justify-between gap-2">
        <label class="swap swap-rotate" :class="isActive ? 'swap-active' : ''">
          <Icon name="heroicons-arrow-right" class="swap-on h-6 w-6" />
          <Icon name="heroicons-map-pin" class="swap-off h-6 w-6" />
        </label>
        <span class="mx-auto">
          {{ location.name }}
        </span>
        <span v-if="hasCount" class="badge badge-primary h-6 w-6 badge-lg"> {{ count }}</span>
      </h2>
    </div>
  </NuxtLink>
</template>

<script lang="ts" setup>
  import { LocationOut, LocationOutCount, LocationSummary } from "~~/lib/api/types/data-contracts";

  const props = defineProps({
    location: {
      type: Object as () => LocationOutCount | LocationOut | LocationSummary,
      required: true,
    },
    dense: {
      type: Boolean,
      default: false,
    },
  });

  const hasCount = computed(() => {
    return !!(props.location as LocationOutCount).itemCount;
  });

  const count = computed(() => {
    if (hasCount.value) {
      return (props.location as LocationOutCount).itemCount;
    }
  });

  const card = ref(null);
  const isHover = useElementHover(card);
  const { focused } = useFocus(card);

  const isActive = computed(() => isHover.value || focused.value);
</script>
