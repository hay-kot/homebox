<template>
  <div class="dropdown dropdown-left">
    <slot>
      <label tabindex="0" class="btn btn-circle btn-sm join-item">
        <Icon name="mdi-qrcode" />
      </label>
    </slot>
    <div tabindex="0" class="card compact dropdown-content shadow-lg bg-base-100 rounded-box w-[42rem]">
      <div class="flex">
        <div class="card-body basis-1/2">
          <h2 class="text-center">Page URL</h2>
          <img :src="getQRCodeUrl()" />
        </div>
        <div class="divider divider-horizontal">OR</div>
        <div class="card-body basis-1/2 flex flex-col items-center">
          <div class="flex flex-row w-full space-x-1">
            <h2 class="text-center">Page URL</h2>
            <input class="input input-sm input-bordered w-full max-w-xs" type="text">
          </div>


          <button class="btn btn-primary">
            Generate PDF
          </button>

        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
  import { route } from "../../lib/api/base";
  function getQRCodeUrl(): string {
    const currentURL = window.location.href;

    return route(`/qrcode`, { data: encodeURIComponent(currentURL) });
  }
</script>

<style lang="scss" scoped></style>
