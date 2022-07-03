<template>
  <div class="content">
    <div class="description content__description">
      <h1 class="description__header">
        Contribute
      </h1>
      <div class="description__content">
        <p>
          90% of the ingame data is already done by ourselves but it will take some time before we implement it and
          display on the website.
        </p>
        <p>
          Though, there is some data that we need help from the community to gather.
        </p>
      </div>

    </div>
    <div class="table content__table">
      <!-- Books -->
      <div class="table__column">
        <button
          class="button table__button"
          @click="$router.push('/contribute/books')"
        >Books Drop
          Statistics
        </button>
        <p class="table__content">
          We'd like to list where some books are most likely to drop from, e.g. Merciless Nightstalker Armour
          recipes
          usually drop from Grayson the Armourer. We need more verified data on all of the books before we can
          display
          it in the database. If you found a book, please take a screenshot of where exactly you found it and fill
          up
          the form.
        </p>
      </div>

      <!-- Hats/Cloaks -->
      <div class="table__column">
        <button
          class="button table__button"
          @click="$router.push('/contribute/cosmetics')"
        >Hats/Cloaks Drop
          Statistics
        </button>
        <p class="table__content">
          Same goes for Hats and Cloaks. We need verifiable data on how to obtain these precious cosmetic items.
        </p>
      </div>

      <!-- Servant Blood Quality -->
      <div class="table__column">
        <button
          class="button table__button"
          @click="$router.push('/contribute/servants/bloodquality')"
        >
          Servant Blood Quality
        </button>
        <p class="table__content">
          Need to test all blood qualities of humans (all 100 from 1% to 100%) and figure out how it affects the
          range of possible expertise levels. </p>
      </div>
    </div>
    <div class="footer content__footer">
      <div class="footer__content">
        We are also looking for players that can help us gather info on mosters, their behaviour, loot drops and their
        locations. If you are interested, please reach out to <a
          class="footer__link"
          href="#"
          @click="$router.push('/credits/bobrokrot/discord')"
        >Bobrokrot#0111</a> on Discord.
      </div>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@import "@/assets/styles/utility/vars.scss";

.content {
  max-width: $md;
  margin-left: auto;
  margin-right: auto;
  padding: 0 25px;

  &__description {
    margin-top: 25px;
  }

  &__table {
    margin-top: 50px;
  }

  &__footer {
    margin-top: 15px;
  }
}

.description {
  &__header {
    text-align: center;
    font-size: 2.5rem;
    color: white;
    font-weight: bold;
  }

  &__content {
    margin-top: 15px;
    text-align: center;
    font-size: 1.25rem;

    @media (max-width: $sm) {
      font-size: 1rem;
    }
  }
}

.table {
  display: flex;

  &__column {
    display: flex;
    flex-direction: column;
    width: 100%;
    padding: 0 10px;
    text-align: center;
  }

  &__button {
    margin-bottom: 5px;
  }

  &__content {
    @media (max-width: $sm) {
      font-size: 0.8rem;
    }
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

.footer {
  &__content {
    text-align: center;
    font-size: 1.2rem;

    @media (max-width: $sm) {
      font-size: 1rem;
    }
  }

  &__link {
    color: $primary;
  }
}
</style>