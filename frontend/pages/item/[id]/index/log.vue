<script setup lang="ts">
  import DatePicker from "~~/components/Form/DatePicker.vue";
  import { ItemOut } from "~~/lib/api/types/data-contracts";

  const props = defineProps<{
    item: ItemOut;
  }>();

  const api = useUserApi();
  const toast = useNotifier();

  const { data: log, refresh: refreshLog } = useAsyncData(async () => {
    const { data } = await api.items.maintenance.getLog(props.item.id);
    return data;
  });

  const stats = computed(() => {
    if (!log.value) return [];

    return [
      {
        id: "total",
        title: "Total Cost",
        subtitle: "Sum over all entries",
        value: fmtCurrency(log.value.costTotal),
      },
      {
        id: "average",
        title: "Monthly Average",
        subtitle: "Average over all entries",
        value: fmtCurrency(log.value.costAverage),
      },
    ];
  });

  const entry = reactive({
    modal: false,
    name: "",
    date: new Date(),
    description: "",
    cost: "",
  });

  function newEntry() {
    entry.modal = true;
  }

  async function createEntry() {
    const { error } = await api.items.maintenance.create(props.item.id, {
      name: entry.name,
      date: entry.date,
      description: entry.description,
      cost: entry.cost,
    });

    if (error) {
      toast.error("Failed to create entry");
      return;
    }

    entry.modal = false;

    refreshLog();
  }

  async function deleteEntry(id: string) {
    const { error } = await api.items.maintenance.delete(props.item.id, id);

    if (error) {
      toast.error("Failed to delete entry");
      return;
    }

    refreshLog();
  }
</script>

<template>
  <div v-if="log">
    <BaseModal v-model="entry.modal">
      <template #title> Create Entry </template>
      <form @submit.prevent="createEntry">
        <FormTextField v-model="entry.name" autofocus label="Entry Name" />
        <DatePicker v-model="entry.date" label="Date" />
        <FormTextArea v-model="entry.description" label="Notes" />
        <FormTextField v-model="entry.cost" autofocus label="Cost" />
        <div class="py-2 flex justify-end">
          <BaseButton type="submit" class="ml-2">
            <template #icon>
              <Icon name="mdi-post" />
            </template>
            Create
          </BaseButton>
        </div>
      </form>
    </BaseModal>

    <div class="flex">
      <BaseButton class="ml-auto" size="sm" @click="newEntry()">
        <template #icon>
          <Icon name="mdi-post" />
        </template>
        Log Maintenance
      </BaseButton>
    </div>
    <section class="page-layout my-6">
      <div class="main-slot container space-y-6">
        <BaseCard v-for="e in log.entries" :key="e.id">
          <BaseSectionHeader class="p-6 border-b border-b-gray-300">
            <span class="text-base-content">
              {{ e.name }}
            </span>
            <template #description>
              <div class="flex gap-2">
                <div class="badge p-3">
                  <Icon name="mdi-calendar" class="mr-2" />
                  <DateTime :date="e.date" format="human" />
                </div>
                <div class="tooltip tooltip-primary" data-tip="Cost">
                  <div class="badge badge-primary p-3">
                    <Currency :amount="e.cost" />
                  </div>
                </div>
              </div>
            </template>
          </BaseSectionHeader>
          <div class="p-6">
            <Markdown :source="e.description" />
          </div>
          <div class="flex justify-end p-4">
            <BaseButton size="sm" @click="deleteEntry(e.id)">
              <template #icon>
                <Icon name="mdi-delete" />
              </template>
              Delete
            </BaseButton>
          </div>
        </BaseCard>
      </div>
      <div class="side-slot space-y-6">
        <div v-for="stat in stats" :key="stat.id" class="stats block shadow-xl border-l-primary">
          <div class="stat">
            <div class="stat-title">{{ stat.title }}</div>
            <div class="stat-value text-primary">{{ stat.value }}</div>
            <div class="stat-desc">{{ stat.subtitle }}</div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
  .page-layout {
    display: grid;
    grid-template-columns: auto minmax(0, 1fr);
    grid-template-rows: auto;
    grid-template-areas: "side main";
    gap: 1rem;
  }

  .side-slot {
    grid-area: side;
  }

  .main-slot {
    grid-area: main;
  }

  .grid {
    display: grid;
  }
</style>
