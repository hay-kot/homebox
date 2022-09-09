import { defineNuxtConfig } from "nuxt";

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
  target: "static",
  ssr: false,
  modules: ["@nuxtjs/tailwindcss", "@pinia/nuxt", "@vueuse/nuxt"],
  meta: {
    title: "Homebox",
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.svg" }],
  },
  vite: {
    server: {
      proxy: {
        "/api": "http://localhost:7745",
      },
    },
    plugins: [],
  },
});
