import { ComputedRef } from "vue";
import { DaisyTheme } from "./use-preferences";

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
      htmlEl.value.setAttribute("data-theme", newTheme);
    }

    themeRef.value = newTheme;
  };

  const htmlEl = ref<HTMLElement>(null);

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
