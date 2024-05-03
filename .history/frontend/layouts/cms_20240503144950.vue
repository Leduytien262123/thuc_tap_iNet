<template>
  <div class="w-full">
    <CmsTrangChuDefault :response="response[0]" />
    <CmsTrangChuHeader
      v-model:menu="response[0].Metadata.data_menus"
      v-model:form="response[0].Metadata.data_form"
      :response="response"
    />
    <div class="bg-[#F3F3F4]">
      <CmsTrangChuBanner
        v-model:form="response[0].Metadata.data_banner"
        :response="response"
      />
      <CmsTrangChuThongSo
        v-model:item="response[0].Metadata.data_thongso"
        :response="response"
      />
      <CmsTrangChuTsChonChungToi
        v-model:item="response[0].Metadata.data_tschonchungtoi"
        v-model:item2="response[0].Metadata.tschonchungtoi"
        :response="response"
      />
      <CmsTrangChuKhoaHocNoiBat
        v-model:course="response[0].Metadata.data_khoahocnoibat"
        :response="response"
      />
      <CmsTrangChuGoiKhoaHocTieuBieu
        v-model:course="response[0].Metadata.data_goikhoahoctieubieu"
        :response="response"
      />
      <CmsTrangChuDanhMucKhoaHoc
        v-model:category="response[0].Metadata.data_danhmuckhoahoc"
        :response="response"
      />
      <CmsTrangChuBaiVietNoiBat
        v-model:article="response[0].Metadata.data_baivietnoibat"
        :response="response"
      />
      <CmsTrangChuGiangVienTieuBieu
        v-model:teacher="response[0].Metadata.data_giangvientieubieu"
        :response="response"
      />
      <CmsTrangChuHvNoiGiVeChungToi
        v-model:testimonial="response[0].Metadata.data_hvnoigivechungtoi"
        :response="response"
      />
      <CmsTrangChuDoiTacTieuBieu
        v-model:partner="response[0].Metadata.data_doitactieubieu"
        :response="response"
      />
      <TrangChuNhanUuDaiNgay />
    </div>
    <CmsTrangChuFooter
      v-model:logoUrls="response[0].Metadata.data_footer[0].logoUrl"
      v-model:contents="response[0].Metadata.data_footer[1].content"
      v-model:contact="response[0].Metadata.data_footer[2].contacts"
      v-model:categories1s="response[0].Metadata.data_footer[3].categories1"
      v-model:categories2s="response[0].Metadata.data_footer[4].categories2"
      v-model:socialMedias="response[0].Metadata.data_footer[5].socialMedia"
      :response="response"
    />
  </div>
</template>

<script setup>
const { data } = await useFetch("http://localhost:8000/page");
const response = data.value;

const tschonchungtoi = ref(response[0]?.Metadata?.tschonchungtoi ?? {});
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
const data_footer = response[0]?.Metadata?.data_footer ?? {};
</script>
