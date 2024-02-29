<script setup lang="ts">
  import { ViewType } from "~~/composables/use-preferences";
  import { ItemSummary } from "~~/lib/api/types/data-contracts";

  type Props = {
    view?: ViewType;
    items: ItemSummary[];
  };

  const preferences = useViewPreferences();

  const props = defineProps<Props>();
  const viewSet = computed(() => {
    return !!props.view;
  });

  const itemView = computed(() => {
    return props.view ?? preferences.value.itemDisplayView;
  });

  function setViewPreference(view: ViewType) {
    preferences.value.itemDisplayView = view;
  }
</script>

<template>
  <section>
    <BaseSectionHeader class="mb-2 flex justify-between items-center">
      {{ $t("item.selectable.title") }}
      <template #description>
        <div v-if="!viewSet" class="dropdown dropdown-hover dropdown-left">
          <label tabindex="0" class="btn btn-ghost m-1">
            <Icon name="mdi-dots-vertical" class="h-7 w-7" />
          </label>
          <ul tabindex="0" class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-32">
            <li>
              <button @click="setViewPreference('card')">
                <Icon name="mdi-card-text-outline" class="h-5 w-5" />
                {{ $t("item.selectable.card") }}
              </button>
            </li>
            <li>
              <button @click="setViewPreference('table')">
                <Icon name="mdi-table" class="h-5 w-5" />
                {{ $t("item.selectable.table") }}
              </button>
            </li>
          </ul>
        </div>
      </template>
    </BaseSectionHeader>

    <template v-if="itemView === 'table'">
      <ItemViewTable :items="items" />
    </template>
    <template v-else>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        <ItemCard v-for="item in items" :key="item.id" :item="item" />
        <div class="first:block hidden text-lg">{{ $t("item.selectable.empty") }}</div>
      </div>
    </template>
  </section>
</template>

<style scoped></style>
