<template>
  <div class="form-control">
    <label class="label">
      <span class="label-text">{{ label }}</span>
    </label>
    <textarea class="textarea textarea-bordered h-24" v-model="value" :placeholder="placeholder" />
    <label v-if="limit" class="label">
      <span class="label-text-alt"></span>
      <span class="label-text-alt"> {{ valueLen }}/{{ limit }}</span>
    </label>
  </div>
</template>

<script lang="ts" setup>
  const emit = defineEmits(['update:modelValue']);
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
    limit: {
      type: [Number, String],
      default: null,
    },
    placeholder: {
      type: String,
      default: '',
    },
  });

  const value = useVModel(props, 'modelValue', emit);
  const valueLen = computed(() => {
    return value.value ? value.value.length : 0;
  });
</script>
