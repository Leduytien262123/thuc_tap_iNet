<template>
  <div
    class="w-[98%] sticky top-0 z-[100] m-auto bg-[#3A434E] mt-auto rounded-md p-2 flex items-center"
  >
    <div class="w-1/3">
      <select
        id="pageSelect"
        class="p-2 rounded-md w-[45%] border border-gray-300 ml-2"
        v-model="selectedPage"
      >
        <option
          v-for="(option, index) in options"
          :key="option"
          :value="option"
          :disabled="index == 0"
        >
          {{ option.text }}
        </option>
      </select>
      <select
        class="rounded-md p-2 ml-3 w-[45%] border border-gray-300"
        v-model="selectedSubPage"
      >
        <option v-for="option in options" :key="option" :value="option">
          {{ option.text }}
        </option>
      </select>
    </div>
    <div class="p-2 w-[40%] flex m-auto">
      <div class="m-auto">
        <button class="">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            viewBox="0 0 320 512"
            class="bg-blue-500 text-white w-9 h-8 p-2 rounded-lg hover:bg-sky-300"
          >
            <path
              d="M272 0H48C21.5 0 0 21.5 0 48v416c0 26.5 21.5 48 48 48h224c26.5 0 48-21.5 48-48V48c0-26.5-21.5-48-48-48zM160 480c-17.7 0-32-14.3-32-32s14.3-32 32-32s32 14.3 32 32zm112-108c0 6.6-5.4 12-12 12H60c-6.6 0-12-5.4-12-12V60c0-6.6 5.4-12 12-12h200c6.6 0 12 5.4 12 12v312z"
              fill="currentColor"
            ></path>
          </svg>
        </button>
        <button class="ml-3">
          <svg
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            x="0px"
            y="0px"
            viewBox="0 0 512 512"
            enable-background="new 0 0 512 512"
            xml:space="preserve"
            class="bg-blue-500 text-white w-9 h-8 py-2 px-0.5 rounded-lg hover:bg-sky-300"
          >
            <g>
              <g>
                <path
                  fill="white"
                  width="10px"
                  d="M409,39c-4.5-4.5-10.6-7-16.9-7H119.9c-6.4,0-12.4,2.5-16.9,7l0,0c-4.5,4.5-7,10.6-7,16.9v400.1
			c0,6.4,2.5,12.4,7,16.9l0,0c4.5,4.5,10.6,7,16.9,7h272.1c6.4,0,12.4-2.5,16.9-7l0,0c4.5-4.5,7-10.6,7-16.9V55.9
			C416,49.6,413.5,43.5,409,39L409,39z M255.6,48.7c3.9,0,7,3.1,7,7c0,3.9-3.1,7-7,7c-3.9,0-7-3.1-7-7
			C248.6,51.9,251.8,48.7,255.6,48.7z M256,470c-7.7,0-14-6.5-14-14.1c0-7.5,6.2-14,14-14c7.7,0,14.1,6.4,14.1,14
			C270,463.5,263.7,470,256,470z M400,432H112V80h288V432z"
                ></path>
              </g>
            </g>
          </svg>
        </button>
        <button class="ml-3">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            viewBox="0 0 24 24"
            class="bg-blue-500 text-white w-9 h-8 p-2 rounded-lg hover:bg-sky-300"
          >
            <path
              d="M23 18h-1V5c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v13H1c-.55 0-1 .45-1 1s.45 1 1 1h22c.55 0 1-.45 1-1s-.45-1-1-1zm-9.5 0h-3c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h3c.28 0 .5.22.5.5s-.22.5-.5.5zm6.5-3H4V6c0-.55.45-1 1-1h14c.55 0 1 .45 1 1v9z"
              fill="currentColor"
            ></path>
          </svg>
        </button>
      </div>
    </div>
    <div class="w-[20%] flex justify-end">
      <div class="mr-3">
        <button
          class="py-1.5 px-3 bg-green-500 text-white rounded-md transition-transform transform-gpu motion-safe:hover:scale-110"
          @click="saveData"
        >
          Save
        </button>
      </div>
      <div class="mr-3">
        <button
          class="py-1.5 px-3 bg-green-500 rounded-md text-white transition-transform transform-gpu motion-safe:hover:scale-110"
          @click="saveAndExit"
        >
          Save & Exit
        </button>
      </div>
      <div class="mr-2">
        <button
          class="py-1.5 px-3 bg-red-500 text-white rounded-md transition-transform transform-gpu motion-safe:hover:scale-110"
          @click="cancel"
        >
          <nuxt-link to="/">Cancel</nuxt-link>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

// const message = useMessage();

const options = [
  { value: "", text: "Chọn danh sách page" },
  { value: "home", text: "Trang chủ" },
  { value: "banner", text: "Banner" },
];

const selectedPage = ref({ value: "", text: "Chọn danh sách page" });
const selectedSubPage = ref(options[1]);

const subPageOptions = ["Trang chủ"];

// const { data } = await useFetch("http://localhost:8000/page");

// const datas = data.value;

// const form = ref("");
// const menus = ref("");

// datas.forEach((item) => {
//   if (item.Metadata) {
//     form.value = item.Metadata.data_form;
//     menus.value = item.Metadata.data_menus;
//   }
// });

const saveData = async () => {
  const responsePatch = await useFetch("http://localhost:8000/page", {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json",
    },
    body: {
      key: "home_page",
      Label: "Trang",
      Link: "/",
      Metadata: {
        data_form: [
          {
            id: 1,
            logo: "https://static.ladipage.net/5b33575d031f5281797296b9/mschool-20201019033206.png",
            dangKy: "Đăng ký",
            tieuDe: "Khóa học giao tiếp Ưu đãi sắp kết thúc. Mua ngay",
            dangNhap: "Đăng nhập",
            linkDangKy: "",
            linkTieuDe: "",
            linkDangNhap: "",
          },
        ],
        data_menus: [
          {
            id: 1,
            link: "/",
            name: "Trang chủ",
          },
          {
            id: 2,
            link: "",
            name: "Tin tức",
          },
          {
            id: 3,
            link: "/khoa-hoc",
            name: "Khóa học",
          },
          {
            id: 4,
            link: "",
            name: "Kích hoạt khóa học",
          },
          {
            id: 5,
            link: "",
            name: "Liên hệ",
          },
        ],
        data_banner: {
          title: "Học online",
          button1: "Đăng ký ngay",
          button2: "Xem khóa học",
          subtitle: "thật dễ cùng mSchool",
          description: "Chương trình tiên tiến - Cải thiện kỹ năng",
          linkButton1: "",
          linkButton2: "",
          button1Color: "#5236FF",
          button2Color: "#FFFFFF",
          backgroundImage:
            "https://mobifonehanoi.vn/medias/634/1652839097_1618902872287_m_school.jpg",
        },
        data_footer: [
          {
            logoUrl:
              "https://cdn.mobiedu.vn/mskill/uploads/mb3/1698306009-mschool-240x132.png",
          },
          {
            content: {
              content1: "Thông tin liên hệ",
              content2: "Theo dõi chúng tôi",
            },
          },
          {
            contacts: [
              {
                id: 1,
                link: "",
                text: "hello@mschool.com",
                iconUrl:
                  "https://cdn-icons-png.flaticon.com/128/2099/2099199.png",
              },
              {
                id: 2,
                link: "",
                text: "+0123456789",
                iconUrl:
                  "https://cdn-icons-png.flaticon.com/128/483/483947.png",
              },
              {
                id: 3,
                link: "",
                text: "247 Cầu Giấy, Dịch Vọng, Cầu Giấy, Hà Nội",
                iconUrl:
                  "https://cdn-icons-png.flaticon.com/128/447/447031.png",
              },
            ],
          },
          {
            categories1: [
              {
                items: [
                  {
                    id: 1,
                    link: "",
                    text: "Về chúng tôi",
                  },
                  {
                    id: 2,
                    link: "",
                    text: "Khóa học",
                  },
                  {
                    id: 3,
                    link: "",
                    text: "Tin tức",
                  },
                  {
                    id: 4,
                    link: "",
                    text: "Liên hệ",
                  },
                ],
                title: "Khám phá",
              },
            ],
          },
          {
            categories2: [
              {
                items: [
                  {
                    id: 1,
                    link: "",
                    text: "Điều khoản dịch vụ",
                  },
                  {
                    id: 2,
                    link: "",
                    text: "Chính sách đổi và hoàn tiền",
                  },
                  {
                    id: 3,
                    link: "",
                    text: "Chính sách bảo mật",
                  },
                  {
                    id: 4,
                    link: "",
                    text: "Chính sách thanh toán",
                  },
                ],
                title: "Hướng dẫn",
              },
            ],
          },
          {
            socialMedia: [
              {
                id: 1,
                alt: "Facebook",
                icon: "https://cdn-icons-png.flaticon.com/128/15047/15047495.png",
                link: "",
              },
              {
                id: 2,
                alt: "X",
                icon: "https://cdn-icons-png.flaticon.com/128/14417/14417460.png",
                link: "",
              },
              {
                id: 3,
                alt: "In",
                icon: "https://cdn-icons-png.flaticon.com/128/1384/1384014.png",
                link: "",
              },
              {
                id: 4,
                alt: "Youtube",
                icon: "https://cdn-icons-png.flaticon.com/128/1384/1384028.png",
                link: "",
              },
            ],
          },
        ],
        data_thongso: [
          {
            id: 1,
            label: "Giảng viên",
            value: 300,
          },
          {
            id: 2,
            label: "Học viên",
            value: 20000,
          },
          {
            id: 3,
            label: "Bài giảng",
            value: 10000,
          },
        ],
        baivietnoibat: {
          title: "Bài viết",
          subTitle: "nổi bật",
        },
        khoahocnoibat: {
          title: "Khóa học",
          subTitle: "nổi bật",
        },
        danhmuckhoahoc: {
          title: "Danh mục",
          subTitle: "khóa học",
        },
        doitactieubieu: {
          title: "Đối tác",
          subTitle: "tiêu biểu",
        },
        tschonchungtoi: {
          title: "Tại sao",
          subTitle: "chọn chúng tôi",
        },
        data_vechungtoi: [
          {
            id: 1,
            img1: "https://media-cdn-v2.laodong.vn/storage/newsportal/2023/8/26/1233821/Giai-Nhi-1--Nang-Tre.jpg",
            img2: "https://d1hjkbq40fs2x4.cloudfront.net/2016-01-31/files/1045-2.jpg",
            name: "Về chúng tôi",
            text: "mSchool giúp nhà trường, giáo viên tạo website giảng dạy online, với hệ thống lớp học ảo, tổ chức thi online và quản lý trường học toàn diện. Giải pháp tạo website dạy học dễ dàng - bảo mật chỉ và tùy biến theo nhu cầu khách hàng.",
            title: "mSchool - Giải pháp dạy & học toàn diện",
          },
        ],
        giangvientieubieu: {
          title: "Giảng viên",
          subTitle: "tiêu biểu",
        },
        hvnoigivechungtoi: {
          title: "Học viên nói gì",
          subTitle: "về chúng tôi",
        },
        data_baivietnoibat: [
          {
            id: 1,
            link: "",
            tag1: "",
            tag2: "",
            image:
              "https://image.luatvietnam.vn/uploaded/1200x675twebp/images/original/2020/12/15/cach-xep-loai-hoc-luc-cap-3_1512095022.jpg",
            nameTitle: "10 Top tips for marking your Saas product sticky",
            description:
              "It is a long established fact that a reader will be distracted by the readable content of a page from when looking at it layout. The point of using Lorem Ipsum",
          },
          {
            id: 2,
            link: "",
            tag1: "Marketing",
            tag2: "Analysis",
            image:
              "https://image.luatvietnam.vn/uploaded/1200x675twebp/images/original/2020/12/15/cach-xep-loai-hoc-luc-cap-3_1512095022.jpg",
            nameTitle: "10 Top tips for marking your Saas product sticky",
            description:
              "It is a long established fact that a reader will be distracted by the readable content of a page from when looking at it layout. The point of using Lorem Ipsum",
          },
        ],
        data_khoahocnoibat: [
          {
            id: 1,
            link: "",
            name: "Lê Duy Tiến",
            image:
              "https://cdn.tgdd.vn/Products/Images/44/313619/Slider/vi-vn-lenovo-yoga-7-14irl8-i7-82yl006bvn-slider-6.jpg",
            price: "199000",
            avatar:
              "https://img.lovepik.com/free-png/20211107/lovepik-confident-workplace-business-male-image-png-image_400441790_wh1200.png",
            lessons: 20,
            oldPrice: "500000",
            nameTitle: "Lập trình hướng đối tượng",
          },
          {
            id: 2,
            link: "",
            name: "Lê Duy Tiến",
            image:
              "https://cdn.tgdd.vn/Products/Images/44/313619/Slider/vi-vn-lenovo-yoga-7-14irl8-i7-82yl006bvn-slider-6.jpg",
            price: "0",
            avatar:
              "https://img.lovepik.com/free-png/20211107/lovepik-confident-workplace-business-male-image-png-image_400441790_wh1200.png",
            lessons: 20,
            oldPrice: "500000",
            nameTitle: "Lập trình hướng đối tượng",
          },
          {
            id: 3,
            link: "",
            name: "Lê Duy Tiến",
            image:
              "https://cdn.tgdd.vn/Products/Images/44/313619/Slider/vi-vn-lenovo-yoga-7-14irl8-i7-82yl006bvn-slider-6.jpg",
            price: "199000",
            avatar:
              "https://img.lovepik.com/free-png/20211107/lovepik-confident-workplace-business-male-image-png-image_400441790_wh1200.png",
            lessons: 20,
            oldPrice: "500000",
            nameTitle: "Lập trình hướng đối tượng",
          },
          {
            id: 4,
            link: "",
            name: "Lê Duy Tiến",
            image:
              "https://cdn.tgdd.vn/Products/Images/44/313619/Slider/vi-vn-lenovo-yoga-7-14irl8-i7-82yl006bvn-slider-6.jpg",
            price: "0",
            avatar:
              "https://img.lovepik.com/free-png/20211107/lovepik-confident-workplace-business-male-image-png-image_400441790_wh1200.png",
            lessons: 20,
            oldPrice: "500000",
            nameTitle: "Lập trình hướng đối tượng",
          },
          {
            id: 5,
            link: "",
            name: "Lê Duy Tiến",
            image:
              "https://cdn.tgdd.vn/Products/Images/44/313619/Slider/vi-vn-lenovo-yoga-7-14irl8-i7-82yl006bvn-slider-6.jpg",
            price: "199000",
            avatar:
              "https://img.lovepik.com/free-png/20211107/lovepik-confident-workplace-business-male-image-png-image_400441790_wh1200.png",
            lessons: 20,
            oldPrice: "500000",
            nameTitle: "Lập trình hướng đối tượng",
          },
          {
            id: 6,
            link: "",
            name: "Lê Duy Tiến",
            image:
              "https://cdn.tgdd.vn/Products/Images/44/313619/Slider/vi-vn-lenovo-yoga-7-14irl8-i7-82yl006bvn-slider-6.jpg",
            price: "0",
            avatar:
              "https://img.lovepik.com/free-png/20211107/lovepik-confident-workplace-business-male-image-png-image_400441790_wh1200.png",
            lessons: 20,
            oldPrice: "500000",
            nameTitle: "Lập trình hướng đối tượng",
          },
          {
            id: 7,
            link: "",
            name: "Lê Duy Tiến",
            image:
              "https://cdn.tgdd.vn/Products/Images/44/313619/Slider/vi-vn-lenovo-yoga-7-14irl8-i7-82yl006bvn-slider-6.jpg",
            price: "199000",
            avatar:
              "https://img.lovepik.com/free-png/20211107/lovepik-confident-workplace-business-male-image-png-image_400441790_wh1200.png",
            lessons: 20,
            oldPrice: "500000",
            nameTitle: "Lập trình hướng đối tượng",
          },
          {
            id: 8,
            link: "",
            name: "Lê Duy Tiến",
            image:
              "https://cdn.tgdd.vn/Products/Images/44/313619/Slider/vi-vn-lenovo-yoga-7-14irl8-i7-82yl006bvn-slider-6.jpg",
            price: "0",
            avatar:
              "https://img.lovepik.com/free-png/20211107/lovepik-confident-workplace-business-male-image-png-image_400441790_wh1200.png",
            lessons: 20,
            oldPrice: "500000",
            nameTitle: "Lập trình hướng đối tượng",
          },
        ],
        goikhoahoctieubieu: {
          title: "Gói khóa học",
          subTitle: "tiêu biểu",
        },
        data_danhmuckhoahoc: [
          {
            id: 1,
            icon: "/icon/marketing.png",
            link: "",
            name1: "Digtal",
            name2: "Marketing",
          },
          {
            id: 2,
            icon: "/icon/development.png",
            link: "",
            name1: "Web",
            name2: "Development",
          },
          {
            id: 3,
            icon: "/icon/art.png",
            link: "",
            name1: "Art &",
            name2: "Humanities",
          },
          {
            id: 4,
            icon: "/icon/personal.png",
            link: "",
            name1: "Personal",
            name2: "Development",
          },
          {
            id: 5,
            icon: "/icon/software.png",
            link: "",
            name1: "IT and",
            name2: "Software",
          },
          {
            id: 6,
            icon: "/icon/marketing.png",
            link: "",
            name1: "Digtal",
            name2: "Marketing",
          },
          {
            id: 7,
            icon: "/icon/development.png",
            link: "",
            name1: "Web",
            name2: "Development",
          },
          {
            id: 8,
            icon: "/icon/art.png",
            link: "",
            name1: "Art &",
            name2: "Humanities",
          },
          {
            id: 9,
            icon: "/icon/personal.png",
            link: "",
            name1: "Personal",
            name2: "Development",
          },
          {
            id: 10,
            icon: "/icon/software.png",
            link: "",
            name1: "IT and",
            name2: "Software",
          },
        ],
        data_doitactieubieu: [
          {
            id: 1,
            link: "",
            name: "Zapier",
            image:
              "https://e7.pngegg.com/pngimages/296/156/png-clipart-zapier-black-logo-tech-companies.png",
          },
          {
            id: 2,
            link: "",
            name: "Spotify",
            image:
              "https://w7.pngwing.com/pngs/803/51/png-transparent-spotify-hd-logo.png",
          },
          {
            id: 3,
            link: "",
            name: "Zoom",
            image:
              "https://joyfullearningcenter.com/wp-content/uploads/2020/05/zoom-logo.png",
          },
          {
            id: 4,
            link: "",
            name: "Amazon",
            image:
              "https://upload.wikimedia.org/wikipedia/commons/a/a9/Amazon_logo.svg",
          },
          {
            id: 5,
            link: "",
            name: "Adobe",
            image:
              "https://upload.wikimedia.org/wikipedia/commons/8/8d/Adobe_Corporate_Logo.png",
          },
          {
            id: 6,
            link: "",
            name: "Notion",
            image:
              "https://i0.wp.com/get.site/wp-content/uploads/2021/10/notion-logo.png?ssl=1",
          },
          {
            id: 7,
            link: "",
            name: "Netflix",
            image:
              "https://logolook.net/wp-content/uploads/2021/11/Netflix-Logo.png",
          },
        ],
        data_tschonchungtoi: [
          {
            id: 1,
            alt: "Bóng đèn",
            link: "",
            imageSrc: "/img/bongden.png",
            nameTitle: "Update mới nhất",
            description:
              "Long established fact that a reader will be distracted by the readable content of a page when looking at its layout",
          },
          {
            id: 2,
            alt: "Bài giảng",
            link: "",
            imageSrc: "/img/baigiang.png",
            nameTitle: "Bài giảng hấp dẫn",
            description:
              "Long established fact that a reader will be distracted by the readable content of a page when looking at its layout",
          },
          {
            id: 3,
            alt: "Giảng viên",
            link: "",
            imageSrc: "/img/giangvien.png",
            nameTitle: "Giảng viên chất lượng",
            description:
              "Long established fact that a reader will be distracted by the readable content of a page when looking at its layout",
          },
          {
            id: 4,
            alt: "Nội dung",
            link: "",
            imageSrc: "/img/noidung.png",
            nameTitle: "Nội dung đa dạng",
            description:
              "Long established fact that a reader will be distracted by the readable content of a page when looking at its layout",
          },
        ],
        data_khoahoctieubieu: [
          {
            id: 1,
            link: "",
            name: "Nguyễn Tú Nam",
            image: "/img/hoctap.jpg",
            price: "199000",
            oldPrice: "500000",
            nameTitle: "Web Design",
            description:
              "Learn the fundamentals of web design, including HTML, CSS, and responsive design principles. Develop the skills to create visually appealing and user-friendly websites.",
          },
          {
            id: 2,
            link: "",
            name: "Nguyễn Tú Nam",
            image: "/img/hoctap.jpg",
            price: "199000",
            oldPrice: "500000",
            nameTitle: "Web Design",
            description:
              "Learn the fundamentals of web design, including HTML, CSS, and responsive design principles. Develop the skills to create visually appealing and user-friendly websites.",
          },
          {
            id: 3,
            link: "",
            name: "Nguyễn Tú Nam",
            image: "/img/hoctap.jpg",
            price: "199000",
            oldPrice: "500000",
            nameTitle: "Web Design",
            description:
              "Learn the fundamentals of web design, including HTML, CSS, and responsive design principles. Develop the skills to create visually appealing and user-friendly websites.",
          },
        ],
        data_sumenhvatamnhin: [
          {
            id: 1,
            img: "https://www.vietnamworks.com/hrinsider/wp-content/uploads/2023/12/hinh-anh-thien-nhien-3d-tuyet-dep-003.jpg",
            name: "Sứ mệnh",
            text: "<div>Lorem ipsum dolor sit amet, consectetur adipiscing Pharetra Ld eu aliquet diam lorem viverra at justo. Nulla odio nequefjf gravida in pharetra egestas. Ac id sagittis at morbi interdum nibh diam sagittis et. </div><br><div>Lorem ipsum dolor sit amet, consectetur adipiscing Pharetra id eu aliquet diam lorem viverra at justo. Nulla odio nequesg gravida in pharetra egestas.</div>",
          },
          {
            id: 2,
            img: "https://www.vietnamworks.com/hrinsider/wp-content/uploads/2023/12/hinh-anh-thien-nhien-3d-tuyet-dep-003.jpg",
            name: "Tầm nhìn",
            text: "<div>Lorem ipsum dolor sit amet, consectetur adipiscing Pharetra Ld eu aliquet diam lorem viverra at justo. Nulla odio nequefjf gravida in pharetra egestas. Ac id sagittis at morbi interdum nibh diam sagittis et. </div><br><div>Lorem ipsum dolor sit amet, consectetur adipiscing Pharetra id eu aliquet diam lorem viverra at justo. Nulla odio nequesg gravida in pharetra egestas.</div>",
          },
        ],
        data_giangvientieubieu: [
          {
            id: 1,
            name: "Nguyễn Tú Nam",
            image:
              "https://htmediagroup.vn/wp-content/uploads/2021/06/Anh-doanh-nhan-19.jpg",
            iconFb: "https://cdn-icons-png.flaticon.com/128/15047/15047495.png",
            iconIg: "/icon/logo-instagram.svg",
            iconIn: "/icon/logo-linkedin.svg",
            linkFb: "",
            linkIg: "",
            linkIn: "",
            description: "Trung tâm phát triển Kỹ năng sống",
          },
          {
            id: 2,
            name: "Nguyễn Tú Nam",
            image:
              "https://htmediagroup.vn/wp-content/uploads/2021/06/Anh-doanh-nhan-19.jpg",
            iconFb: "https://cdn-icons-png.flaticon.com/128/15047/15047495.png",
            iconIg: "/icon/logo-instagram.svg",
            iconIn: "/icon/logo-linkedin.svg",
            linkFb: "",
            linkIg: "",
            linkIn: "",
            description: "Trung tâm phát triển Kỹ năng sống",
          },
          {
            id: 3,
            name: "Nguyễn Tú Nam",
            image:
              "https://htmediagroup.vn/wp-content/uploads/2021/06/Anh-doanh-nhan-19.jpg",
            iconFb: "https://cdn-icons-png.flaticon.com/128/15047/15047495.png",
            iconIg: "/icon/logo-instagram.svg",
            iconIn: "/icon/logo-linkedin.svg",
            linkFb: "",
            linkIg: "",
            linkIn: "",
            description: "Trung tâm phát triển Kỹ năng sống",
          },
          {
            id: 4,
            name: "Nguyễn Tú Nam",
            image:
              "https://htmediagroup.vn/wp-content/uploads/2021/06/Anh-doanh-nhan-19.jpg",
            iconFb: "https://cdn-icons-png.flaticon.com/128/15047/15047495.png",
            iconIg: "/icon/logo-instagram.svg",
            iconIn: "/icon/logo-linkedin.svg",
            linkFb: "",
            linkIg: "",
            linkIn: "",
            description: "Trung tâm phát triển Kỹ năng sống",
          },
          {
            id: 5,
            name: "Nguyễn Tú Nam",
            image:
              "https://htmediagroup.vn/wp-content/uploads/2021/06/Anh-doanh-nhan-19.jpg",
            iconFb: "https://cdn-icons-png.flaticon.com/128/15047/15047495.png",
            iconIg: "/icon/logo-instagram.svg",
            iconIn: "/icon/logo-linkedin.svg",
            linkFb: "",
            linkIg: "",
            linkIn: "",
            description: "Trung tâm phát triển Kỹ năng sống",
          },
        ],
        data_hvnoigivechungtoi: [
          {
            id: 1,
            name: "Lê Minh Anh",
            image:
              "https://bizweb.dktcdn.net/100/438/408/files/anh-chan-dung-dep-yodyvn1.jpg?v=1683537734987",
            content:
              "The web design course provided a solid foundation for me. The instructors were knowledgeable and supportive, and the interactive learning environment was engaging. I highly recommend it!",
          },
          {
            id: 2,
            name: "Lê Minh Tuấn",
            image:
              "https://bizweb.dktcdn.net/100/438/408/files/anh-chan-dung-dep-yodyvn1.jpg?v=1683537734987",
            content:
              "The web design course provided a solid foundation for me. The instructors were knowledgeable and supportive, and the interactive learning environment was engaging. I highly recommend it!",
          },
          {
            id: 3,
            name: "Lê Minh Khôi",
            image:
              "https://bizweb.dktcdn.net/100/438/408/files/anh-chan-dung-dep-yodyvn1.jpg?v=1683537734987",
            content:
              "The web design course provided a solid foundation for me. The instructors were knowledgeable and supportive, and the interactive learning environment was engaging. I highly recommend it!",
          },
          {
            id: 4,
            name: "Lê Minh Anh",
            image:
              "https://bizweb.dktcdn.net/100/438/408/files/anh-chan-dung-dep-yodyvn1.jpg?v=1683537734987",
            content:
              "The web design course provided a solid foundation for me. The instructors were knowledgeable and supportive, and the interactive learning environment was engaging. I highly recommend it!",
          },
        ],
        data_goikhoahoctieubieu: [
          {
            id: 1,
            link: "",
            name: "Nguyễn Tú Nam",
            image: "/img/hoctap.jpg",
            price: "199000",
            oldPrice: "500000",
            nameTitle: "Web Design",
            description:
              "Learn the fundamentals of web design, including HTML, CSS, and responsive design principles. Develop the skills to create visually appealing and user-friendly websites.",
          },
          {
            id: 2,
            link: "",
            name: "Nguyễn Tú Nam",
            image: "/img/hoctap.jpg",
            price: "199000",
            oldPrice: "500000",
            nameTitle: "Web Design",
            description:
              "Learn the fundamentals of web design, including HTML, CSS, and responsive design principles. Develop the skills to create visually appealing and user-friendly websites.",
          },
          {
            id: 3,
            link: "",
            name: "Nguyễn Tú Nam",
            image: "/img/hoctap.jpg",
            price: "199000",
            oldPrice: "500000",
            nameTitle: "Web Design",
            description:
              "Learn the fundamentals of web design, including HTML, CSS, and responsive design principles. Develop the skills to create visually appealing and user-friendly websites.",
          },
        ],
      },
    },
  });
};

const saveAndExit = () => {};

const cancel = () => {};

watch(
  () => selectedPage.value,
  (newValue) => {
    const result = options.find((o) => o.value == newValue.value);
    selectedSubPage.value = result;
  }
);
</script>

<style>
select:focus {
  outline: none;
  border: 2px solid #007bff;
  box-shadow: 0 0 5px #007bff;
}
</style>
