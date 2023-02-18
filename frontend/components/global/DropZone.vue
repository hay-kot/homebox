<template>
  <div
    ref="el"
    class="h-24 w-full border-2 border-primary border-dashed grid place-content-center"
    :class="isOverDropZone ? 'bg-primary bg-opacity-10' : ''"
  >
    <slot />
  </div>
</template>

<script setup lang="ts">
  defineProps({
    modelValue: {
      type: Boolean,
      required: false,
    },
  });

  const emit = defineEmits(["update:modelValue", "drop"]);

  const el = ref<HTMLDivElement>();
  const { isOverDropZone } = useDropZone(el, files => {
    emit("drop", files);
  });

  watch(isOverDropZone, () => {
    emit("update:modelValue", isOverDropZone.value);
  });
</script>

<style scoped></style>
