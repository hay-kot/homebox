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
  });

  const selectedIdx = ref(-1);
  const internalSelected = useVModel(props, "modelValue", emit);

  watch(selectedIdx, newVal => {
    internalSelected.value = props.items[newVal];
  });

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  function compare(a: any, b: any): boolean {
    if (props.value != null) {
      return a[props.value] === b[props.value];
    }
    return a === b;
  }

  watch(
    internalSelected,
    () => {
      const idx = props.items.findIndex(item => compare(item, internalSelected.value));
      selectedIdx.value = idx;
    },
    {
      immediate: true,
    }
  );
</script>
