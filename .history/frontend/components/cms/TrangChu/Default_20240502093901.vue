<template>
  <div
    class="w-[98%] sticky top-0 z-[100] m-auto bg-[#3A434E] mt-auto rounded-md p-2 flex items-center"
  >
    <div class="w-1/3">
      <select
        id="pageSelect"
        class="p-2 rounded-md w-[45%] border border-gray-300 ml-2"
        v-model="selectedPage"
      >
        <option
          v-for="(option, index) in options"
          :key="option"
          :value="option"
          :disabled="index == 0"
        >
          {{ option.text }}
        </option>
      </select>
      <select
        class="rounded-md p-2 ml-3 w-[45%] border border-gray-300"
        v-model="selectedSubPage"
      >
        <option v-for="option in options" :key="option" :value="option">
          {{ option.text }}
        </option>
      </select>
    </div>
    <div class="p-2 w-[40%] flex m-auto">
      <div class="m-auto">
        <button class="">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            viewBox="0 0 320 512"
            class="bg-blue-500 text-white w-9 h-8 p-2 rounded-lg hover:bg-sky-300"
          >
            <path
              d="M272 0H48C21.5 0 0 21.5 0 48v416c0 26.5 21.5 48 48 48h224c26.5 0 48-21.5 48-48V48c0-26.5-21.5-48-48-48zM160 480c-17.7 0-32-14.3-32-32s14.3-32 32-32s32 14.3 32 32zm112-108c0 6.6-5.4 12-12 12H60c-6.6 0-12-5.4-12-12V60c0-6.6 5.4-12 12-12h200c6.6 0 12 5.4 12 12v312z"
              fill="currentColor"
            ></path>
          </svg>
        </button>
        <button class="ml-3">
          <svg
            version="1.1"
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            x="0px"
            y="0px"
            viewBox="0 0 512 512"
            enable-background="new 0 0 512 512"
            xml:space="preserve"
            class="bg-blue-500 text-white w-9 h-8 py-2 px-0.5 rounded-lg hover:bg-sky-300"
          >
            <g>
              <g>
                <path
                  fill="white"
                  width="10px"
                  d="M409,39c-4.5-4.5-10.6-7-16.9-7H119.9c-6.4,0-12.4,2.5-16.9,7l0,0c-4.5,4.5-7,10.6-7,16.9v400.1
			c0,6.4,2.5,12.4,7,16.9l0,0c4.5,4.5,10.6,7,16.9,7h272.1c6.4,0,12.4-2.5,16.9-7l0,0c4.5-4.5,7-10.6,7-16.9V55.9
			C416,49.6,413.5,43.5,409,39L409,39z M255.6,48.7c3.9,0,7,3.1,7,7c0,3.9-3.1,7-7,7c-3.9,0-7-3.1-7-7
			C248.6,51.9,251.8,48.7,255.6,48.7z M256,470c-7.7,0-14-6.5-14-14.1c0-7.5,6.2-14,14-14c7.7,0,14.1,6.4,14.1,14
			C270,463.5,263.7,470,256,470z M400,432H112V80h288V432z"
                ></path>
              </g>
            </g>
          </svg>
        </button>
        <button class="ml-3">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            viewBox="0 0 24 24"
            class="bg-blue-500 text-white w-9 h-8 p-2 rounded-lg hover:bg-sky-300"
          >
            <path
              d="M23 18h-1V5c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v13H1c-.55 0-1 .45-1 1s.45 1 1 1h22c.55 0 1-.45 1-1s-.45-1-1-1zm-9.5 0h-3c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h3c.28 0 .5.22.5.5s-.22.5-.5.5zm6.5-3H4V6c0-.55.45-1 1-1h14c.55 0 1 .45 1 1v9z"
              fill="currentColor"
            ></path>
          </svg>
        </button>
      </div>
    </div>
    <div class="w-[20%] flex justify-end">
      <div class="mr-3">
        <button
          class="py-1.5 px-3 bg-green-500 text-white rounded-md transition-transform transform-gpu motion-safe:hover:scale-110"
          @click="saveData"
        >
          Save
        </button>
      </div>
      <div class="mr-3">
        <button
          class="py-1.5 px-3 bg-green-500 rounded-md text-white transition-transform transform-gpu motion-safe:hover:scale-110"
          @click="saveAndExit"
        >
          Save & Exit
        </button>
      </div>
      <div class="mr-2">
        <button
          class="py-1.5 px-3 bg-red-500 text-white rounded-md transition-transform transform-gpu motion-safe:hover:scale-110"
          @click="cancel"
        >
          <nuxt-link to="/">Cancel</nuxt-link>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

// const message = useMessage();

const options = [
  { value: "", text: "Chọn danh sách page" },
  { value: "home", text: "Trang chủ" },
  { value: "banner", text: "Banner" },
];

const selectedPage = ref({ value: "", text: "Chọn danh sách page" });
const selectedSubPage = ref(options[1]);

const subPageOptions = ["Trang chủ"];

const { data } = await useFetch("http://localhost:8000/page");
const response = data.value;

response.forEach((item) => {
  if (item.Metadata) {
    // console.log(item.Metadata);
  } else {
    // console.log(
    //   "Không tìm thấy thông tin Metadata trong một số phần tử của response."
    // );
  }
});

// Tạo một mảng để lưu thông tin chỉnh sửa
let editData = [];

const saveData = async () => {
  // Duyệt qua mảng editData để tạo các yêu cầu PATCH
  for (const data of editData) {
    const patchResponse = await fetch("http://localhost:8000/page", {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data.body),
    });

    if (!patchResponse.ok) {
      throw new Error("Failed to save data");
    }
  }

  console.log("Lưu thành công");
};

// Gọi hàm saveData với await để đợi cho các yêu cầu PATCH được thực hiện
await saveData();

const saveAndExit = () => {};

const cancel = () => {};

watch(
  () => selectedPage.value,
  (newValue) => {
    const result = options.find((o) => o.value == newValue.value);
    selectedSubPage.value = result;
  }
);
</script>

<style>
select:focus {
  outline: none;
  border: 2px solid #007bff;
  box-shadow: 0 0 5px #007bff;
}
</style>
