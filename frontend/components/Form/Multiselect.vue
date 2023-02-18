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
        style="display: inline"
        class="dropdown-content mb-1 menu shadow border border-gray-400 rounded bg-base-100 w-full z-[9999] max-h-60 overflow-y-scroll"
      >
        <li
          v-for="(obj, idx) in items"
          :key="idx"
          :class="{
            bordered: selected[idx],
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
      type: Array as () => any[],
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
    selectFirst: {
      type: Boolean,
      default: false,
    },
  });

  const value = useVModel(props, "modelValue", emit);

  const selected = computed<Record<number, boolean>>(() => {
    const obj: Record<number, boolean> = {};
    value.value.forEach(itm => {
      const idx = props.items.findIndex(item => item[props.name] === itm.name);
      obj[idx] = true;
    });
    return obj;
  });

  function toggle(index: number) {
    const item = props.items[index];
    if (selected.value[index]) {
      value.value = value.value.filter(itm => itm.name !== item.name);
    } else {
      value.value = [...value.value, item];
    }
  }
</script>
