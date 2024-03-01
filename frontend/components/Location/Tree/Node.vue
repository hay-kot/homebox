<script setup lang="ts">
  import { useTreeState } from "./tree-state";
  import type { TreeItem } from "~~/lib/api/types/data-contracts";
  import MdiChevronDown from "~icons/mdi/chevron-down";
  import MdiChevronRight from "~icons/mdi/chevron-right";
  import MdiMapMarker from "~icons/mdi/map-marker";
  import MdiPackageVariant from "~icons/mdi/package-variant";

  type Props = {
    treeId: string;
    item: TreeItem;
  };
  const props = withDefaults(defineProps<Props>(), {});

  const link = computed(() => {
    return props.item.type === "location" ? `/location/${props.item.id}` : `/item/${props.item.id}`;
  });

  const hasChildren = computed(() => {
    return props.item.children.length > 0;
  });

  const state = useTreeState(props.treeId);

  const openRef = computed({
    get() {
      return state.value[nodeHash.value] ?? false;
    },
    set(value: boolean) {
      state.value[nodeHash.value] = value;
    },
  });

  const nodeHash = computed(() => {
    // converts a UUID to a short hash
    return props.item.id.replace(/-/g, "").substring(0, 8);
  });
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
          <MdiChevronRight name="mdi-chevron-right" class="h-6 w-6 swap-off" />
          <MdiChevronDown name="mdi-chevron-down" class="h-6 w-6 swap-on" />
        </label>
      </div>
      <MdiMapMarker v-if="item.type === 'location'" class="h-4 w-4" />
      <MdiPackageVariant v-else class="h-4 w-4" />
      <NuxtLink class="hover:link text-lg" :to="link" @click.stop>{{ item.name }} </NuxtLink>
    </div>
    <div v-if="openRef" class="ml-4">
      <LocationTreeNode v-for="child in item.children" :key="child.id" :item="child" :tree-id="treeId" />
    </div>
  </div>
</template>

<style scoped></style>
