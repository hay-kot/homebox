import { useAuthStore } from "~~/stores/auth";

export default defineNuxtRouteMiddleware(async () => {
  const auth = useAuthStore();
  const api = useUserApi();

  if (!auth.self) {
    const { data, error } = await api.user.self();
    if (error) {
      navigateTo("/");
    }

    auth.$patch({ self: data.item });
  }
});
