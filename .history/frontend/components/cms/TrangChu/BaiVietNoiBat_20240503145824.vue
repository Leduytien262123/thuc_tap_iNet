<template>
  <div class="w-[72%] m-auto mb-[110px] relative">
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
    <div class="w-full grid grid-cols-2 gap-x-8">
      <template v-for="article in articles.slice(0, 2)" :key="article.id">
        <div class="bg-white rounded-xl relative">
          <div class="w-[87%] m-auto">
            <div class="my-[6.5%] relative h-[280px]">
              <img class="rounded-xl w-full h-full" :src="article.image" />
              <div class="absolute bottom-4 left-0 z-10 flex">
                <nuxt-link :to="article.tag1">
                  <button
                    class="px-8 py-2.5 ml-4 font-medium text-sm bg-white rounded-3xl transition-transform transform-gpu motion-safe:hover:scale-110 hover:bg-sky-300"
                  >
                    Marketing
                  </button></nuxt-link
                >
                <nuxt-link :to="article.tag2">
                  <button
                    class="px-8 py-2.5 ml-4 font-medium text-sm bg-white rounded-3xl transition-transform transform-gpu motion-safe:hover:scale-110 hover:bg-sky-300"
                  >
                    Analysis
                  </button></nuxt-link
                >
              </div>
            </div>
            <p class="mt-2 mb-4">{{ currentDate }}</p>
            <b class="text-xl">{{ article.nameTitle }}</b>
            <p class="text-[#797979] mt-3 mb-6 text-[13.5px]">
              {{ article.description }}
            </p>
            <div class="mb-[6.5%]">
              <nuxt-link
                :to="article.link"
                class="text-blue-600 underline text-sm"
                >Xem thêm</nuxt-link
              >
            </div>
          </div>
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
                    >Tên bài viết:</label
                  >
                  <input
                    type="text"
                    id="menuLessons"
                    v-model="newArticles.newNameTitle"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Nhập tên bài giảng"
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
                    v-model="newArticles.newDescription"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Nhập tên giảng viên"
                    @keydown.enter.prevent="handleEnterKey"
                  ></textarea>
                </div>
                <div class="mb-6">
                  <label
                    for="menuLink1"
                    class="block text-gray-800 font-semibold"
                    >Đường dẫn Marketing:</label
                  >
                  <input
                    type="text"
                    id="menuLink1"
                    v-model="newArticles.newTag1"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Nhập đường dẫn "
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-6">
                  <label
                    for="menuLink2"
                    class="block text-gray-800 font-semibold"
                    >Đường dẫn AnalySis:</label
                  >
                  <input
                    type="text"
                    id="menuLink2"
                    v-model="newArticles.newTag2"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Nhập đường dẫn "
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-6">
                  <label
                    for="menuLink3"
                    class="block text-gray-800 font-semibold"
                    >Link xem thêm:</label
                  >
                  <input
                    type="text"
                    id="menuLink3"
                    v-model="newArticles.newLink"
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
import { ref, watchEffect } from "vue";
import {
  useMessage,
  NCollapse,
  NCollapseItem,
  NButton,
  NCard,
  NModal,
} from "naive-ui";

const title = ref("Bài viết");
const subTitle = ref("nổi bật");
const editMode = ref(false);
const editAdd = ref(false);

const message = useMessage();

const currentDate = ref("");

watchEffect(() => {
  const date = new Date();
  const day = date.getDate();
  const month = date.getMonth() + 1;
  const year = date.getFullYear();
  currentDate.value = `${day}/${month}/${year}`;
});

const article = defineModel("article");

const props = defineProps({
  response: {
    type: Object,
    default: {},
  },
});

const articles = ref(article.value);

const handleEnterKey = (event) => {
  if (event.key === "Enter") {
    saveChanges();
  }
};

const editData = ref({
  editTitle: title.value,
  editSubTitle: subTitle.value,
  editArticles: [],
});

articles.value.forEach((article, index) => {
  editData.value.editArticles.push({
    id: article.id,
    editNameTitle: article.nameTitle,
    editDescription: article.description,
    editImage: article.image,
    editTag1: article.tag1,
    editTag2: article.tag2,
    editLink: article.link,
  });
});

const saveChanges = () => {
  editMode.value = false;
  title.value = editData.value.editTitle;
  subTitle.value = editData.value.editSubTitle;
  articles.value = [];
  editData.value.editArticles.forEach((article, index) => {
    articles.value.push({
      id: article.id,
      nameTitle: article.editNameTitle,
      description: article.editDescription,
      image: article.editImage,
      tag1: article.editTag1,
      tag2: article.editTag2,
      link: article.editLink,
    });
  });
  article.value = articles.value;
  message.success("Lưu thành công");
};

const cancelEdit = () => {
  editMode.value = false;
  editData.value.editTitle = title.value;
  editData.value.editSubTitle = subTitle.value;
  editData.value.editArticles = [];
  articles.value.forEach((article, index) => {
    editData.value.editArticles.push({
      id: article.id,
      editNameTitle: article.nameTitle,
      editDescription: article.description,
      editImage: article.image,
      editTag1: article.tag1,
      editTag2: article.tag2,
      editLink: article.link,
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

// ------------------------------------------

const newArticles = ref({
  newImage: "",
  newNameTitle: "",
  newDescription: "",
  newTag1: "",
  newTag2: "",
  newLink: "",
});

const uploadImage = (itemKey, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      newArticles.value.newImage = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const newImage = (event) => {
  uploadImage("newImage", event);
};

function addNewDanhMucKhoaHoc() {
  if (
    // true
    !newArticles.value.newImage ||
    !newArticles.value.newNameTitle ||
    !newArticles.value.newDescription
  ) {
    alert("Vui lòng nhập đầy đủ thông tin");
    return; // Dừng hàm
  }

  editData.value.editArticles.push({
    id: editData.value.editArticles.length + 1,
    editImage: newArticles.value.newImage,
    editNameTitle: newArticles.value.newNameTitle,
    editDescription: newArticles.value.newDescription,
    editTag1: newArticles.value.newTag1,
    editTag2: newArticles.value.newTag2,
    editLink: newArticles.value.newLink,
  });

  newArticles.value = {
    newImage: "",
    newNameTitle: "",
    newDescription: "",
    newTag1: "",
    newTag2: "",
    newLink: "",
  };

  message.success("Thêm mới thành công");

  editAdd.value = false;
}

function editAddNewDanhMucKhoaHoc() {
  newArticles.value.newImage = "";
  newArticles.value.newNameTitle = "";
  newArticles.value.newDescription = "";
  newArticles.value.newTag1 = "";
  newArticles.value.newTag2 = "";
  newArticles.value.newLink = "";

  editAdd.value = false;
}

function remove(id) {
  editData.value.editArticles = editData.value.editArticles.filter(function (
    item
  ) {
    return item.id !== id;
  });
}
</script>
