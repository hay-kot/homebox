<template>
  <div ref="label" class="dropdown dropdown-end dropdown-top w-full">
    <FormTextField v-model="dateText" tabindex="0" label="Date" :inline="inline" readonly />
    <div tabindex="0" class="mt-1 card compact dropdown-content shadow bg-base-100 rounded-box w-64" @blur="resetTime">
      <div class="card-body">
        <div class="flex justify-between items-center">
          <button class="btn btn-xs" @click="prevMonth">
            <Icon class="h-5 w-5" name="mdi-arrow-left"></Icon>
          </button>
          <p class="text-center">{{ month }} {{ year }}</p>
          <button class="btn btn-xs" @click="nextMonth">
            <Icon class="h-5 w-5" name="mdi-arrow-right"></Icon>
          </button>
        </div>
        <div class="grid grid-cols-7 gap-2">
          <div v-for="d in daysIdx" :key="d">
            <p class="text-center">
              {{ d }}
            </p>
          </div>
          <template v-for="day in days">
            <button
              v-if="day.number != ''"
              :key="day.number"
              class="text-center btn-xs btn btn-outline"
              @click="select($event, day.date)"
            >
              {{ day.number }}
            </button>
            <div v-else :key="`${day.number}-empty`"></div>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  const emit = defineEmits(["update:modelValue", "update:text"]);

  const props = defineProps({
    modelValue: {
      type: Date,
      required: false,
      default: null,
    },
    inline: {
      type: Boolean,
      default: false,
    },
  });

  const selected = useVModel(props, "modelValue", emit);
  const dateText = computed(() => {
    if (!validDate(selected.value)) {
      return "";
    }

    if (selected.value) {
      return selected.value.toLocaleDateString();
    }

    return "";
  });

  const time = ref(new Date());
  function resetTime() {
    time.value = new Date();
  }

  const label = ref<HTMLElement>();
  onClickOutside(label, () => {
    resetTime();
  });

  const month = computed(() => {
    return time.value.toLocaleString("default", { month: "long" });
  });

  const year = computed(() => {
    return time.value.getFullYear();
  });

  function nextMonth() {
    time.value.setMonth(time.value.getMonth() + 1);
    time.value = new Date(time.value);
  }

  function prevMonth() {
    time.value.setMonth(time.value.getMonth() - 1);
    time.value = new Date(time.value);
  }

  const daysIdx = computed(() => {
    return ["Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"];
  });

  function select(e: MouseEvent, day: Date) {
    selected.value = day;
    // @ts-ignore - this is a vue3 bug
    e.target.blur();
    resetTime();
  }

  type DayEntry = {
    number: number | string;
    date: Date;
  };

  function daysInMonth(month: number, year: number) {
    return new Date(year, month, 0).getDate();
  }

  const days = computed<DayEntry[]>(() => {
    const days = [];

    const totalDays = daysInMonth(time.value.getMonth() + 1, time.value.getFullYear());

    const firstDay = new Date(time.value.getFullYear(), time.value.getMonth(), 1).getDay();

    for (let i = 0; i < firstDay; i++) {
      days.push({
        number: "",
        date: new Date(),
      });
    }

    for (let i = 1; i <= totalDays; i++) {
      days.push({
        number: i,
        date: new Date(time.value.getFullYear(), time.value.getMonth(), i),
      });
    }

    return days;
  });
</script>
