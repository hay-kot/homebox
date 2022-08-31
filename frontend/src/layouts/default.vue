<script setup lang="ts">
	import Toast from '@/components/App/Toast.vue';
</script>
<template>
	<Toast />
	<header>
		<app-header />
	</header>

	<main
		class="p-8 dark:bg-gray-800 dark:text-white bg-white text-gray-800 min-h-screen"
	>
		<router-view />
	</main>
</template>
