<template>
  <div class="w-full">
    <CmsTrangChuDefault :response="response[0]" />
    <CmsTrangChuHeader
      v-model:menu="response[0].Metadata.data_menus"
      v-model:form="response[0].Metadata.data_menus"
      :response="response"
    />
    <div class="bg-[#F3F3F4]">
      <CmsTrangChuBanner
        v-model:banner="response[0].Metadata.data_banner"
        :response="response"
      />
      <CmsTrangChuThongSo
        v-model:thongso="response[0].Metadata.data_thongso"
        :response="response"
      />
      <CmsTrangChuTsChonChungToi
        v-model:form="response[0].Metadata.data_tschonchungtoi"
        :response="response"
      />
      <CmsTrangChuKhoaHocNoiBat
        v-model:form="response[0].Metadata.data_khoahocnoibat"
        :response="response"
      />
      <CmsTrangChuGoiKhoaHocTieuBieu
        v-model:form="response[0].Metadata.data_goikhoahoctieubieu"
        :response="response"
      />
      <CmsTrangChuDanhMucKhoaHoc
        v-model:form="response[0].Metadata.data_danhmuckhoahoc"
        :response="response"
      />
      <CmsTrangChuBaiVietNoiBat
        v-model:form="response[0].Metadata.data_baivietnoibat"
        :response="response"
      />
      <CmsTrangChuGiangVienTieuBieu
        v-model:form="response[0].Metadata.data_giangvientieubieu"
        :response="response"
      />
      <CmsTrangChuHvNoiGiVeChungToi
        v-model:form="response[0].Metadata.data_hvnoigivechungtoi"
        :response="response"
      />
      <CmsTrangChuDoiTacTieuBieu
        v-model:form="response[0].Metadata.data_doitactieubieu"
        :response="response"
      />
      <TrangChuNhanUuDaiNgay />
    </div>
    <CmsTrangChuFooter
      v-model:form="response[0].Metadata.data_footer"
      :response="response"
    />
  </div>
</template>

<script setup>
const { data } = await useFetch("http://localhost:8000/page");
const response = data.value;

const data_baivietnoibat = ref(response[0]?.Metadata?.data_baivietnoibat ?? {});
const data_banner = ref(response[0]?.Metadata?.data_banner ?? {});
const data_danhmuckhoahoc = ref(
  response[0]?.Metadata?.data_danhmuckhoahoc ?? {}
);
const data_doitactieubieu = ref(
  response[0]?.Metadata?.data_doitactieubieu ?? {}
);
const data_form = ref(response[0]?.Metadata?.data_form ?? {});
const data_giangvientieubieu = ref(
  response[0]?.Metadata?.data_giangvientieubieu ?? {}
);
const data_goikhoahoctieubieu = ref(
  response[0]?.Metadata?.data_goikhoahoctieubieu ?? {}
);
const data_hvnoigivechungtoi = ref(
  response[0]?.Metadata?.data_hvnoigivechungtoi ?? {}
);
const data_khoahocnoibat = ref(response[0]?.Metadata?.data_khoahocnoibat ?? {});
const data_khoahoctieubieu = ref(
  response[0]?.Metadata?.data_khoahoctieubieu ?? {}
);
const data_menus = ref(response[0]?.Metadata?.data_menus ?? {});
const data_sumenhvatamnhin = ref(
  response[0]?.Metadata?.data_sumenhvatamnhin ?? {}
);
const data_thongso = ref(response[0]?.Metadata?.data_thongso ?? {});
const data_tschonchungtoi = ref(
  response[0]?.Metadata?.data_tschonchungtoi ?? {}
);
const data_vechungtoi = ref(response[0]?.Metadata?.data_vechungtoi ?? {});
watch(
  () => response[0].Metadata.data_menus,
  () => response[0].Metadata.data_form,
  () => response[0].Metadata.data_baivietnoibat,
  () => response[0].Metadata.data_banner,
  () => response[0].Metadata.data_danhmuckhoahoc,
  () => response[0].Metadata.data_doitactieubieu,
  () => response[0].Metadata.data_giangvientieubieu,
  () => response[0].Metadata.data_goikhoahoctieubieu,
  () => response[0].Metadata.data_hvnoigivechungtoi,
  () => response[0].Metadata.data_khoahocnoibat,
  () => response[0].Metadata.data_khoahoctieubieu,
  () => response[0].Metadata.data_sumenhvatamnhin,
  () => response[0].Metadata.data_sumenhvatamnhin,
  () => response[0].Metadata.data_thongso,
  () => response[0].Metadata.data_tschonchungtoi,
  () => response[0].Metadata.data_vechungtoi
);
</script>
