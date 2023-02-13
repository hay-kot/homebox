<template>
  <NuxtLayout>
    <Html lang="en" :data-theme="theme || 'homebox'" />
    <Link rel="icon" type="image/svg" href="/favicon.svg"></Link>
    <NuxtPage />
  </NuxtLayout>
</template>

<script lang="ts" setup>
  const { theme } = useTheme();
</script>
