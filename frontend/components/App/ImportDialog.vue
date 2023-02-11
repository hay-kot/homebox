<template>
  <BaseModal v-model="dialog">
    <template #title> Import CSV File </template>
    <p>
      Import a CSV file containing your items, labels, and locations. See documentation for more information on the
      required format.
    </p>

    <form @submit.prevent="submitCsvFile">
      <div class="flex flex-col gap-2 py-6">
        <input ref="importRef" type="file" class="hidden" accept=".csv,.tsv" @change="setFile" />

        <BaseButton type="button" @click="uploadCsv">
          <Icon class="h-5 w-5 mr-2" name="mdi-upload" />
          Upload
        </BaseButton>
        <p class="text-center pt-4 -mb-5">
          {{ importCsv?.name }}
        </p>
      </div>

      <div class="modal-action">
        <BaseButton type="submit" :disabled="!importCsv"> Submit </BaseButton>
      </div>
    </form>
  </BaseModal>
</template>

<script setup lang="ts">
  type Props = {
    modelValue: boolean;
  };

  const props = withDefaults(defineProps<Props>(), {
    modelValue: false,
  });

  const emit = defineEmits(["update:modelValue"]);

  const dialog = useVModel(props, "modelValue", emit);

  const api = useUserApi();
  const toast = useNotifier();

  const importCsv = ref<File | null>(null);
  const importLoading = ref(false);
  const importRef = ref<HTMLInputElement>();
  whenever(
    () => !dialog.value,
    () => {
      importCsv.value = null;
    }
  );

  function setFile(e: Event) {
    const result = e.target as HTMLInputElement;
    if (!result.files || result.files.length === 0) {
      return;
    }

    importCsv.value = result.files[0];
  }

  function uploadCsv() {
    importRef.value?.click();
  }

  const eventBus = useEventBus();

  async function submitCsvFile() {
    if (!importCsv.value) {
      toast.error("Please select a file to import.");
      return;
    }

    importLoading.value = true;

    const { error } = await api.items.import(importCsv.value);

    if (error) {
      toast.error("Import failed. Please try again later.");
    }

    // Reset
    dialog.value = false;
    importLoading.value = false;
    importCsv.value = null;

    if (importRef.value) {
      importRef.value.value = "";
    }

    eventBus.emit(EventTypes.InvalidStores);

    toast.success("Import successful!");
  }
</script>
