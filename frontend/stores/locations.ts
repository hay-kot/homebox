import { defineStore } from "pinia";
import { LocationsApi } from "~~/lib/api/classes/locations";
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
        this.client.locations.getAll({ filterChildren: true }).then(result => {
          if (result.error) {
            console.error(result.error);
            return;
          }

          this.parents = result.data.items;
        });
      }
      return state.parents ?? [];
    },
    allLocations(state): LocationOutCount[] {
      if (state.Locations === null) {
        this.client.locations.getAll({ filterChildren: false }).then(result => {
          if (result.error) {
            console.error(result.error);
            return;
          }

          this.Locations = result.data.items;
        });
      }
      return state.Locations ?? [];
    },
  },
  actions: {
    async refreshParents(): ReturnType<LocationsApi["getAll"]> {
      const result = await this.client.locations.getAll({ filterChildren: true });
      if (result.error) {
        return result;
      }

      this.parents = result.data.items;
      return result;
    },
    async refreshChildren(): ReturnType<LocationsApi["getAll"]> {
      const result = await this.client.locations.getAll({ filterChildren: false });
      if (result.error) {
        return result;
      }

      this.Locations = result.data.items;
      return result;
    },
  },
});
