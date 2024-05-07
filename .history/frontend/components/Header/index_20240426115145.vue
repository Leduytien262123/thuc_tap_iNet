<template>
  <div class="bg-[#F3F3F4] w-full mx-auto">
    <div v-for="item in form" :key="item.id" class="w-[98%] m-auto">
      <nuxt-link :to="item.linkTieuDe" class="mx-auto">
        <button
          class="border-none py-3 text-base w-full mx-auto bg-[#5236ff] text-white rounded-md mt-4 hover:bg-sky-500 flex items-center cursor-pointer"
        >
          <div class="mx-auto">{{ item.tieuDe }}</div>
        </button>
      </nuxt-link>
      <div class="w-full flex relative">
        <div class="w-[80%] mx-auto flex py-3">
          <div class="flex items-center justify-between w-full">
            <div class="flex items-center">
              <div class="logo ml-5">
                <img :src="item.logo" alt="Logo" class="w-40 h-20 px-4 py-4" />
              </div>
              <div class="text-base ml-8" v-for="menu in menus" :key="menu.id">
                <nuxt-link :to="menu.link">{{ menu.name }}</nuxt-link>
              </div>
              <div class="ml-5 flex">
                <div class="text-base mt-1">
                  <nuxt-link to="#">
                    <img
                      class="w-5 h-5"
                      src="https://cdn-icons-png.flaticon.com/512/1170/1170627.png"
                    />
                  </nuxt-link>
                </div>
              </div>
            </div>
            <div class="ml-4 mr-10">
              <nuxt-link :to="item.linkDangNhap" class="text-base mr-7">{{
                item.dangNhap
              }}</nuxt-link>
              <nuxt-link
                :to="item.linkDangKy"
                class="bg-[#5236ff] py-3 px-7 border-none rounded-md cursor-pointer text-white hover:bg-sky-500"
                >{{ item.dangKy }}</nuxt-link
              >
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  response: Object,
});

const form = ref([]);
const menus = ref([]);

if (props.response && props.response[0]?.Metadata?.data_form) {
  form.value = props.response[0].Metadata.data_form;
}

if (props.response && props.response[0]?.Metadata?.data_menus) {
  menus.value = props.response[0].Metadata.data_menus;
}

// const { data } = await useFetch("http://localhost:8000/page");
// const header = data.value;

// let form = ref("");
// let menus = ref("");

// header.forEach((item) => {
//   if (item.Metadata) {
//     if (item.Metadata.data_form) {
//       form = item.Metadata.data_form;
//     }
//     if (item.Metadata.data_menus) {
//       menus = item.Metadata.data_menus;
//     }
//   }
// });

// console.log(form);
// console.log(menus);
</script>
