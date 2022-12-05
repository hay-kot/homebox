<script setup lang="ts">
  import { AnyDetail, Detail, Details } from "~~/components/global/DetailsSection/types";
  import { ItemAttachment } from "~~/lib/api/types/data-contracts";

  definePageMeta({
    middleware: ["auth"],
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
    attachments: ItemAttachment[];
    warranty: ItemAttachment[];
    manuals: ItemAttachment[];
    receipts: ItemAttachment[];
  };

  type Photo = {
    src: string;
  };

  const photos = computed<Photo[]>(() => {
    return (
      item.value?.attachments.reduce((acc, cur) => {
        if (cur.type === "photo") {
          acc.push({
            src: api.authURL(`/items/${item.value.id}/attachments/${cur.id}`),
          });
        }
        return acc;
      }, [] as Photo[]) || []
    );
  });

  const attachments = computed<FilteredAttachments>(() => {
    if (!item.value) {
      return {
        attachments: [],
        manuals: [],
        warranty: [],
        receipts: [],
      };
    }

    return item.value.attachments.reduce(
      (acc, attachment) => {
        if (attachment.type === "photo") {
          return acc;
        }
        if (attachment.type === "warranty") {
          acc.warranty.push(attachment);
        } else if (attachment.type === "manual") {
          acc.manuals.push(attachment);
        } else if (attachment.type === "receipt") {
          acc.receipts.push(attachment);
        } else {
          acc.attachments.push(attachment);
        }
        return acc;
      },
      {
        attachments: [] as ItemAttachment[],
        warranty: [] as ItemAttachment[],
        manuals: [] as ItemAttachment[],
        receipts: [] as ItemAttachment[],
      }
    );
  });

  const assetID = computed<Details>(() => {
    if (item.value?.assetId === "000-000") {
      return [];
    }

    return [
      {
        name: "Asset ID",
        text: item.value?.assetId,
      },
    ];
  });

  const itemDetails = computed<Details>(() => {
    return [
      {
        name: "Description",
        type: "markdown",
        text: item.value?.description,
      },
      {
        name: "Quantity",
        text: item.value?.quantity,
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
        type: "markdown",
        text: item.value?.notes,
      },
      ...assetID.value,
      ...item.value.fields.map(field => {
        /**
         * Support Special URL Syntax
         */
        const url = maybeUrl(field.textValue);
        if (url.isUrl) {
          return {
            type: "link",
            name: field.name,
            text: url.text,
            href: url.url,
          } as AnyDetail;
        }

        return {
          name: field.name,
          text: field.textValue,
        };
      }),
    ];
  });

  const showAttachments = computed(() => {
    if (preferences.value?.showEmpty) {
      return true;
    }

    return (
      attachments.value.attachments.length > 0 ||
      attachments.value.warranty.length > 0 ||
      attachments.value.manuals.length > 0 ||
      attachments.value.receipts.length > 0
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

    if (attachments.value.attachments.length > 0) {
      push("Attachments");
    }

    if (attachments.value.warranty.length > 0) {
      push("Warranty");
    }

    if (attachments.value.manuals.length > 0) {
      push("Manuals");
    }

    if (attachments.value.receipts.length > 0) {
      push("Receipts");
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
    const details: Details = [
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
      type: "markdown",
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

  const purchaseDetails = computed<Details>(() => {
    return [
      {
        name: "Purchased From",
        text: item.value?.purchaseFrom || "",
      },
      {
        name: "Purchase Price",
        text: item.value?.purchasePrice || "",
        type: "currency",
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

  const soldDetails = computed<Details>(() => {
    return [
      {
        name: "Sold To",
        text: item.value?.soldTo || "",
      },
      {
        name: "Sold Price",
        text: item.value?.soldPrice || "",
        type: "currency",
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

  const refDialog = ref<HTMLDialogElement>();
  const dialoged = reactive({
    src: "",
  });

  function openDialog(img: Photo) {
    refDialog.value.showModal();
    dialoged.src = img.src;
  }

  function closeDialog() {
    refDialog.value.close();
  }

  const refDialogBody = ref<HTMLDivElement>();
  onClickOutside(refDialogBody, () => {
    closeDialog();
  });
</script>

<template>
  <BaseContainer v-if="item" class="pb-8">
    <dialog ref="refDialog" class="z-[999] fixed bg-transparent">
      <div ref="refDialogBody" class="relative">
        <div class="absolute right-0 -mt-3 -mr-3 sm:-mt-4 sm:-mr-4 space-x-1">
          <a class="btn btn-sm sm:btn-md btn-primary btn-circle" :href="dialoged.src" download>
            <Icon class="h-5 w-5" name="mdi-download" />
          </a>
          <button class="btn btn-sm sm:btn-md btn-primary btn-circle" @click="closeDialog()">
            <Icon class="h-5 w-5" name="mdi-close" />
          </button>
        </div>

        <img class="max-w-[80vw] max-h-[80vh]" :src="dialoged.src" />
      </div>
    </dialog>
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
              <div v-if="item.parent" class="text-sm breadcrumbs pb-0">
                <ul class="text-base-content/70">
                  <li>
                    <NuxtLink :to="`/item/${item.parent.id}`"> {{ item.parent.name }}</NuxtLink>
                  </li>
                  <li>{{ item.name }}</li>
                </ul>
              </div>
              <template #description>
                <div class="flex flex-wrap gap-2 mt-3">
                  <NuxtLink ref="badge" class="badge p-3" :to="`/location/${item.location.id}`">
                    <Icon name="heroicons-map-pin" class="mr-2 swap-on"></Icon>
                    {{ item.location.name }}
                  </NuxtLink>
                  <template v-if="item.labels && item.labels.length > 0">
                    <LabelChip v-for="label in item.labels" :key="label.id" class="badge-primary" :label="label" />
                  </template>
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

        <BaseCard v-if="photos && photos.length > 0">
          <template #title> Photos </template>
          <div
            class="container border-t border-gray-300 p-4 flex flex-wrap gap-2 mx-auto max-h-[500px] overflow-y-scroll scroll-bg"
          >
            <button v-for="(img, i) in photos" :key="i" @click="openDialog(img)">
              <img class="rounded max-h-[200px]" :src="img.src" />
            </button>
          </div>
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
            <template #receipts>
              <ItemAttachmentsList
                v-if="attachments.receipts.length > 0"
                :attachments="attachments.receipts"
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

    <section class="my-6 px-3">
      <BaseSectionHeader v-if="item && item.children && item.children.length > 0"> Child Items </BaseSectionHeader>
      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <ItemCard v-for="child in item.children" :key="child.id" :item="child" />
      </div>
    </section>
  </BaseContainer>
</template>

<style lang="css" scoped>
  /* Style dialog background */
  dialog::backdrop {
    background: rgba(0, 0, 0, 0.5);
  }

  .scroll-bg::-webkit-scrollbar {
    width: 0.5rem;
  }

  .scroll-bg::-webkit-scrollbar-thumb {
    border-radius: 0.25rem;
    @apply bg-base-300;
  }
</style>
