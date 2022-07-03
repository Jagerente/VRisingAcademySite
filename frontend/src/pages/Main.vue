<template>
  <div class="container">
    <!-- Logo -->
    <img
      class="container__logo"
      draggable="false"
      :src="
        require('@/assets/images/ui/main_page/VAlogo.webp')
      "
    >
    <!-- Text -->
    <div class="description container__description">
      <div class="block description__block">Welcome to the academy, V Academy!
        <p class="block__text">
          This website is Work In Progress, and we are
          working
          hard to fill up the
          database with missing items and develop new features.
        </p>
      </div>
      <div class="description__block">
        <button
          class="button block__button"
          @click="gotoBoosty()"
        >
          Support Us!
        </button>
        <p class="block__text">
          Support us monetarily to pay for server hosting and keep this site
          ad-free.
        </p>
      </div>
      <div class="description__block">
        <button
          @click="$router.push('/contribute')"
          class="button block__button"
        >
          Contribute
        </button>
        <p class="block__text">
          If you want to help fill up the database, there is a number of topics we
          seek data about.
        </p>
      </div>
    </div>
    <div class="changelog container__changelog">
      <h2 class="changelog__header">Changelog</h2>
      <div class="log changelog__log">
        <h3 class="log__title">16/06/2022 "Show on Map" Update</h3>
        <div class="log__body">
          <p>— Some items now have a "Show on Map" button, that opens Mapgenie with highlighted areas where the chosen
            item can be farmed. It shows all areas where ingame map tooltips suggest that an item can be gathered.</p>
          <p>— Max Stack: all items now show how many instances of an item can be stored in a single stack (on official
            servers).</p>
          <p>— Added info on Razer items which can be obtained in Razer giveaways.</p>
          <p>— Added Castle Heart upgrades as separte items as well as their characteristics (on official servers).</p>
          <p>— Updated Spells description due to today's Balance Patch.</p>
        </div>
      </div>
      <div class="log changelog__log">
        <h3 class="log__title">12/06/2022 — Blood Update</h3>
        <div class="log__body">
          <p>
            — Added a page for <a
              class="log__link"
              href="bloodtypes"
              target="_blank"
            >blood types</a>.
          </p>
          <p> — This page is interactive and you can see how blood quality affects the abilities of a certain blood
            type.</p>
          <p>
            — Added text formating to <a
              class="log__link"
              href="spells"
              target="_blank"
            >spells</a>. Descriptions should be now easier to read.
          </p>
          <p>
            — Added Chaos Burn description to all related <a
              class="log__link"
              href="spells"
              target="_blank"
            >spells</a> based on <a
              class="log__link"
              href="https://docs.google.com/spreadsheets/d/1kdUng_E6RC0l7MxWQ7VaqZpTDz7DwDI6yTZHD4ONt5A/edit?usp=sharing"
              target="_blank"
            >this testing.</a>
          </p>
          <p>
            — Scourge Pendant's list of abilities was incorrect (it doesn't have any abilities), it has been fixed
            now.
          </p>
        </div>
      </div>
      <div class="log changelog__log">
        <h3 class="log__title">09/06/2022 Salvage Update</h3>
        <div class="log__body">
          <p>— Salvageable items now display items that they can be salvaged into. Now you can witness the true power of
            The Fish!</p>
          <p>— Items now also display whether you can obtain them from salvaging other items.</p>
          <p>— Note that weapons, armour and jewellery are not included since their salvage results depend on their
            durability.</p>
          <p>— Other fixes:</p>
          <p>— Sanguine Crossbow's crafting time is fixed.</p>
          <p>— Royal Braziers' recipes had all incorrectly required Regular Amethyst. This has been fixed.</p>
          <p>— Tainted Heart, Unsullied Heart and Exquisite Heart's descriptions have been updated to include the
            phrase "Refined into X by using a Blood Press" that were added into the game in Patch 0.5.41237.</p>
        </div>
      </div>
      <div class="log changelog__log">
        <h3 class="log__title">07/06/2022 Items Update</h3>
        <div class="log__body">
          <p>— Added Structures and Blueprints (Books)</p>
          <p>— Structures display their Variants which essentially represent the same structure with the same recipe
            and
            is unlocked at the same time but with a different color/look.</p>
          <p>— Please note that currently Variants do not have recipes displayed. Please reference the recipe of their
            "parent" Variant.</p>
          <p>— Reworked Items Catalogue. Items now are split into more distinct categories.</p>
          <p>— Items are divided into what we call Sets for easier navigation.</p>
          <p>— Added Recipes' crafting times.</p>
          <p>— Added options to display Recipes' time and cost discounts in Confined Rooms and with Matching Floor.</p>
          <p>— Reworked search behaviour. Search by tag now is only enabled when the full name of tags is written. You
            can now also search by tag only ignoring names by clicking on the tag in the Information section.</p>
          <p>— Some minor fixes to ordering of several items. Notably, now cloak are ordered by their tier and
            alphabetically within their tier, hats are also order alphabetically, fish is ordered by their tier (as per
            their salvage returns).</p>
          <p>— Added several missing items, including Soul Shards, Acorn, Razer Hood.</p>
          <p>— Several major fixes to items' information:</p>
          <p>— Major Explosive Box and Minor Explosive Box had their images swapped.</p>
          <p>— Copper Castle Key and General's Soul Reaper recipes displayed discounted amount of resources instead of
            base which was inconsitent with the other recipes (now users have the option to display both base and
            discounted
            versions).</p>
          <p>— Siege Golem Stone had an incorrect amount of resources in its recipe.</p>
          <p>— Added Contribute page.</p>
        </div>
      </div>
      <div class="log changelog__log">
        <h3 class="log__title">31/05/2022 Spells Update!</h3>
        <div class="log__body">
          <p>— Compete catalogue of all spells including weapon abilities</p>
          <p>— For each spell their cooldown, cast time and number of charges is specified (according to the ingame
            tooltip).</p>
          <p>— For vampire spells cast times were determined manually by recording a video of spell cast and counting
            frames, approximate values are shown on the website.</p>
        </div>
      </div>
      <div class="log changelog__log">
        <h3 class="log__title">24/05/2022</h3>
        <div class="log__body">
          <p>— added a variety (34) of hats and capes (big thanks to u/Difinitus)</p>
          <p>— added Blood Key</p>
          <p>— some minor fixes to tags of particular items</p>
        </div>
      </div>
      <div class="log changelog__log">
        <h3 class="log__title">21/05/2022</h3>
        <div class="log__body">
          <p>— Launched the website with the folowing features:</p>
          <p>— items catalouge excluding sturctures and the majority of hats and cloaks</p>
          <p>— search items by name</p>
          <p>— search items by tags</p>
          <p>— clickable tags to initiate and instant search by tag name</p>
          <p>— display all recipes that a chosen item can be crafted from</p>
          <p>— display all items that a chosen item is reagent for</p>
          <p>— clickable items in the "recipe" and "reagent for" sections to display their information</p>
        </div>
      </div>
    </div>
  </div>
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
@import "@/assets/styles/utility/vars.scss";

.container {
  display: flex;
  flex-direction: column;
  max-width: $lg;
  margin-left: auto;
  margin-right: auto;
  padding: 15px 20px;

  &__logo {
    width: 500px;
    margin-left: auto;
    margin-right: auto;

    @media (max-width: $sm) {
      width: 300px;

    }
  }

  &__description {
    text-align: center;
  }
}

.description {
  max-width: $md;
  margin-left: auto;
  margin-right: auto;

  &__block {
    font-size: 1.2rem;
  }
}

.block {
  margin-bottom: 15px;

  &__button {
    margin-bottom: 15px;
  }

  &__text {
    margin-bottom: 15px;
  }
}

.changelog {
  &__header {
    color: white;
    font-size: 2rem;
  }
}

.log {
  margin-top: 15px;

  &__title {
    color: white;
    font-size: 1.3rem;
  }

  &__body {
    margin-top: 15px;
    line-height: 1.25rem;
  }

  &__link {
    color: $primary;
  }
}

.button {
  color: white;
  font-size: 0.85rem;
  background-color: $primary;
  padding: 7px 20px;
  border-radius: 5px;
  margin-left: auto;
  margin-right: auto;

  @media (max-width: $sm) {
    font-size: 0.65rem;
  }
}
</style>