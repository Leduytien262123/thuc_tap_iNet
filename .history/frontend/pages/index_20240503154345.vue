<template>
  <div class="w-full mx-auto bg-[#F3F3F4] relative z-10">
    <TrangChuBanner :response="response" />
    <TrangChuThongSo :response="response" />
    <TrangChuTsChonChungToi
      v-model:item2="response[0].Metadata.tschonchungtoi"
      :response="response"
    />
    <TrangChuKhoaHocNoiBat
      v-model:course2="response[0].Metadata.khoahocnoibat"
      :response="response"
    />
    <TrangChuGoiKhoaHocTieuBieu
      v-model:course2="response[0].Metadata.goikhoahoctieubieu"
      :response="response"
    />
    <TrangChuDanhMucKhoaHoc
      v-model:category2="response[0].Metadata.danhmuckhoahoc"
      :response="response"
    />
    <TrangChuBaiVietNoiBat
      v-model:article2="response[0].Metadata.baivietnoibat"
      :response="response"
    />
    <TrangChuGiangVienTieuBieu
      v-model:teacher2="response[0].Metadata.giangvientieubieu"
      :response="response"
    />
    <TrangChuHvNoiGiVeChungToi
      v-model:testimonial2="response[0].Metadata.hvnoigivechungtoi"
      :response="response"
    />
    <TrangChuDoiTacTieuBieu :response="response" />
    <TrangChuNhanUuDaiNgay :response="response" />
  </div>
</template>

<script setup>
const { data } = await useFetch("http://localhost:8000/page");
const response = data.value;

const tschonchungtoi = ref(response[0]?.Metadata?.tschonchungtoi ?? {});
const khoahocnoibat = ref(response[0]?.Metadata?.khoahocnoibat ?? {});
const goikhoahoctieubieu = ref(response[0]?.Metadata?.goikhoahoctieubieu ?? {});
const danhmuckhoahoc = ref(response[0]?.Metadata?.danhmuckhoahoc ?? {});
const baivietnoibat = ref(response[0]?.Metadata?.baivietnoibat ?? {});
const giangvientieubieu = ref(response[0]?.Metadata?.giangvientieubieu ?? {});
const hvnoigivechungtoi = ref(response[0]?.Metadata?.hvnoigivechungtoi ?? {});
const doitactieubieu = ref(response[0]?.Metadata?.doitactieubieu ?? {});
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
