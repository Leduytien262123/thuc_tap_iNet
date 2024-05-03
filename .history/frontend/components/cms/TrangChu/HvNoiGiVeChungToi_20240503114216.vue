<template>
  <div class="w-[72%] relative m-auto mb-[100px]">
    <div class="flex items-center justify-between mb-8">
      <h2 class="text-[40px] flex font-semibold">
        {{ title }}
        <p class="ml-3 text-[#6046FF]">{{ subTitle }}</p>
      </h2>
    </div>
    <div class="w-full">
      <div class="w-full mx-auto flex justify-between items-center">
        <div class="w-full relative flex justify-between items-center">
          <button @click="mySwiper.slidePrev()">
            <img
              class="absolute -left-16 top-1/2 -translate-y-1/2 w-6 h-6 opacity-50 hover:opacity-100"
              src="/icons/chevron-back-outline.svg"
            />
          </button>
          <Swiper
            :modules="modules"
            @swiper="(s) => (mySwiper = s)"
            :slides-per-view="1"
            :breakpoints="{
              300: {
                slidesPerView: 1.175,
              },
              400: {
                slidesPerView: 1.515,
              },
              768: {
                slidesPerView: 3,
                spaceBetween: 20,
              },
              1024: {
                slidesPerView: 3,
              },
            }"
            :space-between="25"
            loop
            :autoplay="{
              delay: 2000,
              disableOnInteraction: false,
              pauseOnMouseEnter: true,
            }"
            :speed="1000"
            class="!z-0 rounded-lg"
          >
            <SwiperSlide
              v-for="(testimonial, index) of testimonials"
              :key="index"
            >
              <div class="w-full flex mb-6 justify-between">
                <div class="bg-white rounded-lg justify-center items-center">
                  <div class="w-[78%] m-auto mt-8 mb-8">
                    <p>{{ testimonial.content }}</p>
                  </div>
                  <hr />
                  <div class="w-[78%] m-auto mt-6 mb-6 flex items-center">
                    <div class="w-[50px] h-[55px] rounded-lg overflow-hidden">
                      <img
                        class="w-full h-full rounded-lg object-cover"
                        :src="testimonial.image"
                      />
                    </div>
                    <b class="text-lg ml-4">{{ testimonial.name }}</b>
                  </div>
                </div>
              </div>
            </SwiperSlide>
          </Swiper>
          <button @click="mySwiper.slideNext()">
            <img
              class="absolute -right-16 top-1/2 -translate-y-1/2 w-6 h-6 opacity-50 hover:opacity-100"
              src="/icons/chevron-forward-outline.svg"
            />
          </button>
        </div>
      </div>
    </div>
    <div class="absolute -right-[165px] top-1/2 -translate-y-1/2 mt-11">
      <n-button
        @click="editMode = true"
        type="success"
        class="bg-green-600 px-6 py-3"
      >
        Sửa
      </n-button>
    </div>
    <n-modal v-model:show="editMode">
      <n-card
        style="width: 600px"
        title="Chỉnh sửa"
        :bordered="false"
        size="huge"
        role="dialog"
        aria-modal="true"
      >
        <template #header-extra> {{ title + " " + subTitle }} </template>
        <div class="w-full">
          <div class="bg-white p-8 pr-0 rounded-lg">
            <h2 class="text-2xl font-semibold mb-4">Chỉnh sửa</h2>
            <div class="mb-4">
              <label for="editTitle" class="block font-semibold"
                >Tiêu đề chính:</label
              >
              <input
                type="text"
                id="editTitle"
                class="w-full border rounded p-2"
                v-model="editData.editTitle"
                @keydown.enter.prevent="handleEnterKey"
              />
            </div>
            <div class="mb-4">
              <label for="editSubTitle" class="block font-semibold"
                >Tiêu đề phụ:</label
              >
              <input
                type="text"
                id="editSubTitle"
                class="w-full border rounded p-2"
                v-model="editData.editSubTitle"
                @keydown.enter.prevent="handleEnterKey"
              />
            </div>
            <n-collapse
              arrow-placement="right"
              v-for="item in editData.editTestimonials"
              :key="item.id"
              class="mt-3 flex items-center"
            >
              <n-collapse-item
                :title="'Chỉnh sửa nội dung học viên:' + ' ' + item.editName"
                :name="item.id"
                class="w-full"
              >
                <div class="mb-4">
                  <label :for="'editName' + item.id" class="block font-semibold"
                    >Tên học viên</label
                  >
                  <input
                    :id="'editName' + item.id"
                    class="w-full border rounded p-2"
                    v-model="item.editName"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label for="editImage" class="block font-semibold"
                    >Ảnh:</label
                  >
                  <input
                    type="file"
                    id="editImage"
                    accept="image/*"
                    class="w-full border rounded p-2 border-gray-500"
                    @change="handleEditImage(item, $event)"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label
                    :for="'editContent' + item.id"
                    class="block font-semibold"
                    >Nội dung:</label
                  >
                  <textarea
                    :id="'editContent' + item.id"
                    class="w-full border rounded p-2"
                    v-model="item.editContent"
                    @keydown.enter.prevent="handleEnterKey"
                  >
                  </textarea>
                </div>
              </n-collapse-item>
              <button @click="remove(item.id)" class="self-start ml-auto">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  xmlns:xlink="http://www.w3.org/1999/xlink"
                  viewBox="0 0 48 48"
                  class="w-8 h-8 text-red-500 items-center ml-2 mt-1"
                >
                  <g fill="none">
                    <path
                      d="M24 6.75a6.25 6.25 0 0 1 6.246 6.02l.004.231L37 13a1.75 1.75 0 0 1 .144 3.494L37 16.5h-1.167l-1.627 21.57A4.25 4.25 0 0 1 29.968 42H18.032a4.25 4.25 0 0 1-4.238-3.93L12.166 16.5H11a1.75 1.75 0 0 1-1.744-1.607l-.006-.143a1.75 1.75 0 0 1 1.607-1.744L11 13h6.75c0-3.298 2.555-6 5.794-6.234l.227-.012L24 6.75zm3.75 13a1.25 1.25 0 0 0-1.244 1.122L26.5 21v12l.006.128a1.25 1.25 0 0 0 2.488 0L29 33V21l-.006-.128a1.25 1.25 0 0 0-1.244-1.122zm-7.5 0a1.25 1.25 0 0 0-1.244 1.122L19 21v12l.006.128a1.25 1.25 0 0 0 2.488 0L21.5 33V21l-.006-.128a1.25 1.25 0 0 0-1.244-1.122zm3.918-9.495L24 10.25a2.75 2.75 0 0 0-2.745 2.582l-.005.169l5.5-.001a2.75 2.75 0 0 0-2.582-2.745z"
                      fill="currentColor"
                    ></path>
                  </g>
                </svg>
              </button>
              <hr class="border-black my-2" />
            </n-collapse>
          </div>
        </div>

        <template #footer>
          <div class="text-right mt-10">
            <button
              class="px-4 py-2 bg-blue-400 text-white rounded hover:bg-blue-600"
              @click="saveChanges"
            >
              Lưu
            </button>
            <button
              class="px-4 py-2 bg-gray-300 text-gray-800 rounded ml-2 hover:bg-gray-500"
              @click="cancelEdit"
            >
              Hủy
            </button>
          </div>
          <div class="mt-8">
            <n-button
              class="bg-green-600 px-6 py-3"
              @click="editAdd = true"
              type="success"
            >
              Thêm
            </n-button>
          </div>
        </template>

        <div class="">
          <n-modal v-model:show="editAdd">
            <n-card
              style="width: 600px"
              title="Thêm mới"
              :bordered="false"
              size="huge"
              role="dialog"
              aria-modal="true"
            >
              <template #header-extra>
                {{ title + " " + subTitle }}
              </template>
              <div class="h-full pr-8 overflow-y-auto">
                <div class="mb-4">
                  <label for="editIcon" class="block font-semibold">Ảnh:</label>
                  <input
                    type="file"
                    id="editIcon"
                    accept="image/*"
                    class="w-full border rounded p-2 border-gray-500"
                    @change="newImage($event)"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label
                    for="menuLessons"
                    class="block text-gray-800 font-semibold"
                    >Tên học viên viên:</label
                  >
                  <input
                    type="text"
                    id="menuLessons"
                    v-model="newTestimonials.newName"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Nhập tên giảng viên"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label
                    for="menuOldPrice"
                    class="block text-gray-800 font-semibold"
                    >Nội dung:</label
                  >
                  <textarea
                    type="text"
                    id="menuOldPrice"
                    v-model="newTestimonials.newContent"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Mô tả giảng viên"
                    @keydown.enter.prevent="handleEnterKey"
                  ></textarea>
                </div>
              </div>
              <template #footer>
                <div class="text-center">
                  <button
                    @click="addNewHvNoiGiVeChungToi"
                    class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-6 rounded-md mr-3"
                  >
                    Thêm mới
                  </button>
                  <button
                    @click="editAddNewHvNoiGiVeChungToi"
                    class="bg-gray-500 hover:bg-gray-600 text-white font-semibold py-2 px-6 rounded-md"
                  >
                    Hủy
                  </button>
                </div>
              </template>
            </n-card>
          </n-modal>
        </div>
      </n-card>
    </n-modal>
  </div>
</template>

<script setup>
import { Swiper, SwiperSlide } from "swiper/vue";
import {
  useMessage,
  NCollapse,
  NCollapseItem,
  NButton,
  NCard,
  NModal,
} from "naive-ui";
import { Autoplay, Pagination } from "swiper/modules";
import "swiper/css";

const message = useMessage();

const modules = [Autoplay, Pagination];
const mySwiper = ref();

const handleEnterKey = (event) => {
  if (event.key === "Enter") {
    saveChanges();
  }
};

const title = ref("Học viên nói gì");
const subTitle = ref("về chúng tôi");
const editMode = ref(false);
const editAdd = ref(false);

const testimonial = defineModel("testimonial");

const props = defineProps({
  response: {
    type: Object,
    default: {},
  },
});

const testimonials = ref(testimonial.value);

// if (props.response && props.response[0]?.Metadata?.data_hvnoigivechungtoi) {
//   testimonials.value = props.response[0].Metadata.data_hvnoigivechungtoi;
// }

const editData = ref({
  editTitle: title.value,
  editSubTitle: subTitle.value,
  editTestimonials: [],
});

testimonials.value.forEach((testimonial, index) => {
  editData.value.editTestimonials.push({
    id: testimonial.id,
    editName: testimonial.name,
    editImage: testimonial.image,
    editContent: testimonial.content,
  });
});

const saveChanges = () => {
  editMode.value = false;
  title.value = editData.value.editTitle;
  subTitle.value = editData.value.editSubTitle;
  testimonials.value = [];
  editData.value.editTestimonials.forEach((testimonial, index) => {
    testimonials.value.push({
      id: testimonial.id,
      name: testimonial.editName,
      image: testimonial.editImage,
      content: testimonial.editContent,
    });
  });
  testimonial.value = testimonials.value;
  message.success("Lưu thành công");
};

const cancelEdit = () => {
  editMode.value = false;
  editData.value.editTitle = title.value;
  editData.value.editSubTitle = subTitle.value;
  editData.value.editTestimonials = [];
  testimonials.value.forEach((testimonial, index) => {
    editData.value.editTestimonials.push({
      id: testimonial.id,
      editName: testimonial.name,
      editImage: testimonial.image,
      editContent: testimonial.content,
    });
  });
};

const handleEditImage = (item, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      item.editImage = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

// ------------------------------------

const newTestimonials = ref({
  newName: "",
  newImage: "",
  newContent: "",
});

const uploadImage = (itemKey, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      newTestimonials.value.newImage = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const newImage = (event) => {
  uploadImage("newImage", event);
};

function addNewHvNoiGiVeChungToi() {
  // Kiểm tra nếu có trường nào không được nhập liệu
  if (
    // true
    !newTestimonials.value.newName ||
    !newTestimonials.value.newImage ||
    !newTestimonials.value.newContent
  ) {
    // console.log(newTestimonials.value);
    // Hiển thị thông báo yêu cầu nhập đầy đủ các trường
    alert("Vui lòng nhập đầy đủ thông tin");
    return; // Dừng hàm
  }

  // Nếu tất cả các trường đều được nhập đầy đủ, tiến hành lưu dữ liệu
  editData.value.editTestimonials.push({
    id: editData.value.editTestimonials.length + 1,
    editName: newTestimonials.value.newName,
    editImage: newTestimonials.value.newImage,
    editContent: newTestimonials.value.newContent,
  });

  // Reset giá trị của newTestimonials
  newTestimonials.value = {
    newName: "",
    newImage: "",
    newContent: "",
  };

  message.success("Thêm mới thành công");

  editAdd.value = false;
}

function editAddNewHvNoiGiVeChungToi() {
  // Reset giá trị của các biến ref mới
  newTestimonials.value.newName = "";
  newTestimonials.value.newImage = "";
  newTestimonials.value.newContent = "";

  editAdd.value = false;
}

function remove(id) {
  editData.value.editTestimonials = editData.value.editTestimonials.filter(
    function (item) {
      return item.id !== id;
    }
  );
}
</script>
