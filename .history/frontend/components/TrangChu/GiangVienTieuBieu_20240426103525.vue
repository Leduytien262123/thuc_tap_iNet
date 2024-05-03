<template>
  <div class="w-full mb-[110px]">
    <div class="w-[72%] m-auto">
      <div class="flex items-center justify-between mb-8">
        <h2 class="text-[40px] flex font-semibold">
          Giảng viên
          <p class="ml-3 text-[#6046FF]">tiêu biểu</p>
        </h2>
        <button
          class="bg-white text-black rounded-md px-5 py-3 text-center block ml-auto"
        >
          Xem thêm
        </button>
      </div>
    </div>
    <div class="w-full">
      <div class="w-[80%] mx-auto flex justify-between items-center">
        <div class="w-[90%] relative ml-[5%] flex justify-between items-center">
          <button @click="mySwiper.slidePrev()">
            <img
              class="absolute -left-16 top-1/2 -translate-y-1/2 w-6 h-6 opacity-50 hover:opacity-100"
              src="/icons/chevron-back-outline.svg"
            />
          </button>
          <Swiper
            :modules="modules"
            @swiper="(s) => (mySwiper = s)"
            :slides-per-view="1"
            :breakpoints="{
              300: {
                slidesPerView: 1.175,
              },
              400: {
                slidesPerView: 1.515,
              },
              768: {
                slidesPerView: 3,
                spaceBetween: 20,
              },
              1024: {
                slidesPerView: 4,
              },
            }"
            :space-between="30"
            loop
            :autoplay="{
              delay: 2000,
              disableOnInteraction: false,
              pauseOnMouseEnter: true,
            }"
            :speed="1000"
            class="!z-0 rounded-lg"
          >
            <SwiperSlide v-for="(teacher, index) of teachers" :key="index">
              <div
                class="w-full py-0.5 bg-white rounded-lg text-center hover:bg-[#EEEEFF]"
              >
                <div
                  class="mx-auto mt-8 mb-8 flex flex-col items-center scroll-snap-start"
                >
                  <div
                    class="rounded-full w-[110px] h-[110px] mb-4 flex items-center justify-center overflow-hidden"
                  >
                    <img
                      class="w-full h-full object-cover"
                      :src="teacher.image"
                      :alt="teacher.name"
                    />
                  </div>
                  <b class="text-lg">{{ teacher.name }}</b>
                  <p class="text-sm mt-1">{{ teacher.description }}</p>
                  <div class="flex justify-center items-center mt-3">
                    <nuxt-link
                      v-for="(link, linkIndex) in teacher.links"
                      :key="linkIndex"
                      :to="link.url"
                    >
                      <img
                        class="w-3.5 h-3.5 mx-1.5"
                        :src="link.icon"
                        :alt="link.name"
                      />
                    </nuxt-link>
                  </div>
                </div>
              </div>
            </SwiperSlide>
          </Swiper>
          <button @click="mySwiper.slideNext()">
            <img
              class="absolute -right-16 top-1/2 -translate-y-1/2 w-6 h-6 opacity-50 hover:opacity-100"
              src="/icons/chevron-forward-outline.svg"
            />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Swiper, SwiperSlide } from "swiper/vue";
import { Autoplay, Pagination } from "swiper/modules";
import "swiper/css";

const modules = [Autoplay, Pagination];
const mySwiper = ref();

const props = defineProps({
  response: Object,
});

const teachers = ref([]);

if (props.response && props.response[0]?.Metadata?.data_giangvientieubieu) {
  teachers.value = props.response[0].Metadata.data_giangvientieubieu;
}

// const { data } = await useFetch("http://localhost:8000/page");

// const giangVienTieuBieu = data.value;

// let teachers = ref([]);

// giangVienTieuBieu.forEach((item) => {
//   if (item.Metadata) {
//     if (item.Metadata.data_giangvientieubieu) {
//       teachers.value = item.Metadata.data_giangvientieubieu;
//     }
//   }
// });
</script>
