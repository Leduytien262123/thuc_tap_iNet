export default {
  // Cấu hình khác của bạn ở đây
  devtools: { enabled: true },
  modules: ["@nuxtjs/tailwindcss"],
  middleware: ["domReady"],
  css: ["~/assets/css/tailwind.css", "~/assets/css/main.css"],
};
