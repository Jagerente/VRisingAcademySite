<template>
  <div class="d-flex flex-column h-100">
    <my-navbar />
    <router-view />
    <my-footer v-if="false" />
  </div>
</template>

<script>
import MyNavbar from "@/components/MyNavbar.vue";
import MyFooter from "./components/MyFooter.vue";

export default {
  components: {
    MyNavbar,
    MyFooter
  },
};
</script>

<style scoped>
</style>