import { defineStore } from "pinia";
import { LocationOutCount } from "~~/lib/api/types/data-contracts";

export const useLocationStore = defineStore("locations", {
  state: () => ({
    parents: null as LocationOutCount[] | null,
    Locations: null as LocationOutCount[] | null,
    client: useUserApi(),
  }),
  getters: {
    /**
     * locations represents the locations that are currently in the store. The store is
     * synched with the server by intercepting the API calls and updating on the
     * response
     */
    parentLocations(state): LocationOutCount[] {
      if (state.parents === null) {
        Promise.resolve(this.refreshParents());
      }
      return state.parents;
    },
    allLocations(state): LocationOutCount[] {
      if (state.Locations === null) {
        Promise.resolve(this.refreshChildren());
      }
      return state.Locations;
    },
  },
  actions: {
    async refreshParents(): Promise<LocationOutCount[]> {
      const result = await this.client.locations.getAll({ filterChildren: true });
      if (result.error) {
        return result;
      }

      this.parents = result.data.items;
      return result;
    },
    async refreshChildren(): Promise<LocationOutCount[]> {
      const result = await this.client.locations.getAll({ filterChildren: false });
      if (result.error) {
        return result;
      }

      this.Locations = result.data.items;
      return result;
    },
  },
});
