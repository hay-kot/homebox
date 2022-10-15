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
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      type: [Object, String] as any,
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
    valueKey: {
      type: String,
      default: null,
    },
    value: {
      type: String,
      default: "",
    },
  });

  const selectedIdx = ref(-1);

  const internalSelected = useVModel(props, "modelValue", emit);

  watch(selectedIdx, newVal => {
    internalSelected.value = props.items[newVal];
  });

  watch(internalSelected, newVal => {
    if (props.valueKey) {
      emit("update:value", newVal[props.valueKey]);
    }
  });

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  function compare(a: any, b: any): boolean {
    if (a === b) {
      return true;
    }

    if (!a || !b) {
      return false;
    }

    return JSON.stringify(a) === JSON.stringify(b);
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
