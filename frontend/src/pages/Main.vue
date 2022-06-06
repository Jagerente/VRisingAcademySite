<template>
  <div class="content">
    <div class="d-flex flex-column my-auto h-100">
      <!-- Logo -->
      <img class="mx-auto" style="" draggable="false" :src="
        require('@/assets/images/ui/main_page/VAlogo.webp')
      ">
      <!-- Text -->
      <div class="d-flex flex-column info mx-auto">
        <p class="text-center">Welcome to the academy, V Academy!
          <br>This website is Work In Progress, and we are
          working
          hard to fill up the
          database with missing items and develop new features.
          <br>If you want to contribute or leave feedback you can do
          it in <a
            href="https://docs.google.com/forms/d/e/1FAIpQLSdNlAa_hyXcHsMTJy59alBP5hxhiWsebMHNr9M52He0s5mUsQ/viewform?usp=sf_link"
            target="blank">this form.</a>
        </p>

        <div class="d-flex flex-column w-100 justify-content-center">
          <!-- <Popper hover content="<3"> -->
          <button @click="gotoBoosty()" class="btn btn-primary my-2">Support
            Us!</button>
          <p class="text-center text-white">Support us monetarily to pay for server hosting and keep this site
            ad-free.</p>
          <!-- </Popper> -->
        </div>
        <p>
        <h3 class="text-white">Changelog</h3>
        <h5 class="mt-2 mb-0 text-white">31/05/2022 Spells Update!</h5>
        <br>— Compete catalogue of all spells including weapon abilities
        <br>— For each spell their cooldown, cast time and number of charges is specified (according to the ingame
        tooltip).
        <br>— For vampire spells cast times were determined manually by recording a video of spell cast and counting
        frames, approximate values are shown on the website.
        <h5 class="mt-2 mb-0 text-white">24/05/2022</h5>
        <br>— added a variety (34) of hats and capes (big thanks to u/Difinitus)
        <br>— added Blood Key
        <br>— some minor fixes to tags of particular items
        <h5 class="mt-2 mb-0 text-white">21/05/2022</h5>
        <br>— Launched the website with the folowing features:
        <br>— items catalouge excluding sturctures and the majority of hats and cloaks
        <br>— search items by name
        <br>— search items by tags
        <br>— clickable tags to initiate and instant search by tag name
        <br>— display all recipes that a chosen item can be crafted from
        <br>— display all items that a chosen item is reagent for
        <br>— clickable items in the "recipe" and "reagent for" sections to display their information
        </p>
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
  methods: {
    gotoBoosty() {
      window.open('https://boosty.to/vrisingacademy', '_blank')
    }
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