<script setup lang="ts"></script>
<template>
  <div>
    <AppToast />
    <AppHeader />
    <main class="p-8 dark:bg-gray-800 dark:text-white bg-white text-gray-800 min-h-screen">
      <slot></slot>
    </main>
  </div>
</template>
