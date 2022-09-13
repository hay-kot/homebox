import { defineStore } from "pinia";
import { LocationCount } from "~~/lib/api/types/data-contracts";

export const useLocationStore = defineStore("locations", {
  state: () => ({
    allLocations: null as LocationCount[] | null,
    client: useUserApi(),
  }),
  getters: {
    /**
     * locations represents the locations that are currently in the store. The store is
     * synched with the server by intercepting the API calls and updating on the
     * response
     */
    locations(state): LocationCount[] {
      if (state.allLocations === null) {
        Promise.resolve(this.refresh());
      }
      return state.allLocations;
    },
  },
  actions: {
    async refresh(): Promise<LocationCount[]> {
      const result = await this.client.locations.getAll();
      if (result.error) {
        return result;
      }

      this.allLocations = result.data.items;
      return result;
    },
  },
});
