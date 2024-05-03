<template>
  <!-- Phần hiển thị dữ liệu -->
  <div class="w-full flex">
    <!-- Các ô hiển thị dữ liệu -->
    <div class="w-[14%]"></div>
    <div class="w-[72%] h-auto bg-[#F9F9F9] m-auto rounded-xl flex">
      <div
        v-for="item in items"
        :key="item.id"
        class="w-1/3 h-[250px] flex items-center justify-center"
      >
        <div class="text-center animate-fade-in">
          <span class="text-6xl flex font-bold mb-3"
            >{{ item.value }}
            <p class="text-[blue] font-bold">+</p></span
          >
          <span class="text-xl">{{ item.label }}</span>
        </div>
      </div>
    </div>

    <!-- Nút sửa -->
    <div class="m-auto w-[14%] relative flex items-center justify-center">
      <div class="absolute right-8">
        <!-- v-if="!editMode" rounded-md transition-transform transform-gpu motion-safe:hover:scale-110 text-white-->
        <n-button
          @click="editMode = true"
          type="success"
          class="px-6 py-3 bg-green-600"
        >
          Sửa
        </n-button>
      </div>
    </div>

    <!-- Form chỉnh sửa -->
    <n-modal v-model:show="editMode">
      <n-card
        style="width: 600px"
        title="Chỉnh sửa"
        :bordered="false"
        size="huge"
        role="dialog"
        aria-modal="true"
      >
        <div class="bg-white p-8 rounded-lg w-full">
          <h2 class="text-2xl font-semibold mb-4">Chỉnh sửa</h2>
          <div class="mb-4" v-for="(item, index) in editData" :key="index">
            <label :for="'editField' + index" class="block font-semibold"
              >{{ item.label }}:</label
            >
            <input
              :type="item.type"
              :id="'editField' + index"
              class="w-full border rounded p-2 border-gray-500"
              v-model="item.value"
              @keydown.enter.prevent="handleEnterKey"
            />
          </div>
          <div class="text-right mt-10">
            <button
              class="px-4 py-2 bg-blue-400 text-white rounded hover:bg-blue-600"
              @click="saveChanges"
            >
              Lưu
            </button>
            <button
              class="px-4 py-2 bg-gray-300 text-gray-800 rounded ml-2 hover:bg-[#717f94]"
              @click="cancelEdit"
            >
              Hủy
            </button>
          </div>
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
  NInput,
} from "naive-ui";

const message = useMessage();

const editMode = ref(false);

const props = defineProps({
  response: Object,
});

const items = ref([]);

if (props.response && props.response[0]?.Metadata?.data_thongso) {
  items.value = props.response[0].Metadata.data_thongso;
}

const editData = ref([
  { label: "Tên giảng viên", value: items.value[0].label, type: "text" },
  { label: "Số giảng viên", value: items.value[0].value, type: "text" },
  { label: "Tên học viên", value: items.value[1].label, type: "text" },
  { label: "Số học viên", value: items.value[1].value, type: "text" },
  { label: "Tên bài giảng", value: items.value[2].label, type: "text" },
  { label: "Số bài giảng", value: items.value[2].value, type: "text" },
]);

const saveChanges = () => {
  editMode.value = false;
  items.value[0].label = editData.value[0].value;
  items.value[0].value = editData.value[1].value;
  items.value[1].label = editData.value[2].value;
  items.value[1].value = editData.value[3].value;
  items.value[2].label = editData.value[4].value;
  items.value[2].value = editData.value[5].value;
  message.success("Lưu thành công");
};

const cancelEdit = () => {
  editMode.value = false;
  editData.value[0].value = items.value[0].label;
  editData.value[1].value = items.value[0].value;
  editData.value[2].value = items.value[1].label;
  editData.value[3].value = items.value[1].value;
  editData.value[4].value = items.value[2].label;
  editData.value[5].value = items.value[2].value;
};

const handleEnterKey = (event) => {
  if (event.key === "Enter") {
    saveChanges();
  }
};
</script>
