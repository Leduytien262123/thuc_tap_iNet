<template>
  <div class="bg-[#F3F3F4]">
    <KhoaHocVeChungToi
      v-model:about="response[0].Metadata.data_vechungtoi"
      :response="response"
    />
    <TrangChuThongSo :response="response" />
    <KhoaHocSuMenhVaTamNhin />
    <TrangChuNhanUuDaiNgay />
  </div>
</template>

<script setup>
const { data } = await useFetch("http://localhost:8000/page");
const response = data.value;

const data_thongso = response[0]?.Metadata?.data_thongso ?? {};
</script>
