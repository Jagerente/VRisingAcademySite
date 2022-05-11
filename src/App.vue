<template>
  <navbar />
  <div>
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
.test {
  background-color: aqua;
  border-color: red;
  border: 3px dashed red;
}
/* * {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
} */

// body {
//   background-color: #1a1927;
//   color: #a9a8ae;
// }

// @import "../node_modules/bootstrap/scss/bootstrap";
// @import "assets/vrising.scss";
</style>