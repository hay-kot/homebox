<script setup lang="ts">
  import { useTreeState } from "./tree-state";
  import { ItemSummary, TreeItem } from "~~/lib/api/types/data-contracts";

  type Props = {
    type?: "location" | "item";
    treeId: string;
    item: TreeItem;
  };
  const props = withDefaults(defineProps<Props>(), {
    type: "location",
  });

  const link = computed(() => {
    return props.type === "location" ? `/location/${props.item.id}` : `/item/${props.item.id}`;
  });

  const hasChildren = computed(() => {
    return props.item.children.length > 0;
  });

  const state = useTreeState(props.treeId);

  const openRef = computed({
    get() {
      return state.value[nodeHash.value] ?? false;
    },
    set(value) {
      state.value[nodeHash.value] = value;
    },
  });

  const nodeHash = computed(() => {
    // converts a UUID to a short hash
    return props.item.id.replace(/-/g, "").substring(0, 8);
  });

  const api = useUserApi();

  const hasFetched = ref(false);
  const items = ref<ItemSummary[]>([]);

  async function fetchItems() {
    const { data, error } = await api.items.getAll({
      locations: [props.item.id],
    });

    if (error) {
      return;
    }

    hasFetched.value = true;
    items.value = data.items;

    console.log("fetched items", items.value);
  }

  watch(
    openRef,
    async value => {
      if (value && !hasFetched.value && props.type === "location") {
        await fetchItems();
      }
    },
    { immediate: true }
  );
</script>

<template>
  <div>
    <div
      class="node flex items-center gap-1 rounded p-1"
      :class="{
        'cursor-pointer hover:bg-base-200': hasChildren,
      }"
      @click="openRef = !openRef"
    >
      <div
        class="p-1/2 rounded mr-1 flex items-center justify-center"
        :class="{
          'hover:bg-base-200': hasChildren,
        }"
      >
        <div v-if="!hasChildren" class="h-6 w-6"></div>
        <label
          v-else
          class="swap swap-rotate"
          :class="{
            'swap-active': openRef,
          }"
        >
          <Icon name="mdi-chevron-right" class="h-6 w-6 swap-off" />
          <Icon name="mdi-chevron-down" class="h-6 w-6 swap-on" />
        </label>
      </div>
      <Icon name="mdi-map-marker" class="h-4 w-4" />
      <NuxtLink class="hover:link text-lg" :to="link" @click.stop>{{ item.name }} </NuxtLink>
    </div>
    <div v-if="openRef" class="ml-4">
      <LocationTreeNode v-for="child in item.children" :key="child.id" :item="child" :tree-id="treeId" />
    </div>
  </div>
</template>

<style scoped></style>
