<template>
  <div class="w-[72%] m-auto mt-20 flex mb-12 relative">
    <div class="w-[35%]">
      <img class="w-36 h-24 ml-3 mr-32 mt-4 mb-5 flex" :src="logoUrl" />
      <div
        v-for="contact in contacts"
        :key="contact.id"
        class="flex items-center gap-2"
      >
        <img class="w-4" :src="contact.iconUrl" />
        <nuxt-link :to="contact.link" class="py-1.5 text-[#2c2934] text-base">{{
          contact.text
        }}</nuxt-link>
      </div>
    </div>
    <div v-for="category in categories1" :key="category.id" class="w-[18%]">
      <span class="py-2"
        ><b>{{ category.title }}</b></span
      >
      <div class="mt-1">
        <div v-for="item in category.items" :key="item.id" class="py-1 mt-2.5">
          <nuxt-link :to="item.link" class="text-base">{{
            item.text
          }}</nuxt-link>
        </div>
      </div>
    </div>
    <div v-for="category in categories2" :key="category.id" class="w-[25%]">
      <span class="py-2"
        ><b>{{ category.title }}</b></span
      >
      <div class="mt-1">
        <div v-for="item in category.items" :key="item.id" class="py-1 mt-2.5">
          <nuxt-link :to="item.link" class="text-base">{{
            item.text
          }}</nuxt-link>
        </div>
      </div>
    </div>
    <div class="w-[20%]">
      <span class="py-2"
        ><b>{{ content.content2 }}</b></span
      >
      <div class="flex mt-4">
        <div v-for="social in socialMedia" :key="social.id" class="px-2">
          <nuxt-link :to="social.link" class="">
            <img class="w-8" :src="social.icon" :alt="social.alt" />
          </nuxt-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  response: Object,
});

const logoUrl = ref([]);
const content = ref([]);
const contacts = ref([]);
const categories1 = ref([]);
const categories2 = ref([]);
const socialMedia = ref([]);

if (props.response && props.response[0]?.Metadata?.data_footer[0].logoUrl) {
  logoUrl.value = props.response[0].Metadata.data_footer[0].logoUrl;
}

if (props.response && props.response[0]?.Metadata?.data_footer[1].content) {
  content.value = props.response[0].Metadata.data_footer[1].content;
}

if (props.response && props.response[0]?.Metadata?.data_footer[2].contacts) {
  contacts.value = props.response[0].Metadata.data_footer[2].contacts;
}

if (props.response && props.response[0]?.Metadata?.data_footer[3].categories1) {
  categories1.value = props.response[0].Metadata.data_footer[3].categories1;
}

if (props.response && props.response[0]?.Metadata?.data_footer[4].categories2) {
  categories2.value = props.response[0].Metadata.data_footer[4].categories2;
}

if (props.response && props.response[0]?.Metadata?.data_footer[5].socialMedia) {
  socialMedia.value = props.response[0].Metadata.data_footer[5].socialMedia;
}

// const logoUrl = ref(
//   "https://cdn.mobiedu.vn/mskill/uploads/mb3/1698306009-mschool-240x132.png"
// );

// const content = ref({
//   content1: "Thông tin liên hệ",
//   content2: "Theo dõi chúng tôi",
// });

// const contacts = ref([
//   {
//     id: 1,
//     iconUrl: "https://cdn-icons-png.flaticon.com/128/2099/2099199.png",
//     text: "hello@mschool.com",
//     link: "",
//   },
//   {
//     id: 2,
//     iconUrl: "https://cdn-icons-png.flaticon.com/128/483/483947.png",
//     text: "+0123456789",
//     link: "",
//   },
//   {
//     id: 3,
//     iconUrl: "https://cdn-icons-png.flaticon.com/128/447/447031.png",
//     text: "247 Cầu Giấy, Dịch Vọng, Cầu Giấy, Hà Nội",
//     link: "",
//   },
// ]);

// const categories1 = ref([
//   {
//     title: "Khám phá",
//     items: [
//       { id: 1, text: "Về chúng tôi", link: "" },
//       { id: 2, text: "Khóa học", link: "" },
//       { id: 3, text: "Tin tức", link: "" },
//       { id: 4, text: "Liên hệ", link: "" },
//     ],
//   },
// ]);
// const categories2 = ref([
//   {
//     title: "Hướng dẫn",
//     items: [
//       { id: 1, text: "Điều khoản dịch vụ", link: "" },
//       { id: 2, text: "Chính sách đổi và hoàn tiền", link: "" },
//       { id: 3, text: "Chính sách bảo mật", link: "" },
//       { id: 4, text: "Chính sách thanh toán", link: "" },
//     ],
//   },
// ]);

// const socialMedia = ref([
//   {
//     id: 1,
//     icon: "https://cdn-icons-png.flaticon.com/128/15047/15047495.png",
//     alt: "Facebook",
//     link: "",
//   },
//   {
//     id: 2,
//     icon: "https://cdn-icons-png.flaticon.com/128/14417/14417460.png",
//     alt: "X",
//     link: "",
//   },
//   {
//     id: 3,
//     icon: "https://cdn-icons-png.flaticon.com/128/1384/1384014.png",
//     alt: "In",
//     link: "",
//   },
//   {
//     id: 4,
//     icon: "https://cdn-icons-png.flaticon.com/128/1384/1384028.png",
//     alt: "Youtube",
//     link: "",
//   },
// ]);
</script>
