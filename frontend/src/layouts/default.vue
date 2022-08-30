<script setup lang="ts"></script>
<template>
	<header>
		<app-header />
	</header>

	<main
		class="
			p-8
			dark:bg-gray-800 dark:text-white
			bg-white
			text-gray-800
			min-h-screen
		"
	>
		<router-view />
	</main>
</template>
