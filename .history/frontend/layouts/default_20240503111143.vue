<template>
  <div class="w-full">
    <Header :response="response" />
    <slot />
    <Footer :response="response" />
  </div>
</template>

<script setup>
const { data } = await useFetch("http://localhost:8000/page");
const response = data.value;

const data_form = response[0]?.Metadata?.data_form ?? {};
const data_menus = response[0]?.Metadata?.data_menus ?? {};
const data_footer = response[0]?.Metadata?.data_footer ?? {};
</script>
