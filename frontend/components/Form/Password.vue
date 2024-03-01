<template>
  <div class="relative">
    <FormTextField v-model="value" placeholder="Password" :label="label" :type="inputType"> </FormTextField>
    <button
      type="button"
      class="inline-flex p-1 ml-1 justify-center mt-auto mb-3 tooltip absolute top-11 right-3"
      data-tip="Toggle Password Show"
      @click="toggle()"
    >
      <MdiEye name="mdi-eye" class="h-5 w-5" />
    </button>
  </div>
</template>

<script setup lang="ts">
  import MdiEye from "~icons/mdi/eye";

  type Props = {
    modelValue: string;
    placeholder?: string;
    label: string;
  };

  const props = withDefaults(defineProps<Props>(), {
    placeholder: "Password",
    label: "Password",
  });

  const [hide, toggle] = useToggle(true);

  const inputType = computed(() => {
    return hide.value ? "password" : "text";
  });

  const value = useVModel(props, "modelValue");
</script>
