<template>
  <div v-if="!inline" class="form-control w-full">
    <label class="label">
      <span class="label-text"> Date </span>
    </label>
    <input ref="input" v-model="selected" type="date" class="input input-bordered w-full" />
  </div>
  <div v-else class="sm:grid sm:grid-cols-4 sm:items-start sm:gap-4">
    <label class="label">
      <span class="label-text"> Date </span>
    </label>
    <input v-model="selected" type="date" class="input input-bordered col-span-3 w-full mt-2" />
  </div>
</template>

<script setup lang="ts">
  const emit = defineEmits(["update:modelValue", "update:text"]);

  const props = defineProps({
    modelValue: {
      type: Date,
      required: false,
      default: null,
    },
    inline: {
      type: Boolean,
      default: false,
    },
  });

  const selected = computed({
    get() {
      // return modelValue as string as YYYY-MM-DD or null
      if (validDate(props.modelValue)) {
        return props.modelValue ? props.modelValue.toISOString().split("T")[0] : null;
      }

      return null;
    },
    set(value: string | null) {
      // emit update:modelValue with a Date object or null
      emit("update:modelValue", value ? new Date(value) : null);
    },
  });
</script>

<style class="scoped">
  ::-webkit-calendar-picker-indicator {
    filter: invert(1);
  }
</style>
