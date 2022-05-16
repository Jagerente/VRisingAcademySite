<template>
  <div class="container-fluid vh-100 p-0">
    <navbar />
    <router-view></router-view>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar.vue";
export default {
  components: {
    Navbar,
  },
};
</script>

<style lang="scss">


</style>