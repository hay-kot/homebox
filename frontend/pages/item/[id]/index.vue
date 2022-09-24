<script setup lang="ts">
  import { ItemAttachment } from "~~/lib/api/types/data-contracts";

  definePageMeta({
    layout: "home",
  });

  const route = useRoute();
  const api = useUserApi();
  const toast = useNotifier();

  const itemId = computed<string>(() => route.params.id as string);
  const preferences = useViewPreferences();

  const { data: item, refresh } = useAsyncData(itemId.value, async () => {
    const { data, error } = await api.items.get(itemId.value);
    if (error) {
      toast.error("Failed to load item");
      navigateTo("/home");
      return;
    }
    return data;
  });
  onMounted(() => {
    refresh();
  });

  type FilteredAttachments = {
    photos: ItemAttachment[];
    attachments: ItemAttachment[];
    warranty: ItemAttachment[];
    manuals: ItemAttachment[];
  };

  const attachments = computed<FilteredAttachments>(() => {
    if (!item.value) {
      return {
        photos: [],
        attachments: [],
        manuals: [],
        warranty: [],
      };
    }

    return item.value.attachments.reduce(
      (acc, attachment) => {
        if (attachment.type === "photo") {
          acc.photos.push(attachment);
        } else if (attachment.type === "warranty") {
          acc.warranty.push(attachment);
        } else if (attachment.type === "manual") {
          acc.manuals.push(attachment);
        } else {
          acc.attachments.push(attachment);
        }
        return acc;
      },
      {
        photos: [] as ItemAttachment[],
        attachments: [] as ItemAttachment[],
        warranty: [] as ItemAttachment[],
        manuals: [] as ItemAttachment[],
      }
    );
  });

  const itemSummary = computed(() => {
    return {
      Description: item.value?.description || "",
      "Serial Number": item.value?.serialNumber || "",
      "Model Number": item.value?.modelNumber || "",
      Manufacturer: item.value?.manufacturer || "",
      Notes: item.value?.notes || "",
      Insured: item.value?.insured ? "Yes" : "No",
    };
  });

  const showAttachments = computed(() => {
    if (preferences.value?.showEmpty) {
      return true;
    }

    return (
      attachments.value.photos.length > 0 ||
      attachments.value.attachments.length > 0 ||
      attachments.value.warranty.length > 0 ||
      attachments.value.manuals.length > 0
    );
  });

  const itemAttachments = computed(() => {
    const val: Record<string, string> = {};

    if (preferences.value.showEmpty) {
      return {
        Photos: "",
        Manuals: "",
        Warranty: "",
        Attachments: "",
      };
    }

    if (attachments.value.photos.length > 0) {
      val.Photos = "";
    }

    if (attachments.value.manuals.length > 0) {
      val.Manuals = "";
    }

    if (attachments.value.warranty.length > 0) {
      val.Warranty = "";
    }

    if (attachments.value.attachments.length > 0) {
      val.Attachments = "";
    }

    return val;
  });

  const showWarranty = computed(() => {
    if (preferences.value.showEmpty) {
      return true;
    }
    return validDate(item.value?.warrantyExpires);
  });

  const warrantyDetails = computed(() => {
    const payload = {
      "Lifetime Warranty": item.value?.lifetimeWarranty ? "Yes" : "No",
    };

    if (showWarranty.value) {
      payload["Warranty Expires"] = item.value?.warrantyExpires || "";
    }

    payload["Warranty Details"] = item.value?.warrantyDetails || "";

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
      "Purchased From": item.value?.purchaseFrom || "",
      "Purchased Price": item.value?.purchasePrice ? fmtCurrency(item.value.purchasePrice) : "",
      "Purchased At": item.value?.purchaseTime || "",
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
      "Sold To": item.value?.soldTo || "",
      "Sold Price": item.value?.soldPrice ? fmtCurrency(item.value.soldPrice) : "",
      "Sold At": item.value?.soldTime || "",
    };
  });

  const confirm = useConfirm();

  async function deleteItem() {
    const confirmed = await confirm.reveal("Are you sure you want to delete this item?");

    if (!confirmed.data) {
      return;
    }

    const { error } = await api.items.delete(itemId.value);
    if (error) {
      toast.error("Failed to delete item");
      return;
    }
    toast.success("Item deleted");
    navigateTo("/home");
  }
</script>

<template>
  <BaseContainer v-if="item" class="pb-8">
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
              <p class="text-sm text-gray-600 font-bold pb-0 mb-0">
                {{ item.location.name }} - Quantity {{ item.quantity }}
              </p>
              <template #after>
                <div v-if="item.labels && item.labels.length > 0" class="flex flex-wrap gap-3 mt-3">
                  <LabelChip v-for="label in item.labels" :key="label.id" class="badge-primary" :label="label" />
                </div>
                <div class="modal-action mt-3">
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
        </BaseDetails>
        <BaseDetails v-if="showAttachments" :details="itemAttachments">
          <template #title> Attachments </template>
          <template #Manuals>
            <ItemAttachmentsList
              v-if="attachments.manuals.length > 0"
              :attachments="attachments.manuals"
              :item-id="item.id"
            />
          </template>
          <template #Attachments>
            <ItemAttachmentsList
              v-if="attachments.attachments.length > 0"
              :attachments="attachments.attachments"
              :item-id="item.id"
            />
          </template>
          <template #Warranty>
            <ItemAttachmentsList
              v-if="attachments.warranty.length > 0"
              :attachments="attachments.warranty"
              :item-id="item.id"
            />
          </template>
          <template #Photos>
            <ItemAttachmentsList
              v-if="attachments.photos.length > 0"
              :attachments="attachments.photos"
              :item-id="item.id"
            />
          </template>
        </BaseDetails>
        <BaseDetails v-if="showPurchase" :details="purchaseDetails">
          <template #title> Purchase Details </template>
          <template #PurchasedAt>
            <DateTime :date="item.purchaseTime" />
          </template>
        </BaseDetails>
        <BaseDetails v-if="showWarranty" :details="warrantyDetails">
          <template #title> Warranty </template>
          <template #WarrantyExpires>
            <DateTime :date="item.warrantyExpires" />
          </template>
        </BaseDetails>
        <BaseDetails v-if="showSold" :details="soldDetails">
          <template #title> Sold </template>
          <template #SoldAt>
            <DateTime :date="item.soldTime" />
          </template>
        </BaseDetails>
      </div>
    </section>
  </BaseContainer>
</template>
