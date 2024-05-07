<template>
  <div class="w-[72%] m-auto pb-[100px]">
    <div
      v-for="content in contents"
      :key="content.id"
      class="odd:flex-row even:flex-row-reverse w-full flex items-center mt-24"
    >
      <div class="w-[43%]">
        <h2 class="text-[40px] font-bold mb-4">{{ content.name }}</h2>
        <div v-html="content.text" class="text-base text-[#797979]" />
      </div>
      <div class="w-[6%]"></div>
      <div class="w-[53%]">
        <div
          class="w-full h-[300px] rounded-[50px] flex justify-center items-center overflow-hidden"
        >
          <img :src="content.img" alt="" class="w-full h-full" />
        </div>
      </div>
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

const content = defineModel("content");
const props = defineProps({
  response: {
    type: Object,
    default: {},
  },
});

const contents = ref(content.value);
</script>
