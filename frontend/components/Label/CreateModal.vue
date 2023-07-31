<template>
  <BaseModal v-model="modal">
    <template #title> Create Label </template>
    <div @keyup="keySubmit">
      <FormTextField
        ref="locationNameRef"
        v-model="form.name"
        :trigger-focus="focused"
        :autofocus="true"
        label="Label Name"
      />
      <FormTextArea v-model="form.description" label="Label Description" />
      <div class="modal-action">
        <div class="flex justify-center">
          <BaseButton class="rounded-r-none" type="submit" :loading="loading" @click.prevent="create(true)">
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
      use <kbd class="kbd kbd-xs">Shift</kbd> + <kdb class="kbd kbd-xs"> Enter </kdb> to create and add another
    </p>
  </BaseModal>
</template>

<script setup lang="ts">
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
    color: "", // Future!
  });

  function reset() {
    form.name = "";
    form.description = "";
    form.color = "";
    focused.value = false;
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

  async function create(close: boolean) {
    const { error, data } = await api.labels.create(form);
    if (error) {
      toast.error("Couldn't create label");
      return;
    }

    toast.success("Label created");
    reset();

    if (close) {
      modal.value = false;
      navigateTo(`/label/${data.id}`);
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
