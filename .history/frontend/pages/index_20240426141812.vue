<template>
  <div class="w-full mx-auto bg-[#F3F3F4] relative z-10">
    <TrangChuBanner :response="response" />
    <TrangChuThongSo :response="response" />
    <TrangChuTsChonChungToi :response="response" />
    <TrangChuKhoaHocNoiBat :response="response" />
    <TrangChuGoiKhoaHocTieuBieu :response="response" />
    <TrangChuDanhMucKhoaHoc :response="response" />
    <TrangChuBaiVietNoiBat :response="response" />
    <TrangChuGiangVienTieuBieu :response="response" />
    <TrangChuHvNoiGiVeChungToi :response="response" />
    <TrangChuDoiTacTieuBieu :response="response" />
    <TrangChuNhanUuDaiNgay :response="response" />
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
const data_giangvientieubieu =
  response[0]?.Metadata?.data_giangvientieubieu ?? {};
const data_goikhoahoctieubieu =
  response[0]?.Metadata?.data_goikhoahoctieubieu ?? {};
const data_hvnoigivechungtoi =
  response[0]?.Metadata?.data_hvnoigivechungtoi ?? {};
const data_khoahocnoibat = response[0]?.Metadata?.data_khoahocnoibat ?? {};
const data_khoahoctieubieu = response[0]?.Metadata?.data_khoahoctieubieu ?? {};
const data_sumenhvatamnhin = response[0]?.Metadata?.data_sumenhvatamnhin ?? {};
const data_thongso = response[0]?.Metadata?.data_thongso ?? {};
const data_tschonchungtoi = response[0]?.Metadata?.data_tschonchungtoi ?? {};
const data_vechungtoi = response[0]?.Metadata?.data_vechungtoi ?? {};
</script>
