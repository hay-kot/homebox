<template>
  <BaseModal v-model="modal">
    <template #title> Create Label </template>
    <form @submit.prevent="create">
      <FormTextField
        ref="locationNameRef"
        v-model="form.name"
        :trigger-focus="focused"
        :autofocus="true"
        label="Label Name"
      />
      <FormTextField v-model="form.description" label="Label Description" />
      <div class="modal-action">
        <BaseButton type="submit" :loading="loading"> Create </BaseButton>
      </div>
    </form>
  </BaseModal>
</template>

<script setup lang="ts">
  const props = defineProps({
    modelValue: {
      type: Boolean,
      required: true,
    },
  });

  const modal = useVModel(props, 'modelValue');
  const loading = ref(false);
  const focused = ref(false);
  const form = reactive({
    name: '',
    description: '',
    color: '', // Future!
  });

  function reset() {
    form.name = '';
    form.description = '';
    form.color = '';
    focused.value = false;
    modal.value = false;
    loading.value = false;
  }

  whenever(
    () => modal.value,
    () => {
      focused.value = true;
    }
  );

  const api = useUserApi();
  const toast = useNotifier();

  async function create() {
    const { error } = await api.labels.create(form);
    if (error) {
      toast.error("Couldn't create label");
      return;
    }

    toast.success('Label created');
    reset();
  }
</script>
