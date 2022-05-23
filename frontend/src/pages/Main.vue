<template>
  <div class="content h-100">
    <div class="d-flex flex-column my-auto h-100" style="padding-top: 50px;">
      <!-- Logo -->
      <img class="mx-auto" draggable="false" :src="
        require('@/assets/images/ui/main_page/VAlogo.webp')
      ">
      <!-- Text -->
      <h3 class="text-center">Welcome to the academy, V Academy!
        <br>This website is Work In Progress, and we are
        working
        hard to fill up the
        database with missing items and develop new features.
        <br>If you want to contribute or leave feedback you can do
        it in <a
          href="https://docs.google.com/forms/d/e/1FAIpQLSdNlAa_hyXcHsMTJy59alBP5hxhiWsebMHNr9M52He0s5mUsQ/viewform?usp=sf_link"
          target="blank">this form.</a>
      </h3>
      <!-- Footer -->
      <my-footer class="mt-auto" />
    </div>
  </div>
</template>

<script>
import MyFooter from "@/components/MyFooter";

export default {
  components: {
    MyFooter
  }
};
</script>

<style lang="scss" scoped>
* {
  user-select: none;
}

@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';

@include media-breakpoint-down(sm) {
  .content {
    width: 90%;
  }

  img {
    width: 300px;
    margin-bottom: -75px;
  }
}

@include media-breakpoint-up(sm) {
  .content {
    width: 45%;
  }

  img {
    width: 100%;
    max-width: 500px;
    margin-bottom: -175px;
  }

  h3 {
    margin-top: 100px;
    font-size: 2em;
  }
}

.content {
  padding-left: auto;
  padding-right: auto;
  margin-left: auto;
  margin-right: auto;
}
</style>