<script setup lang="ts">
  definePageMeta({
    layout: 'home',
  });
  useHead({
    title: 'Homebox | Home',
  });

  const api = useUserApi();

  const { data: locations } = useAsyncData('locations', async () => {
    const { data } = await api.locations.getAll();
    return data.items;
  });

  const { data: labels } = useAsyncData('labels', async () => {
    const { data } = await api.labels.getAll();
    return data.items;
  });
</script>

<template>
  <BaseContainer class="space-y-16">
    <section>
      <BaseSectionHeader class="mb-5"> Storage Locations </BaseSectionHeader>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        <NuxtLink
          :to="`/location/${l.id}`"
          class="card bg-primary text-primary-content hover:-translate-y-1 focus:-translate-y-1 transition duration-300"
          v-for="l in locations"
        >
          <div class="card-body p-4">
            <h2 class="flex items-center gap-2">
              <Icon name="heroicons-map-pin" class="h-5 w-5 text-white" height="25" />
              {{ l.name }}
              <!-- <span class="badge badge-accent badge-lg ml-auto text-accent-content text-lg">0</span> -->
            </h2>
          </div>
        </NuxtLink>
      </div>
    </section>

    <section>
      <BaseSectionHeader class="mb-5"> Labels </BaseSectionHeader>
      <div class="flex gap-2">
        <LabelChip v-for="label in labels" :label="label" />
      </div>
    </section>
  </BaseContainer>
</template>
