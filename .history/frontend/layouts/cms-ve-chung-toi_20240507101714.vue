<template>
  <div class="w-full">
    <CmsTrangChuDefault :response="response[0]" />
    <CmsTrangChuHeader
      v-model:menu="response[0].Metadata.data_menus"
      v-model:form="response[0].Metadata.data_form"
      :response="response"
    />
    <div class="bg-[#F3F3F4]">
      <CmsVeChungToiTVeChungToi
        v-model:about="response[0].Metadata.data_vechungtoi"
        :response="response"
      />
      <CmsTrangChuThongSo
        v-model:item="response[0].Metadata.data_thongso"
        :response="response"
      />
      <CmsVeChungToiSuMenhVaTamNhin
        v-model:content="response[0].Metadata.data_sumenhvatamnhin"
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
const { data } = await useFetch("http://10.50.20.163:4000/page");
const response = data.value;

const data_form = ref(response[0]?.Metadata?.data_form ?? {});
const data_menus = ref(response[0]?.Metadata?.data_menus ?? {});
const data_sumenhvatamnhin = ref(
  response[0]?.Metadata?.data_sumenhvatamnhin ?? {}
);
const data_thongso = ref(response[0]?.Metadata?.data_thongso ?? {});
const data_vechungtoi = ref(response[0]?.Metadata?.data_vechungtoi ?? {});
const data_footer = response[0]?.Metadata?.data_footer ?? {};
</script>
