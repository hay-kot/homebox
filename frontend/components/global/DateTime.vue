<template>
  {{ value }}
</template>

<script setup lang="ts">
  type DateTimeFormat = "relative" | "long" | "short" | "human";

  const props = defineProps({
    date: {
      type: [Date, String],
      required: true,
    },
    format: {
      type: String as () => DateTimeFormat,
      default: "relative",
    },
  });

  const value = computed(() => {
    return fmtDate(props.date, props.format);
  });
</script>
