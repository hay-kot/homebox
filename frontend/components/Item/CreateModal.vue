<template>
  <BaseModal v-model="modal">
    <template #title> Create Item </template>
    <form @submit.prevent="create">
      <FormSelect label="Location" v-model="form.location" :items="locations ?? []" select-first />
      <FormTextField
        :trigger-focus="focused"
        ref="locationNameRef"
        :autofocus="true"
        label="Item Name"
        v-model="form.name"
      />
      <FormTextField label="Item Description" v-model="form.description" />
      <FormMultiselect label="Labels" v-model="form.labels" :items="labels ?? []" />
      <div class="modal-action">
        <BaseButton ref="submitBtn" type="submit" :loading="loading">
          <template #icon>
            <Icon name="mdi-package-variant" class="swap-off" />
            <Icon name="mdi-package-variant-closed" class="swap-on" />
          </template>
          Create
        </BaseButton>
      </div>
    </form>
  </BaseModal>
</template>

<script setup lang="ts">
  import { type Location } from '~~/lib/api/classes/locations';
  const props = defineProps({
    modelValue: {
      type: Boolean,
      required: true,
    },
  });

  const submitBtn = ref(null);

  const modal = useVModel(props, 'modelValue');
  const loading = ref(false);
  const focused = ref(false);
  const form = reactive({
    location: {} as Location,
    name: '',
    description: '',
    color: '', // Future!
    labels: [],
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

  const { data: locations } = useAsyncData(async () => {
    const { data } = await api.locations.getAll();
    return data.items;
  });

  const { data: labels } = useAsyncData(async () => {
    const { data } = await api.labels.getAll();
    return data.items;
  });

  async function create() {
    if (!form.location) {
      return;
    }

    const out = {
      name: form.name,
      description: form.description,
      locationId: form.location.id as string,
      labelIds: form.labels.map(l => l.id) as string[],
    };

    const { data, error } = await api.items.create(out);
    if (error) {
      toast.error("Couldn't create label");
      return;
    }

    toast.success('Item created');
    reset();
  }
</script>
