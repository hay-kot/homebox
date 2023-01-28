export default defineNuxtPlugin(nuxtApp => {
  nuxtApp.hook("page:finish", () => {
    document.body.scrollTo({ top: 0 });
  });
});
