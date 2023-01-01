export default defineNuxtPlugin(nuxtApp => {
  nuxtApp.hook("page:finish", () => {
    console.log(document.body);
    document.body.scrollTo({ top: 0 });
    console.log("page:finish");
  });
});
