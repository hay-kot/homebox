import { defineStore } from "pinia";
import { ItemOut } from "~~/lib/api/types/data-contracts";

export const useItemStore = defineStore("items", {
  state: () => ({
    allItems: null as ItemOut[] | null,
    client: useUserApi(),
  }),
  getters: {
    /**
     * items represents the items that are currently in the store. The store is
     * synched with the server by intercepting the API calls and updating on the
     * response.
     */
    items(state): ItemOut[] {
      if (state.allItems === null) {
        Promise.resolve(this.refresh());
      }
      return state.allItems;
    },
  },
  actions: {
    async refresh(): Promise<ItemOut[]> {
      const result = await this.client.items.getAll();
      if (result.error) {
        return result;
      }

      this.allItems = result.data.items;
      return result;
    },
  },
});
