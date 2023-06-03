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

    <BaseContainer v-if="label" class="space-y-6 mb-16">
      <section>
        <BaseSectionHeader v-if="label">
          <Icon name="mdi-package-variant" class="mr-2 -mt-1 text-base-content" />
          <span class="text-base-content">
            {{ label ? label.name : "" }}
          </span>

          <template #description>
            <Markdown class="text-lg" :source="label.description"> </Markdown>
          </template>
        </BaseSectionHeader>

        <div class="flex gap-3 flex-wrap mb-6 text-sm italic">
          <div>
            Created
            <DateTime :date="label?.createdAt" />
          </div>
          <div>
            <Icon name="mdi-circle-small" />
          </div>
          <div>
            Last Updated
            <DateTime :date="label?.updatedAt" />
          </div>
        </div>

        <div class="flex flex-wrap items-center justify-between mb-6 mt-3">
          <div class="btn-group">
            <PageQRCode class="dropdown-right" />
            <BaseButton class="ml-auto" size="sm" @click="openUpdate">
              <Icon class="mr-1" name="mdi-pencil" />
              Edit
            </BaseButton>
          </div>
          <BaseButton class="btn btn-sm" @click="confirmDelete()">
            <Icon name="mdi-delete" class="mr-2" />
            Delete
          </BaseButton>
        </div>
      </section>
    </BaseContainer>

    <section v-if="label && label.items">
      <ItemViewSelectable :items="label.items" />
    </section>
  </BaseContainer>
</template>
