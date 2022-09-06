<template>
  <BaseContainer class="pb-8">
    <section class="px-3">
      <BaseSectionHeader v-if="item" dark class="mb-5">
        <Icon name="mdi-package-variant" class="-mt-1" />
        {{ item.name }}
        <template #description>
          <div class="flex flex-wrap gap-3 mt-3">
            <LabelChip class="badge-primary" v-for="label in item.labels" :label="label"></LabelChip>
          </div>
        </template>
      </BaseSectionHeader>
      <div class="flex justify-between items-center">
        <div class="form-control">
          <label class="label cursor-pointer">
            <input type="checkbox" v-model.checked="preferences.showEmpty" class="checkbox" />
            <span class="label-text ml-4"> Show Empty </span>
          </label>
        </div>
        <BaseButton size="sm" :to="`/item/${itemId}/edit`">
          <template #icon>
            <Icon name="mdi-pencil" />
          </template>
          Edit
        </BaseButton>
      </div>
      <div class="grid grid-cols-1 gap-3">
        <BaseDetails :details="itemSummary">
          <template #title> Item Summary </template>
          <template #Attachments>
            <ul role="list" class="divide-y divide-gray-400 rounded-md border border-gray-400">
              <li class="flex items-center justify-between py-3 pl-3 pr-4 text-sm">
                <div class="flex w-0 flex-1 items-center">
                  <Icon name="mdi-paperclip" class="h-5 w-5 flex-shrink-0 text-gray-400" aria-hidden="true" />
                  <span class="ml-2 w-0 flex-1 truncate">User Guide.pdf</span>
                </div>
                <div class="ml-4 flex-shrink-0">
                  <a href="#" class="font-medium">Download</a>
                </div>
              </li>
              <li class="flex items-center justify-between py-3 pl-3 pr-4 text-sm">
                <div class="flex w-0 flex-1 items-center">
                  <Icon name="mdi-paperclip" class="h-5 w-5 flex-shrink-0 text-gray-400" aria-hidden="true" />
                  <span class="ml-2 w-0 flex-1 truncate">Purchase Receipts.pdf</span>
                </div>
                <div class="ml-4 flex-shrink-0">
                  <a href="#" class="font-medium">Download</a>
                </div>
              </li>
            </ul>
          </template>
        </BaseDetails>
        <BaseDetails :details="purchaseDetails" v-if="showPurchase">
          <template #title> Purchase Details </template>
        </BaseDetails>
        <BaseDetails :details="soldDetails" v-if="showSold">
          <template #title> Sold Details </template>
        </BaseDetails>
      </div>
    </section>
  </BaseContainer>
</template>

<script setup lang="ts">
  definePageMeta({
    layout: 'home',
  });

  const route = useRoute();
  const api = useUserApi();
  const toast = useNotifier();

  const itemId = computed<string>(() => route.params.id as string);
  const preferences = useViewPreferences();

  const { data: item } = useAsyncData(async () => {
    const { data, error } = await api.items.get(itemId.value);
    if (error) {
      toast.error('Failed to load item');
      navigateTo('/home');
      return;
    }
    return data;
  });

  const itemSummary = computed(() => {
    return {
      Name: item.value?.name || '',
      Description: item.value?.description || '',
      'Serial Number': item.value?.serialNumber || '',
      'Model Number': item.value?.modelNumber || '',
      Manufacturer: item.value?.manufacturer || '',
      Notes: item.value?.notes || '',
      Attachments: '', // TODO: Attachments
    };
  });

  const showPurchase = computed(() => {
    if (preferences.value.showEmpty) {
      return true;
    }
    return item.value?.purchaseFrom || item.value?.purchasePrice;
  });

  const purchaseDetails = computed(() => {
    return {
      'Purchased From': item.value?.purchaseFrom || '',
      'Purchased Price': item.value?.purchasePrice || '',
      'Purchased At': item.value?.purchaseTime || '',
    };
  });

  const showSold = computed(() => {
    if (preferences.value.showEmpty) {
      return true;
    }

    return item.value?.soldTo || item.value?.soldPrice;
  });

  const soldDetails = computed(() => {
    return {
      'Sold To': item.value?.soldTo || '',
      'Sold Price': item.value?.soldPrice || '',
      'Sold At': item.value?.soldTime || '',
    };
  });
</script>

<style scoped></style>
