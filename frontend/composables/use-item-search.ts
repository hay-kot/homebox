import type { ItemSummary, LabelSummary, LocationSummary } from "~~/lib/api/types/data-contracts";
import type { UserClient } from "~~/lib/api/user";

type SearchOptions = {
  immediate?: boolean;
};

export function useItemSearch(client: UserClient, opts?: SearchOptions) {
  const query = ref("");
  const locations = ref<LocationSummary[]>([]);
  const labels = ref<LabelSummary[]>([]);
  const results = ref<ItemSummary[]>([]);
  const includeArchived = ref(false);

  watchDebounced(query, search, { debounce: 250, maxWait: 1000 });
  async function search() {
    const locIds = locations.value.map(l => l.id);
    const labelIds = labels.value.map(l => l.id);

    const { data, error } = await client.items.getAll({
      q: query.value,
      locations: locIds,
      labels: labelIds,
      includeArchived: includeArchived.value,
    });
    if (error) {
      return;
    }
    results.value = data.items;
  }

  if (opts?.immediate) {
    search();
  }

  return {
    query,
    results,
    locations,
    labels,
  };
}
