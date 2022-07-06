<template>
  <div class="app__container">
    <my-navbar />
    <router-view />
    <my-footer v-show="false" />
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