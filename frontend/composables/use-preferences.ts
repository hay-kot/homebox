import { Ref } from 'vue';

export type LocationViewPreferences = {
  showDetails: boolean;
};

/**
 * useLocationViewPreferences loads the view preferences from local storage and hydrates
 * them. These are reactive and will update the local storage when changed.
 */
export function useLocationViewPreferences(): Ref<LocationViewPreferences> {
  const results = useLocalStorage(
    'homebox/preferences/location',
    {
      showDetails: true,
    },
    { mergeDefaults: true }
  );

  // casting is required because the type returned is removable, however since we
  // use `mergeDefaults` the result _should_ always be present.
  return results as unknown as Ref<LocationViewPreferences>;
}
