<template>
  <BaseModal v-model="modal">
    <template #title> Create Item </template>
    <form @submit.prevent="create">
      <FormSelect v-model="form.location" label="Location" :items="locations ?? []" />
      <FormTextField
        ref="locationNameRef"
        v-model="form.name"
        :trigger-focus="focused"
        :autofocus="true"
        label="Item Name"
      />
      <FormTextField v-model="form.description" label="Item Description" />
      <FormMultiselect v-model="form.labels" label="Labels" :items="labels ?? []" />
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
  import { ItemCreate, LocationOut } from "~~/lib/api/types/data-contracts";
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";

  const props = defineProps({
    modelValue: {
      type: Boolean,
      required: true,
    },
  });

  const api = useUserApi();
  const toast = useNotifier();

  const locationsStore = useLocationStore();
  const locations = computed(() => locationsStore.locations);

  const labelStore = useLabelStore();
  const labels = computed(() => labelStore.labels);

  const submitBtn = ref(null);

  const modal = useVModel(props, "modelValue");
  const loading = ref(false);
  const focused = ref(false);
  const form = reactive({
    location: locations.value && locations.value.length > 0 ? locations.value[0] : ({} as LocationOut),
    name: "",
    description: "",
    color: "", // Future!
    labels: [],
  });

  function reset() {
    form.name = "";
    form.description = "";
    form.color = "";
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

  async function create() {
    if (!form.location) {
      return;
    }

    const out: ItemCreate = {
      parentId: undefined,
      name: form.name,
      description: form.description,
      locationId: form.location.id as string,
      labelIds: form.labels.map(l => l.id) as string[],
    };

    const { error, data } = await api.items.create(out);
    if (error) {
      toast.error("Couldn't create item");
      return;
    }

    toast.success("Item created");
    reset();
    navigateTo(`/item/${data.id}`);
  }
</script>
