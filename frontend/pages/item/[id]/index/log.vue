<script setup lang="ts">
  import DatePicker from "~~/components/Form/DatePicker.vue";
  import { StatsFormat } from "~~/components/global/StatCard/types";
  import { ItemOut, MaintenanceEntry } from "~~/lib/api/types/data-contracts";

  const props = defineProps<{
    item: ItemOut;
  }>();

  const api = useUserApi();
  const toast = useNotifier();

  const { data: log, refresh: refreshLog } = useAsyncData(async () => {
    const { data } = await api.items.maintenance.getLog(props.item.id);
    return data;
  });

  const count = computed(() => {
    if (!log.value) return 0;
    return log.value.entries.length;
  });
  const stats = computed(() => {
    if (!log.value) return [];

    return [
      {
        id: "count",
        title: "Total Entries",
        value: count.value || 0,
        type: "number" as StatsFormat,
      },
      {
        id: "total",
        title: "Total Cost",
        value: log.value.costTotal || 0,
        type: "currency" as StatsFormat,
      },
      {
        id: "average",
        title: "Monthly Average",
        value: log.value.costAverage || 0,
        type: "currency" as StatsFormat,
      },
    ];
  });

  const entry = reactive({
    id: null as string | null,
    modal: false,
    name: "",
    date: new Date(),
    description: "",
    cost: "",
  });

  function newEntry() {
    entry.modal = true;
  }

  function resetEntry() {
    entry.id = null;
    entry.name = "";
    entry.date = new Date();
    entry.description = "";
    entry.cost = "";
  }

  async function createEntry() {
    if (entry.id) {
      editEntry();
      resetEntry();
      return;
    }

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
    resetEntry();
  }

  const confirm = useConfirm();

  async function deleteEntry(id: string) {
    const result = await confirm.open("Are you sure you want to delete this entry?");
    if (result.isCanceled) {
      return;
    }

    const { error } = await api.items.maintenance.delete(props.item.id, id);

    if (error) {
      toast.error("Failed to delete entry");
      return;
    }
    refreshLog();
  }

  function openEditDialog(e: MaintenanceEntry) {
    entry.modal = true;
    entry.id = e.id;
    entry.name = e.name;
    entry.date = new Date(e.date);
    entry.description = e.description;
    entry.cost = e.cost;
  }

  async function editEntry() {
    if (!entry.id) {
      return;
    }

    const { error } = await api.items.maintenance.update(props.item.id, entry.id, {
      name: entry.name,
      date: entry.date,
      description: entry.description,
      cost: entry.cost,
    });

    if (error) {
      toast.error("Failed to update entry");
      return;
    }

    entry.modal = false;

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
          <BaseButton type="submit" class="ml-2 mt-2">
            <template #icon>
              <Icon name="mdi-post" />
            </template>
            {{ entry.id ? "Update" : "Create" }}
          </BaseButton>
        </div>
      </form>
    </BaseModal>

    <section class="space-y-6">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <StatCard
          v-for="stat in stats"
          :key="stat.id"
          class="stats block shadow-xl border-l-primary"
          :title="stat.title"
          :value="stat.value"
          :type="stat.type"
        />
      </div>
      <div class="flex">
        <BaseButton class="ml-auto" size="sm" @click="newEntry()">
          <template #icon>
            <Icon name="mdi-post" />
          </template>
          Log Maintenance
        </BaseButton>
      </div>
      <div class="container space-y-6">
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
          <div class="flex justify-end p-4 gap-1">
            <BaseButton size="sm" @click="openEditDialog(e)">
              <template #icon>
                <Icon name="mdi-edit" />
              </template>
              Edit
            </BaseButton>
            <BaseButton size="sm" @click="deleteEntry(e.id)">
              <template #icon>
                <Icon name="mdi-delete" />
              </template>
              Delete
            </BaseButton>
          </div>
        </BaseCard>
        <div class="hidden first:block">
          <button
            type="button"
            class="relative block w-full rounded-lg border-2 border-dashed border-base-content p-12 text-center"
            @click="newEntry()"
          >
            <Icon name="mdi-wrench-clock" class="h-16 w-16"></Icon>
            <span class="mt-2 block text-sm font-medium text-gray-900"> Create Your First Entry </span>
          </button>
        </div>
      </div>
    </section>
  </div>
</template>
