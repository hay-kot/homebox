import { defineNuxtConfig } from "nuxt/config";

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
  ssr: false,
  modules: ["@nuxtjs/tailwindcss", "@pinia/nuxt", "@vueuse/nuxt"],
  nitro: {
    devProxy: {
      "/api": {
        target: "http://localhost:7745/api",
        changeOrigin: true,
      },
    },
  },
  css: ["@/assets/css/main.css"],
  plugins: [],
});
