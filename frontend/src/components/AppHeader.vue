<script setup lang="ts">
	defineEmits(['toggleSidebar']);

	const { availableLocales } = useI18n();

	const preferedDark = usePreferredDark();
	const isDark = useStorage('isDark', preferedDark.value);
	const body = ref<HTMLBodyElement | null>(null);

	const toggleDarkMode = () => {
		if (body.value) {
			if (isDark.value) {
				body.value.classList.remove('dark');
			} else {
				body.value.classList.add('dark');
			}
		}
		isDark.value = !isDark.value;
	};

	onMounted(async () => {
		await nextTick();

		body.value = document.querySelector('body') as HTMLBodyElement;
		if (body.value) {
			if (isDark.value) body.value.classList.add('dark');
		}
	});
</script>

<template>
	<header>
		<nav
			class="
				w-full
				bg-white
				text-gray-800
				dark:bg-gray-800 dark:text-white
				py-4
				px-8
				shadow-md
				dark:shadow-md
				flex
				items-center
				border-b border-gray-400/50
			"
		>
			<router-link :to="{ name: 'home' }">
				<div class="font-bold lg:text-xl md:text-lg text-md">Vitailse</div>
			</router-link>
			<div class="ml-auto flex items-center h-full">
				<select
					id="language"
					v-model="$i18n.locale"
					class="py-1 focus:outline-none rounded dark:text-gray-800"
				>
					<option
						v-for="locale in availableLocales"
						:key="locale"
						:value="locale"
					>
						{{ locale }}
					</option>
				</select>
				<button
					class="mx-5 cursor-pointer focus:outline-none"
					@click="toggleDarkMode"
				>
					<icon:bx:bx-moon class="w-6 h-6" v-if="!isDark" />
					<icon:bx:bxs-moon class="w-6 h-6" v-else />
				</button>
				<a href="https://github.com/zynth17/vitailse">
					<icon-akar-icons:github-fill />
				</a>
			</div>
		</nav>
	</header>
</template>

<style scoped></style>
