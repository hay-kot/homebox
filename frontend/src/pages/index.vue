<script setup lang="ts">
	import TextField from '@/components/Form/TextField.vue';
	useHead({
		title: 'Homebox | Organize and Tag Your Stuff',
	});

	const registerFields = [
		{
			label: "What's your name?",
			value: '',
		},
		{
			label: "What's your email?",
			value: '',
		},
		{
			label: 'Name your group',
			value: '',
		},
		{
			label: 'Set your password',
			value: '',
			type: 'password',
		},
		{
			label: 'Confirm your password',
			value: '',
			type: 'password',
		},
	];

	function registerUser() {
		// Print Values of registerFields

		for (let i = 0; i < registerFields.length; i++) {
			console.log(registerFields[i].label, registerFields[i].value);
		}
	}

	const loginFields = [
		{
			label: 'Email',
			value: '',
		},
		{
			label: 'Password',
			value: '',
			type: 'password',
		},
	];

	const registerForm = ref(false);

	function toggleLogin() {
		registerForm.value = !registerForm.value;
	}
</script>

<template>
	<header class="sm:px-6 py-2 lg:p-14 sm:py-6">
		<h2
			class="mt-1 text-4xl font-bold tracking-tight text-gray-200 sm:text-5xl lg:text-6xl"
		>
			Homebox
		</h2>
		<p class="ml-1 text-lg text-gray-400">
			Track, Organize, and Manage your Shit.
		</p>
	</header>
	<div class="grid p-6 sm:place-items-center min-h-[50vh]">
		<Transition name="slide-fade">
			<form v-if="registerForm" @submit.prevent="registerUser">
				<div
					class="card w-max-[500px] md:w-[500px] bg-base-100 shadow-xl"
				>
					<div class="card-body">
						<h2 class="card-title">Register</h2>
						<TextField
							v-for="field in registerFields"
							v-model="field.value"
							:label="field.label"
							:key="field.label"
							:type="field.type"
						/>
						<div class="card-actions justify-end">
							<button type="submit" class="btn btn-primary mt-2">
								Register
							</button>
						</div>
					</div>
				</div>
				<div class="text-center mt-2">
					<button @click="toggleLogin">Already a User? Login</button>
				</div>
			</form>
			<div v-else>
				<div
					class="card w-max-[500px] md:w-[500px] bg-base-100 shadow-xl"
				>
					<div class="card-body">
						<h2 class="card-title">Login</h2>
						<TextField
							v-for="field in loginFields"
							v-model="field.value"
							:label="field.label"
							:key="field.label"
							:type="field.type"
						/>
						<div class="card-actions justify-end mt-2">
							<button class="btn btn-primary">Login</button>
						</div>
					</div>
				</div>
				<div class="text-center mt-2">
					<button @click="toggleLogin">Not a User? Register</button>
				</div>
			</div>
		</Transition>
	</div>
	<div class="min-w-full absolute bottom-0 z-[-1]">
		<svg
			xmlns="http://www.w3.org/2000/svg"
			preserveAspectRatio="none"
			viewBox="0 0 1440 320"
		>
			<path
				fill="#0099ff9b"
				fill-opacity="1"
				d="M0,32L30,42.7C60,53,120,75,180,80C240,85,300,75,360,80C420,85,480,107,540,128C600,149,660,171,720,160C780,149,840,107,900,90.7C960,75,1020,85,1080,122.7C1140,160,1200,224,1260,234.7C1320,245,1380,203,1410,181.3L1440,160L1440,320L1410,320C1380,320,1320,320,1260,320C1200,320,1140,320,1080,320C1020,320,960,320,900,320C840,320,780,320,720,320C660,320,600,320,540,320C480,320,420,320,360,320C300,320,240,320,180,320C120,320,60,320,30,320L0,320Z"
			/>
		</svg>
		<div class="bg-[#0099ff9b] flex-col flex min-h-[32vh]">
			<div class="mt-auto mx-auto mb-8">
				<p class="text-center text-gray-200">
					&copy; 2022 Contents. All Rights Reserved. Haybytes LLC
				</p>
			</div>
		</div>
	</div>
</template>

<route lang="yaml">
name: home
</route>

<style lang="css">
	.slide-fade-enter-active {
		transition: all 0.2s ease-out;
	}

	.slide-fade-enter-from,
	.slide-fade-leave-to {
		position: absolute;
		transform: translateX(20px);
		opacity: 0;
	}
</style>
