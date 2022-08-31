<template>
	<div class="force-above fixed top-2 right-2 w-[300px]">
		<TransitionGroup name="notify" tag="div">
			<div
				v-for="(notify, index) in notifications.slice(0, 4)"
				:key="notify.id"
				class="my-2 w-[300px] rounded-md p-3 text-sm text-white opacity-75"
				:class="{
					'bg-primary': notify.type === 'info',
					'bg-red-600': notify.type === 'error',
					'bg-green-600': notify.type === 'success',
				}"
				@click="dropNotification(index)"
			>
				<div class="flex gap-1">
					<template v-if="notify.type == 'info'">
						<Icon
							icon="mdi-information-outline"
							class="h-5 w-5"
							height="25"
						/>
					</template>
					<template v-if="notify.type == 'success'">
						<Icon
							icon="mdi-check-circle-outline"
							class="h-5 w-5"
							height="25"
						/>
					</template>

					<template v-if="notify.type == 'error'">
						<Icon
							icon="mdi-alert-circle-outline"
							class="h-5 w-5"
							height="25"
						/>
					</template>
					{{ notify.message }}
				</div>
			</div>
		</TransitionGroup>
	</div>
</template>

<script setup lang="ts">
	import { Icon } from '@iconify/vue';
	import { useNotifications } from '@/composables/use-notifier';

	const { notifications, dropNotification } = useNotifications();
</script>

<style scoped>
	.force-above {
		z-index: 9999;
	}

	.notify-move,
	.notify-enter-active,
	.notify-leave-active {
		transition: all 0.5s ease;
	}
	.notify-enter-from,
	.notify-leave-to {
		opacity: 0;
		transform: translateY(-30px);
	}
	.notify-leave-active {
		position: absolute;
		transform: translateY(30px);
	}
</style>
