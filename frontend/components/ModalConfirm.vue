<template>
  <BaseModal v-model="isRevealed" readonly @cancel="cancel(false)">
    <template #title> Confirm </template>
    <div>
      <p>{{ text }}</p>
    </div>
    <div class="modal-action">
      <BaseButton type="submit" @click="confirm(true)"> Confirm </BaseButton>
    </div>
  </BaseModal>
</template>

<script setup lang="ts">
  const { text, isRevealed, confirm, cancel } = useConfirm();
</script>
