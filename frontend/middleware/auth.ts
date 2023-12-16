export default defineNuxtRouteMiddleware(async () => {
  const ctx = useAuthContext();
  const api = useUserApi();

  if (!ctx.isAuthorized()) {
    if (window.location.pathname !== "/") {
      console.debug("[middleware/auth] isAuthorized returned false, redirecting to /");
      return navigateTo("/");
    }
  }

  if (!ctx.user) {
    console.log("Fetching user data");
    const { data, error } = await api.user.self();
    if (error) {
      if (window.location.pathname !== "/") {
        console.debug("[middleware/user] user is null and fetch failed, redirecting to /");
        return navigateTo("/");
      }
    }

    ctx.user = data.item;
  }
});
