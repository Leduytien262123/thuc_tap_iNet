<template>
  <div class="w-[98%] mb-[100px]">
    <div class="text-[45px] flex justify-center mb-8 mt-20">
      <b class="flex items-center">
        <p class="text-blue-600 mr-3">{{ title }}</p>
        {{ subTitle }}
      </b>
    </div>
    <div class="w-full flex">
      <div class="w-[72%] m-auto rounded-xl flex">
        <div class="w-[48%] flex-">
          <div v-for="item in topItems" :key="item.id">
            <nuxt-link :to="item.link">
              <div
                class="bg-[#F9F9F9] h-[200px] mb-8 rounded-3xl flex flex-col"
              >
                <div class="inline-block mt-auto mb-auto">
                  <div class="flex items-center">
                    <img
                      class="w-[100px] h-[130px] rounded-xl mr-4 ml-6"
                      :src="item.imageSrc"
                      :alt="item.alt"
                    />
                    <div class="flex flex-col mr-6 ml-3">
                      <b class="mr-2 text-xl">{{ item.nameTitle }}</b>
                      <p class="text-base text-gray-600 mt-1">
                        {{ item.description }}
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </nuxt-link>
          </div>
        </div>
        <div class="w-[4%]"></div>
        <div class="w-[48%]">
          <div v-for="item in bottomItems" :key="item.id">
            <nuxt-link :to="item.link">
              <div
                class="bg-[#F9F9F9] h-[200px] mb-8 rounded-3xl flex flex-col"
              >
                <div class="inline-block mt-auto mb-auto">
                  <div class="flex items-center">
                    <img
                      class="w-[100px] h-[130px] rounded-xl mr-4 ml-6"
                      :src="item.imageSrc"
                      :alt="item.alt"
                    />
                    <div class="flex flex-col mr-6 ml-3">
                      <b class="mr-2 text-xl">{{ item.nameTitle }}</b>
                      <p class="text-base text-gray-600 mt-1">
                        {{ item.description }}
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </nuxt-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const item2 = defineModel("item2");
const title = ref(item2.value.title);
const subTitle = ref(item2.value.subTitle);

const props = defineProps({
  response: Object,
});

const items = ref([]);

if (props.response && props.response[0]?.Metadata?.data_tschonchungtoi) {
  items.value = props.response[0].Metadata.data_tschonchungtoi;
}

const topItems = ref(items.value.slice(0, 2));
const bottomItems = ref(items.value.slice(2));
</script>
