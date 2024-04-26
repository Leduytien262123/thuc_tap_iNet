<template>
  <div class="banner relative w-[98%] flex flex-initial mx-auto">
    <img
      :src="form.backgroundImage"
      alt=""
      class="w-[100%] h-[600px] mx-auto rounded-lg mt-1 blur-none"
    />
    <div
      class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 items-center text-center justify-center inline"
    >
      <div class="bg-[#e9e9e9bf] p-4 rounded-md flex items-center gap-2">
        <div
          class="bg-[#FFF4E5] w-12 h-12 flex items-center justify-center rounded-md"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="ionicon w-6 h-6"
            viewBox="0 0 512 512"
          >
            <path d="M432 208H288l32-192L80 304h144l-32 192z" />
          </svg>
        </div>
        <span class="text-[#5236FF] text-4xl font-bold mr-1">{{
          form.title
        }}</span>
        <span class="text-black text-4xl font-bold mr-1">{{
          form.subtitle
        }}</span>
      </div>
      <div class="text-white text-3xl text-center mt-4">
        {{ form.description }}
      </div>
      <div class="mt-14">
        <nuxt-link :to="form.linkButton1">
          <button
            :style="{ backgroundColor: form.button1Color }"
            class="text-white px-5 py-3 rounded-md mr-3 text-center transition-transform transform-gpu motion-safe:hover:scale-110"
            @click="registerNow"
          >
            {{ form.button1 }}
          </button></nuxt-link
        >
        <nuxt-link :to="form.linkButton2">
          <button
            :style="{ backgroundColor: form.button2Color }"
            class="text-black px-5 py-3 rounded-md text-center transition-transform transform-gpu motion-safe:hover:scale-110"
            @click="viewCourses"
          >
            {{ form.button2 }}
          </button></nuxt-link
        >
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  response: Object,
});

const form = ref([]);

if (props.response && props.response[0]?.Metadata?.data_banner) {
  form.value = props.response[0].Metadata.data_banner;
}

// const { data: resHomeData } = await useFetch("http://localhost:8000/page");

// const banner = resHomeData.value;

// let form = ref({
//   title: "",
//   subtitle: "",
//   description: "",
//   backgroundImage: "",
//   button1: "",
//   linkButton1: "",
//   button1Color: "",
//   button2: "",
//   linkButton2: "",
//   button2Color: "",
// });

// banner.forEach((item) => {
//   if (item.Metadata && item.Metadata.data_banner) {
//     const bannerData = item.Metadata.data_banner;
//     form.value = {
//       title: bannerData.title,
//       subtitle: bannerData.subtitle,
//       description: bannerData.description,
//       backgroundImage: bannerData.backgroundImage,
//       button1: bannerData.button1,
//       linkButton1: bannerData.linkButton1,
//       button1Color: bannerData.button1Color,
//       button2: bannerData.button2,
//       linkButton2: bannerData.linkButton2,
//       button2Color: bannerData.button2Color,
//     };
//   }
// });
</script>
