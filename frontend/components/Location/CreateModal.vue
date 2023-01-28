<template>
  <BaseModal v-model="modal">
    <template #title> Create Location </template>
    <form @submit.prevent="create">
      <FormTextField
        ref="locationNameRef"
        v-model="form.name"
        :trigger-focus="focused"
        :autofocus="true"
        label="Location Name"
      />
      <FormTextArea v-model="form.description" label="Location Description" />
      <FormAutocomplete
        v-model="form.parent"
        v-model:search="form.search"
        :items="locations"
        item-text="display"
        item-value="id"
        item-search="name"
        label="Parent Location"
      >
        <template #display="{ item }">
          <div>
            <div>
              {{ item.name }}
            </div>
            <div v-if="item.name != item.display" class="text-xs mt-1">{{ item.display }}</div>
          </div>
        </template>
      </FormAutocomplete>
      <div class="modal-action">
        <BaseButton type="submit" :loading="loading"> Create </BaseButton>
      </div>
    </form>
  </BaseModal>
</template>

<script setup lang="ts">
  import { LocationSummary } from "~~/lib/api/types/data-contracts";
  const props = defineProps({
    modelValue: {
      type: Boolean,
      required: true,
    },
  });

  const locations = await useFlatLocations();

  const modal = useVModel(props, "modelValue");
  const loading = ref(false);
  const focused = ref(false);
  const form = reactive({
    name: "",
    search: "",
    description: "",
    parent: null as LocationSummary | null,
  });

  whenever(
    () => modal.value,
    () => {
      focused.value = true;
    }
  );

  function reset() {
    form.name = "";
    form.description = "";
    form.search = "";
    form.parent = null;
    focused.value = false;
    modal.value = false;
    loading.value = false;
  }

  const api = useUserApi();
  const toast = useNotifier();

  async function create() {
    loading.value = true;

    const { data, error } = await api.locations.create({
      name: form.name,
      description: form.description,
      parentId: form.parent ? form.parent.id : null,
    });

    if (error) {
      toast.error("Couldn't create location");
    }

    if (data) {
      toast.success("Location created");
    }
    reset();
    navigateTo(`/location/${data.id}`);
  }
</script>
