<script setup lang="ts">
	import { useUserApi } from '@/composables/use-api';
	useHead({
		title: 'Homebox | Home',
	});

	const links = [
		{
			name: 'Home',
			href: '/home',
		},
		{
			name: 'Logout',
			href: '/logout',
			last: true,
		},
	];
	const api = useUserApi();

	const user = ref({});

	onMounted(async () => {
		const { data } = await api.self();

		if (data) {
			user.value = data.item;
		}
	});
</script>

<template>
	<section class="max-w-7xl mx-auto">
		<header class="sm:px-6 py-2 lg:p-14 sm:py-6">
			<h2
				class="mt-1 text-4xl font-bold tracking-tight text-gray-200 sm:text-5xl lg:text-6xl"
			>
				Homebox
			</h2>
			<div class="ml-1 text-lg text-gray-400 space-x-2 italic">
				<template v-for="link in links">
					<router-link
						class="hover:text-base-content transition-color duration-200"
						:to="link.href"
					>
						{{ link.name }}
					</router-link>
					<span v-if="!link.last"> / </span>
				</template>
			</div>
		</header>
	</section>
	<section class="max-w-7xl mx-auto sm:px-6 lg:px-14">
		{{ user }}
	</section>
</template>

<route lang="yaml">
name: home
</route>
