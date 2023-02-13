<template>
  <div>
    <AppImportDialog v-model="modals.import" />
    <BaseContainer class="flex flex-col gap-4 mb-6">
      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-database" class="mr-2 -mt-1" />
            <span> Import / Export </span>
            <template #description>
              Import and export your inventory to and from a CSV file. This is useful for migrating your inventory to a
              new instance of Homebox.
            </template>
          </BaseSectionHeader>
          <div class="border-t border-gray-300 divide-gray-300 divide-y">
            <DetailAction @action="modals.import = true">
              <template #title>Import Inventory</template>
              Imports the standard CSV format for Homebox. This will <b>not</b> overwrite any existing items in your
              inventory. It will only add new items.
            </DetailAction>
            <!-- <DetailAction>
              <template #title>Export Inventory</template>
              Exports the standard CSV format for Homebox. This will export all items in your inventory.
            </DetailAction> -->
          </div>
        </template>
      </BaseCard>
      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-file-chart" class="mr-2 -mt-1" />
            <span> Reports </span>
            <template #description> Generate different reports for your inventory. </template>
          </BaseSectionHeader>
          <div class="border-t border-gray-300 divide-gray-300 divide-y">
            <DetailAction @action="navigateTo('/reports/label-generator')">
              <template #title>Asset ID Labels</template>
              Generates a printable PDF of labels for a range of Asset ID. These are not specific to your invetory so
              your are able to print labels ahead of time and apply them to your inventory when you receive them.
              <template #button>
                Label Generator
                <Icon name="mdi-arrow-right" class="ml-2" />
              </template>
            </DetailAction>
          </div>
        </template>
      </BaseCard>
      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-warning" class="mr-2 -mt-1" />
            <span> Inventory Actions </span>
            <template #description>
              Apply Actions to your inventory in bulk. These are irreversible actions. <b>Be careful</b>
            </template>
          </BaseSectionHeader>
          <div class="border-t border-gray-300 divide-gray-300 divide-y">
            <DetailAction @action="ensureAssetIDs">
              <template #title>Ensure Asset IDs</template>
              Ensures that all items in your inventory have a valid asset_id field. This is done by finding the highest
              current asset_id field in the database and applying the next value to each item that has an unset asset_id
              field. This is done in order of the created_at field.
            </DetailAction>
            <DetailAction @click="resetItemDateTimes">
              <template #title> Zero Item Date Times</template>
              Resets the time value for all date time fields in your inventory to the beginning of the date. This is to
              fix a bug that was introduced early on in the development of the site that caused the time value to be
              stored with the time which caused issues with date fields displaying accurate values.
              <a class="link" href="https://github.com/hay-kot/homebox/issues/236" target="_blank">
                See Github Issue #236 for more details
              </a>
            </DetailAction>
          </div>
        </template>
      </BaseCard>
    </BaseContainer>
  </div>
</template>

<script setup lang="ts">
  definePageMeta({
    middleware: ["auth"],
  });
  useHead({
    title: "Homebox | Profile",
  });

  const modals = ref({
    item: false,
    location: false,
    label: false,
    import: false,
  });

  const api = useUserApi();
  const confirm = useConfirm();
  const notify = useNotifier();

  async function ensureAssetIDs() {
    const { isCanceled } = await confirm.open(
      "Are you sure you want to ensure all assets have an ID? This can take a while and cannot be undone."
    );

    if (isCanceled) {
      return;
    }

    const result = await api.actions.ensureAssetIDs();

    if (result.error) {
      notify.error("Failed to ensure asset IDs.");
      return;
    }

    notify.success(`${result.data.completed} assets have been updated.`);
  }

  async function resetItemDateTimes() {
    const { isCanceled } = await confirm.open(
      "Are you sure you want to reset all date and time values? This can take a while and cannot be undone."
    );

    if (isCanceled) {
      return;
    }

    const result = await api.actions.resetItemDateTimes();

    if (result.error) {
      notify.error("Failed to reset date and time values.");
      return;
    }

    notify.success(`${result.data.completed} assets have been updated.`);
  }
</script>

<style scoped></style>
