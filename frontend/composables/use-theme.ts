import type { ComputedRef } from "vue";
import type { DaisyTheme } from "~~/lib/data/themes";

export interface UseTheme {
  theme: ComputedRef<DaisyTheme>;
  setTheme: (theme: DaisyTheme) => void;
}

const themeRef = ref<DaisyTheme>("garden");

export function useTheme(): UseTheme {
  const preferences = useViewPreferences();
  themeRef.value = preferences.value.theme;

  const setTheme = (newTheme: DaisyTheme) => {
    preferences.value.theme = newTheme;

    if (htmlEl) {
      htmlEl.value?.setAttribute("data-theme", newTheme);
    }

    themeRef.value = newTheme;
  };

  const htmlEl = ref<HTMLElement | null>();

  onMounted(() => {
    if (htmlEl.value) {
      return;
    }

    htmlEl.value = document.querySelector("html");
  });

  const theme = computed(() => {
    return themeRef.value;
  });

  return { theme, setTheme };
}

export function useIsDark() {
  const theme = useTheme();

  const darkthemes = [
    "synthwave",
    "retro",
    "cyberpunk",
    "valentine",
    "halloween",
    "forest",
    "aqua",
    "black",
    "luxury",
    "dracula",
    "business",
    "night",
    "coffee",
  ];

  return computed(() => {
    return darkthemes.includes(theme.theme.value);
  });
}
