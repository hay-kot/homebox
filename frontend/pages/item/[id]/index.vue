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
      Description: item.value?.description || '',
      'Serial Number': item.value?.serialNumber || '',
      'Model Number': item.value?.modelNumber || '',
      Manufacturer: item.value?.manufacturer || '',
      Notes: item.value?.notes || '',
      Attachments: '', // TODO: Attachments
    };
  });

  const showWarranty = computed(() => {
    if (preferences.value.showEmpty) {
      return true;
    }
    return item.value?.warrantyExpires !== undefined;
  });

  const warrantyDetails = computed(() => {
    const payload = {};

    if (item.value.lifetimeWarranty) {
      payload['Lifetime Warranty'] = 'Yes';
    } else {
      payload['Warranty Expires'] = item.value?.warrantyExpires || '';
    }

    payload['Warranty Details'] = item.value?.warrantyDetails || '';

    return payload;
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

  const confirm = useConfirm();

  async function deleteItem() {
    const confirmed = await confirm.reveal('Are you sure you want to delete this item?');

    if (!confirmed.data) {
      return;
    }

    const { error } = await api.items.delete(itemId.value);
    if (error) {
      toast.error('Failed to delete item');
      return;
    }
    toast.success('Item deleted');
    navigateTo('/home');
  }
</script>

<template>
  <BaseContainer class="pb-8">
    <section class="px-3">
      <div class="flex justify-between items-center">
        <div class="form-control"></div>
      </div>
      <div class="grid grid-cols-1 gap-3">
        <BaseDetails :details="itemSummary">
          <template #title>
            <BaseSectionHeader v-if="item" class="pb-0">
              <Icon name="mdi-package-variant" class="-mt-1 mr-2 text-gray-600" />
              <span class="text-gray-600">
                {{ item.name }}
              </span>
              <template #after>
                <div class="flex flex-wrap gap-3 mt-3">
                  <LabelChip v-for="label in item.labels" :key="label.id" class="badge-primary" :label="label" />
                </div>
                <div class="modal-action">
                  <label class="label cursor-pointer mr-auto">
                    <input v-model="preferences.showEmpty" type="checkbox" class="toggle toggle-primary" />
                    <span class="label-text ml-4"> Show Empty </span>
                  </label>
                  <BaseButton size="sm" :to="`/item/${itemId}/edit`">
                    <template #icon>
                      <Icon name="mdi-pencil" />
                    </template>
                    Edit
                  </BaseButton>
                  <BaseButton size="sm" @click="deleteItem">
                    <template #icon>
                      <Icon name="mdi-delete" />
                    </template>
                    Delete
                  </BaseButton>
                </div>
              </template>
            </BaseSectionHeader>
          </template>
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
        <BaseDetails v-if="showPurchase" :details="purchaseDetails">
          <template #title> Purchase Details </template>
        </BaseDetails>
        <BaseDetails v-if="showWarranty" :details="warrantyDetails">
          <template #title> Warranty </template>
        </BaseDetails>
        <BaseDetails v-if="showSold" :details="soldDetails">
          <template #title> Sold </template>
        </BaseDetails>
      </div>
    </section>
  </BaseContainer>
</template>
