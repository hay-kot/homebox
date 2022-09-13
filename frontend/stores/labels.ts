import { defineStore } from "pinia";
import { LabelOut } from "~~/lib/api/types/data-contracts";

export const useLabelStore = defineStore("labels", {
  state: () => ({
    allLabels: null as LabelOut[] | null,
    client: useUserApi(),
  }),
  getters: {
    /**
     * labels represents the labels that are currently in the store. The store is
     * synched with the server by intercepting the API calls and updating on the
     * response.
     */
    labels(state): LabelOut[] {
      if (state.allLabels === null) {
        Promise.resolve(this.refresh());
      }
      return state.allLabels;
    },
  },
  actions: {
    async refresh(): Promise<LabelOut[]> {
      const result = await this.client.labels.getAll();
      if (result.error) {
        return result;
      }

      this.allLabels = result.data.items;
      return result;
    },
  },
});
