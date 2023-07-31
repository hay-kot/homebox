<template>
  <BaseModal v-model="modal">
    <template #title> Create Location </template>
    <div @keyup="keySubmit">
      <FormTextField
        ref="locationNameRef"
        v-model="form.name"
        :trigger-focus="focused"
        :autofocus="true"
        label="Location Name"
      />
      <FormTextArea v-model="form.description" label="Location Description" />
      <LocationSelector v-model="form.parent" />
      <div class="modal-action">
        <div class="flex justify-center">
          <BaseButton class="rounded-r-none" type="submit" :loading="loading" @click="create(true)">
            Create
          </BaseButton>
          <div class="dropdown dropdown-top">
            <label tabindex="0" class="btn rounded-l-none rounded-r-xl">
              <Icon class="h-5 w-5" name="mdi-chevron-down" />
            </label>
            <ul tabindex="0" class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-64">
              <li>
                <button @click.prevent="create(false)">Create and Add Another</button>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
    <p class="text-sm text-center mt-4">
      use <kbd class="kbd kbd-xs">Shift</kbd> + <kbd class="kbd kbd-xs"> Enter </kbd> to create and add another
    </p>
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

  const modal = useVModel(props, "modelValue");
  const loading = ref(false);
  const focused = ref(false);
  const form = reactive({
    name: "",
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
    form.parent = null;
    focused.value = false;
    loading.value = false;
  }

  const api = useUserApi();
  const toast = useNotifier();

  async function create(close: boolean) {
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

    if (close) {
      modal.value = false;
      navigateTo(`/location/${data.id}`);
    }
  }

  async function keySubmit(e: KeyboardEvent) {
    // Shift + Enter
    if (e.shiftKey && e.key === "Enter") {
      e.preventDefault();
      await create(false);
      focused.value = true;
    } else if (e.key === "Enter") {
      e.preventDefault();
      await create(true);
    }
  }
</script>
