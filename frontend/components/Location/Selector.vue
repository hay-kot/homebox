<template>
  <FormAutocomplete
    v-model="value"
    v-model:search="form.search"
    :items="locations"
    item-text="display"
    item-value="id"
    item-search="name"
    label="Parent Location"
  >
    <template #display="{ item }">
      <div>
        <div>
          {{ item.name }}
        </div>
        <div v-if="item.name != item.display" class="text-xs mt-1">{{ item.display }}</div>
      </div>
    </template>
  </FormAutocomplete>
</template>

<script lang="ts" setup>
  import { LocationSummary } from "~~/lib/api/types/data-contracts";

  type Props = {
    modelValue?: LocationSummary | null;
  };

  const props = defineProps<Props>();

  const value = useVModel(props, "modelValue");

  const locations = await useFlatLocations();

  const form = ref({
    parent: null as LocationSummary | null,
    search: "",
  });

  // Whenever parent goes from value to null reset search
  watch(
    () => value.value,
    () => {
      if (!value.value) {
        form.value.search = "";
      }
    }
  );
</script>
