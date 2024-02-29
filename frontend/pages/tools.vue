<template>
  <div>
    <AppImportDialog v-model="modals.import" />
    <BaseContainer class="flex flex-col gap-4 mb-6">
      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-file-chart" class="mr-2 -mt-1" />
            <span> {{ $t("tools.report.title") }} </span>
            <template #description> {{ $t("tools.report.desp") }} </template>
          </BaseSectionHeader>
        </template>
        <div class="border-t px-6 pb-3 border-gray-300 divide-gray-300 divide-y">
          <DetailAction @action="navigateTo('/reports/label-generator')">
            <template #title> {{ $t("tools.report.asset.title") }} </template>
            {{ $t("tools.report.asset.desp") }}
            <template #button>
              {{ $t("tools.report.asset.button") }}
              <Icon name="mdi-arrow-right" class="ml-2" />
            </template>
          </DetailAction>
          <DetailAction @action="getBillOfMaterials()">
            <template #title>{{ $t("tools.report.bom.title") }}</template>
            {{ $t("tools.report.bom.desp") }}
            <template #button> {{ $t("tools.report.bom.button") }} </template>
          </DetailAction>
        </div>
      </BaseCard>
      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-database" class="mr-2 -mt-1" />
            <span> {{ $t("tools.import_export.title") }} </span>
            <template #description>
              {{ $t("tools.import_export.desp") }}
            </template>
          </BaseSectionHeader>
        </template>
        <div class="border-t px-6 pb-3 border-gray-300 divide-gray-300 divide-y">
          <DetailAction @action="modals.import = true">
            <template #title> {{ $t("tools.import_export.import.title") }} </template>
            {{ $t("tools.import_export.import.desp") }}
          </DetailAction>
          <DetailAction @action="getExportTSV()">
            <template #title>{{ $t("tools.import_export.export.title") }}</template>
            {{ $t("tools.import_export.export.desp") }}
          </DetailAction>
        </div>
      </BaseCard>
      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-warning" class="mr-2 -mt-1" />
            <span> {{ $t("tools.inventory.title") }} </span>
            <template #description>
              {{ $t("tools.inventory.desp") }}
            </template>
          </BaseSectionHeader>
        </template>
        <div class="border-t px-6 pb-3 border-gray-300 divide-gray-300 divide-y">
          <DetailAction @action="ensureAssetIDs">
            <template #title> {{ $t("tools.inventory.ensureAssetIDs.title") }} </template>
            {{ $t("tools.inventory.ensureAssetIDs.desp") }}
          </DetailAction>
          <DetailAction @action="ensureImportRefs">
            <template #title> {{ $t("tools.inventory.ensureImportRefs.title") }} </template>
            {{ $t("tools.inventory.ensureImportRefs.desp") }}
          </DetailAction>
          <DetailAction @action="resetItemDateTimes">
            <template #title> {{ $t("tools.inventory.resetItemDateTimes.title") }}</template>
            {{ $t("tools.inventory.resetItemDateTimes.desp") }}
            <a class="link" href="https://github.com/hay-kot/homebox/issues/236" target="_blank">
              See Github Issue #236 for more details.
            </a>
          </DetailAction>
          <DetailAction @action="setPrimaryPhotos">
            <template #title> {{ $t("tools.inventory.setPrimaryPhotos.title") }}</template>
            {{ $t("tools.inventory.setPrimaryPhotos.desp") }}
             <a class="link" href="https://github.com/hay-kot/homebox/pull/576">See GitHub PR #576</a>
          </DetailAction>
        </div>
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

  function getBillOfMaterials() {
    const url = api.reports.billOfMaterialsURL();
    window.open(url, "_blank");
  }

  function getExportTSV() {
    const url = api.items.exportURL();
    window.open(url, "_blank");
  }

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

  async function ensureImportRefs() {
    const { isCanceled } = await confirm.open(
      "Are you sure you want to ensure all assets have an import_ref? This can take a while and cannot be undone."
    );

    if (isCanceled) {
      return;
    }

    const result = await api.actions.ensureImportRefs();

    if (result.error) {
      notify.error("Failed to ensure import refs.");
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

  async function setPrimaryPhotos() {
    const { isCanceled } = await confirm.open(
      "Are you sure you want to set primary photos? This can take a while and cannot be undone."
    );

    if (isCanceled) {
      return;
    }

    const result = await api.actions.setPrimaryPhotos();

    if (result.error) {
      notify.error("Failed to set primary photos.");
      return;
    }

    notify.success(`${result.data.completed} assets have been updated.`);
  }
</script>

<style scoped></style>
