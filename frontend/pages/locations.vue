<script setup lang="ts">
  import { useTreeState } from "~~/components/Location/Tree/tree-state";

  definePageMeta({
    middleware: ["auth"],
  });

  useHead({
    title: "Homebox | Items",
  });

  const api = useUserApi();

  const { data: tree } = useAsyncData(async () => {
    const { data, error } = await api.locations.getTree();

    if (error) {
      return [];
    }

    return data.items;
  });

  const locationTreeId = "locationTree";

  const treeState = useTreeState(locationTreeId);

  const route = useRouter();

  onMounted(() => {
    // set tree state from query params
    const query = route.currentRoute.value.query;

    if (query && query[locationTreeId]) {
      console.log("setting tree state from query params");
      const data = JSON.parse(query[locationTreeId] as string);

      for (const key in data) {
        treeState.value[key] = data[key];
      }
    }
  });

  watch(
    treeState,
    () => {
      // Push the current state to the URL
      route.replace({ query: { [locationTreeId]: JSON.stringify(treeState.value) } });
    },
    { deep: true }
  );
</script>

<template>
  <BaseContainer class="mb-16">
    <BaseSectionHeader> Locations </BaseSectionHeader>

    <LocationTreeRoot v-if="tree" :locs="tree" :tree-id="locationTreeId" />
  </BaseContainer>
</template>
