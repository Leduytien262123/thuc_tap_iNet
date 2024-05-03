<template>
  <div class="w-full mb-[100px]">
    <div class="w-[72%] m-auto">
      <div class="flex items-center justify-between mb-8">
        <h2 class="text-[40px] flex font-semibold">
          {{ title }}
          <p class="ml-3 text-[#6046FF]">{{ subTitle }}</p>
        </h2>
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
                slidesPerView: 3,
              },
            }"
            :space-between="25"
            loop
            :autoplay="{
              delay: 2000,
              disableOnInteraction: false,
              pauseOnMouseEnter: true,
            }"
            :speed="1000"
            class="!z-0 rounded-lg"
          >
            <SwiperSlide
              v-for="(testimonial, index) of testimonials"
              :key="index"
            >
              <div class="w-full flex mb-6 justify-between">
                <div class="bg-white rounded-lg justify-center items-center">
                  <div class="w-[78%] m-auto mt-8 mb-8">
                    <p>{{ testimonial.content }}</p>
                  </div>
                  <hr />
                  <div class="w-[78%] m-auto mt-6 mb-6 flex items-center">
                    <div class="w-[50px] h-[55px] rounded-lg overflow-hidden">
                      <img
                        class="w-full h-full rounded-lg object-cover"
                        :src="testimonial.image"
                      />
                    </div>
                    <b class="text-lg ml-4">{{ testimonial.name }}</b>
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

const testimonial2 = defineModel("testimonial2");

const title = ref(testimonial2.value.title);
const subTitle = ref(testimonial2.value.subTitle);

const modules = [Autoplay, Pagination];
const mySwiper = ref();

const props = defineProps({
  response: Object,
});

const testimonials = ref([]);

if (props.response && props.response[0]?.Metadata?.data_hvnoigivechungtoi) {
  testimonials.value = props.response[0].Metadata.data_hvnoigivechungtoi;
}
</script>
