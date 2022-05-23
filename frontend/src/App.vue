<template>
  <div class="container-fluid vh-100 p-0 m-0">
    <my-navbar />
    <router-view />
  </div>
</template>

<script>
import MyNavbar from "@/components/MyNavbar.vue";
export default {
  components: {
    MyNavbar,
  },
};
</script>

<style>
</style>