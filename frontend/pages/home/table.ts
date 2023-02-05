import { UserClient } from "~~/lib/api/user";

export function itemsTable(api: UserClient) {
  const { data: items } = useAsyncData(async () => {
    const { data } = await api.items.getAll({
      page: 1,
      pageSize: 5,
    });
    return data.items;
  });

  return computed(() => {
    return {
      items: items.value || [],
    };
  });
}
