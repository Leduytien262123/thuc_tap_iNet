<template>
  <div class="bg-[#F3F3F4] w-full mx-auto">
    <div v-for="item in form" :key="item.id" class="w-[98%] m-auto">
      <nuxt-link :to="item.linkTieuDe" class="mx-auto">
        <button
          class="border-none py-3 text-base w-full mx-auto bg-[#5236ff] text-white rounded-md mt-4 hover:bg-sky-500 flex items-center cursor-pointer"
        >
          <div class="mx-auto">{{ item.tieuDe }}</div>
        </button>
      </nuxt-link>
      <div class="w-full flex relative">
        <div class="w-[80%] mx-auto flex py-3">
          <div class="flex items-center justify-between w-full">
            <div class="flex items-center">
              <div class="logo ml-5">
                <img :src="item.logo" alt="Logo" class="w-40 h-20 px-4 py-4" />
              </div>
              <div class="text-base ml-8" v-for="menu in menus" :key="menu.id">
                <nuxt-link :to="menu.link">{{ menu.name }}</nuxt-link>
              </div>
              <div class="ml-5 flex">
                <div class="text-base mt-1">
                  <nuxt-link to="#">
                    <img
                      class="w-5 h-5"
                      src="https://cdn-icons-png.flaticon.com/512/1170/1170627.png"
                    />
                  </nuxt-link>
                </div>
              </div>
            </div>
            <div class="ml-4 mr-10">
              <nuxt-link :to="item.linkDangNhap" class="text-base mr-7">{{
                item.dangNhap
              }}</nuxt-link>
              <nuxt-link
                :to="item.linkDangKy"
                class="bg-[#5236ff] py-3 px-7 border-none rounded-md cursor-pointer text-white hover:bg-sky-500"
                >{{ item.dangKy }}</nuxt-link
              >
            </div>
          </div>
        </div>
        <div class="absolute right-0 top-1/2 transform -translate-y-1/2">
          <n-button
            @click="editMode = true"
            type="success"
            class="bg-green-600 px-6 py-3 mr-4"
          >
            Sửa
          </n-button>
        </div>
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
          <div
            v-for="item in editData.editForm"
            :key="item.id"
            class="bg-white p-8 pr-0 rounded-lg"
          >
            <div class="mb-4">
              <label for="editTitle" class="block font-semibold"
                >Tiêu đề:</label
              >
              <div class="flex items-center">
                <div class="flex-1">
                  <input
                    type="text"
                    id="editTitle"
                    class="w-full border rounded p-2 border-gray-500"
                    v-model="item.editTieuDe"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                  <label for="editSubtitle" class="block font-semibold"
                    >Link:</label
                  >
                  <input
                    id="editSubtitle"
                    type="text"
                    v-model="item.editLinkTieuDe"
                    class="w-full border rounded p-2 border-gray-500"
                    :placeholder="'Nhập link'"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
              </div>
            </div>
            <div class="mb-4">
              <label for="editBackground" class="block font-semibold"
                >Logo:</label
              >
              <div class="flex items-center">
                <input
                  type="file"
                  id="editLogo"
                  accept="image/*"
                  class="w-full border rounded p-2 border-gray-500"
                  @change="handleEditLogo"
                  @keydown.enter.prevent="handleEnterKey"
                />
              </div>
            </div>
            <div class="mb-4 space-y-3">
              <label for="editSubtitle" class="block font-semibold"
                >Menu:</label
              >
              <n-collapse
                arrow-placement="right"
                v-for="item in editData.editMenus"
                :key="item.id"
                class="mt-2 flex items-center"
              >
                <n-collapse-item
                  class="flex-1"
                  :title="item.editName"
                  :name="item.id"
                >
                  <label for="editSubtitle" class="block font-semibold">
                    {{ item.editName }}</label
                  >
                  <input
                    id="editSubtitle"
                    type="text"
                    v-model="item.editName"
                    class="w-full border rounded p-2 border-gray-500"
                    :placeholder="'Nhập tên menu ' + item.editName"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                  <label for="editSubtitle" class="block font-semibold"
                    >Link:</label
                  >
                  <input
                    id="editSubtitle"
                    type="text"
                    v-model="item.editLink"
                    class="w-full border rounded p-2 border-gray-500"
                    :placeholder="'Nhập link ' + item.editName"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </n-collapse-item>
                <button
                  @click="
                    selectedItemId = item.id;
                    showDelete = true;
                  "
                  class="self-start mt-0"
                >
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
                <n-modal
                  v-model:show="showDelete"
                  :mask-closable="false"
                  preset="dialog"
                  title="Bạn chắc chắn xóa?"
                  positive-text="Xóa"
                  positive-text-color="red"
                  type="error"
                  negative-text="Thoát"
                  @positive-click="remove"
                  @negative-click="cancel"
                  class="border text-red-500 bg-white"
                />
              </n-collapse>
            </div>
            <div class="mt-2">
              <label for="editSubtitle1" class="block font-semibold">{{
                item.editDangNhap
              }}</label>
              <input
                id="editSubtitle1"
                type="text"
                v-model="item.editDangNhap"
                class="w-full border rounded p-2 border-gray-500"
                :placeholder="item.editDangNhap"
                @keydown.enter.prevent="handleEnterKey"
              />
              <label for="editSubtitle2" class="block font-semibold"
                >Link</label
              >
              <input
                id="editSubtitle2"
                type="text"
                v-model="item.editLinkDangNhap"
                class="w-full border rounded p-2 border-gray-500"
                :placeholder="'Nhập link ' + item.editDangNhap"
                @keydown.enter.prevent="handleEnterKey"
              />
            </div>

            <div class="mt-2 mb-8">
              <label for="editSubtitle3" class="block font-semibold">{{
                item.editDangKy
              }}</label>
              <input
                id="editSubtitle3"
                type="text"
                v-model="item.editDangKy"
                class="w-full border rounded p-2 border-gray-500"
                :placeholder="item.editDangKy"
                @keydown.enter.prevent="handleEnterKey"
              />
              <label for="editSubtitle4" class="block font-semibold"
                >Link</label
              >
              <input
                id="editSubtitle4"
                type="text"
                v-model="item.editLinkDangKy"
                class="w-full border rounded p-2 border-gray-500"
                :placeholder="'Nhập link ' + item.editDangKy"
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
          <div class="mt-4">
            <n-button
              class="bg-green-600 px-6 py-3"
              @click="editAdd = true"
              type="success"
            >
              Thêm
            </n-button>
          </div>
        </template>

        <div class="w-full">
          <n-modal v-model:show="editAdd">
            <n-card
              style="width: 600px"
              title="Thêm mới"
              :bordered="false"
              size="huge"
              role="dialog"
              aria-modal="true"
            >
              <template #header-extra> </template>
              <div class="h-full pr-8 overflow-y-auto">
                <div class="mb-4">
                  <div class="mb-4">
                    <label
                      for="menuName"
                      class="block text-gray-800 font-semibold"
                      >Tên menu:</label
                    >
                    <input
                      type="text"
                      id="menuName"
                      v-model="newMenuName"
                      class="w-full border rounded p-2 border-gray-500"
                      placeholder="Nhập tên menu"
                      @keydown.enter.prevent="handleEnterKey"
                    />
                  </div>
                  <div class="mb-6">
                    <label
                      for="menuLink"
                      class="block text-gray-800 font-semibold"
                      >Liên kết:</label
                    >
                    <input
                      type="text"
                      id="menuLink"
                      v-model="newMenuLink"
                      class="w-full border rounded p-2 border-gray-500"
                      placeholder="Nhập đường dẫn menu"
                      @keydown.enter.prevent="handleEnterKey"
                    />
                  </div>
                </div>
              </div>
              <template #footer>
                <div class="text-center">
                  <button
                    @click="addNewMenu"
                    class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-6 rounded-md mr-3"
                  >
                    Thêm mới
                  </button>
                  <button
                    @click="editAddNewMenu"
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
import { ref, defineComponent } from "vue";
import {
  useMessage,
  NCollapse,
  NCollapseItem,
  NButton,
  NCard,
  NModal,
  NInput,
} from "naive-ui";

const message = useMessage();
defineEmits(["update:menus"]);
const model = defineModel();
const props = defineProps({
  response: {
    type: Object,
    default: {},
  },
  menu: {
    type: Array,
    default: [],
  },
});
console.log(props.menu);
const form = ref([]);
const menus = ref(props.menu);

if (props.response && props.response[0]?.Metadata?.data_form) {
  form.value = props.response[0].Metadata.data_form;
}

// if (props.response && props.response[0]?.Metadata?.data_menus) {
//   menus.value = props.response[0].Metadata.data_menus;
// }

// const { data: resHomeData } = await useFetch("http://localhost:8000/page");

// const header = resHomeData.value;

// let form = ref([]);
// let menus = ref([]);

// header.forEach((item) => {
//   if (item.Metadata) {
//     if (item.Metadata.data_form) {
//       form.value = item.Metadata.data_form;
//     }
//     if (item.Metadata.data_menus) {
//       menus.value = item.Metadata.data_menus;
//     }
//   }
// });

const editMode = ref(false);

const editData = ref({
  editForm: [],
  editMenus: [],
});

form.value.forEach((form, index) => {
  editData.value.editForm.push({
    id: form.id,
    editTieuDe: form.tieuDe,
    editLinkTieuDe: form.linkTieuDe,
    editLogo: form.logo,
    editDangNhap: form.dangNhap,
    editLinkDangNhap: form.linkDangNhap,
    editDangKy: form.dangKy,
    editLinkDangKy: form.linkDangKy,
  });
});

menus.value.forEach((menu, index) => {
  editData.value.editMenus.push({
    id: menu.id,
    editName: menu.name,
    editLink: menu.link,
  });
});

let temporaryLogo;

const saveChanges = () => {
  editMode.value = false;

  temporaryLogo = null;

  editData.value.editForm.forEach((formData, index) => {
    form.value[index].tieuDe = formData.editTieuDe;
    form.value[index].linkTieuDe = formData.editLinkTieuDe;
    form.value[index].logo = formData.editLogo;
    form.value[index].dangNhap = formData.editDangNhap;
    form.value[index].linkDangNhap = formData.editLinkDangNhap;
    form.value[index].dangKy = formData.editDangKy;
    form.value[index].linkDangKy = formData.editLinkDangKy;
  });

  menus.value = [];
  editData.value.editMenus.forEach((menu, index) => {
    menus.value.push({
      id: menu.id,
      name: menu.editName,
      link: menu.editLink,
    });
  });
  message.success("Lưu thành công");
};

const cancelEdit = () => {
  editMode.value = false;

  editData.value.editForm = [];
  form.value.forEach((form, index) => {
    editData.value.editForm.push({
      id: form.id,
      editTieuDe: form.tieuDe,
      editLinkTieuDe: form.linkTieuDe,
      editLogo: form.logo,
      editDangNhap: form.dangNhap,
      editLinkDangNhap: form.linkDangNhap,
      editDangKy: form.dangKy,
      editLinkDangKy: form.linkDangKy,
    });
  });

  editData.value.editMenus = [];
  menus.value.forEach((menu, index) => {
    editData.value.editMenus.push({
      id: menu.id,
      editName: menu.name,
      editLink: menu.link,
    });
  });
};

const handleEditLogo = (event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      editData.value.editForm.forEach((formData, index) => {
        formData.editLogo = e.target.result;
      });
    };
    reader.readAsDataURL(file);
  }
};

const handleEnterKey = (event) => {
  if (event.key === "Enter") {
    saveChanges();
    // addNewMenu();
  }
};

const editAdd = ref(false);

const newMenuName = ref("");
const newMenuLink = ref("");

function editAddNewMenu() {
  newMenuName.value = "";
  newMenuLink.value = "";
  editAdd.value = false;
}

function addNewMenu() {
  if (newMenuName.value && newMenuLink.value) {
    editData.value.editMenus = [
      ...editData.value.editMenus,
      {
        id: editData.value.editMenus.length + 1,
        editName: newMenuName.value,
        editLink: newMenuLink.value,
      },
    ];

    newMenuName.value = "";
    newMenuLink.value = "";
  }
  message.success("Thêm mới thành công");

  editAdd.value = false;
  console.log(editData.value.editMenus);
}

const showDelete = ref(false);

const cancel = () => {
  message.success("Thoát");
  showDelete.value = false;
};

let selectedItemId = null;

function remove() {
  if (selectedItemId !== null) {
    const index = editData.value.editMenus.findIndex(
      (item) => item.id === selectedItemId
    );
    if (index !== -1) {
      editData.value.editMenus.splice(index, 1);
      message.success("Xóa thành công");
    } else {
      message.error("Không tìm thấy phần tử để xóa");
    }
    console.log(editData.value.editMenus);
  }
  showDelete.value = false;
  selectedItemId = null;
}
</script>
