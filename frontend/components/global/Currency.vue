<template>
  {{ value }}
</template>

<script setup lang="ts">
  type Props = {
    amount: string | number;
  };

  const props = defineProps<Props>();

  const fmt = await useFormatCurrency();

  const value = computed(() => {
    if (!props.amount || props.amount === "0") {
      return fmt(0);
    }

    return fmt(props.amount);
  });
</script>
