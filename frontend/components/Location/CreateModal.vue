<template>
  <BaseModal v-model="modal">
    <template #title> Create Location </template>
    <form @submit.prevent="create">
      <FormTextField
        :trigger-focus="focused"
        ref="locationNameRef"
        :autofocus="true"
        label="Location Name"
        v-model="form.name"
      />
      <FormTextField label="Location Description" v-model="form.description" />
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
  });

  whenever(
    () => modal.value,
    () => {
      focused.value = true;
    }
  );

  function reset() {
    form.name = '';
    form.description = '';
    focused.value = false;
    modal.value = false;
    loading.value = false;
  }

  const api = useUserApi();
  const toast = useNotifier();

  async function create() {
    loading.value = true;

    const { data, error } = await api.locations.create(form);

    if (error) {
      toast.error("Couldn't create location");
    }

    if (data) {
      toast.success('Location created');
      navigateTo(`/location/${data.id}`);
    }

    reset();
  }
</script>
