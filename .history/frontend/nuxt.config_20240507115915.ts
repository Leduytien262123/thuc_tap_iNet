export default {
  devtools: { enabled: true },
  modules: ["@nuxtjs/tailwindcss"],
  middleware: ["domReady"],
  css: ["~/assets/css/tailwind.css", "~/assets/css/main.css"],

  axios: {
    baseURL: process.env.API_URL, // Sử dụng biến môi trường API_URL cho cấu hình Axios
  },
  server: {
    host: "0.0.0.0", // default: localhost
    port: process.env.PORT || 3000, // Sử dụng biến môi trường PORT, hoặc mặc định là 3000
  },
  debug: process.env.DEBUG_MODE === "true", // Chỉ bật chế độ debug nếu DEBUG_MODE được đặt thành 'true'

  head: {
    link: [
      { rel: "icon", href: "/assets/image/mschool.png" }, // Thêm thuộc tính 'rel' và chỉ định biểu tượng favicon
    ],
  },
};
