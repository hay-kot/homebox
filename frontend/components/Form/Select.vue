<template>
  <div class="form-control w-full">
    <label class="label">
      <span class="label-text">{{ label }}</span>
    </label>
    <select v-model="selectedIdx" class="select select-bordered">
      <option disabled selected>Pick one</option>
      <option v-for="(obj, idx) in items" :key="name != '' ? obj[name] : obj" :value="idx">
        {{ name != "" ? obj[name] : obj }}
      </option>
    </select>
    <!-- <label class="label">
      <span class="label-text-alt">Alt label</span>
      <span class="label-text-alt">Alt label</span>
    </label> -->
  </div>
</template>

<script lang="ts" setup>
  const emit = defineEmits(["update:modelValue"]);
  const props = defineProps({
    label: {
      type: String,
      default: "",
    },
    modelValue: {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      type: [Object, String, Boolean] as any,
      default: null,
    },
    items: {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      type: Array as () => any[],
      required: true,
    },
    name: {
      type: String,
      default: "name",
    },
    value: {
      type: String,
      default: null,
      required: false,
    },
    selectFirst: {
      type: Boolean,
      default: false,
    },
  });

  function syncSelect() {
    if (!props.modelValue) {
      if (props.selectFirst) {
        selectedIdx.value = 0;
      }
      return;
    }
    // Check if we're already synced
    if (props.value) {
      if (props.modelValue[props.value] === props.items[selectedIdx.value][props.value]) {
        return;
      }
    } else if (props.modelValue === props.items[selectedIdx.value]) {
      return;
    }

    const idx = props.items.findIndex(item => {
      if (props.value) {
        return item[props.value] === props.modelValue;
      }
      return item === props.modelValue;
    });

    selectedIdx.value = idx;
  }

  watch(() => props.items, syncSelect);
  watch(() => props.modelValue, syncSelect);

  const selectedIdx = ref(0);
  watch(
    () => selectedIdx.value,
    () => {
      if (props.value) {
        emit("update:modelValue", props.items[selectedIdx.value][props.value]);
        return;
      }
      emit("update:modelValue", props.items[selectedIdx.value]);
    }
  );
</script>
