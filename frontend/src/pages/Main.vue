<template>
  <div class="content">
    <div class="d-flex flex-column my-auto h-100">
      <!-- Logo -->
      <img class="mx-auto" style="" draggable="false" :src="
        require('@/assets/images/ui/main_page/VAlogo.webp')
      ">
      <!-- Text -->
      <div class="d-flex info mx-auto">
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
      </div>
    </div>
  </div>
  <my-footer />

</template>

<script>
import MyFooter from "@/components/MyFooter.vue";

export default {
  components: {
    MyFooter,
  },
};
</script>

<style lang="scss" scoped>
* {
  user-select: none;
}

@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';

@include media-breakpoint-down(lg) {
  .info {
    width: 90%;
  }

  img {
    width: 300px;
    margin-bottom: -75px;
    margin-top: -50px;
  }
}

@include media-breakpoint-up(lg) {
  .info {
    width: 45%;
  }

  img {
    width: 100%;
    max-width: 500px;
    margin-bottom: -150px;
    margin-top: -100px;
  }

  h3 {
    font-size: 1.5em;
  }
}

.content {
  padding-left: auto;
  padding-right: auto;
  margin-left: auto;
  margin-right: auto;
}
</style>