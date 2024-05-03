<template>
  <div class="w-[72%] m-auto">
    <div class="flex items-center justify-between mb-8">
      <h2 class="text-[40px] flex font-semibold">
        {{ title }}
        <p class="ml-3 text-[#6046FF]">{{ subTitle }}</p>
      </h2>
    </div>
    <div class="flex justify-center">
      <div
        v-for="(partner, index) in partners"
        :key="index"
        class="w-[14.2%] bg-[#F3F3F4] overflow-hidden"
      >
        <div
          class="px-5 py-3 transition duration-300 transform hover:scale-110 cursor-pointer"
        >
          <img
            class="w-full h-full object-cover"
            :src="partner.image"
            :alt="partner.name"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const partner2 = defineModel("partner2");

const title = ref(partner2.value.title);
const subTitle = ref(partner2.value.subTitle);

const props = defineProps({
  response: Object,
});

const partners = ref([]);

if (props.response && props.response[0]?.Metadata?.data_doitactieubieu) {
  partners.value = props.response[0].Metadata.data_doitactieubieu;
}
</script>
