<template>
  <ul role="list" class="divide-y divide-gray-400 rounded-md border border-gray-400">
    <li
      v-for="attachment in attachments"
      :key="attachment.id"
      class="flex items-center justify-between py-3 pl-3 pr-4 text-sm"
    >
      <div class="flex w-0 flex-1 items-center">
        <Icon name="mdi-paperclip" class="h-5 w-5 flex-shrink-0 text-gray-400" aria-hidden="true" />
        <span class="ml-2 w-0 flex-1 truncate"> {{ attachment.document.title }}</span>
      </div>
      <div class="ml-4 flex-shrink-0">
        <button class="font-medium" @click="getAttachmentUrl(attachment)">Download</button>
      </div>
    </li>
  </ul>
</template>

<script setup lang="ts">
  import { ItemAttachment } from "~~/lib/api/types/data-contracts";

  const props = defineProps({
    attachments: {
      type: Object as () => ItemAttachment[],
      required: true,
    },
    itemId: {
      type: String,
      required: true,
    },
  });

  const api = useUserApi();
  const toast = useNotifier();
  async function getAttachmentUrl(attachment: ItemAttachment) {
    const url = await api.items.getAttachmentUrl(props.itemId, attachment.id);

    if (!url) {
      toast.error("Failed to get attachment url");
      return;
    }

    if (!document) {
      window.open(url, "_blank");
      return;
    }

    const link = document.createElement("a");
    link.href = url;
    link.target = "_blank";
    link.setAttribute("download", attachment.document.title);
    link.click();
  }
</script>

<style scoped></style>
