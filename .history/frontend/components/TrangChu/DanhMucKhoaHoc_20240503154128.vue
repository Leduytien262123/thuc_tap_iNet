<template>
  <div class="w-[72%] m-auto mb-[120px] relative">
    <div class="flex items-center justify-between mb-8">
      <h2 class="text-[40px] flex font-semibold">
        <p class="mr-3 text-[#6046FF]">{{ title }}</p>
        {{ subTitle }}
      </h2>
      <button
        class="bg-white text-black rounded-md px-5 py-3 text-center block ml-auto"
      >
        Xem thêm
      </button>
    </div>
    <div class="w-full grid grid-cols-5 gap-x-5 gap-y-6 mb-6">
      <template v-for="category in categories" :key="category.id">
        <div
          class="pb-10 bg-white rounded-lg block overflow-hidden text-center cursor-pointer transition-transform transform-gpu motion-safe:hover:scale-105 border border-transparent hover:border-[#6046FF] hover:filter"
        >
          <div before class="m-auto w-14 h-14 mt-12 mb-8">
            <img :src="category.icon" alt="" class="w-full h-full" />
          </div>
          <span class="font-medium text-lg mb-8"
            >{{ category.name1 }}<br />
            {{ category.name2 }}
          </span>
        </div>
      </template>
    </div>
  </div>
</template>

<script setup>
const category2 = defineModel("category2");

const title = ref(category2.value.title);
const subTitle = ref(category2.value.subTitle);

const props = defineProps({
  response: Object,
});

const categories = ref([]);

if (props.response && props.response[0]?.Metadata?.data_danhmuckhoahoc) {
  categories.value = props.response[0].Metadata.data_danhmuckhoahoc;
}
</script>
