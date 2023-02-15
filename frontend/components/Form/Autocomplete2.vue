<template>
  <div>
    <Combobox v-model="value">
      <ComboboxLabel class="label">
        <span class="label-text">{{ label }}</span>
      </ComboboxLabel>
      <div class="relative">
        <ComboboxInput
          :display-value="i => extractDisplay(i as SupportValues)"
          class="w-full input input-bordered"
          @change="search = $event.target.value"
        />
        <ComboboxButton class="absolute inset-y-0 right-0 flex items-center rounded-r-md px-2 focus:outline-none">
          <Icon name="mdi-chevron-down" class="w-5 h-5" />
        </ComboboxButton>
        <ComboboxOptions
          v-if="computedItems.length > 0"
          class="absolute dropdown-content z-10 mt-2 max-h-60 w-full overflow-auto rounded-md card bg-base-100 border border-gray-400"
        >
          <ComboboxOption
            v-for="item in computedItems"
            :key="item.id"
            v-slot="{ active, selected }"
            :value="item.value"
            as="template"
          >
            <li
              :class="[
                'relative cursor-default select-none py-2 pl-3 pr-9 duration-75 ease-in-out transition-colors',
                active ? 'bg-primary text-primary-content' : 'text-base-content',
              ]"
            >
              <slot name="display" v-bind="{ item: item, selected, active }">
                <span :class="['block truncate', selected && 'font-semibold']">
                  {{ item.display }}
                </span>
                <span
                  v-if="selected"
                  :class="[
                    'absolute inset-y-0 right-0 flex text-primary items-center pr-4',
                    active ? 'text-primary-content' : 'bg-primary',
                  ]"
                >
                  <Icon name="mdi-check" class="h-5 w-5" aria-hidden="true" />
                </span>
              </slot>
            </li>
          </ComboboxOption>
        </ComboboxOptions>
      </div>
    </Combobox>
  </div>
</template>

<script setup lang="ts">
  import {
    Combobox,
    ComboboxInput,
    ComboboxOptions,
    ComboboxOption,
    ComboboxButton,
    ComboboxLabel,
  } from "@headlessui/vue";

  type SupportValues = string | { [key: string]: any };

  type ComboItem = {
    display: string;
    value: SupportValues;
    id: number;
  };

  type Props = {
    label: string;
    modelValue: SupportValues | null | undefined;
    items: string[] | object[];
    display?: string;
    multiple?: boolean;
  };

  const emit = defineEmits(["update:modelValue", "update:search"]);
  const props = withDefaults(defineProps<Props>(), {
    label: "",
    modelValue: "",
    display: "text",
    multiple: false,
  });

  const search = ref("");
  const value = useVModel(props, "modelValue", emit);

  function extractDisplay(item?: SupportValues): string {
    if (!item) {
      return "";
    }

    if (typeof item === "string") {
      return item;
    }

    if (props.display in item) {
      return item[props.display] as string;
    }

    // Try these options as well
    const fallback = ["name", "title", "display", "value"];
    for (let i = 0; i < fallback.length; i++) {
      const key = fallback[i];
      if (key in item) {
        return item[key] as string;
      }
    }

    return "";
  }

  const computedItems = computed<ComboItem[]>(() => {
    const list: ComboItem[] = [];

    for (let i = 0; i < props.items.length; i++) {
      const item = props.items[i];

      const out: Partial<ComboItem> = {
        id: i,
        value: item,
      };

      switch (typeof item) {
        case "string":
          out.display = item;
          break;
        case "object":
          // @ts-ignore - up to the user to provide a valid display key
          out.display = item[props.display] as string;
          break;
        default:
          out.display = "";
          break;
      }

      if (search.value && out.display) {
        const foldSearch = search.value.toLowerCase();
        const foldDisplay = out.display.toLowerCase();

        if (foldDisplay.startsWith(foldSearch)) {
          list.push(out as ComboItem);
        }

        continue;
      }

      list.push(out as ComboItem);
    }

    return list;
  });
</script>
