<template>
  <my-card title="Information" :custom="false">
    <!-- <template #left>
      <icon-bar :left="true" />
    </template>
    <template #right>
      <icon-bar :left="false" />
    </template> -->

    <div v-if="selectedItem === null">
      <div></div>
      <h1 class="text-center">Select an item.</h1>
      <h5 class="text-center">This module is still W.I.P.</h5>
    </div>

    <div class="d-flex flex-column overflow p-2" v-else>
      <div class="d-flex justify-content-between">
        <div class="d-flex flex-column">
          <!-- Title -->
          <p class="h-1">
            {{ selectedItem.name }}
          </p>
          <!-- Tags -->
          <div class="d-flex flex-wrap tags mb-2">
            <input v-for="(tag, i) in selectedItem.tags" type="image" class="tag px-2" :value="tag" :key="i"
              @click="updateSearchQuery({ query: tag, type: 2 })" />
          </div>
          <div v-if="selectedItem.mainStat" class="d-flex stats justify-content-between">
            <div class="d-flex flex-column w-100">
              <!-- Gear Level -->
              <p class="h-2" v-if="selectedItem.gearLevel">
                <span class="text-white">{{ selectedItem.gearLevel }}</span> Gear Level
              </p>
              <!-- Main Stat -->
              <p class="h-3"><span class="text-white">+{{ selectedItem.mainStat }}</span> {{ getStat(selectedItem) }}
              </p>
              <!-- Sets -->
              <div v-if="selectedItem.set.description" class="my-2">
                <p class="h-3 text-white">
                  {{ selectedItem.set.name }}
                </p>
                <p class="h-3" v-for="str in selectedItem.set.description.split('\\n')">
                  {{ str }}
                </p>
              </div>
              <!-- Bonus Stats -->
              <div v-if="selectedItem.bonusStats.length">
                <p class="h-3" v-for="(stat, i) in selectedItem.bonusStats" :key="i">
                  <span class="text-white">{{ stat.split(' ')[0] }}</span>
                  <span>
                    {{ ' ' + stat.split(' ').slice(1).join(' ') }}
                  </span>
                </p>
              </div>
              <!-- Durability -->
              <p class="h-3" v-if="selectedItem.durability">
                <span class="text-white">Durability: {{ selectedItem.durability }}</span>
              </p>
            </div>
          </div>
        </div>
        <!-- Preview -->
        <div class="d-flex justify-content-center flex-fill mx-2">
          <item-preview class="preview-lg" :item="selectedItem" :button="false" />
        </div>
      </div>
      <!-- Description -->
      <div v-if="selectedItem.description" class="block mt-2 py-2">
        <p class="h-3">{{ selectedItem.description }}</p>
      </div>
      <!-- Variants -->
      <div class="d-flex flex-column" v-if="selectedItem.variants.length">
        <p class="h-2">Variants</p>
        <div class="d-flex flex-wrap block py-2">
          <item-preview v-for="itemId in selectedItem.variants" class="preview-sm m-1"
            :item="this.items.find(item => { return item.id == itemId })" :button="true" />
        </div>
      </div>
      <div class="d-flex flex-column">
        <!-- Recipes -->
        <div v-if="selectedItem.recipes.length" class="d-flex flex-column">
          <p class="h-2 mt-1">Recipes</p>
          <div class="d-flex block" :class="i < selectedItem.recipes.length - 1 ? 'mb-2' : ''"
            v-for="(recipeId, i) in selectedItem.recipes">
            <div class="d-flex">
              <my-recipe v-for="input in this.recipes.find(recipe => { return recipe.id == recipeId }).ingredients"
                :item="input" />
            </div>
            <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" fill="#a8a9ae" class="m-2"
              viewBox="0 0 16 16">
              <path
                d="m12.14 8.753-5.482 4.796c-.646.566-1.658.106-1.658-.753V3.204a1 1 0 0 1 1.659-.753l5.48 4.796a1 1 0 0 1 0 1.506z" />
            </svg>
            <div class="d-flex flex-fill">
              <my-recipe v-for="output in this.recipes.find(recipe => { return recipe.id == recipeId }).results"
                :item="output" />
            </div>
          </div>
        </div>
        <!-- Reagent For -->
        <div v-if="selectedItem.reagentFor.length" class="d-flex flex-column">
          <p class="h-2">Reagent for</p>
          <div class="d-flex block flex-wrap">
            <div v-for="recipeId in selectedItem.reagentFor">
              <item-preview class="preview-sm m-1"
                v-for="output in this.recipes.find(recipe => { return recipe.id == recipeId }).results"
                :item="this.items.find(item => { return item.id == output.itemId })" :button="true" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </my-card>
</template>

<script>
import { mapState, mapActions } from "vuex";

export default {
  methods: {
    ...mapActions({
      updateSearchQuery: "items/updateSearchQuery",
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
  },
  computed: {
    ...mapState({
      selectedItem: (state) => state.items.selectedItem,
      searchQuery: (state) => state.items.searchQuery,
      items: (state) => state.items.items,
      recipes: (state) => state.items.recipes
    }),
  },

};
</script>

<style scoped lang="scss">
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';

@include media-breakpoint-down(sm) {
  .tag {
    margin-top: 5px;
    font-size: 0.7rem;
  }
}

p {
  margin: 0;
  padding: 0;
}

.h-1 {
  font-size: 2rem;
}

.h-2 {
  font-size: 1.5rem;
}

.h-3 {
  font-size: 1rem;
}

.tag {
  margin-top: 5px;
  font-size: 0.8rem;
}

@include media-breakpoint-up(sm) {
  .tag {
    margin-top: 5px;
  }
}

.tag {
  background: #ae1d1d;
  border-radius: 100px;
  text-transform: capitalize;
  margin-right: 5px;
  border: none;
  color: silver;
}

.tag:hover {
  box-shadow: 0 0 5px black;
  transition: box-shadow 0.05s ease-in-out;
}

.block {
  background: #14131b;
  border-radius: 10px;
  margin-left: -10px;
  margin-right: -10px;
  padding-left: 15px;
  padding-right: 15px;
}
</style>