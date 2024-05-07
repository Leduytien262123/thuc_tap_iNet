<template>
  <div class="relative w-full">
    <div class="w-[80%] mx-auto flex justify-between items-center">
      <button
        class="absolute bottom-[14%] right-[8%] bg-[#6046FF] hover:bg-sky-500 p-3 rounded-full"
        @click="scrollToTop"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="ionicon w-6 h-6 text-white fill-current"
          viewBox="0 0 512 512"
        >
          <path
            fill="none"
            stroke="currentColor"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="48"
            d="M112 328l144-144 144 144"
          />
        </svg>
      </button>
    </div>
    <div class="w-[72%] mx-auto mt-12">
      <div class="relative rounded-lg overflow-hidden">
        <img
          src="/bg/bg-cuoicung.png"
          alt=""
          class="w-full h-[370px] rounded-lg"
        />
        <div
          class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 text-center"
        >
          <span class="text-3xl text-[#6046FF] font-bold block"
            >Nhận ưu đãi ngay</span
          >
          <input
            v-model="email"
            type="email"
            class="text-[#6046FF] mt-7 mb-8 w-full text-base py-4 pl-6 rounded-xl border focus:outline-none focus:border-[#6046FF] placeholder-[#6046FF]"
            placeholder="Email"
          />
          <button
            @click="sendEmail"
            class="bg-[#6046FF] rounded-lg py-4 px-20 text-lg text-white hover:bg-[#432CBB]"
          >
            Gửi
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const email = ref("");

const scrollToTop = () => {
  window.scrollTo({ top: 0, behavior: "smooth" });
};

const sendEmail = async () => {
  try {
    if (!validateEmail(email.value)) {
      alert("Vui lòng nhập một địa chỉ email hợp lệ.");
      return;
    }
    const response = await fetch("http://10.50.20.163:4000/page", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer YOUR_SENDGRID_API_KEY_HERE",
      },
      body: JSON.stringify({
        personalizations: [
          {
            to: [{ email: email.value }],
            subject: "Subject of the email",
          },
        ],
        from: { email: "leduytien262@gmail.com" },
        content: [
          {
            type: "text/plain",
            value: "Hi cảm ơn bạn đã đăng ký khóa học",
          },
        ],
      }),
    });
    const data = await response.json();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
};

const validateEmail = (email) => {
  // Sử dụng một biểu thức chính quy đơn giản để kiểm tra địa chỉ email hợp lệ
  const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return regex.test(email);
};
</script>
