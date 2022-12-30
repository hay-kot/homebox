<template>
  <DoughnutChart
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
  import { Doughnut as DoughnutChart } from "vue-chartjs";
  import { Chart as ChartJS, Title, Tooltip, Legend, CategoryScale, LinearScale, ArcElement } from "chart.js";
  import { TChartData } from "vue-chartjs/dist/types";

  ChartJS.register(Title, Tooltip, Legend, CategoryScale, LinearScale, ArcElement);

  export default defineComponent({
    name: "BarChart",
    components: {
      DoughnutChart,
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
        type: Object as () => TChartData<"doughnut", number[], unknown>,
        default: () => {
          return {
            labels: ["Red", "Blue", "Yellow"],
            datasets: [
              {
                label: "My First Dataset",
                data: [300, 50, 100],
                backgroundColor: ["rgb(255, 99, 132)", "rgb(54, 162, 235)", "rgb(255, 205, 86)"],
                hoverOffset: 4,
              },
            ],
          };
        },
      },
    },
    data() {
      return {
        chartOptions: {
          responsive: false,
          // Legend on the left
          plugins: {
            legend: {
              position: "bottom",
            },
            // Display percentage
            // tooltip: {
            //   callbacks: {
            //     label: context => {
            //       const label = context.dataset?.label || "";
            //       const value = context.parsed.y;
            //       return `${label}: ${value}%`;
            //     },
            //   },
            // },
          },
        },
      };
    },
  });
</script>
