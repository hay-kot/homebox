<script setup lang="ts">
  import type { AnyDetail, Details } from "~~/components/global/DetailsSection/types";

  definePageMeta({
    middleware: ["auth"],
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

  const details = computed<Details>(() => {
    const details = [
      {
        name: "Name",
        text: label.value?.name,
      } as AnyDetail,
      {
        name: "Description",
        type: "markdown",
        text: label.value?.description,
      } as AnyDetail,
    ];

    if (preferences.value.showDetails) {
      return [
        ...details,
        {
          name: "Created",
          text: label.value?.createdAt,
          type: "date",
        } as AnyDetail,
        {
          name: "Updated",
          text: label.value?.updatedAt,
          type: "date",
        } as AnyDetail,
        {
          name: "Database ID",
          text: label.value?.id,
        } as AnyDetail,
      ];
    }

    return details;
  });

  const confirm = useConfirm();

  async function confirmDelete() {
    const { isCanceled } = await confirm.open(
      "Are you sure you want to delete this label? This action cannot be undone."
    );

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
        <FormTextArea v-model="updateData.description" label="Label Description" />
        <div class="modal-action">
          <BaseButton type="submit" :loading="updating"> Update </BaseButton>
        </div>
      </form>
    </BaseModal>

    <BaseCard class="mb-16">
      <template #title>
        <BaseSectionHeader>
          <Icon name="mdi-tag" class="mr-2 -mt-1 text-base-content" />
          <span class="text-base-content">
            {{ label ? label.name : "" }}
          </span>
        </BaseSectionHeader>
      </template>

      <template #title-actions>
        <div class="flex flex-wrap mt-2 gap-2">
          <div class="form-control max-w-[160px]">
            <label class="label cursor-pointer">
              <input v-model="preferences.showDetails" type="checkbox" class="toggle toggle-primary" />
              <span class="label-text ml-2"> Detailed View </span>
            </label>
          </div>
          <BaseButton class="ml-auto" size="sm" @click="openUpdate">
            <Icon class="mr-1" name="mdi-pencil" />
            Edit
          </BaseButton>
          <BaseButton size="sm" @click="confirmDelete">
            <Icon class="mr-1" name="mdi-delete" />
            Delete
          </BaseButton>
          <PageQRCode />
        </div>
      </template>

      <DetailsSection :details="details" />
    </BaseCard>

    <section v-if="label && label.items && label.items.length > 0">
      <ItemViewSelectable :items="label.items" />
    </section>
  </BaseContainer>
</template>
