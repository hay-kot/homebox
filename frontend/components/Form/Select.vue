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
  const emit = defineEmits(["update:modelValue", "update:value"]);
  const props = defineProps({
    label: {
      type: String,
      default: "",
    },
    modelValue: {
      type: [Object, String] as any,
      default: null,
    },
    items: {
      type: Array as () => any[],
      required: true,
    },
    name: {
      type: String,
      default: "name",
    },
    valueKey: {
      type: String,
      default: null,
    },
    value: {
      type: String,
      default: "",
    },
    compareKey: {
      type: String,
      default: null,
    },
  });

  const selectedIdx = ref(-1);

  const internalSelected = useVModel(props, "modelValue", emit);
  const internalValue = useVModel(props, "value", emit);

  watch(
    selectedIdx,
    newVal => {
      if (newVal === -1) {
        return;
      }

      if (props.value) {
        internalValue.value = props.items[newVal][props.valueKey];
      }

      internalSelected.value = props.items[newVal];
    },
    { immediate: true }
  );

  watch(
    [internalSelected, () => props.value],
    () => {
      if (props.valueKey) {
        const idx = props.items.findIndex(item => compare(item, internalValue.value));
        selectedIdx.value = idx;
        return;
      }
      const idx = props.items.findIndex(item => compare(item, internalSelected.value));
      selectedIdx.value = idx;
    },
    { immediate: true }
  );

  function compare(a: any, b: any): boolean {
    if (a === b) {
      return true;
    }

    if (props.valueKey) {
      return a[props.valueKey] === b;
    }

    // Try compare key
    if (props.compareKey && a && b) {
      return a[props.compareKey] === b[props.compareKey];
    }

    return JSON.stringify(a) === JSON.stringify(b);
  }
</script>
