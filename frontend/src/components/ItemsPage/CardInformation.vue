<template>
  <my-card
    title="Information"
    :custom="false"
    class="information"
  >
    <div v-if="this.selectedItem === null">
      <h1 class="text-center">Select an item.</h1>
      <h5 class="text-center">This module is still W.I.P.</h5>
    </div>

    <div
      class="information__block-column"
      v-else
    >
      <!-- Title -->
      <div class="information__header">
        {{ this.selectedItem.name }}
      </div>
      <div class="block information__block">
        <div class="block__left">
          <!-- Tags -->
          <div class="information__tags">
            <button
              v-for="(tag, i) in this.selectedItem.tags"
              class="tag"
              :key="i"
              @click="updateSearchQuery({ query: tag, type: 2 })"
            >
              {{ tag }}
            </button>

          </div>
          <!-- Stats -->
          <div v-if="this.selectedItem.mainStat">
            <!-- Gear Level -->
            <div
              class="stat"
              v-if="this.selectedItem.gearLevel"
            >
              <span class="stat__value">{{ this.selectedItem.gearLevel }}</span>
              Gear Level
            </div>
            <!-- Main Stat -->
            <div class="stat">
              <span class="stat__value">+{{ this.selectedItem.mainStat }}</span>
              {{ getStat(this.selectedItem) }}
            </div>
            <!-- Sets -->
            <div
              v-if="this.selectedItem.set.description"
              class="stat"
            >
              <div class="set__name">
                {{ this.selectedItem.set.name }}
              </div>
              <Markdown
                class="stat"
                breaks
                html
                xhtmlOut
                :source="this.selectedItem.set.description.replace('\\n', '\n')"
              />
            </div>
            <!-- Bonus Stats -->
            <div v-if="this.selectedItem.bonusStats.length">
              <div
                class="stat"
                v-for="(stat, i) in this.selectedItem.bonusStats"
                :key="i"
              >
                <span class="stat__value">{{ stat.split(' ')[0] }}</span>
                {{ ' ' + stat.split(' ').slice(1).join(' ') }}
              </div>
            </div>
            <!-- Durability -->
            <div
              class="stat"
              v-if="this.selectedItem.durability"
            >
              <span class="stat__value">Durability: {{ this.selectedItem.durability }}</span>
            </div>
            <!-- Max Stack -->
          </div>
          <div
            class="stat"
            v-if="this.selectedItem.maxStack"
          >
            <span class="stat__value">Max Stack: {{ this.selectedItem.maxStack }}</span>
          </div>
        </div>
        <!-- Preview -->
        <div class="block__right">
          <ItemPreview
            class="item__image-preview"
            :item="this.selectedItem"
            :button="false"
          />
          <!-- Show on map -->
          <button
            v-if="this.selectedItem.locations.length"
            class="button"
            @click="this.openUrl(`https://mapgenie.io/v-rising/maps/vardoran?locationIds=${this.mapGenieLocations}`)"
          >Show on map</button>
        </div>
      </div>
      <!-- Description -->
      <Markdown
        :source="this.selectedItem.description"
        class="block__card"
        html
        xhtmlOut
        v-if="this.selectedItem.description"
      />
      <!-- Variants -->
      <div
        class="information__block-column"
        v-if="this.selectedItem.variants.length"
      >
        <div class="variants__title">Variants</div>
        <div class="item variants__list block__card">
          <ItemPreview
            v-for="itemId in selectedItem.variants"
            class="item-link"
            :item="this.items.find(item => { return item.id === itemId })"
            :button="true"
          />
        </div>
      </div>
      <!-- Recipes -->
      <div class="recipes">
        <div
          v-if="this.selectedItem.recipes.length"
          class="recipes__ingridients"
        >
          <div class="ingridients">
            <div class="ingridients__title">Recipes</div>
            <div class="ingridients__options">
              <MySwitch
                label="Matching Floor"
                :update="this.updateMatchingFloor"
                :checked="this.matchingFloor"
              />
              <MySwitch
                label="Confined Room"
                :update="this.updateConfinedRoom"
                :checked="this.confinedRoom"
              />
            </div>
          </div>

          <MyRecipe
            class="recipes__recipe block__card"
            v-for="(recipeId, i) in this.selectedItem.recipes"
            :recipe="this.recipes.find(recipe => { return recipe.id == recipeId })"
          />
        </div>
        <!-- Reagent For -->
        <div
          v-if="this.selectedItem.reagentFor.length"
          class="reagents"
        >
          <div class="reagents__title">Reagent for</div>
          <div class="reagents__list block__card">
            <div
              class="item"
              v-for="recipeId in selectedItem.reagentFor"
            >
              <ItemPreview
                class="item__image-link"
                v-for="output in this.recipes.find(recipe => { return recipe.id == recipeId }).results"
                :item="this.items.find(item => { return item.id == output.itemId })"
                :button="true"
              />
            </div>
          </div>
        </div>
        <!-- Salvageables -->
        <div class="salvageables">
          <div v-if="this.selectedItem.salvageables.length">
            <div class="salvageables__title">Salvageable For</div>
            <div class="salvageables__list block__card ">
              <div
                class="item"
                v-for="output in this.salvageables.find(salvageable => { return salvageable.id === this.selectedItem.salvageables[0] }).results"
              >
                <ItemPreview
                  class="item__image-link"
                  :item="this.items.find(item => { return item.id === output.itemId })"
                  :button="true"
                />
                <div class="item__text">
                  {{ output.amount }}
                </div>
              </div>
            </div>
          </div>
          <div v-if="this.selectedItem.salvageableOf.length">
            <p class="salvageables__title">Salvageable From</p>
            <div class="salvageables__list block__card item">
              <ItemPreview
                class="item__image-link"
                v-for="input in this.selectedItem.salvageableOf	"
                :item="this.items.find(item => { return item.id === this.salvageables.find(salvageable => { return salvageable.id === input }).itemId })"
                :button="true"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </my-card>
</template>

<script>
import { mapState, mapActions, mapGetters } from "vuex";
import Markdown from "vue3-markdown-it";
import ItemPreview from "@/components/ItemsPage/ItemPreview.vue";
import MyRecipe from "@/components/ItemsPage/MyRecipe.vue";

export default {
  components: {
    Markdown,
    ItemPreview,
    MyRecipe,
  },

  methods: {
    ...mapActions({
      updateSearchQuery: "items/updateSearchQuery",
      updateMatchingFloor: "items/updateMatchingFloor",
      updateConfinedRoom: "items/updateConfinedRoom",
    }),
    getStat(item) {
      switch (item.type.id) {
        case 1:
          return "Physical Power";
        case 2:
        case 5:
          return "Max Health"
        case 3:
          return "Spell Power"
        default:
          console.error("Wrong Item type:", item.typeid, item);
          return '';
      }
    },
    openUrl(url) {
      window.open(url, '_blank')
    }
  },
  computed: {
    ...mapState({
      selectedItem: (state) => state.items.selectedItem,
      searchQuery: (state) => state.items.searchQuery,
      items: (state) => state.items.items,
      locations: (state) => state.items.locations,
      recipes: (state) => state.items.recipes,
      salvageables: (state) => state.items.salvageables,
      matchingFloor: (state) => state.items.matchingFloor,
      confinedRoom: (state) => state.items.confinedRoom,
    }),
    ...mapGetters({
      mapGenieLocations: "items/mapGenieLocations",
    })
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/styles/utility/vars.scss";

.information {
  width: 100%;


  &__block {
    display: flex;
    justify-content: space-between;

    &:not(:last-child) {
      margin-bottom: $m1;
    }

    &-column {
      flex-direction: column;

      &:not(:last-child) {
        margin-bottom: $m1;
      }
    }

    .block {
      &__right {
        display: flex;
        flex-direction: column;
      }
    }
  }

  &__header {
    color: white;
    font-size: 32px;
    margin-bottom: $m1;
  }

  &__tags {
    margin-bottom: $m1;

    &>.tag {
      margin-top: 5px;
      font-size: 0.8rem;
      background: $primary;
      border-radius: 15px;
      text-transform: capitalize;
      border: none;
      color: white;
      transition: box-shadow 0.1s ease-in-out;
      padding-left: $m1;
      padding-right: $m1;
      padding-top: 3px;
      padding-bottom: 3px;

      &:not(:last-child) {
        margin-right: 5px;
      }

      &:hover {
        box-shadow: 0 0 3px black;
      }
    }
  }

  .block__card {
    background: $dark;
    border-radius: 10px;
    margin-bottom: $m1;
    padding-left: 15px;
    padding-right: 15px;
    padding-top: 15px;
    padding-bottom: 15px;
  }

  .stat {
    margin-bottom: 5px;
    font-size: 1rem;

    &__value {
      color: white;
    }

    .set {
      &__name {
        color: white;
      }
    }
  }
}

.variants {
  &__title {
    color: white;
    font-size: 1.5rem;
    margin-bottom: $m1;
    user-select: none;
  }

  &__list {
    display: flex;
    flex-wrap: wrap;
  }
}

.recipes {
  display: flex;
  flex-direction: column;

  &__ingridients {
    .ingridients {

      display: flex;
      justify-content: space-between;

      &__title {
        color: white;
        font-size: 1.5rem;
        margin-bottom: $m1;
        user-select: none;
      }

      &__options {
        margin-bottom: $m1;
        display: flex;
        flex-direction: column;
      }
    }

  }

  .reagents {
    display: flex;
    flex-direction: column;

    &__title {
      color: white;
      font-size: 1.5rem;
      margin-bottom: $m1;
      user-select: none;
    }

    &__list {
      display: flex;
      flex-wrap: wrap;
    }
  }



  .salvageables {
    display: flex;
    flex-direction: column;

    &__title {
      color: white;
      font-size: 1.5rem;
      margin-bottom: $m1;
      user-select: none;
    }

    &__list {
      display: flex;
      flex-wrap: wrap;
    }
  }
}

.item {
  position: relative;

  &__image {
    &-preview {
      background: none;
      $item-size: 10rem;
      width: $item-size;
      height: $item-size;
      border: none;
      pointer-events: none;

      &:hover {
        box-shadow: none;
      }

      @media (min-width: $lg) {
        $item-size: 13rem;
        width: $item-size;
        height: $item-size;
      }
    }

    &-link {
      position: relative;
      $item-size: 2.5rem;
      width: $item-size;
      height: $item-size;
      background: black;
      border: 1px solid black;
    }
  }

  &__text {
    position: absolute;
    color: white;
    font-family: sans-serif;
    pointer-events: none;
    bottom: 0px;
    right: 0px;
    font-size: 16px;
    user-select: none;
  }
}

.title {
  color: white;
  font-size: 1.5rem;
  margin-bottom: $m1;
}

.button {
  background-color: $primary;
  border: 1px solid $primary;
  border-radius: 5px;
  color: white;
  padding-left: 15px;
  padding-right: 15px;
  padding-top: 5px;
  padding-bottom: 5px;
  width: 100%;
  transition: box-shadow 0.1s ease-in-out;

  &:hover {
    box-shadow: 0 0 3px black;
  }
}
</style>
