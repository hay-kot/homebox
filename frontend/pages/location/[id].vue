<script setup lang="ts">
  import { AnyDetail, Details } from "~~/components/global/DetailsSection/types";
  import { LocationSummary, LocationUpdate } from "~~/lib/api/types/data-contracts";
  import { useLocationStore } from "~~/stores/locations";

  definePageMeta({
    middleware: ["auth"],
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

    if (data.parent) {
      parent.value = locations.value.find(l => l.id === data.parent.id);
    }

    return data;
  });

  const details = computed<Details>(() => {
    const details = [
      {
        name: "Name",
        text: location.value?.name ?? "",
      },
      {
        name: "Description",
        type: "markdown",
        text: location.value?.description ?? "",
      } as AnyDetail,
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

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const parent = ref<LocationSummary | any>({});
</script>

<template>
  <div>
    <BaseModal v-model="updateModal">
      <template #title> Update Location </template>
      <form v-if="location" @submit.prevent="update">
        <FormTextField v-model="updateData.name" :autofocus="true" label="Location Name" />
        <FormTextArea v-model="updateData.description" label="Location Description" />
        <FormAutocomplete v-model="parent" :items="locations" item-text="name" item-value="id" label="Parent" />
        <div class="modal-action">
          <BaseButton type="submit" :loading="updating"> Update </BaseButton>
        </div>
      </form>
    </BaseModal>
    <BaseContainer class="space-y-10 mb-16">
      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-map-marker" class="mr-2 -mt-1 text-base-content" />
            <span class="text-base-content">
              {{ location ? location.name : "" }}
            </span>
            <div v-if="location && location.parent" class="text-sm breadcrumbs pb-0">
              <ul class="text-base-content/70">
                <li>
                  <NuxtLink :to="`/location/${location.parent.id}`"> {{ location.parent.name }}</NuxtLink>
                </li>
                <li>{{ location.name }}</li>
              </ul>
            </div>
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
            <PageQRCode />
          </div>
        </template>

        <DetailsSection :details="details" />
      </BaseCard>

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
