<template>
  <div class="force-above fixed top-2 right-2 w-[300px]">
    <TransitionGroup name="notify" tag="div">
      <div
        v-for="(notify, index) in notifications.slice(0, 4)"
        :key="notify.id"
        class="my-2 w-[300px] rounded-md p-3 text-sm text-white"
        :class="{
          'bg-primary': notify.type === 'info',
          'bg-red-600': notify.type === 'error',
          'bg-green-600': notify.type === 'success',
        }"
        @click="dropNotification(index)"
      >
        <div class="flex gap-1" style="background-color: #008933;">
          <template v-if="notify.type == 'success'">
            <Icon name="heroicons-check" class="h-5 w-5" />
          </template>
          <template v-if="notify.type == 'info'">
            <Icon name="heroicons-information-circle" class="h-5 w-5" />
          </template>

          <template v-if="notify.type == 'error'">
            <Icon name="heroicons-bell-alert" class="h-5 w-5" />
          </template>
          {{ notify.message }}
        </div>
      </div>
    </TransitionGroup>
  </div>
</template>

<script setup lang="ts">
  import { useNotifications } from "@/composables/use-notifier";

  const { notifications, dropNotification } = useNotifications();
</script>

<style scoped>
  .force-above {
    z-index: 9999;
  }

  .notify-move,
  .notify-enter-active,
  .notify-leave-active {
    transition: all 0.5s ease;
  }
  .notify-enter-from,
  .notify-leave-to {
    opacity: 0;
    transform: translateY(-30px);
  }
  .notify-leave-active {
    position: absolute;
    transform: translateY(30px);
  }
</style>
