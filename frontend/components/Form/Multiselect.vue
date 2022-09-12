<template>
  <div ref="menu" class="form-control w-full">
    <label class="label">
      <span class="label-text">{{ label }}</span>
    </label>
    <div class="dropdown dropdown-top sm:dropdown-end">
      <div tabindex="0" class="w-full min-h-[48px] flex gap-2 p-4 flex-wrap border border-gray-400 rounded-lg">
        <span v-for="itm in value" :key="name != '' ? itm[name] : itm" class="badge">
          {{ name != "" ? itm[name] : itm }}
        </span>
      </div>
      <ul
        tabindex="0"
        class="dropdown-content mb-1 menu shadow border border-gray-400 rounded bg-base-100 w-full z-[9999] max-h-60 overflow-y-scroll scroll-bar"
      >
        <li
          v-for="(obj, idx) in items"
          :key="idx"
          :class="{
            bordered: selectedIndexes[idx],
          }"
        >
          <button type="button" @click="toggle(idx)">
            {{ name != "" ? obj[name] : obj }}
          </button>
        </li>
      </ul>
    </div>
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
      type: Array as () => any[],
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
    selectFirst: {
      type: Boolean,
      default: false,
    },
  });

  const selectedIndexes = ref<Record<number, boolean>>({});

  function toggle(index: number) {
    selectedIndexes.value[index] = !selectedIndexes.value[index];

    const item = props.items[index];

    if (selectedIndexes.value[index]) {
      value.value = [...value.value, item];
    } else {
      value.value = value.value.filter(itm => itm !== item);
    }
  }

  watchOnce(
    () => props.items,
    () => {
      if (props.selectFirst && props.items.length > 0) {
        value.value = props.items[0];
      }
    }
  );

  const value = useVModel(props, "modelValue", emit);
</script>
