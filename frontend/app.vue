<template>
  <NuxtLayout>
    <Html lang="en" :data-theme="theme || 'homebox'" />
    <NuxtPage />
  </NuxtLayout>
</template>

<script lang="ts" setup>
  const { theme } = useTheme();
</script>
