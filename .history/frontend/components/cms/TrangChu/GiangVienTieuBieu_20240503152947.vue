<template>
  <div class="w-[72%] m-auto relative mb-[110px]">
    <div class="flex items-center justify-between mb-8">
      <h2 class="text-[40px] flex font-semibold">
        {{ title }}
        <p class="ml-3 text-[6046FF]">{{ subTitle }}</p>
      </h2>
      <button
        class="bg-white text-black rounded-md px-5 py-3 text-center block ml-auto"
      >
        Xem thêm
      </button>
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
              300: { slidesPerView: 1.175 },
              400: { slidesPerView: 1.515 },
              768: { slidesPerView: 3, spaceBetween: 20 },
              1024: { slidesPerView: 4 },
            }"
            :space-between="30"
            loop
            :autoplay="{
              delay: 2000,
              disableOnInteraction: false,
              pauseOnMouseEnter: true,
            }"
            :speed="1000"
            class="!z-0 rounded-lg"
          >
            <SwiperSlide v-for="(teacher, index) of teachers" :key="index">
              <div
                class="w-full py-0.5 bg-white rounded-lg text-center hover:bg-[EEEEFF]"
              >
                <div
                  class="mx-auto mt-8 mb-8 flex flex-col items-center scroll-snap-start"
                >
                  <div
                    class="rounded-full w-[110px] h-[110px] mb-4 flex items-center justify-center overflow-hidden"
                  >
                    <img
                      class="w-full h-full object-cover"
                      :src="teacher.image"
                      :alt="teacher.name"
                    />
                  </div>
                  <b class="text-lg">{{ teacher.name }}</b>
                  <p class="text-sm mt-1">{{ teacher.description }}</p>
                  <div class="flex justify-center items-center mt-3">
                    <a :href="teacher.linkFb">
                      <img
                        class="w-3.5 h-3.5 mx-1.5"
                        :src="teacher.iconFb"
                        :alt="'Facebook ' + teacher.name"
                      />
                    </a>
                    <a :href="teacher.linkIg">
                      <img
                        class="w-3.5 h-3.5 mx-1.5"
                        :src="teacher.iconIg"
                        :alt="'Instagram ' + teacher.name"
                      />
                    </a>
                    <a :href="teacher.linkIn">
                      <img
                        class="w-3.5 h-3.5 mx-1.5"
                        :src="teacher.iconIn"
                        :alt="'LinkedIn ' + teacher.name"
                      />
                    </a>
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
              v-for="item in editData.editTeachers"
              :key="item.id"
              class="mt-3 flex"
            >
              <n-collapse-item
                :title="'Chỉnh sửa giảng viên:' + ' ' + item.editName"
                :name="item.id"
                class="w-full"
              >
                <div class="mb-4">
                  <label :for="'editName' + item.id" class="block font-semibold"
                    >Tên giảng viên</label
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
                    :for="'editDescription' + item.id"
                    class="block font-semibold"
                    >Mô tả:</label
                  >
                  <textarea
                    :id="'editDescription' + item.id"
                    class="w-full border rounded p-2"
                    v-model="item.editDescription"
                    @keydown.enter.prevent="handleEnterKey"
                  >
                  </textarea>
                </div>
                <div class="mb-4">
                  <label for="editIconfb" class="block font-semibold"
                    >Icon Facebook:</label
                  >
                  <input
                    type="file"
                    id="editIconfb"
                    accept="image/*"
                    class="w-full border rounded p-2 border-gray-500"
                    @change="handleEditFbIcon(item, $event)"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label
                    :for="'editFacebook' + item.id"
                    class="block font-semibold"
                    >Link Facebook:</label
                  >
                  <input
                    type="text"
                    :id="'editFacebook' + item.id"
                    class="w-full border rounded p-2"
                    v-model="item.editFb"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label for="editIconIg" class="block font-semibold"
                    >Icon Instagram:</label
                  >
                  <input
                    type="file"
                    id="editIconIg"
                    accept="image/*"
                    class="w-full border rounded p-2 border-gray-500"
                    @change="handleEditIgIcon(item, $event)"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label
                    :for="'editInstagram' + item.id"
                    class="block font-semibold"
                    >Link Instagram:</label
                  >
                  <input
                    type="text"
                    :id="'editInstagram' + item.id"
                    class="w-full border rounded p-2"
                    v-model="item.editIg"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label for="editIconIn" class="block font-semibold"
                    >Icon LinkedIn:</label
                  >
                  <input
                    type="file"
                    id="editIconIn"
                    accept="image/*"
                    class="w-full border rounded p-2 border-gray-500"
                    @change="handleEditInIcon(item, $event)"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label
                    :for="'editLinkedin' + item.id"
                    class="block font-semibold"
                    >Link LinkedIn:</label
                  >
                  <input
                    type="text"
                    :id="'editLinkedin' + item.id"
                    class="w-full border rounded p-2"
                    v-model="item.editIn"
                    @keydown.enter.prevent="handleEnterKey"
                  />
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
                    >Tên giảng viên:</label
                  >
                  <input
                    type="text"
                    id="menuLessons"
                    v-model="newTeachers.newName"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Nhập tên giảng viên"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label
                    for="menuOldPrice"
                    class="block text-gray-800 font-semibold"
                    >Mô tả:</label
                  >
                  <textarea
                    type="text"
                    id="menuOldPrice"
                    v-model="newTeachers.newDescription"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Mô tả giảng viên"
                    @keydown.enter.prevent="handleEnterKey"
                  ></textarea>
                </div>
                <div class="mb-4">
                  <label for="editIcon" class="block font-semibold"
                    >Icon FaceBook:</label
                  >
                  <input
                    type="file"
                    id="editIcon"
                    accept="image/*"
                    class="w-full border rounded p-2 border-gray-500"
                    @change="newIconFb($event)"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-6">
                  <label
                    for="menuLink1"
                    class="block text-gray-800 font-semibold"
                    >Đường dẫn Facebook:</label
                  >
                  <input
                    type="text"
                    id="menuLink1"
                    v-model="newTeachers.newLinkFb"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Nhập đường dẫn "
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label for="editIcon" class="block font-semibold"
                    >Icon Instagram:</label
                  >
                  <input
                    type="file"
                    id="editIcon"
                    accept="image/*"
                    class="w-full border rounded p-2 border-gray-500"
                    @change="newIconIg($event)"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-6">
                  <label
                    for="menuLink1"
                    class="block text-gray-800 font-semibold"
                    >Đường dẫn Instagram:</label
                  >
                  <input
                    type="text"
                    id="menuLink1"
                    v-model="newTeachers.newLinkIg"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Nhập đường dẫn "
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label for="editIcon" class="block font-semibold"
                    >Icon Linkedln:</label
                  >
                  <input
                    type="file"
                    id="editIcon"
                    accept="image/*"
                    class="w-full border rounded p-2 border-gray-500"
                    @change="newIconIn($event)"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-6">
                  <label
                    for="menuLink1"
                    class="block text-gray-800 font-semibold"
                    >Đường dẫn Linkedln:</label
                  >
                  <input
                    type="text"
                    id="menuLink1"
                    v-model="newTeachers.newLinkIn"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Nhập đường dẫn "
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
              </div>
              <template #footer>
                <div class="text-center">
                  <button
                    @click="addNewDanhMucKhoaHoc"
                    class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-6 rounded-md mr-3"
                  >
                    Thêm mới
                  </button>
                  <button
                    @click="editAddNewDanhMucKhoaHoc"
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
// import { ref, reactive } from "vue";
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

const modules = [Autoplay, Pagination];
const mySwiper = ref();

const handleEnterKey = (event) => {
  if (event.key === "Enter") {
    saveChanges();
  }
};

const teacher2 = defineModel("teacher2");

const title = ref(teacher2.value.title);
const subTitle = ref(teacher2.value.subTitle);
const editMode = ref(false);
const editAdd = ref(false);

const message = useMessage();

const teacher = defineModel("teacher");

const props = defineProps({
  response: {
    type: Object,
    default: {},
  },
});

const teachers = ref(teacher.value);

// if (props.response && props.response[0]?.Metadata?.data_giangvientieubieu) {
//   teachers.value = props.response[0].Metadata.data_giangvientieubieu;
// }

const editData = ref({
  editTitle: title.value,
  editSubTitle: subTitle.value,
  editTeachers: [],
});

teachers.value.forEach((teacher, index) => {
  editData.value.editTeachers.push({
    id: teacher.id,
    editName: teacher.name,
    editImage: teacher.image,
    editDescription: teacher.description,
    editFb: teacher.linkFb,
    editFbIcon: teacher.iconFb,
    editIg: teacher.linkIg,
    editIgIcon: teacher.iconIg,
    editIn: teacher.linkIn,
    editInIcon: teacher.iconIn,
  });
});

const saveChanges = () => {
  editMode.value = false;
  title.value = editData.value.editTitle;
  teacher2.value.title = title.value;
  subTitle.value = editData.value.editSubTitle;
  teacher2.value.title = title.value;

  teachers.value = [];
  editData.value.editTeachers.forEach((teacher, index) => {
    teachers.value.push({
      id: teacher.id,
      name: teacher.editName,
      image: teacher.editImage,
      description: teacher.editDescription,
      linkFb: teacher.editFb,
      iconFb: teacher.editFbIcon,
      linkIg: teacher.editIg,
      iconIg: teacher.editIgIcon,
      linkIn: teacher.editIn,
      iconIn: teacher.editInIcon,
    });
  });
  teacher.value = teachers.value;
  message.success("Lưu thành công");
};

const cancelEdit = () => {
  editMode.value = false;
  editData.value.editTitle = title.value;
  editData.value.editSubTitle = subTitle.value;
  editData.value.editTeachers = [];
  teachers.value.forEach((teacher, index) => {
    editData.value.editTeachers.push({
      id: teacher.id,
      editName: teacher.name,
      editImage: teacher.image,
      editDescription: teacher.description,
      editFb: teacher.linkFb,
      editFbIcon: teacher.iconFb,
      editIg: teacher.linkIg,
      editIgIcon: teacher.iconIg,
      editIn: teacher.linkIn,
      editInIcon: teacher.iconIn,
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

const handleEditFbIcon = (item, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      item.editFbIcon = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const handleEditIgIcon = (item, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      item.editIgIcon = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const handleEditInIcon = (item, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      item.editInIcon = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

// ------------------------------------------

const newTeachers = ref({
  newName: "",
  newImage: "",
  newDescription: "",
  newIconFb: "",
  newLinkFb: "",
  newIconIg: "",
  newLinkIg: "",
  newIconIn: "",
  newLinkIn: "",
});

const uploadImage = (itemKey, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      newTeachers.value.newImage = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const newImage = (event) => {
  uploadImage("newImage", event);
};

const uploadIconFb = (itemKey, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      newTeachers.value.newIconFb = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const newIconFb = (event) => {
  uploadIconFb("newIconFb", event);
};

const uploadIconIg = (itemKey, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      newTeachers.value.newIconIg = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const newIconIg = (event) => {
  uploadIconIg("newIconIg", event);
};

const uploadIconIn = (itemKey, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      newTeachers.value.newIconIn = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const newIconIn = (event) => {
  uploadIconIn("newIconIn", event);
};

function addNewDanhMucKhoaHoc() {
  // Kiểm tra nếu có trường nào không được nhập liệu
  if (
    // true
    !newTeachers.value.newName ||
    !newTeachers.value.newImage ||
    !newTeachers.value.newDescription ||
    !newTeachers.value.newIconFb ||
    !newTeachers.value.newLinkFb
  ) {
    // console.log(newTeachers.value);
    // Hiển thị thông báo yêu cầu nhập đầy đủ các trường
    alert("Vui lòng nhập đầy đủ thông tin");
    return; // Dừng hàm
  }

  // Nếu tất cả các trường đều được nhập đầy đủ, tiến hành lưu dữ liệu
  editData.value.editTeachers.push({
    id: editData.value.editTeachers.length + 1,
    editName: newTeachers.value.newName,
    editImage: newTeachers.value.newImage,
    editDescription: newTeachers.value.newDescription,
    editFbIcon: newTeachers.value.newIconFb,
    editFb: newTeachers.value.newLinkFb,
    editIgIcon: newTeachers.value.newIconIg,
    editIg: newTeachers.value.newLinkIg,
    editInIcon: newTeachers.value.newIconIn,
    editIn: newTeachers.value.newLinkIn,
  });

  // Reset giá trị của newTeachers
  newTeachers.value = {
    newName: "",
    newImage: "",
    newDescription: "",
    newIconFb: "",
    newLinkFb: "",
    newIconIg: "",
    newLinkIg: "",
    newIconIn: "",
    newLinkIn: "",
  };

  message.success("Thêm mới thành công");

  editAdd.value = false;
}

function editAddNewDanhMucKhoaHoc() {
  // Reset giá trị của các biến ref mới
  newTeachers.value.newName = "";
  newTeachers.value.newImage = "";
  newTeachers.value.newDescription = "";
  newTeachers.value.newIconFb = "";
  newTeachers.value.newLinkFb = "";
  newTeachers.value.newIconIg = "";
  newTeachers.value.newLinkIg = "";
  newTeachers.value.newIconIn = "";
  newTeachers.value.newLinkIn = "";

  editAdd.value = false;
}

function remove(id) {
  editData.value.editTeachers = editData.value.editTeachers.filter(function (
    item
  ) {
    return item.id !== id;
  });
}
</script>
