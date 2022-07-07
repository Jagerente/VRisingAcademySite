<template>
  <div class="app__container">
    <Navbar />
    <router-view />
    <my-footer v-show="false" />
  </div>
</template>

<script>
import Navbar from "@/components/Navbar/Navbar.vue";
import MyFooter from "@/components/MyFooter.vue";

export default {
  components: {
    Navbar,
    MyFooter
  },
};
</script>

<style lang="scss">
body {
  font-family: $mainFontFamily;
  background-color: $background;
  color: $text-color;
}

.app__container {
  display: flex;
  flex-direction: column;
}
</style>