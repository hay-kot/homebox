export default defineNuxtRouteMiddleware(async () => {
  const ctx = useAuthContext();
  const api = useUserApi();

  if (!ctx.user) {
    const { data, error } = await api.user.self();
    if (error) {
      navigateTo("/");
    }

    ctx.user = data.item;
  }
});
