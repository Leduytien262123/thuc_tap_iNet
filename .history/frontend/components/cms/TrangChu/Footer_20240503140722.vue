<template>
  <div class="w-[72%] m-auto mt-20 flex mb-12 relative">
    <div class="w-[35%]">
      <img class="w-36 h-24 ml-3 mr-32 mt-4 mb-5 flex" :src="logoUrl" />
      <div
        v-for="contact in contacts"
        :key="contact.id"
        class="flex items-center gap-2 text-lg"
      >
        <img class="w-4" :src="contact.iconUrl" />
        <nuxt-link :to="contact.link" class="py-1.5 text-[#2c2934] text-base">{{
          contact.text
        }}</nuxt-link>
      </div>
    </div>
    <div v-for="category in categories1" :key="category.id" class="w-[18%]">
      <span class="py-2"
        ><b class="text-lg">{{ category.title }}</b></span
      >
      <div class="mt-1">
        <div v-for="item in category.items" :key="item.id" class="py-1 mt-2.5">
          <nuxt-link :to="item.link" class="text-base">{{
            item.text
          }}</nuxt-link>
        </div>
      </div>
    </div>
    <div v-for="category in categories2" :key="category.id" class="w-[25%]">
      <span class="py-2"
        ><b class="text-lg">{{ category.title }}</b></span
      >
      <div class="mt-1">
        <div v-for="item in category.items" :key="item.id" class="py-1 mt-2.5">
          <nuxt-link :to="item.link" class="text-base">{{
            item.text
          }}</nuxt-link>
        </div>
      </div>
    </div>
    <div class="w-[20%]">
      <span class="py-2"
        ><b class="text-lg">{{ content.content2 }}</b></span
      >
      <div class="flex mt-4">
        <div v-for="social in socialMedia" :key="social.id" class="px-2">
          <nuxt-link :to="social.link" class="">
            <img class="w-8" :src="social.icon" :alt="social.alt" />
          </nuxt-link>
        </div>
      </div>
    </div>
    <div class="absolute -right-[165px] top-1/2 -translate-y-1/2 -mt-20">
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
            <div class="mb-4">
              <label for="editImage" class="block font-semibold">Logo:</label>
              <input
                type="file"
                id="editImage"
                accept="image/*"
                class="w-full border rounded p-2 border-gray-500"
                @change="handleEditLogoUrl($event)"
                @keydown.enter.prevent="handleEnterKey"
              />
            </div>
            <n-collapse arrow-placement="right" class="mt-3 flex items-center">
              <n-collapse-item
                :title="'Chỉnh sửa:' + ' ' + content.content1"
                class="w-full"
              >
                <n-collapse
                  arrow-placement="right"
                  v-for="item in editData.editContacts"
                  :key="item.id"
                  class="mt-3 flex items-center"
                >
                  <n-collapse-item
                    :title="item.editText"
                    :name="item.id"
                    class="w-full"
                  >
                    <div class="mb-4">
                      <label for="editImageUrl" class="block font-semibold"
                        >Icon:</label
                      >
                      <input
                        type="file"
                        id="editImageUrl"
                        accept="image/*"
                        class="w-full border rounded p-2 border-gray-500"
                        @change="(event) => handleEditIconUrl(event, item.id)"
                        @keydown.enter.prevent="handleEnterKey"
                      />
                    </div>
                    <div class="mb-4">
                      <label
                        :for="'editName' + item.id"
                        class="block font-semibold"
                        >Name</label
                      >
                      <input
                        :id="'editName' + item.id"
                        class="w-full border rounded p-2"
                        v-model="item.editText"
                        @keydown.enter.prevent="handleEnterKey"
                      />
                    </div>
                    <div class="mb-4">
                      <label
                        :for="'editName' + item.id"
                        class="block font-semibold"
                        >Link</label
                      >
                      <input
                        :id="'editName' + item.id"
                        class="w-full border rounded p-2"
                        v-model="item.editLinkContacts"
                        @keydown.enter.prevent="handleEnterKey"
                      />
                    </div>
                  </n-collapse-item>
                  <button
                    @click="removeContacts(item.id)"
                    class="self-start ml-auto"
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
                </n-collapse>
                <button
                  @click="editAddContacts = true"
                  type="success"
                  class="self-start ml-auto float-right mr-0 mt-4"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    xmlns:xlink="http://www.w3.org/1999/xlink"
                    viewBox="0 0 24 24"
                    class="w-8 h-8 text-green-500 items-center ml-2 mt-2"
                  >
                    <path
                      d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 16H5V5h14v14zm-8-2h2v-4h4v-2h-4V7h-2v4H7v2h4z"
                      fill="currentColor"
                    ></path>
                  </svg>
                </button>
              </n-collapse-item>
            </n-collapse>

            <n-collapse
              arrow-placement="right"
              v-for="item in editData.editCategories1"
              :key="item.id"
              class="mt-3 flex items-center"
            >
              <n-collapse-item
                :title="'Chỉnh sửa:' + ' ' + item.editTitle"
                :name="item.id"
                class="w-full"
              >
                <n-collapse
                  v-for="child in item.items"
                  :key="child.id"
                  class="flex items-center mb-3"
                >
                  <n-collapse-item
                    :title="child.editName"
                    :name="child.id"
                    class="w-full"
                  >
                    <div class="w-full">
                      <div class="mb-4">
                        <label
                          :for="'editName' + child.id"
                          class="block font-semibold"
                          >Name</label
                        >
                        <input
                          :id="'editName' + child.id"
                          class="w-full border rounded p-2"
                          v-model="child.editName"
                          @keydown.enter.prevent="handleEnterKey"
                        />
                      </div>
                      <div class="mb-4">
                        <label
                          :for="'editName' + child.id"
                          class="block font-semibold"
                          >Link</label
                        >
                        <input
                          :id="'editName' + child.id"
                          class="w-full border rounded p-2"
                          v-model="child.editLink"
                          @keydown.enter.prevent="handleEnterKey"
                        />
                      </div>
                    </div>
                  </n-collapse-item>
                  <button
                    @click="removeCategories1(child.id)"
                    class="self-start ml-auto"
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
                </n-collapse>
                <button
                  @click="editAddCategories1 = true"
                  type="success"
                  class="self-start ml-auto float-right mr-0 mt-4"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    xmlns:xlink="http://www.w3.org/1999/xlink"
                    viewBox="0 0 24 24"
                    class="w-8 h-8 text-green-500 items-center ml-2 mt-2"
                  >
                    <path
                      d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 16H5V5h14v14zm-8-2h2v-4h4v-2h-4V7h-2v4H7v2h4z"
                      fill="currentColor"
                    ></path>
                  </svg>
                </button>
              </n-collapse-item>
            </n-collapse>

            <n-collapse
              arrow-placement="right"
              v-for="item in editData.editCategories2"
              :key="item.id"
              class="mt-3 flex items-center"
            >
              <n-collapse-item
                :title="'Chỉnh sửa:' + ' ' + item.editTitle"
                :name="item.id"
                class="w-full"
              >
                <n-collapse
                  v-for="child in item.items"
                  :key="child.id"
                  class="flex items-center"
                >
                  <n-collapse-item
                    :title="child.editName"
                    :name="child.id"
                    class="w-full"
                  >
                    <div class="w-full">
                      <div class="mb-4">
                        <label
                          :for="'editName' + child.id"
                          class="block font-semibold"
                          >Name</label
                        >
                        <input
                          :id="'editName' + child.id"
                          class="w-full border rounded p-2"
                          v-model="child.editName"
                          @keydown.enter.prevent="handleEnterKey"
                        />
                      </div>
                      <div class="mb-4">
                        <label
                          :for="'editName' + child.id"
                          class="block font-semibold"
                          >Link</label
                        >
                        <input
                          :id="'editName' + child.id"
                          class="w-full border rounded p-2"
                          v-model="child.editLink"
                          @keydown.enter.prevent="handleEnterKey"
                        />
                      </div>
                    </div>
                  </n-collapse-item>
                  <button
                    @click="removeCategories2(child.id)"
                    class="self-start ml-auto mt-5"
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
                </n-collapse>
                <button
                  @click="editAddCategories2 = true"
                  type="success"
                  class="self-start ml-auto float-right mr-0 mt-4"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    xmlns:xlink="http://www.w3.org/1999/xlink"
                    viewBox="0 0 24 24"
                    class="w-8 h-8 text-green-500 items-center ml-2 mt-2"
                  >
                    <path
                      d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 16H5V5h14v14zm-8-2h2v-4h4v-2h-4V7h-2v4H7v2h4z"
                      fill="currentColor"
                    ></path>
                  </svg>
                </button>
              </n-collapse-item>
            </n-collapse>

            <n-collapse arrow-placement="right" class="mt-3">
              <n-collapse-item
                :title="'Chỉnh sửa:' + ' ' + content.content2"
                class="w-full"
              >
                <n-collapse
                  arrow-placement="right"
                  v-for="item in editData.editSocialMedia"
                  :key="item.id"
                  class="mt-3 flex items-center"
                >
                  <n-collapse-item
                    :title="item.editAlt"
                    :name="item.id"
                    class="w-full"
                  >
                    <div class="mb-4">
                      <label
                        :for="'editName' + item.id"
                        class="block font-semibold"
                        >Name</label
                      >
                      <input
                        :id="'editName' + item.id"
                        class="w-full border rounded p-2"
                        v-model="item.editAlt"
                        @keydown.enter.prevent="handleEnterKey"
                      />
                    </div>
                    <div class="mb-4">
                      <label
                        :for="'editName' + item.id"
                        class="block font-semibold"
                        >Link</label
                      >
                      <input
                        :id="'editName' + item.id"
                        class="w-full border rounded p-2"
                        v-model="item.editLinkSocialMedia"
                        @keydown.enter.prevent="handleEnterKey"
                      />
                    </div>
                    <div class="mb-4">
                      <label for="editIcon" class="block font-semibold"
                        >Icon:</label
                      >
                      <input
                        type="file"
                        id="editIcon"
                        accept="image/*"
                        class="w-full border rounded p-2 border-gray-500"
                        @change="(event) => handleEditIcon(event, item.editAlt)"
                        @keydown.enter.prevent="handleEnterKey"
                      />
                    </div>
                  </n-collapse-item>

                  <button
                    @click="removeSocialMedia(item.id)"
                    class="self-start ml-auto"
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
                </n-collapse>
                <button
                  @click="editAddSocialMedia = true"
                  type="success"
                  class="self-start ml-auto float-right mr-0 mt-4"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    xmlns:xlink="http://www.w3.org/1999/xlink"
                    viewBox="0 0 24 24"
                    class="w-8 h-8 text-green-500 items-center ml-2 mt-2"
                  >
                    <path
                      d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 16H5V5h14v14zm-8-2h2v-4h4v-2h-4V7h-2v4H7v2h4z"
                      fill="currentColor"
                    ></path>
                  </svg>
                </button>
              </n-collapse-item>
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

        <div class="">
          <n-modal v-model:show="editAddContacts">
            <n-card
              style="width: 600px"
              title="Thêm mới"
              :bordered="false"
              size="huge"
              role="dialog"
              aria-modal="true"
            >
              <template #header-extra></template>
              <div class="h-full pr-8 overflow-y-auto">
                <div class="mb-4">
                  <label for="editIcon" class="block font-semibold"
                    >Icon:</label
                  >
                  <input
                    type="file"
                    id="editIcon"
                    accept="image/*"
                    class="w-full border rounded p-2 border-gray-500"
                    @change="newIconUrl($event)"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
                <div class="mb-4">
                  <label
                    for="menuLessons"
                    class="block text-gray-800 font-semibold"
                    >Name:</label
                  >
                  <input
                    type="text"
                    id="menuLessons"
                    v-model="newContacts.newText"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Name"
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
                    v-model="newContacts.newLinkContacts"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Liên kết đến đối tác"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
              </div>
              <template #footer>
                <div class="text-center">
                  <button
                    @click="addNewContacts"
                    class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-6 rounded-md mr-3"
                  >
                    Thêm mới
                  </button>
                  <button
                    @click="editAddNewContacts"
                    class="bg-gray-500 hover:bg-gray-600 text-white font-semibold py-2 px-6 rounded-md"
                  >
                    Hủy
                  </button>
                </div>
              </template>
            </n-card>
          </n-modal>
        </div>

        <!-- -------------------------------------------------------------------------------------------------- -->

        <div class="">
          <n-modal v-model:show="editAddCategories1">
            <n-card
              style="width: 600px"
              title="Thêm mới"
              :bordered="false"
              size="huge"
              role="dialog"
              aria-modal="true"
            >
              <template #header-extra></template>
              <div class="h-full pr-8 overflow-y-auto">
                <div class="mb-4">
                  <label
                    for="menuLessons"
                    class="block text-gray-800 font-semibold"
                    >Name:</label
                  >
                  <input
                    type="text"
                    id="menuLessons"
                    v-model="newCategories1.newName"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Name"
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
                    v-model="newCategories1.newLink"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Liên kết đến đối tác"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
              </div>
              <template #footer>
                <div class="text-center">
                  <button
                    @click="addNewCategories1"
                    class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-6 rounded-md mr-3"
                  >
                    Thêm mới
                  </button>
                  <button
                    @click="editAddNewCategories1"
                    class="bg-gray-500 hover:bg-gray-600 text-white font-semibold py-2 px-6 rounded-md"
                  >
                    Hủy
                  </button>
                </div>
              </template>
            </n-card>
          </n-modal>
        </div>

        <!-- -------------------------------------------------------------------------------------------------- -->

        <div class="">
          <n-modal v-model:show="editAddCategories2">
            <n-card
              style="width: 600px"
              title="Thêm mới"
              :bordered="false"
              size="huge"
              role="dialog"
              aria-modal="true"
            >
              <template #header-extra></template>
              <div class="h-full pr-8 overflow-y-auto">
                <div class="mb-4">
                  <label
                    for="menuLessons"
                    class="block text-gray-800 font-semibold"
                    >Name:</label
                  >
                  <input
                    type="text"
                    id="menuLessons"
                    v-model="newCategories2.newName"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Name"
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
                    v-model="newCategories2.newLink"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Liên kết đến đối tác"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
              </div>
              <template #footer>
                <div class="text-center">
                  <button
                    @click="addNewCategories2"
                    class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-6 rounded-md mr-3"
                  >
                    Thêm mới
                  </button>
                  <button
                    @click="editAddNewCategories2"
                    class="bg-gray-500 hover:bg-gray-600 text-white font-semibold py-2 px-6 rounded-md"
                  >
                    Hủy
                  </button>
                </div>
              </template>
            </n-card>
          </n-modal>
        </div>

        <!-- -------------------------------------------------------------------------------------------------- -->

        <div class="">
          <n-modal v-model:show="editAddSocialMedia">
            <n-card
              style="width: 600px"
              title="Thêm mới"
              :bordered="false"
              size="huge"
              role="dialog"
              aria-modal="true"
            >
              <template #header-extra></template>
              <div class="h-full pr-8 overflow-y-auto">
                <div class="mb-4">
                  <label for="editIcon" class="block font-semibold"
                    >Icon:</label
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
                    >Alt:</label
                  >
                  <input
                    type="text"
                    id="menuLessons"
                    v-model="newSocialMedia.newAlt"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Alt"
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
                    v-model="newSocialMedia.newLink"
                    class="w-full border rounded p-2 border-gray-500"
                    placeholder="Liên kết đến đối tác"
                    @keydown.enter.prevent="handleEnterKey"
                  />
                </div>
              </div>
              <template #footer>
                <div class="text-center">
                  <button
                    @click="addNewSocialMedia"
                    class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-6 rounded-md mr-3"
                  >
                    Thêm mới
                  </button>
                  <button
                    @click="editAddNewSocialMedia"
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

const handleEnterKey = (event) => {
  if (event.key === "Enter") {
    saveChanges();
  }
};

const logoUrls = defineModel("logoUrls");
const contents = defineModel("contents");
const contact = defineModel("contact");
const categories1s = defineModel("categories1s");
const categories2s = defineModel("categories2s");
const socialMedias = defineModel("socialMedias");

const props = defineProps({
  response: {
    type: Object,
    default: {},
  },
});

const logoUrl = ref(logoUrls.value);
const content = ref(contents.value);
const contacts = ref(props.response[0].Metadata.data_footer[2].contacts);
const categories1 = ref(props.response[0].Metadata.data_footer[3].categories1);
const categories2 = ref(props.response[0].Metadata.data_footer[4].categories2);
const socialMedia = ref(props.response[0].Metadata.data_footer[5].socialMedia);

const editMode = ref(false);
const editAddContacts = ref(false);
const editAddCategories1 = ref(false);
const editAddCategories2 = ref(false);
const editAddSocialMedia = ref(false);

const editData = ref({
  editLogoUrl: logoUrl.value,
  editContent1: content.value.content1,
  editContent2: content.value.content2,
  editContacts: [],
  editCategories1: [],
  editCategories2: [],
  editSocialMedia: [],
});

contacts.value.forEach((contact) => {
  editData.value.editContacts.push({
    id: contact.id,
    editIconUrl: contact.iconUrl,
    editText: contact.text,
    editLinkContacts: contact.link,
  });
});

categories1.value.forEach((category) => {
  editData.value.editCategories1.push({
    editTitle: category.title,
    items: category.items.map((item) => ({
      id: item.id,
      editName: item.text,
      editLink: item.link,
    })),
  });
});

categories2.value.forEach((category) => {
  editData.value.editCategories2.push({
    editTitle: category.title,
    items: category.items.map((item) => ({
      id: item.id,
      editName: item.text,
      editLink: item.link,
    })),
  });
});

// console.log(editData.value);

socialMedia.value.forEach((socialMedia) => {
  editData.value.editSocialMedia.push({
    id: socialMedia.id,
    editIcon: socialMedia.icon,
    editAlt: socialMedia.alt,
    editLinkSocialMedia: socialMedia.link,
  });
});

const saveChanges = () => {
  editMode.value = false;
  logoUrl.value = editData.value.editLogoUrl;
  logoUrls.value = logoUrl.value;
  content.value.content1 = editData.value.editContent1;
  content.value.content2 = editData.value.editContent2;
  contacts.value = editData.value.editContacts.map((contact) => ({
    id: contact.id,
    iconUrl: contact.editIconUrl,
    text: contact.editText,
    link: contact.editLinkContacts,
  }));
  contact.value = contacts.value;
  categories1.value = editData.value.editCategories1.map((category) => ({
    title: category.editTitle,
    items: category.items.map((item) => ({
      id: item.id,
      text: item.editName,
      link: item.editLink,
    })),
  }));
  categories1s.value = categories1.value;
  console.log(categories1.value);
  categories2.value = editData.value.editCategories2.map((category) => ({
    title: category.editTitle,
    items: category.items.map((item) => ({
      id: item.id,
      text: item.editName,
      link: item.editLink,
    })),
  }));
  // console.log(editData.value);
  socialMedia.value = editData.value.editSocialMedia.map((socialMedia) => ({
    id: socialMedia.id,
    icon: socialMedia.editIcon,
    alt: socialMedia.editAlt,
    link: socialMedia.editLinkSocialMedia,
  }));
  message.success("Lưu thành công");
};

const cancelEdit = () => {
  editMode.value = false;

  editData.value.editLogoUrl = logoUrl.value;
  editData.value.editContent1 = content.value.content1;
  editData.value.editContent2 = content.value.content2;
  editData.value.editContacts = contacts.value.map((contact) => ({
    id: contact.id,
    editIconUrl: contact.iconUrl,
    editText: contact.text,
    editLinkContacts: contact.link,
  }));
  editData.value.editCategories1 = categories1.value.map((category) => ({
    editTitle: category.title,
    items: category.items.map((item) => ({
      id: item.id,
      editName: item.text,
      editLink: item.link,
    })),
  }));
  editData.value.editCategories2 = categories2.value.map((category) => ({
    editTitle: category.title,
    items: category.items.map((item) => ({
      id: item.id,
      editName: item.text,
      editLink: item.link,
    })),
  }));
  editData.value.editSocialMedia = socialMedia.value.map((socialMedia) => ({
    id: socialMedia.id,
    editIcon: socialMedia.icon,
    editAlt: socialMedia.alt,
    editLinkSocialMedia: socialMedia.link,
  }));
};

const handleEditLogoUrl = (event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      editData.value.editLogoUrl = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const handleEditIconUrl = (event, target) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      let index = editData.value.editContacts.findIndex(
        (x) => x.editAlt === target
      );
      if (index != -1) {
        editData.value.editContacts[index].editIconUrl = e.target.result;
      }
    };
    reader.readAsDataURL(file);
  }
};

const handleEditIcon = (event, target) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      let index = editData.value.editSocialMedia.findIndex(
        (x) => x.editAlt === target
      );
      if (index != -1) {
        editData.value.editSocialMedia[index].editIcon = e.target.result;
      }
    };
    reader.readAsDataURL(file);
  }
};

// -------------------------------------------------------------------------------------

const uploadIconUrl = (itemKey, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      newContacts.value.newIconUrl = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const newIconUrl = (event) => {
  uploadIconUrl("newIconUrl", event);
};

const uploadIcon = (itemKey, event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      newSocialMedia.value.newIcon = e.target.result;
    };
    reader.readAsDataURL(file);
  }
};

const newIcon = (event) => {
  uploadIcon("newIcon", event);
};

const newContacts = ref({
  newIconUrl: "",
  newText: "",
  newLinkContacts: "",
});

function addNewContacts() {
  if (!newContacts.value.newIconUrl || !newContacts.value.newText) {
    alert("Vui lòng nhập đầy đủ thông tin");
    return; // Dừng hàm
  }

  editData.value.editContacts.push({
    id: editData.value.editContacts.length + 1,
    editIconUrl: newContacts.value.newIconUrl,
    editText: newContacts.value.newText,
    editLinkContacts: newContacts.value.newLinkContacts,
  });

  newContacts.value = {
    newIconUrl: "",
    newText: "",
    newLinkContacts: "",
  };
  message.success("Thêm mới thành công");

  editAddContacts.value = false;
}

const newCategories1 = ref({
  newName: "",
  newLink: "",
});

function addNewCategories1() {
  if (!newCategories1.value.newName) {
    alert("Vui lòng nhập đầy đủ thông tin");
    return;
  }
  editData.value.editCategories1[0].items.push({
    id: editData.value.editCategories1[0].items.length + 1,
    editName: newCategories1.value.newName,
    editLink: newCategories1.value.newLink,
  });

  newCategories1.value = {
    newName: "",
    newLink: "",
  };
  message.success("Thêm mới thành công");

  editAddCategories1.value = false;
}

const newCategories2 = ref({
  newName: "",
  newLink: "",
});

function addNewCategories2() {
  console.log(addNewCategories2);

  if (!newCategories2.value.newName) {
    alert("Vui lòng nhập đầy đủ thông tin");
    return;
  }
  // console.log(editAddContacts);

  editData.value.editCategories2[0].items.push({
    id: editData.value.editCategories2[0].items.length + 1,
    editName: newCategories2.value.newName,
    editLink: newCategories2.value.newLink,
  });

  newCategories2.value = {
    newName: "",
    newLink: "",
  };
  message.success("Thêm mới thành công");

  editAddCategories2.value = false;
}

const newSocialMedia = ref({
  newIcon: "",
  newAlt: "",
  newLink: "",
});

function addNewSocialMedia() {
  if (!newSocialMedia.value.newIcon || !newSocialMedia.value.newAlt) {
    alert("Vui lòng nhập đầy đủ thông tin");
    return;
  }

  editData.value.editSocialMedia.push({
    id: editData.value.editSocialMedia.length + 1,
    editIcon: newSocialMedia.value.newIcon,
    editAlt: newSocialMedia.value.newAlt,
    editLinkSocialMedia: newSocialMedia.value.newLink,
  });

  newSocialMedia.value = {
    newIcon: "",
    newAlt: "",
    newLink: "",
  };
  message.success("Thêm mới thành công");

  editAddSocialMedia.value = false;
}

function editAddNewContacts() {
  newContacts.value.newIconUrl = "";
  newContacts.value.newText = "";
  newContacts.value.newLinkContacts = "";

  editAddContacts.value = false;
}

function editAddNewCategories1() {
  newCategories1.value.newName = "";
  newCategories1.value.newLink = "";

  editAddCategories1.value = false;
}

function editAddNewCategories2() {
  newCategories2.value.newName = "";
  newCategories2.value.newLink = "";

  editAddCategories2.value = false;
}

function editAddNewSocialMedia() {
  newSocialMedia.value.newIcon = "";
  newSocialMedia.value.newAlt = "";
  newSocialMedia.value.newLink = "";

  editAddSocialMedia.value = false;
}

function removeContacts(id) {
  editData.value.editContacts = editData.value.editContacts.filter(function (
    item
  ) {
    return item.id !== id;
  });
}

function removeCategories1(id) {
  editData.value.editCategories1[0].items =
    editData.value.editCategories1[0].items.filter(function (item) {
      return item.id !== id;
    });
}

function removeCategories2(id) {
  editData.value.editCategories2[0].items =
    editData.value.editCategories2[0].items.filter(function (item) {
      return item.id !== id;
    });
}

function removeSocialMedia(id) {
  editData.value.editSocialMedia = editData.value.editSocialMedia.filter(
    function (item) {
      return item.id !== id;
    }
  );
}
</script>
