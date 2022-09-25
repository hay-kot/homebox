<script setup lang="ts">
  import { Detail, DateDetail } from "~~/components/global/DetailsSection/types";

  definePageMeta({
    layout: "home",
  });

  const route = useRoute();
  const api = useUserApi();
  const toast = useNotifier();

  const preferences = useViewPreferences();

  const locationId = computed<string>(() => route.params.id as string);

  const { data: location } = useAsyncData(locationId.value, async () => {
    const { data, error } = await api.locations.get(locationId.value);
    if (error) {
      toast.error("Failed to load location");
      navigateTo("/home");
      return;
    }
    return data;
  });

  const details = computed<(Detail | DateDetail)[]>(() => {
    const details = [
      {
        name: "Name",
        text: location.value?.name,
      },
      {
        name: "Description",
        text: location.value?.description,
      },
    ];

    if (preferences.value.showDetails) {
      return [
        ...details,
        {
          name: "Created",
          text: location.value?.createdAt,
          type: "date",
        },
        {
          name: "Updated",
          text: location.value?.updatedAt,
          type: "date",
        },
        {
          name: "Database ID",
          text: location.value?.id,
        },
      ];
    }

    return details;
  });

  const confirm = useConfirm();

  async function confirmDelete() {
    const { isCanceled } = await confirm.open(
      "Are you sure you want to delete this location? This action cannot be undone."
    );
    if (isCanceled) {
      return;
    }

    const { error } = await api.locations.delete(locationId.value);
    if (error) {
      toast.error("Failed to delete location");
      return;
    }

    toast.success("Location deleted");
    navigateTo("/home");
  }

  const updateModal = ref(false);
  const updating = ref(false);
  const updateData = reactive({
    name: "",
    description: "",
  });

  function openUpdate() {
    updateData.name = location.value?.name || "";
    updateData.description = location.value?.description || "";
    updateModal.value = true;
  }

  async function update() {
    updating.value = true;
    const { error, data } = await api.locations.update(locationId.value, updateData);

    if (error) {
      toast.error("Failed to update location");
      return;
    }

    toast.success("Location updated");
    location.value = data;
    updateModal.value = false;
    updating.value = false;
  }
</script>

<template>
  <BaseContainer>
    <BaseModal v-model="updateModal">
      <template #title> Update Location </template>
      <form v-if="location" @submit.prevent="update">
        <FormTextField v-model="updateData.name" :autofocus="true" label="Location Name" />
        <FormTextArea v-model="updateData.description" label="Location Description" />
        <div class="modal-action">
          <BaseButton type="submit" :loading="updating"> Update </BaseButton>
        </div>
      </form>
    </BaseModal>

    <BaseCard class="mb-16">
      <template #title>
        <BaseSectionHeader>
          <Icon name="mdi-map-marker" class="mr-2 text-gray-600" />
          <span class="text-gray-600">
            {{ location ? location.name : "" }}
          </span>
        </BaseSectionHeader>
      </template>

      <template #title-actions>
        <div class="flex mt-2 gap-2">
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
        </div>
      </template>

      <DetailsSection :details="details" />
    </BaseCard>

    <section v-if="location">
      <BaseSectionHeader class="mb-5"> Items </BaseSectionHeader>
      <div class="grid gap-2 grid-cols-1 sm:grid-cols-2">
        <ItemCard v-for="item in location.items" :key="item.id" :item="item" />
      </div>
    </section>
  </BaseContainer>
</template>
