<template>
  {{ value }}
</template>

<script setup lang="ts">
  type DateTimeFormat = "relative" | "long" | "short" | "human";

  type Props = {
    date?: Date | string;
    format?: DateTimeFormat;
  };

  const props = withDefaults(defineProps<Props>(), {
    date: undefined,
    format: "relative",
  });

  const value = computed(() => {
    if (!props.date || !validDate(props.date)) {
      return "";
    }

    return fmtDate(props.date, props.format);
  });
</script>
