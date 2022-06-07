<template>
  <div class="container">
    <div class="d-flex flex-column">
      <h1 class="mt-3 text-center">
        Contribute
      </h1>
      <div class="mx-auto description" style="width: 600px">
        <p>
          90% of the ingame data is already done by ourselves but it will take some time before we implement it and
          display
          on the website.
        </p>
        <p>
          Though, there is some data that we need help from the community to gather.
        </p>
      </div>

      <div class="row mt-5">
        <!-- Books -->
        <div class="col">
          <div class="d-flex flex-column">
            <button class="btn btn-primary mx-auto" @click="$router.push('/contribute/books')">Books Drop
              Statistics
            </button>
            <p class="description">
              We'd like to list where some books are most likely to drop from, e.g. Merciless Nightstalker Armour
              recipes
              usually drop from Grayson the Armourer. We need more verified data on all of the books before we can
              display
              it in the database. If you found a book, please take a screenshot of where exactly you found it and fill
              up
              the form.
            </p>
          </div>
        </div>

        <!-- Hats/Cloaks -->
        <div class="col">
          <div class="d-flex flex-column">
            <button class="btn btn-primary mx-auto" @click="$router.push('/contribute/cosmetics')">Hats/Cloaks Drop
              Statistics
            </button>
            <p class="description">
              Same goes for Hats and Cloaks. We need verifiable data on how to obtain these precious cosmetic items.
            </p>
          </div>
        </div>

        <!-- Servant Blood Quality -->
        <div class="col">
          <div class="d-flex flex-column">
            <button class="btn btn-primary mx-auto" @click="$router.push('/contribute/servants/bloodquality')">
              Servant Blood Quality
            </button>
            <p class="description">
              Need to test all blood qualities of humans (all 100 from 1% to 100%) and figure out how it affects the
              range of possible expertise levels. </p>
          </div>
        </div>
      </div>
      <div class="d-flex flex-column">
        <div class="description mx-auto my-5" style="width: 800px;">
          We are also looking for players that can help us gather info on mosters, their behaviour, loot drops and their
          locations. If you are interested, please reach out to <a href="#"
            @click="$router.push('/credits/bobrokrot/discord')">Bobrokrot#0111</a> on Discord.
        </div>
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
    goto(url) {
      window.open(url, '_blank')
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

.content {
  padding-left: auto;
  padding-right: auto;
  margin-left: auto;
  margin-right: auto;
}

.description {
  text-align: center;
  margin-top: 10px;
}
</style>