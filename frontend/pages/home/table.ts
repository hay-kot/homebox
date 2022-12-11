import { TableHeader } from "~~/components/global/Table.types";

import { UserClient } from "~~/lib/api/user";

export function itemsTable(api: UserClient) {
  const { data: items } = useAsyncData(async () => {
    const { data } = await api.items.getAll({
      page: 1,
      pageSize: 5,
    });
    return data.items;
  });

  const headers = [
    {
      text: "Name",
      sortable: true,
      value: "name",
    },
    {
      text: "Location",
      value: "location.name",
    },
    {
      text: "Warranty",
      value: "warranty",
      align: "center",
    },
    {
      text: "Price",
      value: "purchasePrice",
      align: "center",
    },
  ] as TableHeader[];

  return computed(() => {
    return {
      headers,
      items: items.value || [],
    };
  });
}
