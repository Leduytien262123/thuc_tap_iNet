<template>
  <div class="w-[72%] m-auto pt-4">
    <div v-for="about in aboutUs" :key="about.id" class="">
      <h1 class="text-center text-[40px] font-bold">{{ about.name }}</h1>
      <div class="w-full pt-8">
        <div class="flex">
          <div
            class="w-[62%] h-[280px] mr-[3%] rounded-3xl flex justify-center items-center overflow-hidden"
          >
            <img :src="about.img1" alt="" class="w-full h-full" />
          </div>
          <div
            class="w-[35%] h-[280px] rounded-3xl flex justify-center items-center overflow-hidden"
          >
            <img :src="about.img2" alt="" class="w-full h-full" />
          </div>
        </div>
        <div class="flex pb-2">
          <div class="w-[30%] items-center justify-center overflow-hidden p-8">
            <h6 class="text-2xl font-bold text-center px-4">
              {{ about.title }}
            </h6>
          </div>
          <div
            class="w-[70%] items-center justify-center overflow-hidden py-8 px-10"
          >
            <p class="text-[#4E4E4E] text-sm">
              {{ about.text }}
            </p>
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

const aboutU = defineModel("aboutU");
const props = defineProps({
  response: {
    type: Object,
    default: {},
  },
});

const aboutUs = ref(aboutU.value);
// const { data } = await useFetch("http://localhost:8000/page");
// const vechungtoiData = data.value;

// let aboutUs = ref("");

// vechungtoiData.forEach((item) => {
//   if (item.Metadata) {
//     if (item.Metadata.data_vechungtoi) {
//       aboutUs = item.Metadata.data_vechungtoi;
//     }
//   }
// });
</script>
