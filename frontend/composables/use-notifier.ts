import { useId } from "./use-ids";

interface Notification {
  id: string;
  message: string;
  type: "success" | "error" | "info";
}

const notifications = ref<Notification[]>([]);

function addNotification(notification: Notification) {
  notifications.value.unshift(notification);

  if (notifications.value.length > 4) {
    notifications.value.pop();
  } else {
    setTimeout(() => {
      // Remove notification with ID
      notifications.value = notifications.value.filter(n => n.id !== notification.id);
    }, 5000);
  }
}

export function useNotifications() {
  return {
    notifications,
    dropNotification: (idx: number) => notifications.value.splice(idx, 1),
  };
}

export function useNotifier() {
  return {
    success: (message: string) => {
      addNotification({
        id: useId(),
        message,
        type: "success",
      });
    },
    error: (message: string) => {
      addNotification({
        id: useId(),
        message,
        type: "error",
      });
    },
    info: (message: string) => {
      addNotification({
        id: useId(),
        message,
        type: "info",
      });
    },
  };
}
