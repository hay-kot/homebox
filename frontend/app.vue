<template>
  <NuxtLayout>
    <Html lang="en" :data-theme="theme" />
    <NuxtPage />
  </NuxtLayout>
</template>

<script lang="ts" setup>
  const { theme } = useTheme();
</script>
