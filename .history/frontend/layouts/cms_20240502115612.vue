<template>
  <div class="w-full">
    <CmsTrangChuDefault :response="response" />
    <CmsTrangChuHeader v-model:menu="data_menus" :response="response" />
    <div class="bg-[#F3F3F4]">
      <CmsTrangChuBanner :response="response" />
      <CmsTrangChuThongSo :response="response" />
      <CmsTrangChuTsChonChungToi :response="response" />
      <CmsTrangChuKhoaHocNoiBat :response="response" />
      <CmsTrangChuGoiKhoaHocTieuBieu :response="response" />
      <CmsTrangChuDanhMucKhoaHoc :response="response" />
      <CmsTrangChuBaiVietNoiBat :response="response" />
      <CmsTrangChuGiangVienTieuBieu :response="response" />
      <CmsTrangChuHvNoiGiVeChungToi :response="response" />
      <CmsTrangChuDoiTacTieuBieu :response="response" />
      <TrangChuNhanUuDaiNgay />
    </div>
    <CmsTrangChuFooter :response="response" />
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

const data_baivietnoibat = response[0]?.Metadata?.data_baivietnoibat ?? {};
const data_banner = response[0]?.Metadata?.data_banner ?? {};
const data_danhmuckhoahoc = response[0]?.Metadata?.data_danhmuckhoahoc ?? {};
const data_doitactieubieu = response[0]?.Metadata?.data_doitactieubieu ?? {};
const data_form = response[0]?.Metadata?.data_form ?? {};
const data_giangvientieubieu =
  response[0]?.Metadata?.data_giangvientieubieu ?? {};
const data_goikhoahoctieubieu =
  response[0]?.Metadata?.data_goikhoahoctieubieu ?? {};
const data_hvnoigivechungtoi =
  response[0]?.Metadata?.data_hvnoigivechungtoi ?? {};
const data_khoahocnoibat = response[0]?.Metadata?.data_khoahocnoibat ?? {};
const data_khoahoctieubieu = response[0]?.Metadata?.data_khoahoctieubieu ?? {};
const data_menus = ref(response[0]?.Metadata?.data_menus ?? {});
console.log("data mene: ", data_menus);
const data_sumenhvatamnhin = response[0]?.Metadata?.data_sumenhvatamnhin ?? {};
const data_thongso = response[0]?.Metadata?.data_thongso ?? {};
const data_tschonchungtoi = response[0]?.Metadata?.data_tschonchungtoi ?? {};
const data_vechungtoi = response[0]?.Metadata?.data_vechungtoi ?? {};
watch(
  () => data_menus.value,
  (newValue, oldValue) => {
    console.log("new ", newValue);
  }
);
</script>
