<template>
  <div class="w-[72%] m-auto mb-[110px] relative">
    <div class="flex items-center justify-between mb-8">
      <h2 class="text-[40px] flex font-semibold">
        {{ title }}
        <p class="ml-3 text-[#6046FF]">{{ subTitle }}</p>
      </h2>
      <button
        class="bg-white text-black rounded-md px-5 py-3 text-center block ml-auto"
      >
        Xem thêm
      </button>
    </div>
    <div class="w-full grid grid-cols-2 gap-x-8">
      <template v-for="article in articles.slice(0, 2)" :key="article.id">
        <div class="bg-white rounded-xl relative">
          <div class="w-[87%] m-auto">
            <div class="my-[6.5%] relative h-[280px]">
              <img class="rounded-xl w-full h-full" :src="article.image" />
              <div class="absolute bottom-4 left-0 z-10 flex">
                <nuxt-link :to="article.tag1">
                  <button
                    class="px-8 py-2.5 ml-4 font-medium text-sm bg-white rounded-3xl transition-transform transform-gpu motion-safe:hover:scale-110 hover:bg-sky-300"
                  >
                    Marketing
                  </button></nuxt-link
                >
                <nuxt-link :to="article.tag2">
                  <button
                    class="px-8 py-2.5 ml-4 font-medium text-sm bg-white rounded-3xl transition-transform transform-gpu motion-safe:hover:scale-110 hover:bg-sky-300"
                  >
                    Analysis
                  </button></nuxt-link
                >
              </div>
            </div>
            <p class="mt-2 mb-4">{{ currentDate }}</p>
            <b class="text-xl">{{ article.nameTitle }}</b>
            <p class="text-[#797979] mt-3 mb-6 text-[13.5px]">
              {{ article.description }}
            </p>
            <div class="mb-[6.5%]">
              <nuxt-link
                :to="article.link"
                class="text-blue-600 underline text-sm"
                >Xem thêm</nuxt-link
              >
            </div>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup>
import { ref, watchEffect } from "vue";

const title = ref("Bài viết");
const subTitle = ref("nổi bật");

const currentDate = ref("");

watchEffect(() => {
  const date = new Date();
  const day = date.getDate();
  const month = date.getMonth() + 1;
  const year = date.getFullYear();
  currentDate.value = `${day}/${month}/${year}`;
});

const props = defineProps({
  response: Object,
});

const articles = ref([]);

if (props.response && props.response[0]?.Metadata?.data_baivietnoibat) {
  articles.value = props.response[0].Metadata.data_baivietnoibat;
}

// const { data } = await useFetch("http://localhost:8000/page");

// const baiVietNoiBat = data.value;

// let articles = ref([]);

// baiVietNoiBat.forEach((item) => {
//   if (item.Metadata) {
//     if (item.Metadata.data_baivietnoibat) {
//       articles.value = item.Metadata.data_baivietnoibat;
//     }
//   }
// });
</script>
