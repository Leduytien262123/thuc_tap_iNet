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
    <div class="absolute -right-[100px] top-1/2 -translate-y-1/2 mt-11">
      <n-button
        @click="editMode = true"
        type="success"
        class="bg-green-600 px-6 py-3"
      >
        Sửa
      </n-button>
    </div>
  </div>
</template>

<script setup>
import {
  useMessage,
  NCollapse,
  NCollapseItem,
  NButton,
  NCard,
  NModal,
} from "naive-ui";

const about = defineModel("about");
const props = defineProps({
  response: {
    type: Object,
    default: {},
  },
});

const aboutUs = ref(about.value);
</script>
