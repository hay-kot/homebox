<script setup lang="ts">
  import ActionsDivider from "../../components/Base/ActionsDivider.vue";

  definePageMeta({
    layout: "home",
  });

  const route = useRoute();
  const api = useUserApi();
  const toast = useNotifier();

  const preferences = useViewPreferences();

  const labelId = computed<string>(() => route.params.id as string);

  const { data: label } = useAsyncData(labelId.value, async () => {
    const { data, error } = await api.labels.get(labelId.value);
    if (error) {
      toast.error("Failed to load label");
      navigateTo("/home");
      return;
    }
    return data;
  });

  function maybeTimeAgo(date?: string): string {
    if (!date) {
      return "??";
    }

    const time = new Date(date);

    return `${useTimeAgo(time).value} (${useDateFormat(time, "MM-DD-YYYY").value})`;
  }

  const details = computed(() => {
    const dt = {
      Name: label.value?.name || "",
      Description: label.value?.description || "",
    };

    if (preferences.value.showDetails) {
      dt["Created At"] = maybeTimeAgo(label.value?.createdAt);
      dt["Updated At"] = maybeTimeAgo(label.value?.updatedAt);
      dt["Database ID"] = label.value?.id || "";
      dt["Group Id"] = label.value?.groupId || "";
    }

    return dt;
  });

  const { reveal } = useConfirm();

  async function confirmDelete() {
    const { isCanceled } = await reveal("Are you sure you want to delete this label? This action cannot be undone.");

    if (isCanceled) {
      return;
    }

    const { error } = await api.labels.delete(labelId.value);

    if (error) {
      toast.error("Failed to delete label");
      return;
    }
    toast.success("Label deleted");
    navigateTo("/home");
  }

  const updateModal = ref(false);
  const updating = ref(false);
  const updateData = reactive({
    name: "",
    description: "",
    color: "",
  });

  function openUpdate() {
    updateData.name = label.value?.name || "";
    updateData.description = label.value?.description || "";
    updateModal.value = true;
  }

  async function update() {
    updating.value = true;
    const { error, data } = await api.labels.update(labelId.value, updateData);

    if (error) {
      toast.error("Failed to update label");
      return;
    }

    toast.success("Label updated");
    label.value = data;
    updateModal.value = false;
    updating.value = false;
  }
</script>

<template>
  <BaseContainer>
    <BaseModal v-model="updateModal">
      <template #title> Update Label </template>
      <form v-if="label" @submit.prevent="update">
        <FormTextField v-model="updateData.name" :autofocus="true" label="Label Name" />
        <FormTextField v-model="updateData.description" label="Label Description" />
        <div class="modal-action">
          <BaseButton type="submit" :loading="updating"> Update </BaseButton>
        </div>
      </form>
    </BaseModal>
    <section>
      <BaseSectionHeader class="mb-5" dark>
        {{ label ? label.name : "" }}
      </BaseSectionHeader>
      <BaseDetails class="mb-2" :details="details">
        <template #title> Label Details </template>
      </BaseDetails>
      <div class="form-control ml-auto mr-2 max-w-[130px]">
        <label class="label cursor-pointer">
          <input v-model="preferences.showDetails" type="checkbox" class="toggle" />
          <span class="label-text"> Detailed View </span>
        </label>
      </div>
      <ActionsDivider @delete="confirmDelete" @edit="openUpdate" />
    </section>

    <section v-if="label">
      <BaseSectionHeader class="mb-5"> Items </BaseSectionHeader>
      <div class="grid gap-2 grid-cols-2">
        <ItemCard v-for="item in label.items" :key="item.id" :item="item" />
      </div>
    </section>
  </BaseContainer>
</template>
