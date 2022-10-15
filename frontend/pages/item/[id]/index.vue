<script setup lang="ts">
  import { DateDetail, Detail } from "~~/components/global/DetailsSection/types";
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

  const itemDetails = computed(() => {
    return [
      {
        name: "Description",
        text: item.value?.description,
      },
      {
        name: "Serial Number",
        text: item.value?.serialNumber,
      },
      {
        name: "Model Number",
        text: item.value?.modelNumber,
      },
      {
        name: "Manufacturer",
        text: item.value?.manufacturer,
      },
      {
        name: "Insured",
        text: item.value?.insured ? "Yes" : "No",
      },
      {
        name: "Notes",
        text: item.value?.notes,
      },
    ];
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

  const attachmentDetails = computed(() => {
    const details: Detail[] = [];

    const push = (name: string) => {
      details.push({
        name,
        text: "",
        slot: name.toLowerCase(),
      });
    };

    if (attachments.value.photos.length > 0) {
      push("Photos");
    }

    if (attachments.value.attachments.length > 0) {
      push("Attachments");
    }

    if (attachments.value.warranty.length > 0) {
      push("Warranty");
    }

    if (attachments.value.manuals.length > 0) {
      push("Manuals");
    }

    return details;
  });

  const showWarranty = computed(() => {
    if (preferences.value.showEmpty) {
      return true;
    }
    return validDate(item.value?.warrantyExpires);
  });

  const warrantyDetails = computed(() => {
    const details: (Detail | DateDetail)[] = [
      {
        name: "Lifetime Warranty",
        text: item.value?.lifetimeWarranty ? "Yes" : "No",
      },
    ];

    if (item.value?.lifetimeWarranty) {
      details.push({
        name: "Warranty Expires",
        text: "N/A",
      });
    } else {
      details.push({
        name: "Warranty Expires",
        text: item.value?.warrantyExpires,
        type: "date",
      });
    }

    details.push({
      name: "Warranty Details",
      text: item.value?.warrantyDetails || "",
    });

    return details;
  });

  const showPurchase = computed(() => {
    if (preferences.value.showEmpty) {
      return true;
    }
    return item.value?.purchaseFrom || item.value?.purchasePrice !== "0";
  });

  const purchaseDetails = computed<Array<Detail | DateDetail>>(() => {
    return [
      {
        name: "Purchased From",
        text: item.value?.purchaseFrom || "",
      },
      {
        name: "Purchase Price",
        text: item.value?.purchasePrice ? fmtCurrency(item.value.purchasePrice) : "",
      },
      {
        name: "Purchase Date",
        text: item.value.purchaseTime,
        type: "date",
      },
    ];
  });

  const showSold = computed(() => {
    if (preferences.value.showEmpty) {
      return true;
    }
    return item.value?.soldTo || item.value?.soldPrice !== "0";
  });

  const soldDetails = computed<Array<Detail | DateDetail>>(() => {
    return [
      {
        name: "Sold To",
        text: item.value?.soldTo || "",
      },
      {
        name: "Sold Price",
        text: item.value?.soldPrice ? fmtCurrency(item.value.soldPrice) : "",
      },
      {
        name: "Sold At",
        text: item.value?.soldTime || "",
        type: "date",
      },
    ];
  });

  const confirm = useConfirm();

  async function deleteItem() {
    const confirmed = await confirm.open("Are you sure you want to delete this item?");

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
        <BaseCard>
          <template #title>
            <BaseSectionHeader>
              <Icon name="mdi-package-variant" class="mr-2 -mt-1 text-base-content" />
              <span class="text-base-content">
                {{ item ? item.name : "" }}
              </span>
              <template #description>
                <p class="text-sm text-base-content font-bold pb-0 mb-0">
                  {{ item.location.name }} - Quantity {{ item.quantity }}
                </p>
                <div v-if="item.labels && item.labels.length > 0" class="flex flex-wrap gap-3 mt-3">
                  <LabelChip v-for="label in item.labels" :key="label.id" class="badge-primary" :label="label" />
                </div>
              </template>
            </BaseSectionHeader>
          </template>
          <template #title-actions>
            <div class="modal-action mt-0">
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

          <DetailsSection :details="itemDetails" />
        </BaseCard>

        <BaseCard v-if="showAttachments">
          <template #title> Attachments </template>
          <DetailsSection :details="attachmentDetails">
            <template #manuals>
              <ItemAttachmentsList
                v-if="attachments.manuals.length > 0"
                :attachments="attachments.manuals"
                :item-id="item.id"
              />
            </template>
            <template #attachments>
              <ItemAttachmentsList
                v-if="attachments.attachments.length > 0"
                :attachments="attachments.attachments"
                :item-id="item.id"
              />
            </template>
            <template #warranty>
              <ItemAttachmentsList
                v-if="attachments.warranty.length > 0"
                :attachments="attachments.warranty"
                :item-id="item.id"
              />
            </template>
            <template #photos>
              <ItemAttachmentsList
                v-if="attachments.photos.length > 0"
                :attachments="attachments.photos"
                :item-id="item.id"
              />
            </template>
          </DetailsSection>
        </BaseCard>

        <BaseCard v-if="showPurchase">
          <template #title> Purchase Details </template>
          <DetailsSection :details="purchaseDetails" />
        </BaseCard>

        <BaseCard v-if="showWarranty">
          <template #title> Warranty Details </template>
          <DetailsSection :details="warrantyDetails" />
        </BaseCard>

        <BaseCard v-if="showSold">
          <template #title> Sold Details </template>
          <DetailsSection :details="soldDetails" />
        </BaseCard>
      </div>
    </section>
  </BaseContainer>
</template>
