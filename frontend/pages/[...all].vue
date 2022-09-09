<script setup lang="ts">
  useHead({
    title: "404. Not Found",
  });
  definePageMeta({
    layout: "404",
  });
</script>

<template>
  <h1 class="text-blue-500 font-extrabold flex flex-col text-center">
    <span class="text-7xl">404.</span>
    <span class="text-5xl mt-5"> Page Not Found </span>
  </h1>
</template>
