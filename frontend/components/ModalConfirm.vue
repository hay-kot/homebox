<template>
  <BaseModal v-model="isRevealed" readonly @cancel="cancel(false)">
    <template #title> Confirm {{ title }}</template>
    <div class="mt-3">
      <p>{{ text }}</p>
    </div>
    <div class="modal-action">
      <BaseButton type="submit" @click="confirm(true)"> Confirm </BaseButton>
    </div>
  </BaseModal>
</template>

<script setup lang="ts">
  const { title, text, isRevealed, confirm, cancel } = useConfirm();
</script>
