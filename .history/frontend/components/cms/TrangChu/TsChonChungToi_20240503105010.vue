<template>
  <div class="w-full mb-[100px] m-auto">
    <div class="text-[45px] flex justify-center mb-8 mt-20">
      <b class="flex items-center">
        <p class="text-blue-600 mr-3">{{ title }}</p>
        {{ subTitle }}
      </b>
    </div>
    <div class="w-full flex">
      <div class="w-[14%]"></div>
      <div class="w-[72%] m-auto rounded-xl flex">
        <div class="w-[48%]">
          <div v-for="item in topItems" :key="item.id">
            <nuxt-link :to="item.link">
              <div
                class="bg-[#F9F9F9] h-[200px] mb-8 rounded-3xl flex flex-col"
              >
                <div class="inline-block mt-auto mb-auto w-full">
                  <div class="flex items-center">
                    <img
                      class="w-[100px] h-[130px] rounded-xl mr-4 ml-6"
                      :src="item.imageSrc"
                      :alt="item.alt"
                    />
                    <div class="flex flex-col mr-6 ml-3">
                      <b class="mr-2 text-xl">{{ item.nameTitle }}</b>
                      <p class="text-base text-gray-600 mt-1">
                        {{ item.description }}
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </nuxt-link>
          </div>
        </div>
        <div class="w-[4%]"></div>
        <div class="w-[48%]">
          <div v-for="item in bottomItems" :key="item.id">
            <nuxt-link :to="item.link">
              <div
                class="bg-[#F9F9F9] h-[200px] mb-8 rounded-3xl flex flex-col"
              >
                <div class="inline-block mt-auto mb-auto w-full">
                  <div class="flex items-center w-full">
                    <img
                      class="w-[100px] h-[130px] rounded-xl mr-4 ml-6"
                      :src="item.imageSrc"
                      :alt="item.alt"
                    />
                    <div class="flex flex-col mr-6 ml-3">
                      <b class="mr-2 text-xl">{{ item.nameTitle }}</b>
                      <p class="text-base text-gray-600 mt-1">
                        {{ item.description }}
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </nuxt-link>
          </div>
        </div>
      </div>
      <div class="m-auto w-[14%] relative flex items-center justify-center">
        <div class="absolute right-8 mb-7">
          <n-button
            @click="editMode = true"
            class="px-6 py-3 bg-green-600"
            type="success"
          >
            Sửa
          </n-button>
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
          <div class="bg-white p-8 pr-0 rounded-lg">
            <div class="h-full pr-8 overflow-y-auto">
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
                v-for="item in items"
                :key="item.id"
              >
                <n-collapse-item
                  :title="'Chỉnh sửa ' + item.nameTitle"
                  :name="item.id"
                  class="text-xl"
                >
                  <div>
                    <div class="mb-4">
                      <label
                        :for="'editNameTitle' + item.id"
                        class="block font-semibold"
                        >Tên danh sách:</label
                      >
                      <input
                        type="text"
                        :id="'editNameTitle' + item.id"
                        class="w-full border rounded p-2"
                        v-model="editData['editNameTitle' + item.id]"
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
                        v-model="editData['editDescription' + item.id]"
                        @keydown.enter.prevent="handleEnterKey"
                      ></textarea>
                    </div>
                    <div class="mb-4">
                      <label
                        :for="'editImage' + item.id"
                        class="block font-semibold"
                        >Ảnh/logo:</label
                      >
                      <input
                        type="file"
                        :id="'editImage' + item.id"
                        class="w-full border rounded p-2"
                        @change="(event) => handleEditImg(event, item.id)"
                        @keydown.enter.prevent="handleEnterKey"
                      />
                    </div>
                  </div>
                </n-collapse-item>
              </n-collapse>

              <div class="text-right">
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
            </div>
          </div>
        </n-card>
      </n-modal>
    </div>
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

const editMode = ref(false);
const title = ref("Tại sao");
const subTitle = ref("chọn chúng tôi");
const message = useMessage();

const props = defineProps({
  response: Object,
});

const items = ref([]);

if (props.response && props.response[0]?.Metadata?.data_tschonchungtoi) {
  items.value = props.response[0].Metadata.data_tschonchungtoi;
}

const topItems = ref(items.value.slice(0, 2));
const bottomItems = ref(items.value.slice(2));

const editData = ref({
  editTitle: title.value,
  editSubTitle: subTitle.value,
  editNameTitle1: items.value[0].nameTitle,
  editNameTitle2: items.value[1].nameTitle,
  editNameTitle3: items.value[2].nameTitle,
  editNameTitle4: items.value[3].nameTitle,
  editDescription1: items.value[0].description,
  editDescription2: items.value[1].description,
  editDescription3: items.value[2].description,
  editDescription4: items.value[3].description,
  editImageSrc1: items.value[0].imageSrc,
  editImageSrc2: items.value[1].imageSrc,
  editImageSrc3: items.value[2].imageSrc,
  editImageSrc4: items.value[3].imageSrc,
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
  items.value.forEach((item, index) => {
    item.nameTitle = editData.value[`editNameTitle${index + 1}`];
    item.description = editData.value[`editDescription${index + 1}`];
    item.imageSrc = editData.value[`editImageSrc${index + 1}`];
  });
  message.success("Lưu thành công");
};

const handleEditImg = (event, index) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      editData.value[`editImageSrc${index}`] = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const cancelEdit = () => {
  editMode.value = false;
  editData.value.editTitle = title.value;
  editData.value.editSubTitle = subTitle.value;
  items.value.forEach((item, index) => {
    editData.value[`editNameTitle${index + 1}`] = item.nameTitle;
    editData.value[`editDescription${index + 1}`] = item.description;
    editData.value[`editImageSrc${index + 1}`] = item.imageSrc;
  });
};
</script>
