<template>
  <LineChart
    :chart-options="chartOptions"
    :chart-data="chartData"
    :chart-id="chartId"
    :dataset-id-key="datasetIdKey"
    :css-classes="cssClasses"
    :styles="styles"
    :width="width"
    :height="height"
  />
</template>

<script lang="ts">
  import { Line as LineChart } from "vue-chartjs";
  import {
    Chart as ChartJS,
    PointElement,
    Title,
    Tooltip,
    Legend,
    CategoryScale,
    LinearScale,
    LineElement,
  } from "chart.js";
  import { TChartData } from "vue-chartjs/dist/types";

  ChartJS.register(Title, Tooltip, Legend, CategoryScale, LinearScale, PointElement, LineElement);

  export default defineComponent({
    name: "BarChart",
    components: {
      LineChart,
    },
    props: {
      chartId: {
        type: String,
        default: "bar-chart",
      },
      datasetIdKey: {
        type: String,
        default: "label",
      },
      width: {
        type: Number,
        default: 400,
      },
      height: {
        type: Number,
        default: 400,
      },
      cssClasses: {
        default: "",
        type: String,
      },
      styles: {
        type: Object,
        default: () => {
          return {};
        },
      },
      chartData: {
        type: Object as () => TChartData<"line", number[], unknown>,
        default: () => {
          return {
            labels: ["January", "February", "March"],
            datasets: [{ data: [40, 20, 12] }],
          };
        },
      },
    },
    data() {
      return {
        chartOptions: {
          responsive: true,
          scales: {
            x: {
              display: false,
            },
            y: {
              display: true,
            },
          },
          elements: {
            line: {
              borderWidth: 5,
            },
            point: {
              radius: 4,
            },
          },
        },
      };
    },
  });
</script>
