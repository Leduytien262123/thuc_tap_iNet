<template>
  <div class="relative w-[72%] m-auto mb-[120px]">
    <div class="flex items-center justify-between mb-8">
      <h2 class="text-[40px] flex font-semibold">
        {{ title }}
        <p class="ml-3 text-[#6046FF]">{{ subTitle }}</p>
      </h2>
      <button
        class="bg-white text-black rounded-md px-5 py-3 text-center block ml-auto"
      >
        Xem thêm
      </button>
    </div>
    <div class="w-full h-full mb-6 grid grid-cols-3 gap-6">
      <div
        v-for="(course, index) in courses.slice(0, 3)"
        :key="index"
        class="bg-white rounded-lg block overflow-hidden"
      >
        <div class="w-[90%] m-auto relative">
          <div class="h-[40%]">
            <img
              class="h-[250px] w-full m-auto rounded-md mt-[5%] mb-[8%]"
              :src="course.image"
            />
            <span class="text-xl font-bold">{{ course.nameTitle }}</span>
          </div>
          <div class="h-[30%]">
            <p class="mt-3 text-sm mb-8">
              {{ course.description }}
            </p>
          </div>
          <div class="h-[30%] w-full float-bottom">
            <span class="text-lg">Giảng viên: {{ course.name }}</span>
            <div
              v-if="!course.price"
              class="text-[#6046FF] text-xl font-bold mt-3"
            >
              {{ formatPrice(course.oldPrice) }}
            </div>
            <div
              v-else-if="course.price !== '0'"
              class="flex items-end mb-4 mt-3"
            >
              <span class="text-sm line-through opacity-50 mb-[1.5px]">{{
                formatPrice(course.oldPrice)
              }}</span>
              <span class="text-[#6046FF] text-xl font-bold ml-2">{{
                formatPrice(course.price)
              }}</span>
            </div>
            <div v-else class="flex items-end mb-4 mt-3">
              <span class="text-sm line-through opacity-50 mb-[1.5px]">{{
                formatPrice(course.oldPrice)
              }}</span>
              <span class="text-[#6046FF] text-xl font-bold ml-2"
                >Miễn phí</span
              >
            </div>
            <nuxt-link :to="course.link">
              <button
                class="w-full h-12 mb-[5%] rounded-md border border-[#6046FF] text-[#6046FF] text-base font-semibold transition-transform transform-gpu motion-safe:hover:scale-105"
              >
                Đăng ký ngay
              </button></nuxt-link
            >
          </div>
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
          <!-- fixed z-[110] inset-0 my-auto bg-gray-800 bg-opacity-50 flex justify-center items-center -->
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
                v-for="item in editData.editCourses"
                :key="item.id"
                class="mt-3 flex w-full items-center"
              >
                <n-collapse-item
                  :title="'Chỉnh sửa ' + item.editNameTitle"
                  :name="item.id"
                  class="text-xl w-full"
                >
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
                      :for="'editNameTitle' + item.id"
                      class="block font-semibold"
                      >Tên môn học:</label
                    >
                    <input
                      :id="'editNameTitle' + item.id"
                      class="w-full border rounded p-2"
                      v-model="item.editNameTitle"
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
                    <label
                      :for="'editName' + item.id"
                      class="block font-semibold"
                      >Tên giảng viên:</label
                    >
                    <input
                      :id="'editName' + item.id"
                      class="w-full border rounded p-2"
                      v-model="item.editName"
                      @keydown.enter.prevent="handleEnterKey"
                    />
                  </div>
                  <div class="mb-4">
                    <label
                      :for="'editOldPrice' + item.id"
                      class="block font-semibold"
                      >Giá gốc:</label
                    >
                    <input
                      :id="'editOldPrice' + item.id"
                      class="w-full border rounded p-2"
                      v-model="item.editOldPrice"
                      @keydown.enter.prevent="handleEnterKey"
                    />
                  </div>
                  <div class="mb-4">
                    <label
                      :for="'editPrice' + item.id"
                      class="block font-semibold"
                      >Giá khuyến mãi:</label
                    >
                    <input
                      :id="'editPrice' + item.id"
                      class="w-full border rounded p-2"
                      v-model="item.editPrice"
                      @keydown.enter.prevent="handleEnterKey"
                    />
                  </div>

                  <div class="mb-4">
                    <label
                      :for="'editLink' + item.id"
                      class="block font-semibold"
                      >Đường dẫn:</label
                    >
                    <input
                      :id="'editLink' + item.id"
                      class="w-full border rounded p-2"
                      v-model="item.editLink"
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
          <template #footer
            ><div class="text-right mt-10">
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
                <!-- class="fixed w-[32%] pr-0 h-[85vh] top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 bg-white p-8 rounded-lg shadow-lg border border-gray-300" -->
                <div class="h-full pr-8 overflow-y-auto">
                  <div class="mb-4">
                    <label for="editImage" class="block font-semibold"
                      >Ảnh:</label
                    >
                    <input
                      type="file"
                      id="editImage"
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
                      >Tên khóa học:</label
                    >
                    <input
                      type="text"
                      id="menuLessons"
                      v-model="newCourse.newNameTitle"
                      class="w-full border rounded p-2 border-gray-500"
                      placeholder="Nhập tên bài giảng"
                      @keydown.enter.prevent="handleEnterKey"
                    />
                  </div>
                  <div class="mb-4">
                    <label
                      for="menuNameTitle"
                      class="block text-gray-800 font-semibold"
                      >Mô tả:</label
                    >
                    <textarea
                      type="text"
                      id="menuNameTitle"
                      v-model="newCourse.newDescription"
                      class="w-full border rounded p-2 border-gray-500"
                      placeholder="Mô tả"
                      @keydown.enter.prevent="handleEnterKey"
                    >
                    </textarea>
                  </div>
                  <div class="mb-4">
                    <label
                      for="menuOldPrice"
                      class="block text-gray-800 font-semibold"
                      >Tên giảng viên:</label
                    >
                    <input
                      type="text"
                      id="menuOldPrice"
                      v-model="newCourse.newName"
                      class="w-full border rounded p-2 border-gray-500"
                      placeholder="Nhập tên giảng viên"
                      @keydown.enter.prevent="handleEnterKey"
                    />
                  </div>
                  <div class="mb-4">
                    <label
                      for="menuPrice"
                      class="block text-gray-800 font-semibold"
                      >Nhập giá:</label
                    >
                    <input
                      type="text"
                      id="menuPrice"
                      v-model="newCourse.newOldPrice"
                      class="w-full border rounded p-2 border-gray-500"
                      placeholder="Nhập giá"
                      @keydown.enter.prevent="handleEnterKey"
                    />
                  </div>
                  <div class="mb-4">
                    <label
                      for="menuPrice"
                      class="block text-gray-800 font-semibold"
                      >Nhập giá khuyến mãi:</label
                    >
                    <input
                      type="text"
                      id="menuPrice"
                      v-model="newCourse.newPrice"
                      class="w-full border rounded p-2 border-gray-500"
                      placeholder="Nhập giá khuyến mãi"
                      @keydown.enter.prevent="handleEnterKey"
                    />
                  </div>
                  <div class="mb-6">
                    <label
                      for="menuLink"
                      class="block text-gray-800 font-semibold"
                      >Đường dẫn:</label
                    >
                    <input
                      type="text"
                      id="menuLink"
                      v-model="newCourse.newLink"
                      class="w-full border rounded p-2 border-gray-500"
                      placeholder="Nhập đường dẫn "
                      @keydown.enter.prevent="handleEnterKey"
                    />
                  </div>
                </div>
                <template #footer>
                  <div class="text-center">
                    <button
                      @click="addNewGoiKhoaHoc"
                      class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-6 rounded-md mr-3"
                    >
                      Thêm mới
                    </button>
                    <button
                      @click="editAddNewGoiKhoaHoc"
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
  </div>
</template>

<script setup>
import {
  useMessage,
  NCollapse,
  NCollapseItem,
  NButton,
  NCard,
  NModal,
} from "naive-ui";

const title = ref("Gói khóa học");
const subTitle = ref("tiêu biểu");
const editMode = ref(false);
const editAdd = ref(false);
const message = useMessage();

const course = defineModel("course");

const props = defineProps({
  response: Object,
});

const courses = ref(course.value);

if (props.response && props.response[0]?.Metadata?.data_goikhoahoctieubieu) {
  courses.value = props.response[0].Metadata.data_goikhoahoctieubieu;
}

const formatPrice = (price) => {
  if (!price) return "";
  const formattedPrice = parseFloat(price).toLocaleString("vi-VN", {
    style: "decimal",
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  });
  return formattedPrice + "đ";
};

const editData = ref({
  editTitle: title.value,
  editSubTitle: subTitle.value,
  editCourses: [],
});

courses.value.forEach((course, index) => {
  editData.value.editCourses.push({
    id: course.id,
    editImage: course.image,
    editNameTitle: course.nameTitle,
    editDescription: course.description,
    editOldPrice: course.oldPrice,
    editPrice: course.price,
    editName: course.name,
    editLink: course.link,
  });
});

const handleEnterKey = (event) => {
  if (event.key === "Enter") {
    saveChanges();
  }
};

const saveChanges = () => {
  editMode.value = false;
  title.value = editData.value.editTitle;
  subTitle.value = editData.value.editSubTitle;
  courses.value = [];
  editData.value.editCourses.forEach((course, index) => {
    courses.value.push({
      id: course.id,
      image: course.editImage,
      nameTitle: course.editNameTitle,
      description: course.editDescription,
      oldPrice: course.editOldPrice,
      price: course.editPrice,
      name: course.editName,
      link: course.editLink,
    });
  });
  message.success("Lưu thành công");
};

const cancelEdit = () => {
  editMode.value = false;
  // Khôi phục dữ liệu ban đầu
  editData.value.editTitle = title.value;
  editData.value.editSubTitle = subTitle.value;
  editData.value.editCourses = [];
  courses.value.forEach((course, index) => {
    editData.value.editCourses.push({
      id: course.id,
      editImage: course.image,
      editNameTitle: course.nameTitle,
      editDescription: course.description,
      editOldPrice: course.oldPrice,
      editPrice: course.price,
      editName: course.name,
      editLink: course.link,
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

const uploadImage = (itemKey, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      newCourse.value.newImage = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const newImage = (event) => {
  uploadImage("newImage", event);
};

const newCourse = ref({
  newImage: "",
  newLessons: "",
  newNameTitle: "",
  newOldPrice: "",
  newPrice: "",
  newAvatar: "",
  newName: "",
  newLink: "",
});

function addNewGoiKhoaHoc() {
  // Kiểm tra nếu có trường nào không được nhập liệu
  if (
    !newCourse.value.newImage ||
    !newCourse.value.newNameTitle ||
    !newCourse.value.newDescription ||
    !newCourse.value.newOldPrice ||
    !newCourse.value.newPrice ||
    !newCourse.value.newName
  ) {
    // Hiển thị thông báo yêu cầu nhập đầy đủ các trường
    alert("Vui lòng nhập đầy đủ thông tin");
    return; // Dừng hàm
  }

  // Nếu tất cả các trường đều được nhập đầy đủ, tiến hành lưu dữ liệu
  editData.value.editCourses.push({
    id: editData.value.editCourses.length + 1,
    editImage: newCourse.value.newImage,
    editNameTitle: newCourse.value.newNameTitle,
    newDescription: newCourse.value.newDescription,
    editOldPrice: newCourse.value.newOldPrice,
    editPrice: newCourse.value.newPrice,
    editName: newCourse.value.newName,
    editLink: newCourse.value.newLink,
  });

  // Reset giá trị của newCourse
  newCourse.value = {
    newImage: "",
    newNameTitle: "",
    newDescription: "",
    newOldPrice: "",
    newPrice: "",
    newName: "",
    newLink: "",
  };
  message.success("Thêm mới thành công");

  editAdd.value = false;
}

function editAddNewGoiKhoaHoc() {
  // Reset giá trị của các biến ref mới
  newCourse.value.newImage = "";
  newCourse.value.newNameTitle = "";
  newCourse.value.newDescription = "";
  newCourse.value.newOldPrice = "";
  newCourse.value.newPrice = "";
  newCourse.value.newName = "";
  newCourse.value.newLink = "";

  editAdd.value = false;
}

function remove(id) {
  editData.value.editCourses = editData.value.editCourses.filter(function (
    item
  ) {
    return item.id !== id;
  });
}
</script>
