<template>
	<div class="form-control w-full">
		<label class="label">
			<span class="label-text">{{ label }}</span>
		</label>
		<input
			:type="type"
			v-model="value"
			class="input input-bordered w-full"
		/>
	</div>
</template>

<script lang="ts" setup>
	const props = defineProps({
		modelValue: {
			type: String,
			required: true,
		},
		label: {
			type: String,
			required: true,
		},
		type: {
			type: String,
			default: 'text',
		},
	});

	const value = useVModel(props, 'modelValue');
</script>
