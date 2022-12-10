<template>
  {{ value }}
</template>

<script setup lang="ts">
  type DateTimeFormat = "relative" | "long" | "short" | "human";

  function ordinalIndicator(num: number) {
    if (num > 3 && num < 21) return "th";
    switch (num % 10) {
      case 1:
        return "st";
      case 2:
        return "nd";
      case 3:
        return "rd";
      default:
        return "th";
    }
  }

  const months = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];

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
      case "relative":
        return useTimeAgo(dt).value + useDateFormat(dt, " (MM-DD-YYYY)").value;
      case "long":
        return useDateFormat(dt, "MM-DD-YYYY (dddd)").value;
      case "short":
        return useDateFormat(dt, "MM-DD-YYYY").value;
      case "human":
        // January 1st, 2021
        return `${months[dt.getMonth()]} ${dt.getDate()}${ordinalIndicator(dt.getDate())}, ${dt.getFullYear()}`;
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
