<template>
  <main class="container-fluid vh-100 p-0 m-0">
    <navbar />
    <router-view></router-view>
  </main>

</template>

<script>
import Navbar from "@/components/Navbar.vue";
export default {
  components: {
    Navbar,
  },
};
</script>

<style>
</style>