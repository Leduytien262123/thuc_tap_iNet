<template>
  <div class="banner relative w-[98%] flex flex-initial mx-auto">
    <img
      :src="form.backgroundImage"
      alt=""
      class="w-[100%] h-[600px] mx-auto rounded-lg mt-1 blur-none"
    />
    <div class="absolute top-5 right-0 mr-4">
      <n-button
        @click="editMode = true"
        type="success"
        class="bg-green-600 px-6 py-3"
      >
        Sửa
      </n-button>
    </div>

    <div
      v-if="!editMode"
      class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 items-center text-center justify-center inline"
    >
      <div class="bg-[#e9e9e9bf] p-4 rounded-md flex items-center gap-2">
        <div
          class="bg-[#FFF4E5] w-12 h-12 flex items-center justify-center rounded-md"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="ionicon w-6 h-6"
            viewBox="0 0 512 512"
          >
            <path d="M432 208H288l32-192L80 304h144l-32 192z" />
          </svg>
        </div>
        <span class="text-[#5236FF] text-4xl font-bold mr-1">{{
          form.title
        }}</span>
        <span class="text-black text-4xl font-bold mr-1">{{
          form.subtitle
        }}</span>
      </div>
      <div class="text-white text-3xl text-center mt-4">
        {{ form.description }}
      </div>
      <div class="mt-14">
        <nuxt-link :to="form.linkButton1">
          <button
            :style="{ backgroundColor: form.button1Color }"
            class="text-white px-5 py-3 rounded-md mr-3 text-center transition-transform transform-gpu motion-safe:hover:scale-110"
          >
            {{ form.button1 }}
          </button></nuxt-link
        >
        <nuxt-link :to="form.linkButton2">
          <button
            :style="{ backgroundColor: form.button2Color }"
            class="text-black px-5 py-3 rounded-md text-center transition-transform transform-gpu motion-safe:hover:scale-110"
          >
            {{ form.button2 }}
          </button></nuxt-link
        >
      </div>
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
        <template #header-extra></template>
        <div class="w-full">
          <div class="bg-white p-4 pr-0 rounded-lg">
            <div class="mb-4">
              <label for="editTitle" class="block font-semibold"
                >Tiêu đề:</label
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
              <label for="editSubtitle" class="block font-semibold"
                >Tiêu đề phụ:</label
              >
              <input
                type="text"
                id="editSubtitle"
                class="w-full border rounded p-2"
                v-model="editData.editSubtitle"
                @keydown.enter.prevent="handleEnterKey"
              />
            </div>
            <div class="mb-4">
              <label for="editDescription" class="block font-semibold"
                >Nội dung:</label
              >
              <textarea
                id="editDescription"
                class="w-full border rounded p-2"
                v-model="editData.editDescription"
                @keydown.enter.prevent="handleEnterKey"
              ></textarea>
            </div>
            <div class="mb-4">
              <label for="editButton1Text" class="block font-semibold"
                >Nút: {{ editData.editButton1 }}</label
              >
              <input
                type="text"
                id="editButton1Text"
                class="w-full border rounded p-2"
                v-model="editData.editButton1"
                @keydown.enter.prevent="handleEnterKey"
              />
            </div>
            <div class="mb-4">
              <label for="editButton1Text" class="block font-semibold"
                >Link: {{ editData.editButton1 }}</label
              >
              <input
                type="text"
                id="editButton1Text"
                class="w-full border rounded p-2"
                v-model="editData.editLinkButton1"
                @keydown.enter.prevent="handleEnterKey"
              />
            </div>
            <div class="mb-4">
              <label for="editButton2Text" class="block font-semibold"
                >Nút: {{ editData.editButton2 }}</label
              >
              <input
                type="text"
                id="editButton2Text"
                class="w-full border rounded p-2"
                v-model="editData.editButton2"
                @keydown.enter.prevent="handleEnterKey"
              />
            </div>
            <div class="mb-4">
              <label for="editButton2Text" class="block font-semibold"
                >Link: {{ editData.editButton2 }}</label
              >
              <input
                type="text"
                id="editButton2Text"
                class="w-full border rounded p-2"
                v-model="editData.editLinkButton2"
                @keydown.enter.prevent="handleEnterKey"
              />
            </div>
            <div class="mb-4">
              <label for="editBackground" class="block font-semibold"
                >Ảnh nền:</label
              >
              <input
                type="file"
                id="editBackground"
                accept="image/*"
                class="w-full border rounded p-2"
                @change="handleEditBackgroundChange"
                @keydown.enter.prevent="handleEnterKey"
              />
            </div>
          </div>
        </div>
        <template #footer>
          <div class="text-right mt-0">
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
// import { ref } from "vue";
import {
  useMessage,
  NCollapse,
  NCollapseItem,
  NButton,
  NCard,
  NModal,
} from "naive-ui";

const message = useMessage();

const editMode = ref(false);
const forms = defineModel("form");
const props = defineProps({
  response: {
    type: Object,
    default: {},
  },
});

const form = ref(forms.value);

// if (props.response && props.response[0]?.Metadata?.data_banner) {
//   form.value = props.response[0].Metadata.data_banner;
// }

let temporaryBackgroundImage;

const handleEditBackgroundChange = (event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      editData.value.editBackgroundImage = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const editData = ref({
  editTitle: form.value.title,
  editSubtitle: form.value.subtitle,
  editDescription: form.value.description,
  editBackgroundImage: form.value.backgroundImage,
  editButton1: form.value.button1,
  editLinkButton1: form.value.linkButton1,
  editButton2: form.value.button2,
  editLinkButton2: form.value.linkButton2,
});

const saveChanges = () => {
  editMode.value = false;
  temporaryBackgroundImage = null;
  form.value.title = editData.value.editTitle;
  form.value.subtitle = editData.value.editSubtitle;
  form.value.description = editData.value.editDescription;
  form.value.backgroundImage = editData.value.editBackgroundImage;
  form.value.button1 = editData.value.editButton1;
  form.value.linkButton1 = editData.value.editLinkButton1;
  form.value.button2 = editData.value.editButton2;
  form.value.linkButton2 = editData.value.editLinkButton2;
  message.success("Lưu thành công");
};

const cancelEdit = () => {
  editMode.value = false;
  editData.value.editTitle = form.value.title;
  editData.value.editSubtitle = form.value.subtitle;
  editData.value.editDescription = form.value.description;
  editData.value.editBackgroundImage = form.value.backgroundImage;
  editData.value.editButton1 = form.value.button1;
  editData.value.editLinkButton1 = form.value.linkButton1;
  editData.value.editButton2 = form.value.button2;
  editData.value.editLinkButton2 = form.value.linkButton2;
  if (temporaryBackgroundImage) {
    editData.value.editBackgroundImage = temporaryBackgroundImage;
    temporaryBackgroundImage = null;
  }
};

const handleEnterKey = (event) => {
  if (event.key === "Enter") {
    saveChanges();
  }
};
</script>
