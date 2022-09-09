<script setup lang="ts">
  import ActionsDivider from "../../components/Base/ActionsDivider.vue";

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

  function maybeTimeAgo(date?: string): string {
    if (!date) {
      return "??";
    }

    const time = new Date(date);

    return `${useTimeAgo(time).value} (${useDateFormat(time, "MM-DD-YYYY").value})`;
  }

  const details = computed(() => {
    const dt = {
      Name: location.value?.name || "",
      Description: location.value?.description || "",
    };

    if (preferences.value.showDetails) {
      dt["Created At"] = maybeTimeAgo(location.value?.createdAt);
      dt["Updated At"] = maybeTimeAgo(location.value?.updatedAt);
      dt["Database ID"] = location.value?.id || "";
      dt["Group Id"] = location.value?.groupId || "";
    }

    return dt;
  });

  const { reveal } = useConfirm();

  async function confirmDelete() {
    const { isCanceled } = await reveal("Are you sure you want to delete this location? This action cannot be undone.");

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
        <FormTextField v-model="updateData.description" label="Location Description" />
        <div class="modal-action">
          <BaseButton type="submit" :loading="updating"> Update </BaseButton>
        </div>
      </form>
    </BaseModal>
    <section>
      <BaseSectionHeader class="mb-5" dark>
        {{ location ? location.name : "" }}
      </BaseSectionHeader>
      <BaseDetails class="mb-2" :details="details">
        <template #title> Location Details </template>
      </BaseDetails>
      <div class="form-control ml-auto mr-2 max-w-[130px]">
        <label class="label cursor-pointer">
          <input v-model="preferences.showDetails" type="checkbox" class="toggle" />
          <span class="label-text"> Detailed View </span>
        </label>
      </div>
      <ActionsDivider @delete="confirmDelete" @edit="openUpdate" />
    </section>

    <section v-if="location">
      <BaseSectionHeader class="mb-5"> Items </BaseSectionHeader>
      <div class="grid gap-2 grid-cols-2">
        <ItemCard v-for="item in location.items" :key="item.id" :item="item" />
      </div>
    </section>
  </BaseContainer>
</template>
