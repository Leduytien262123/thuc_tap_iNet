export default {
  devtools: { enabled: true },
  modules: ["@nuxtjs/tailwindcss"],
  middleware: ["domReady"],
  css: ["~/assets/css/tailwind.css", "~/assets/css/main.css"],

  // axios: {
  //   baseURL: process.env.API_URL,
  // },
  // server: {
  //   host: "0.0.0.0",
  //   port: process.env.PORT || 3000,
  // },
  // debug: process.env.DEBUG_MODE === "true",
};
