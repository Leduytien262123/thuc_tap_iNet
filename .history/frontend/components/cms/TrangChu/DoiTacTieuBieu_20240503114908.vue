<template>
  <div class="w-[72%] m-auto relative">
    <div class="flex items-center justify-between mb-8">
      <h2 class="text-[40px] flex font-semibold">
        {{ title }}
        <p class="ml-3 text-[#6046FF]">{{ subTitle }}</p>
      </h2>
    </div>
    <div class="flex justify-center">
      <div
        v-for="(partner, index) in partners.slice(0, 7)"
        :key="index"
        class="w-[14.2%] bg-[#F3F3F4] overflow-hidden"
      >
        <div
          class="px-5 py-3 transition duration-300 transform hover:scale-110 cursor-pointer"
        >
          <nuxt-link :to="partner.link">
            <img
              class="w-full h-full object-cover"
              :src="partner.image"
              :alt="partner.name"
          /></nuxt-link>
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
              v-for="item in editData.editPartners"
              :key="item.id"
              class="mt-3 flex items-center"
            >
              <n-collapse-item
                :title="'Chỉnh sửa đối tác ' + ' ' + item.editName"
                :name="item.id"
                class="w-full"
              >
                <div class="mb-4">
                  <label for="editPartnerName" class="block font-semibold"
                    >Tên đối tác:</label
                  >
                  <input
                    type="text"
                    :id="'editPartnerName'"
                    class="w-full border rounded p-2"
                    v-model="item.editName"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label for="editImage" class="block font-semibold"
                    >Logo đối tác:</label
                  >
                  <input
                    type="file"
                    :id="'editImage'"
                    accept="image/*"
                    class="w-full border rounded p-2 border-gray-500"
                    @change="handleEditImage(item, $event)"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label for="editLink" class="block font-semibold"
                    >Đường dẫn:</label
                  >
                  <input
                    type="text"
                    :id="'editLink'"
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
                    >Logo:</label
                  >
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
                    >Tên đối tác:</label
                  >
                  <input
                    type="text"
                    id="menuLessons"
                    v-model="newPartners.newName"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Nhập tên đối tác"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label
                    for="menuOldPrice"
                    class="block text-gray-800 font-semibold"
                    >Đường dẫn:</label
                  >
                  <input
                    type="text"
                    id="menuOldPrice"
                    v-model="newPartners.newLink"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Liên kết đến đối tác"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
              </div>
              <template #footer>
                <div class="text-center">
                  <button
                    @click="addNewDoiTacTieuBieu"
                    class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-6 rounded-md mr-3"
                  >
                    Thêm mới
                  </button>
                  <button
                    @click="editAddNewDoiTacTieuBieu"
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
import { ref } from "vue";
import {
  useMessage,
  NCollapse,
  NCollapseItem,
  NButton,
  NCard,
  NModal,
} from "naive-ui";

const message = useMessage();

const partner = defineModel("partner");
const props = defineProps({
  response: {
    type: Object,
    default: {},
  },
});

const partners = ref(partner.value);

if (props.response && props.response[0]?.Metadata?.data_doitactieubieu) {
  partners.value = props.response[0].Metadata.data_doitactieubieu;
}

const handleEnterKey = (event) => {
  if (event.key === "Enter") {
    saveChanges();
  }
};

const title = ref("Đối tác");
const subTitle = ref("tiêu biểu");
const editMode = ref(false);
const editAdd = ref(false);

const editData = ref({
  editTitle: title.value,
  editSubTitle: subTitle.value,
  editPartners: [],
});

partners.value.forEach((partner, index) => {
  editData.value.editPartners.push({
    id: partner.id,
    editName: partner.name,
    editImage: partner.image,
    editLink: partner.link,
  });
});

const saveChanges = () => {
  editMode.value = false;
  title.value = editData.value.editTitle;
  subTitle.value = editData.value.editSubTitle;
  partners.value = [];
  editData.value.editPartners.forEach((partner, index) => {
    partners.value.push({
      id: partner.id,
      name: partner.editName,
      image: partner.editImage,
      link: partner.editLink,
    });
  });
  message.success("Lưu thành công");
};

const cancelEdit = () => {
  editMode.value = false;
  editData.value.editTitle = title.value;
  editData.value.editSubTitle = subTitle.value;
  editData.value.editPartners = [];
  partners.value.forEach((partner, index) => {
    editData.value.editPartners.push({
      id: partner.id,
      editName: partner.name,
      editImage: partner.image,
      editLink: partner.link,
    });
  });
};

const handleEditImage = (item, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      item.image = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

// ------------------------------------

const newPartners = ref({
  newName: "",
  newImage: "",
  newLink: "",
});

const uploadImage = (itemKey, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      newPartners.value.newImage = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const newImage = (event) => {
  uploadImage("newImage", event);
};

function addNewDoiTacTieuBieu() {
  // Kiểm tra nếu có trường nào không được nhập liệu
  if (!newPartners.value.newName || !newPartners.value.newImage) {
    // console.log(newPartners.value);
    // Hiển thị thông báo yêu cầu nhập đầy đủ các trường
    alert("Vui lòng nhập đầy đủ thông tin");
    return; // Dừng hàm
  }

  // Nếu tất cả các trường đều được nhập đầy đủ, tiến hành lưu dữ liệu
  editData.value.editPartners.push({
    id: editData.value.editPartners.length + 1,
    editName: newPartners.value.newName,
    editImage: newPartners.value.newImage,
    editLink: newPartners.value.newLink,
  });

  // Reset giá trị của newPartners
  newPartners.value = {
    newName: "",
    newImage: "",
    newLink: "",
  };

  message.success("Thêm mới thành công");

  editAdd.value = false;
}

function editAddNewDoiTacTieuBieu() {
  // Reset giá trị của các biến ref mới
  newPartners.value.newName = "";
  newPartners.value.newImage = "";
  newPartners.value.newLink = "";

  editAdd.value = false;
}

function remove(id) {
  editData.value.editPartners = editData.value.editPartners.filter(function (
    item
  ) {
    return item.id !== id;
  });
}
</script>
