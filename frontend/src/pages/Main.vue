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
        </p>

        <div class="d-flex flex-column justify-content-center">
          <button @click="gotoBoosty()" class="btn btn-primary my-2 mx-auto">Support
            Us!</button>
          <p class="text-center">Support us monetarily to pay for server hosting and keep this site
            ad-free.</p>
        </div>
        <div class="d-flex flex-column w-100 justify-content-center">
          <button @click="$router.push('/contribute')" class="btn btn-primary my-2 mx-auto">Contribute</button>
          <p class="text-center">If you want to help fill up the database, there is a number of topics we
            seek data about.</p>
        </div>

        <p class="mt-5">
        <h3 class="text-white">Changelog</h3>
        <h5 class="mt-2 mb-0 text-white">07/06/2022 Items Update</h5>
        <br>— Added Structures and Blueprints (Books)
        <br>— Structures display their Variants which are structures with the same recipe and essentially represent the
        same
        <br>structure but with a different color/look and are unlocked simultaneously.
        <br>— Reworked Items Catalogue. Items now are split into more distinct categories.
        <br>— Items are divided into what we call Sets for easier navigation.
        <br>— Added Recipes' crafting times.
        <br>— Added options to display Recipes' time and cost discounts in Confined Rooms and with Matching Floor.
        <br>— Reworked search behaviour. Search by tag now is only enabled when the full name of tags is written. You
        can
        now also search by tag only ignoring names by clicking on the tag in the Information section.
        <br>— Added several missing items, including Soul Shards, Acorn, Razer Hood.
        <br>— Some minor fixes to ordering of several items. Notably, now cloak are ordered by their tier and
        alphabetically
        within their tier, hats are also order alphabetically, fish is ordered by their tier (as per their salvage
        returns).
        <br>— A couple of fixes to items' information like Explosive Boxes had their images swapped and some other minor
        errors.
        <br>— Added Contribute page.
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
  <my-footer class="mt-2" />

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