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
        <template #header-extra> </template>
        <div class="w-full">
          <div class="bg-white p-8 pr-0 rounded-lg">
            <div
              v-for="item in editData.editAbout"
              :key="item.id"
              class="mt-3 flex"
            >
              <div class="text-sm w-full">
                <div class="mb-4">
                  <label for="editName" class="block font-semibold">{{
                    about.value.name
                  }}</label>
                  <input
                    id="editName"
                    class="w-full border rounded p-2"
                    v-model="item.editName"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label for="editImg1" class="block font-semibold"
                    >Ảnh 1:</label
                  >
                  <input
                    type="file"
                    id="editImg1"
                    accept="image/*"
                    class="w-full border rounded p-2 border-gray-500"
                    @change="handleEditImg1(item, $event)"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label for="editImg2" class="block font-semibold"
                    >Ảnh 2:</label
                  >
                  <input
                    type="file"
                    id="editImg2"
                    accept="image/*"
                    class="w-full border rounded p-2 border-gray-500"
                    @change="handleEditImg2(item, $event)"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label for="editName" class="block font-semibold"
                    >Tên nội dung:</label
                  >
                  <input
                    id="editName"
                    class="w-full border rounded p-2"
                    v-model="item.editTitle"
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
                    v-model="item.editText"
                    @keydown.enter.prevent="handleEnterKey"
                  >
                  </textarea>
                </div>
              </div>
              <hr class="border-black my-2" />
            </div>
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
import { useMessage, NButton, NCard, NModal } from "naive-ui";

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
  editAbout: [],
});

aboutUs.value.forEach((about, index) => {
  editData.value.editAbout.push({
    id: about.id,
    editImg1: about.img1,
    editImg2: about.img2,
    editName: about.name,
    editText: about.text,
    editTitle: about.title,
  });
});

const saveChanges = () => {
  editMode.value = false;
  aboutUs.value = editData.value.editAbout.map((editAbout) => ({
    id: editAbout.id,
    img1: editAbout.editImg1,
    img2: editAbout.editImg2,
    name: editAbout.editName,
    text: editAbout.editText,
    title: editAbout.editTitle,
  }));
  message.success("Lưu thành công");
};

const cancelEdit = () => {
  editMode.value = false;
  editData.value.editAbout = [];
};

const handleEditImg1 = (item, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      item.editImg1 = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const handleEditImg2 = (item, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      item.editImg2 = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};
</script>
