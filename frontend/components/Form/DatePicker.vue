<template>
  <div v-if="!inline" class="form-control w-full">
    <label class="label">
      <span class="label-text"> {{ label }}</span>
    </label>
    <input ref="input" v-model="selected" type="date" class="input input-bordered w-full" />
  </div>
  <div v-else class="sm:grid sm:grid-cols-4 sm:items-start sm:gap-4">
    <label class="label">
      <span class="label-text"> {{ label }} </span>
    </label>
    <VueDatePicker v-model="selected" :enable-time-picker="false" clearable :dark="isDark" />
  </div>
</template>

<script setup lang="ts">
  // @ts-ignore
  import VueDatePicker from "@vuepic/vue-datepicker";
  import "@vuepic/vue-datepicker/dist/main.css";
  const emit = defineEmits(["update:modelValue", "update:text"]);

  const props = defineProps({
    modelValue: {
      type: Date as () => Date | string | null,
      required: false,
      default: null,
    },
    inline: {
      type: Boolean,
      default: false,
    },
    label: {
      type: String,
      default: "Date",
    },
  });

  const isDark = useIsDark();

  const selected = computed({
    get() {
      // return modelValue as string as YYYY-MM-DD or null

      // String
      if (typeof props.modelValue === "string") {
        // Empty string
        if (props.modelValue === "") {
          return null;
        }

        // Invalid Date string
        if (props.modelValue === "Invalid Date") {
          return null;
        }

        // Valid Date string
        return props.modelValue;
      }

      // Date
      if (props.modelValue instanceof Date) {
        if (props.modelValue.getFullYear() < 1000) {
          return null;
        }

        if (isNaN(props.modelValue.getTime())) {
          return null;
        }

        // Valid Date
        return props.modelValue;
      }

      return null;
    },
    set(value: string | null) {
      // emit update:modelValue with a Date object or null
      console.log("SET", value);
      emit("update:modelValue", value ? new Date(value) : null);
    },
  });
</script>

<style class="scoped">
  ::-webkit-calendar-picker-indicator {
    filter: invert(1);
  }
</style>
