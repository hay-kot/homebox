<script setup lang="ts">
  import { LocationSummary, LocationUpdate } from "~~/lib/api/types/data-contracts";
  import { useLocationStore } from "~~/stores/locations";

  definePageMeta({
    middleware: ["auth"],
  });

  const route = useRoute();
  const api = useUserApi();
  const toast = useNotifier();

  const locationId = computed<string>(() => route.params.id as string);

  const { data: location } = useAsyncData(locationId.value, async () => {
    const { data, error } = await api.locations.get(locationId.value);
    if (error) {
      toast.error("Failed to load location");
      navigateTo("/home");
      return;
    }

    if (data.parent) {
      parent.value = locations.value.find(l => l.id === data.parent.id);
    }

    return data;
  });

  const confirm = useConfirm();

  async function confirmDelete() {
    const { isCanceled } = await confirm.open(
      "Are you sure you want to delete this location and all of its items? This action cannot be undone."
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
  const updateData = reactive<LocationUpdate>({
    id: locationId.value,
    name: "",
    description: "",
    parentId: null,
  });

  function openUpdate() {
    updateData.name = location.value?.name || "";
    updateData.description = location.value?.description || "";
    updateModal.value = true;
  }

  async function update() {
    updating.value = true;
    updateData.parentId = parent.value?.id || null;
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

  const locationStore = useLocationStore();
  const locations = computed(() => locationStore.allLocations);

  const parent = ref<LocationSummary | any>({});
</script>

<template>
  <div>
    <!-- Update Dialog -->
    <BaseModal v-model="updateModal">
      <template #title> Update Location </template>
      <form v-if="location" @submit.prevent="update">
        <FormTextField v-model="updateData.name" :autofocus="true" label="Location Name" />
        <FormTextArea v-model="updateData.description" label="Location Description" />
        <LocationSelector v-model="parent" />
        <div class="modal-action">
          <BaseButton type="submit" :loading="updating"> Update </BaseButton>
        </div>
      </form>
    </BaseModal>

    <BaseContainer v-if="location" class="space-y-6 mb-16">
      <section>
        <BaseSectionHeader v-if="location">
          <Icon name="mdi-package-variant" class="mr-2 -mt-1 text-base-content" />
          <span class="text-base-content">
            {{ location ? location.name : "" }}
          </span>

          <div v-if="location?.parent" class="text-sm breadcrumbs pb-0">
            <ul class="text-base-content/70">
              <li>
                <NuxtLink :to="`/location/${location.parent.id}`"> {{ location.parent.name }}</NuxtLink>
              </li>
              <li>{{ location.name }}</li>
            </ul>
          </div>
          <template #description>
            <Markdown class="text-lg" :source="location.description"> </Markdown>
          </template>
        </BaseSectionHeader>

        <div class="flex gap-3 flex-wrap mb-6 text-sm italic">
          <div>
            Created
            <DateTime :date="location?.createdAt" />
          </div>
          <div>
            <Icon name="mdi-circle-small" />
          </div>
          <div>
            Last Updated
            <DateTime :date="location?.updatedAt" />
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

      <template v-if="location && location.items.length > 0">
        <ItemViewSelectable :items="location.items" />
      </template>

      <section v-if="location && location.children.length > 0">
        <BaseSectionHeader class="mb-5"> Child Locations </BaseSectionHeader>
        <div class="grid gap-2 grid-cols-1 sm:grid-cols-3">
          <LocationCard v-for="item in location.children" :key="item.id" :location="item" />
        </div>
      </section>
    </BaseContainer>
  </div>
</template>
