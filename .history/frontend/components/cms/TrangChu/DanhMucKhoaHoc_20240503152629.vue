<template>
  <div class="w-[72%] m-auto mb-[120px] relative">
    <div class="flex items-center justify-between mb-8">
      <h2 class="text-[40px] flex font-semibold">
        <p class="mr-3 text-[#6046FF]">{{ title }}</p>
        {{ subTitle }}
      </h2>
      <button
        class="bg-white text-black rounded-md px-5 py-3 text-center block ml-auto"
      >
        Xem thêm
      </button>
    </div>
    <div class="w-full grid grid-cols-5 gap-x-5 gap-y-6 mb-6">
      <template v-for="category in categories.slice(0, 10)" :key="category.id">
        <div
          class="pb-10 bg-white rounded-lg block overflow-hidden text-center cursor-pointer transition-transform transform-gpu motion-safe:hover:scale-105 border border-transparent hover:border-[#6046FF] hover:filter"
        >
          <nuxt-link :to="category.link"
            ><div before class="m-auto w-14 h-14 mt-12 mb-8">
              <img :src="category.icon" alt="" class="w-full h-full" />
            </div>
            <span class="font-medium text-lg mb-8"
              >{{ category.name1 }}<br />
              {{ category.name2 }}
            </span></nuxt-link
          >
        </div>
      </template>
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
                v-for="item in editData.editCategories"
                :key="item.id"
                class="mt-3 flex items-center"
              >
                <n-collapse-item
                  :title="'Chỉnh sửa ' + item.editName1 + item.editName2"
                  :name="item.id"
                  class="text-xl w-full"
                >
                  <div class="mb-4">
                    <label for="editIcon" class="block font-semibold"
                      >icon:</label
                    >
                    <input
                      type="file"
                      id="editIcon"
                      accept="image/*"
                      class="w-full border rounded p-2 border-gray-500"
                      @change="handleEditIcon(item, $event)"
                      @keydown.enter.prevent="handleEnterKey"
                    />
                  </div>
                  <div class="mb-4">
                    <label
                      :for="'editName1' + item.id"
                      class="block font-semibold"
                      >Tên dòng trên:</label
                    >
                    <input
                      :id="'editName1' + item.id"
                      class="w-full border rounded p-2"
                      v-model="item.editName1"
                      @keydown.enter.prevent="handleEnterKey"
                    />
                  </div>
                  <div class="mb-4">
                    <label
                      :for="'editName2' + item.id"
                      class="block font-semibold"
                      >Tên dòng dưới:</label
                    >
                    <input
                      :id="'editName2' + item.id"
                      class="w-full border rounded p-2"
                      v-model="item.editName2"
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
                    <label for="editIcon" class="block font-semibold"
                      >Ảnh:</label
                    >
                    <input
                      type="file"
                      id="editIcon"
                      accept="image/*"
                      class="w-full border rounded p-2 border-gray-500"
                      @change="newIcon($event)"
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
                      v-model="newCategories.newName1"
                      class="w-full border rounded p-2 border-gray-500"
                      placeholder="Nhập tên bài giảng"
                      @keydown.enter.prevent="handleEnterKey"
                    />
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
                      v-model="newCategories.newName2"
                      class="w-full border rounded p-2 border-gray-500"
                      placeholder="Nhập tên giảng viên"
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
                      v-model="newCategories.newLink"
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

const category2 = defineModel("category2");

const title = ref(category2.value.title);
const subTitle = ref(category2.value.subTitle);
const editMode = ref(false);
const editAdd = ref(false);
const message = useMessage();

const category = defineModel("category");

const props = defineProps({
  response: {
    type: Object,
    default: {},
  },
});

const categories = ref(category.value);

// if (props.response && props.response[0]?.Metadata?.data_danhmuckhoahoc) {
//   categories.value = props.response[0].Metadata.data_danhmuckhoahoc;
// }

const editData = ref({
  editTitle: title.value,
  editSubTitle: subTitle.value,
  editCategories: [],
});

categories.value.forEach((category, index) => {
  editData.value.editCategories.push({
    id: category.id,
    editIcon: category.icon,
    editName1: category.name1,
    editName2: category.name2,
    editLink: category.link,
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
  categories.value = [];
  editData.value.editCategories.forEach((category, index) => {
    categories.value.push({
      id: category.id,
      icon: category.editIcon,
      name1: category.editName1,
      name2: category.editName2,
      link: category.editLink,
    });
  });
  category.value = categories.value;
  message.success("Lưu thành công");
};

const cancelEdit = () => {
  editMode.value = false;
  editData.value.editTitle = title.value;
  editData.value.editSubTitle = subTitle.value;
  editData.value.editCategories = [];
  categories.value.forEach((category, index) => {
    editData.value.editCategories.push({
      id: category.id,
      editIcon: category.icon,
      editName1: category.name1,
      editName2: category.name2,
      editLink: category.link,
    });
  });
};

const handleEditIcon = (item, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      item.editIcon = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const newCategories = ref({
  newIcon: "",
  newName1: "",
  newName2: "",
  newLink: "",
});

const uploadIcon = (itemKey, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      newCategories.value.newIcon = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const newIcon = (event) => {
  uploadIcon("newIcon", event);
};

function addNewDanhMucKhoaHoc() {
  // Kiểm tra nếu có trường nào không được nhập liệu
  if (
    // true
    !newCategories.value.newIcon ||
    !newCategories.value.newName1 ||
    !newCategories.value.newName2
  ) {
    // console.log(newCategories.value);
    // Hiển thị thông báo yêu cầu nhập đầy đủ các trường
    alert("Vui lòng nhập đầy đủ thông tin");
    return; // Dừng hàm
  }

  // Nếu tất cả các trường đều được nhập đầy đủ, tiến hành lưu dữ liệu
  editData.value.editCategories.push({
    id: editData.value.editCategories.length + 1,
    editIcon: newCategories.value.newIcon,
    editName1: newCategories.value.newName1,
    editName2: newCategories.value.newName2,
    editLink: newCategories.value.newLink,
  });

  // Reset giá trị của newCategories
  newCategories.value = {
    newIcon: "",
    newName1: "",
    newName2: "",
    newLink: "",
  };

  message.success("Thêm mới thành công");

  editAdd.value = false;
}

function editAddNewDanhMucKhoaHoc() {
  // Reset giá trị của các biến ref mới
  newCategories.value.newIcon = "";
  newCategories.value.newName1 = "";
  newCategories.value.newName2 = "";
  newCategories.value.newLink = "";

  editAdd.value = false;
}

function remove(id) {
  editData.value.editCategories = editData.value.editCategories.filter(
    function (item) {
      return item.id !== id;
    }
  );
}
</script>
