<template>
  {{ value }}
</template>

<script setup lang="ts">
  enum DateTimeFormat {
    RELATIVE = "relative",
    LONG = "long",
    SHORT = "short",
  }

  const value = computed(() => {
    if (!props.date) {
      return "";
    }

    const dt = typeof props.date === "string" ? new Date(props.date) : props.date;
    if (!dt) {
      return "";
    }

    if (!validDate(dt)) {
      return "";
    }

    switch (props.format) {
      case DateTimeFormat.RELATIVE:
        return useTimeAgo(dt).value + useDateFormat(dt, " (MM-DD-YYYY)").value;
      case DateTimeFormat.LONG:
        return useDateFormat(dt, "YYYY-MM-DD (dddd)").value;
      case DateTimeFormat.SHORT:
        return useDateFormat(dt, "YYYY-MM-DD").value;
      default:
        return "";
    }
  });

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
</script>
