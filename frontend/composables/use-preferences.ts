import { Ref } from "vue";

export type DaisyTheme =
  | "light"
  | "dark"
  | "cupcake"
  | "bumblebee"
  | "emerald"
  | "corporate"
  | "synthwave"
  | "retro"
  | "cyberpunk"
  | "valentine"
  | "halloween"
  | "garden"
  | "forest"
  | "aqua"
  | "lofi"
  | "pastel"
  | "fantasy"
  | "wireframe"
  | "black"
  | "luxury"
  | "dracula"
  | "cmyk"
  | "autumn"
  | "business"
  | "acid"
  | "lemonade"
  | "night"
  | "coffee"
  | "winter";

export type LocationViewPreferences = {
  showDetails: boolean;
  showEmpty: boolean;
  editorSimpleView: boolean;
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
      editorSimpleView: true,
      theme: "garden",
    },
    { mergeDefaults: true }
  );

  // casting is required because the type returned is removable, however since we
  // use `mergeDefaults` the result _should_ always be present.
  return results as unknown as Ref<LocationViewPreferences>;
}
