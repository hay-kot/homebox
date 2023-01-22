import { Ref } from "vue";
import { DaisyTheme } from "~~/lib/data/themes";

export type LocationViewPreferences = {
  showDetails: boolean;
  showEmpty: boolean;
  editorAdvancedView: boolean;
  theme: DaisyTheme;
};

/**
 * useViewPreferences loads the view preferences from local storage and hydrates
 * them. These are reactive and will update the local storage when changed.
 */
export function useViewPreferences(): Ref<LocationViewPreferences> {
  const results = useLocalStorage(
    "homebox/preferences/location",
    {
      showDetails: true,
      showEmpty: true,
      editorAdvancedView: false,
      theme: "homebox",
    },
    { mergeDefaults: true }
  );

  // casting is required because the type returned is removable, however since we
  // use `mergeDefaults` the result _should_ always be present.
  return results as unknown as Ref<LocationViewPreferences>;
}
