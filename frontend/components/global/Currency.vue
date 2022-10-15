<template>
  {{ value }}
</template>

<script setup lang="ts">
  const props = defineProps({
    amount: {
      type: String,
      required: true,
    },
  });

  const fmt = await useFormatCurrency();

  const value = computed(() => {
    if (!props.amount || props.amount === "0") {
      return "";
    }

    return fmt(props.amount);
  });
</script>
