<template>
  <div class="z-[9999]">
    <input type="checkbox" :id="modalId" class="modal-toggle" v-model="modal" />
    <div class="modal">
      <div class="modal-box relative">
        <button @click="close" :for="modalId" class="btn btn-sm btn-circle absolute right-2 top-2">✕</button>

        <h3 class="font-bold text-lg">
          <slot name="title"></slot>
        </h3>
        <slot> </slot>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  const emit = defineEmits(['cancel', 'update:modelValue']);
  const props = defineProps({
    modelValue: {
      type: Boolean,
      required: true,
    },
    /**
     * in readonly mode the modal only `emits` a "cancel" event to indicate
     * that the modal was closed via the "x" button. The parent component is
     * responsible for closing the modal.
     */
    readonly: {
      type: Boolean,
      default: false,
    },
  });

  function close() {
    if (props.readonly) {
      emit('cancel');
      return;
    }
    modal.value = false;
  }

  const modalId = useId();
  const modal = useVModel(props, 'modelValue', emit);
</script>
