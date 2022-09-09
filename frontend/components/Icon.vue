<script setup lang="ts">
  import type { Ref } from "vue";
  import type { IconifyIcon } from "@iconify/vue";
  import { Icon as Iconify, loadIcon } from "@iconify/vue";

  const nuxtApp = useNuxtApp();
  const props = defineProps({
    name: {
      type: String,
      required: true,
    },
  });

  const icon: Ref<IconifyIcon | null> = ref(null);
  const component = computed(() => nuxtApp.vueApp.component(props.name));

  icon.value = await loadIcon(props.name).catch(() => null);

  watch(
    () => props.name,
    async () => {
      icon.value = await loadIcon(props.name).catch(() => null);
    }
  );
</script>

<template>
  <Iconify v-if="icon" :icon="icon" class="inline-block" />
  <Component :is="component" v-else-if="component" />
  <span v-else>{{ name }}</span>
</template>
