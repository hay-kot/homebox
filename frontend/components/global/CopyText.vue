<template>
  <button class="btn btn-outline btn-square btn-sm" @click="copyText">
    <label
      class="swap swap-rotate"
      :class="{
        'swap-active': copied,
      }"
    >
      <Icon class="swap-off h-5 w-5" name="mdi-content-copy" />
      <Icon class="swap-on h-5 w-5" name="mdi-clipboard" />
    </label>
  </button>
</template>

<script setup lang="ts">
  const props = defineProps({
    text: {
      type: String as () => string,
      default: "",
    },
  });

  const copied = ref(false);

  const { copy } = useClipboard();

  function copyText() {
    copy(props.text);
    copied.value = true;

    setTimeout(() => {
      copied.value = false;
    }, 1000);
  }
</script>

<style scoped></style>
