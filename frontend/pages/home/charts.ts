import { TChartData } from "vue-chartjs/dist/types";
import { UserClient } from "~~/lib/api/user";

export function purchasePriceOverTimeChart(api: UserClient) {
  const { data: timeseries } = useAsyncData(async () => {
    const { data } = await api.stats.totalPriceOverTime();
    return data;
  });

  const primary = useCssVar("--p");

  return computed(() => {
    if (!timeseries.value) {
      return {
        labels: ["Purchase Price"],
        datasets: [
          {
            label: "Purchase Price",
            data: [],
            backgroundColor: primary.value,
            borderColor: primary.value,
          },
        ],
      } as TChartData<"line", number[], unknown>;
    }

    let start = timeseries.value?.valueAtStart;

    return {
      labels: timeseries?.value.entries.map(t => new Date(t.date).toDateString()) || [],
      datasets: [
        {
          label: "Purchase Price",
          data:
            timeseries.value?.entries.map(t => {
              start += t.value;
              return start;
            }) || [],
          backgroundColor: primary.value,
          borderColor: primary.value,
        },
      ],
    } as TChartData<"line", number[], unknown>;
  });
}

export function inventoryByLocationChart(api: UserClient) {
  const { data: donutSeries } = useAsyncData(async () => {
    const { data } = await api.stats.locations();
    return data;
  });

  const primary = useCssVar("--p");
  const secondary = useCssVar("--s");
  const neutral = useCssVar("--n");

  return computed(() => {
    return {
      labels: donutSeries.value?.map(l => l.name) || [],
      datasets: [
        {
          label: "Value",
          data: donutSeries.value?.map(l => l.total) || [],
          backgroundColor: [primary.value, secondary.value, neutral.value],
          borderColor: [primary.value, secondary.value, neutral.value],
          hoverOffset: 4,
        },
      ],
    };
  });
}
