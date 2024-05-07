<template>
  <div class="bg-[#F3F3F4]">
    <KhoaHocVeChungToi />
    <TrangChuThongSo :response="response" />
    <KhoaHocSuMenhVaTamNhin />
    <TrangChuNhanUuDaiNgay />
  </div>
</template>

<script setup>
const { data } = await useFetch("http://localhost:8000/page");
const response = data.value;

response.forEach((item) => {
  if (item.Metadata) {
    // console.log(item.Metadata);
  } else {
    // console.log(
    //   "Không tìm thấy thông tin Metadata trong một số phần tử của response."
    // );
  }
});

const data_thongso = response[0]?.Metadata?.data_thongso ?? {};
</script>
