export default defineNuxtRouteMiddleware(async () => {
  const ctx = useAuthContext();
  const api = useUserApi();

  if (!ctx.isAuthorized()) {
    return navigateTo("/");
  }

  if (!ctx.user) {
    const { data, error } = await api.user.self();
    if (error) {
      return navigateTo("/");
    }

    ctx.user = data.item;
  }
});
