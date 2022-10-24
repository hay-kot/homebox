<template>
  <div ref="menu" class="form-control w-full">
    <label class="label">
      <span class="label-text">{{ label }}</span>
    </label>
    <div class="dropdown dropdown-top sm:dropdown-end">
      <div class="relative">
        <input
          v-model="isearch"
          tabindex="0"
          class="input w-full items-center flex flex-wrap border border-gray-400 rounded-lg"
          @keyup.enter="selectFirst"
        />
        <button
          v-if="!!modelValue && Object.keys(modelValue).length !== 0"
          style="transform: translateY(-50%)"
          class="top-1/2 absolute right-2 btn btn-xs btn-circle no-animation"
          @click="clear"
        >
          x
        </button>
      </div>
      <ul
        tabindex="0"
        style="display: inline"
        class="dropdown-content mb-1 menu shadow border border-gray-400 rounded bg-base-100 w-full z-[9999] max-h-60 overflow-y-scroll"
      >
        <li v-for="(obj, idx) in filtered" :key="idx">
          <button type="button" @click="select(obj)">
            {{ usingObjects ? obj[itemText] : obj }}
          </button>
        </li>
        <li class="hidden first:flex">
          <button disabled>
            {{ noResultsText }}
          </button>
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts" setup>
  type ItemsObject = {
    text?: string;
    value?: string;
    [key: string]: unknown;
  };

  interface Props {
    label: string;
    modelValue: string | ItemsObject;
    items: string[] | ItemsObject[];
    itemText?: keyof ItemsObject;
    itemValue?: keyof ItemsObject;
    search?: string;
    noResultsText?: string;
  }

  const emit = defineEmits(["update:modelValue", "update:search"]);
  const props = withDefaults(defineProps<Props>(), {
    label: "",
    modelValue: "",
    items: () => [],
    itemText: "text",
    itemValue: "value",
    search: "",
    noResultsText: "No Results Found",
  });

  function clear() {
    select(value.value);
  }

  const isearch = ref("");
  watch(isearch, () => {
    internalSearch.value = isearch.value;
  });

  const internalSearch = useVModel(props, "search", emit);
  const value = useVModel(props, "modelValue", emit);

  const usingObjects = computed(() => {
    return props.items.length > 0 && typeof props.items[0] === "object";
  });

  /**
   * isStrings is a type guard function to check if the items are an array of string
   */
  function isStrings(_arr: string[] | ItemsObject[]): _arr is string[] {
    return !usingObjects.value;
  }

  function selectFirst() {
    if (filtered.value.length > 0) {
      select(filtered.value[0]);
    }
  }

  watch(
    value,
    () => {
      if (value.value) {
        if (typeof value.value === "string") {
          isearch.value = value.value;
        } else {
          isearch.value = value.value[props.itemText] as string;
        }
      }
    },
    {
      immediate: true,
    }
  );

  function select(obj: string | ItemsObject) {
    if (isStrings(props.items)) {
      if (obj === value.value) {
        value.value = "";
        return;
      }
      value.value = obj;
    } else {
      if (obj === value.value) {
        value.value = {};
        return;
      }

      value.value = obj;
    }
  }

  const filtered = computed(() => {
    if (!isearch.value || isearch.value === "") {
      return props.items;
    }

    if (isStrings(props.items)) {
      return props.items.filter(item => item.toLowerCase().includes(isearch.value.toLowerCase()));
    } else {
      return props.items.filter(item => {
        if (props.itemText && props.itemText in item) {
          return (item[props.itemText] as string).toLowerCase().includes(isearch.value.toLowerCase());
        }
        return false;
      });
    }
  });
</script>
