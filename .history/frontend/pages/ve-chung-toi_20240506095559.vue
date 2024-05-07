<template>
  <div class="bg-[#F3F3F4]">
    <VeChungToiTVeChungToi
      v-model:about="response[0].Metadata.data_vechungtoi"
      :response="response"
    />
    <TrangChuThongSo
      v-model:item="response[0].Metadata.data_thongso"
      :response="response"
    />
    <VeChungToiSuMenhVaTamNhin
      v-model:content="response[0].Metadata.data_sumenhvatamnhin"
      :response="response"
    />
    <TrangChuNhanUuDaiNgay />
  </div>
</template>

<script setup>
const { data } = await useFetch("http://localhost:8000/page");
const response = data.value;

const data_sumenhvatamnhin = ref(
  response[0]?.Metadata?.data_sumenhvatamnhin ?? {}
);
const data_thongso = ref(response[0]?.Metadata?.data_thongso ?? {});
const data_vechungtoi = ref(response[0]?.Metadata?.data_vechungtoi ?? {});
</script>
