<template>
  <BaseModal v-model="modal">
    <template #title> Create Item </template>
    <form @submit.prevent="create(true)">
      <LocationSelector v-model="form.location" />
      <FormTextField
        ref="locationNameRef"
        v-model="form.name"
        :trigger-focus="focused"
        :autofocus="true"
        label="Item Name"
      />
      <FormTextArea v-model="form.description" label="Item Description" />
      <FormMultiselect v-model="form.labels" label="Labels" :items="labels ?? []" />
      <div class="modal-action">
        <div class="flex justify-center">
          <BaseButton ref="submitBtn" type="submit" class="rounded-r-none" :loading="loading">
            <template #icon>
              <Icon name="mdi-package-variant" class="swap-off h-5 w-5" />
              <Icon name="mdi-package-variant-closed" class="swap-on h-5 w-5" />
            </template>
            Create
          </BaseButton>
          <div class="dropdown dropdown-top">
            <label tabindex="0" class="btn rounded-l-none rounded-r-xl">
              <Icon class="h-5 w-5" name="mdi-chevron-down" />
            </label>
            <ul tabindex="0" class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-64">
              <li><button @click.prevent="create(false)">Create and Add Another</button></li>
            </ul>
          </div>
        </div>
      </div>
    </form>
  </BaseModal>
</template>

<script setup lang="ts">
  import { ItemCreate, LabelOut, LocationOut } from "~~/lib/api/types/data-contracts";
  import { useLabelStore } from "~~/stores/labels";
  import { useLocationStore } from "~~/stores/locations";

  const props = defineProps({
    modelValue: {
      type: Boolean,
      required: true,
    },
  });

  const route = useRoute();

  const labelId = computed(() => {
    if (route.fullPath.includes("/label/")) {
      return route.params.id;
    }
    return null;
  });

  const locationId = computed(() => {
    if (route.fullPath.includes("/location/")) {
      return route.params.id;
    }
    return null;
  });

  const api = useUserApi();
  const toast = useNotifier();

  const locationsStore = useLocationStore();
  const locations = computed(() => locationsStore.allLocations);

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
    labels: [] as LabelOut[],
  });

  whenever(
    () => modal.value,
    () => {
      focused.value = true;

      if (locationId.value) {
        const found = locations.value.find(l => l.id === locationId.value);
        if (found) {
          form.location = found;
        }
      }

      if (labelId.value) {
        form.labels = labels.value.filter(l => l.id === labelId.value);
      }
    }
  );

  async function create(close = false) {
    if (!form.location) {
      return;
    }

    const out: ItemCreate = {
      parentId: null,
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

    // Reset
    form.name = "";
    form.description = "";
    form.color = "";
    focused.value = false;
    loading.value = false;

    if (close) {
      modal.value = false;
      navigateTo(`/item/${data.id}`);
    }
  }
</script>
