<template>
  <div class="relative w-[72%] m-auto mb-[120px]">
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
    <div class="w-full h-full mb-6 grid grid-cols-3 gap-6">
      <div
        v-for="(course, index) in courses.slice(0, 3)"
        :key="index"
        class="bg-white rounded-lg block overflow-hidden"
      >
        <div class="w-[90%] m-auto relative">
          <div class="h-[40%]">
            <img
              class="h-[250px] w-full m-auto rounded-md mt-[5%] mb-[8%]"
              :src="course.image"
            />
            <span class="text-xl font-bold">{{ course.nameTitle }}</span>
          </div>
          <div class="h-[30%]">
            <p class="mt-3 text-sm mb-8">
              {{ course.description }}
            </p>
          </div>
          <div class="h-[30%] w-full float-bottom">
            <span class="text-lg">Giảng viên: {{ course.name }}</span>
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
              <span class="text-[#6046FF] text-xl font-bold ml-2"
                >Miễn phí</span
              >
            </div>
            <nuxt-link :to="course.link">
              <button
                class="w-full h-12 mb-[5%] rounded-md border border-[#6046FF] text-[#6046FF] text-base font-semibold transition-transform transform-gpu motion-safe:hover:scale-105"
              >
                Đăng ký ngay
              </button></nuxt-link
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const course2 = defineModel("course2");

const title = ref(course2.value.title);
const subTitle = ref(course2.value.subTitle);

const props = defineProps({
  response: Object,
});

const courses = ref([]);

if (props.response && props.response[0]?.Metadata?.data_goikhoahoctieubieu) {
  courses.value = props.response[0].Metadata.data_goikhoahoctieubieu;
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
