<template>
  <BaseModal v-model="modal">
    <template #title> Create Label </template>
    <form @submit.prevent="create">
      <FormTextField
        :trigger-focus="focused"
        ref="locationNameRef"
        :autofocus="true"
        label="Label Name"
        v-model="form.name"
      />
      <FormTextField label="Label Description" v-model="form.description" />
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
    const { data, error } = await api.labels.create(form);
    if (error) {
      toast.error("Couldn't create label");
      return;
    }

    toast.success('Label created');
    reset();
  }
</script>
