import type { UserClient } from "~~/lib/api/user";

type StatCard = {
  label: string;
  value: number;
  type: "currency" | "number";
};

export function statCardData(api: UserClient) {
  const { data: statistics } = useAsyncData(async () => {
    const { data } = await api.stats.group();
    return data;
  });

  return computed(() => {
    return [
      {
        label: "Total Value",
        value: statistics.value?.totalItemPrice || 0,
        type: "currency",
      },
      {
        label: "Total Items",
        value: statistics.value?.totalItems || 0,
        type: "number",
      },
      {
        label: "Total Locations",
        value: statistics.value?.totalLocations || 0,
        type: "number",
      },
      {
        label: "Total Labels",
        value: statistics.value?.totalLabels || 0,
        type: "number",
      },
    ] as StatCard[];
  });
}
