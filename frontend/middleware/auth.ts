export default defineNuxtRouteMiddleware(async () => {
  const ctx = useAuthContext();
  const api = useUserApi();

  if (!ctx.isAuthorized()) {
    return navigateTo("/");
  }

  if (!ctx.user) {
    console.log("Fetching user data");
    const { data, error } = await api.user.self();
    if (error) {
      return navigateTo("/");
    }

    ctx.user = data.item;
  }
});
