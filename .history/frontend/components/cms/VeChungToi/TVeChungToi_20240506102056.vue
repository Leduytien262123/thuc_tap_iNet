<template>
  <div class="w-[72%] m-auto pt-4 relative">
    <div v-for="about in aboutUs" :key="about.id" class="">
      <h1 class="text-center text-[40px] font-bold">{{ about.name }}</h1>
      <div class="w-full pt-8">
        <div class="flex">
          <div
            class="w-[62%] h-[280px] mr-[3%] rounded-3xl flex justify-center items-center overflow-hidden"
          >
            <img :src="about.img1" alt="" class="w-full h-full" />
          </div>
          <div
            class="w-[35%] h-[280px] rounded-3xl flex justify-center items-center overflow-hidden"
          >
            <img :src="about.img2" alt="" class="w-full h-full" />
          </div>
        </div>
        <div class="flex pb-2">
          <div class="w-[30%] items-center justify-center overflow-hidden p-8">
            <h6 class="text-2xl font-bold text-center px-4">
              {{ about.title }}
            </h6>
          </div>
          <div
            class="w-[70%] items-center justify-center overflow-hidden py-8 px-10"
          >
            <p class="text-[#4E4E4E] text-sm">
              {{ about.text }}
            </p>
          </div>
        </div>
      </div>
    </div>
    <div class="absolute -right-[185px] top-1/2 -translate-y-1/2">
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
        <template #header-extra> {{ name }} </template>
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
              v-for="item in editData.editArticles"
              :key="item.id"
              class="mt-3 flex"
            >
              <n-collapse-item
                :title="'Chỉnh sửa bài viết:' + ' ' + item.editNameTitle"
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
                  <label :for="'tag1' + item.id" class="block font-semibold"
                    >Đường dẫn Marketing</label
                  >
                  <input
                    :id="'tag1' + item.id"
                    class="w-full border rounded p-2"
                    v-model="item.editTag1"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label :for="'tag2' + item.id" class="block font-semibold"
                    >Đường dẫn Analysis
                  </label>
                  <input
                    :id="'tag2' + item.id"
                    class="w-full border rounded p-2"
                    v-model="item.editTag2"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label
                    :for="'editNameTitle' + item.id"
                    class="block font-semibold"
                    >Tên bài viết:</label
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
                    >Nội dung:</label
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
                  <label :for="'editLink' + item.id" class="block font-semibold"
                    >Link xem thêm:</label
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
        </template>
      </n-card>
    </n-modal>
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

const handleEnterKey = (event) => {
  if (event.key === "Enter") {
    saveChanges();
  }
};

const message = useMessage();

const editMode = ref(false);

const about = defineModel("about");
const props = defineProps({
  response: {
    type: Object,
    default: {},
  },
});

const aboutUs = ref(about.value);

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

const saveChanges = () => {
  editMode.value = false;
  title.value = editData.value.editTitle;
  category2.value.title = title.value;
  subTitle.value = editData.value.editSubTitle;
  category2.value.subTitle = subTitle.value;

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
</script>
