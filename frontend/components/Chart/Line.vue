<template>
  <div ref="el" class="min-h-full flex flex-col">
    {{ styles }}
    <LineChart :chart-options="options" :chart-data="chartData" :styles="styles" />
  </div>
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
      cssClasses: {
        default: "",
        type: String,
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
    setup() {
      const el = ref<HTMLElement | null>(null);

      const calcHeight = ref(0);
      const calcWidth = ref(0);

      function resize() {
        calcHeight.value = el.value?.offsetHeight || 0;
        calcWidth.value = el.value?.offsetWidth || 0;
      }

      onMounted(() => {
        resize();
        window.addEventListener("resize", resize);
      });

      onUnmounted(() => {
        window.removeEventListener("resize", resize);
      });

      const styles = computed(() => {
        return {
          height: `${calcHeight.value}px`,
          width: `${calcWidth.value}px`,
          position: "relative",
        };
      });

      return {
        el,
        parentHeight: calcHeight,
        styles,
      };
    },
    data() {
      return {
        options: {
          responsive: true,
          maintainAspectRatio: false,
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

<style></style>
