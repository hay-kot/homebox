<template>
  <button class="" @click="copyText">
    <label
      class="swap swap-rotate"
      :class="{
        'swap-active': copied,
      }"
    >
      <Icon
        class="swap-off"
        name="mdi-content-copy"
        :style="{
          height: `${iconSize}px`,
          width: `${iconSize}px`,
        }"
      />
      <Icon
        class="swap-on"
        name="mdi-clipboard"
        :style="{
          height: `${iconSize}px`,
          width: `${iconSize}px`,
        }"
      />
    </label>
  </button>
</template>

<script setup lang="ts">
  const props = defineProps({
    text: {
      type: String as () => string,
      default: "",
    },
    iconSize: {
      type: Number as () => number,
      default: 20,
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
