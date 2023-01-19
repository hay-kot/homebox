<template>
  <div class="dropdown dropdown-left">
    <slot>
      <label tabindex="0" class="btn btn-circle btn-sm">
        <Icon name="mdi-qrcode" />
      </label>
    </slot>
    <div tabindex="0" class="card compact dropdown-content shadow-lg bg-base-100 rounded-box w-64">
      <div class="card-body">
        <h2 class="text-center">Page URL</h2>
        <img :src="getQRCodeUrl()" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  const api = useUserApi();

  function getQRCodeUrl(): string {
    const currentURL = window.location.href;

    return `/api/v1/qrcode?data=${encodeURIComponent(currentURL)}&access_token=${api.items.attachmentToken}`;
  }
</script>

<style lang="scss" scoped></style>
