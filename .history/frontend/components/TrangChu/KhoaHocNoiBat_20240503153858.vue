<template>
  <div class="w-[72%] m-auto mb-[100px] relative">
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
    <div class="w-full mx-auto h-auto mb-6 grid grid-cols-4 gap-6">
      <div
        v-for="(course, index) in courses.slice(0, 8)"
        :key="index"
        class="bg-white rounded-lg block overflow-hidden transition-transform duration-300 transform-gpu motion-safe:hover:scale-110"
      >
        <img class="rounded-none w-full h-[130px]" :src="course.image" />
        <div class="w-[86%] m-auto mt-4">
          <div class="flex items-center mb-3">
            <img class="w-4 h-4" src="/img/chayphat.png" />
            <p class="ml-2 text-xs">{{ course.lessons }} bài giảng</p>
          </div>
          <h1 class="text-[17px] font-semibold mb-3">{{ course.nameTitle }}</h1>
          <div
            v-if="!course.price"
            class="text-[#6046FF] text-xl font-bold mt-3"
          >
            {{ formatPrice(course.oldPrice) }}
          </div>
          <div
            v-else-if="course.price !== '0'"
            class="flex items-end mb-4 mt-3"
          >
            <span class="text-sm line-through opacity-50 mb-[1.5px]">{{
              formatPrice(course.oldPrice)
            }}</span>
            <span class="text-[#6046FF] text-xl font-bold ml-2">{{
              formatPrice(course.price)
            }}</span>
          </div>
          <div v-else class="flex items-end mb-4 mt-3">
            <span class="text-sm line-through opacity-50 mb-[1.5px]">{{
              formatPrice(course.oldPrice)
            }}</span>
            <span class="text-[#6046FF] text-xl font-bold ml-2">Miễn phí</span>
          </div>
          <hr />
          <div class="flex mt-4 mb-4">
            <div class="w-[64%] flex">
              <img class="rounded-full w-5 h-5" :src="course.avatar" />
              <span class="text-sm ml-2">{{ course.name }}</span>
            </div>
            <div class="w-[36%]">
              <nuxt-link :to="course.link">
                <button
                  v-if="parseInt(course.price) > 0"
                  class="bg-[#6046FF] text-white ml-0 text-xs rounded-md font-thin px-2 py-1 text-center block hover:bg-sky-500"
                >
                  Mua ngay
                </button>
                <button
                  v-else
                  class="bg-[#6046FF] text-white ml-0 text-xs rounded-md font-thin px-2 py-1 text-center block hover:bg-sky-500"
                >
                  Xem ngay
                </button>
              </nuxt-link>
            </div>
          </div>
        </div>
        <div class="w-[2.4%]"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const course2 = defineModel("course2");

const title = ref(course2.value.title);
const subTitle = ref(course2.value.subTitle);

const props = defineProps({
  response: Object,
});

const courses = ref([]);

if (props.response && props.response[0]?.Metadata?.data_khoahocnoibat) {
  courses.value = props.response[0].Metadata.data_khoahocnoibat;
}

const formatPrice = (price) => {
  if (!price) return "";
  const formattedPrice = parseFloat(price).toLocaleString("vi-VN", {
    style: "decimal",
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  });
  return formattedPrice + "đ";
};
</script>
