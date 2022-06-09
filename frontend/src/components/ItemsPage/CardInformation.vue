<template>
  <my-card title="Information" :custom="false">
    <!-- <template #left>
      <icon-bar :left="true" />
    </template>
    <template #right>
      <icon-bar :left="false" />
    </template> -->

    <div v-if="this.selectedItem === null">
      <div></div>
      <h1 class="text-center">Select an item.</h1>
      <h5 class="text-center">This module is still W.I.P.</h5>
    </div>

    <div class="d-flex flex-column overflow p-2" v-else>
      <div class="d-flex justify-content-between">
        <div class="d-flex flex-column">
          <!-- Title -->
          <p class="h-1">
            {{ this.selectedItem.name }}
          </p>
          <!-- Tags -->
          <div class="d-flex flex-wrap tags mb-2">
            <input v-for="(tag, i) in this.selectedItem.tags" type="button" class="tag px-2" :value="tag" :key="i"
              @click="updateSearchQuery({ query: tag, type: 2 })" />
          </div>
          <div v-if="this.selectedItem.mainStat" class="d-flex stats justify-content-between">
            <div class="d-flex flex-column w-100">
              <!-- Gear Level -->
              <p class="h-2" v-if="this.selectedItem.gearLevel">
                <span class="text-white">{{ this.selectedItem.gearLevel }}</span> Gear Level
              </p>
              <!-- Main Stat -->
              <p class="h-3"><span class="text-white">+{{ this.selectedItem.mainStat }}</span> {{
                  getStat(this.selectedItem)
              }}
              </p>
              <!-- Sets -->
              <div v-if="this.selectedItem.set.description" class="my-2">
                <p class="h-3 text-white">
                  {{ this.selectedItem.set.name }}
                </p>
                <p class="h-3" v-for="str in this.selectedItem.set.description.split('\\n')">
                  {{ str }}
                </p>
              </div>
              <!-- Bonus Stats -->
              <div v-if="this.selectedItem.bonusStats.length">
                <p class="h-3" v-for="(stat, i) in this.selectedItem.bonusStats" :key="i">
                  <span class="text-white">{{ stat.split(' ')[0] }}</span>
                  <span>
                    {{ ' ' + stat.split(' ').slice(1).join(' ') }}
                  </span>
                </p>
              </div>
              <!-- Durability -->
              <p class="h-3" v-if="this.selectedItem.durability">
                <span class="text-white">Durability: {{ this.selectedItem.durability }}</span>
              </p>
            </div>
          </div>
        </div>
        <!-- Preview -->
        <div class="d-flex justify-content-center flex-fill mx-2">
          <item-preview :style="'preview-lg'" :item="this.selectedItem" :button="false" />
        </div>
      </div>
      <!-- Description -->
      <div v-if="this.selectedItem.description" class="block mt-2 py-2">
        <p class="h-3">{{ this.selectedItem.description }}</p>
      </div>
      <!-- Variants -->
      <div class="d-flex flex-column" v-if="this.selectedItem.variants.length">
        <p class="h-2">Variants</p>
        <div class="d-flex flex-wrap block py-2">
          <item-preview v-for="itemId in selectedItem.variants" class="m-1" :style="'preview-sm'"
            :item="this.items.find(item => { return item.id === itemId })" :button="true" />
        </div>
      </div>
      <!-- Recipes -->
      <div class="d-flex flex-column">
        <div v-if="this.selectedItem.recipes.length" class="d-flex flex-column">
          <div class="d-flex justify-content-between mt-1">
            <p class="h-2 mb-auto">Recipes</p>
            <div class="d-flex flex-column mt-auto">
              <div class="form-check form-switch">
                <input @click="this.updateMatchingFloor" class="form-check-input" :checked="this.matchingFloor"
                  type="checkbox" id="flexSwitchMatchingFloor">
                <label class="form-check-label" for="flexSwitchMatchingFloor" style="user-select: none;">Matching
                  Floor</label>
              </div>
              <div class="form-check form-switch">
                <input @click="this.updateConfinedRoom" class="form-check-input" :checked="this.confinedRoom"
                  type="checkbox" id="flexSwitchConfinedRoom">
                <label class="form-check-label" for="flexSwitchConfinedRoom" style="user-select: none;">Confined
                  Room</label>
              </div>
            </div>
          </div>

          <my-recipe :class="i < this.selectedItem.recipes.length - 1 ? 'mb-2' : ''"
            v-for="(recipeId, i) in this.selectedItem.recipes"
            :recipe="this.recipes.find(recipe => { return recipe.id == recipeId })" />
        </div>
        <!-- Reagent For -->
        <div v-if="this.selectedItem.reagentFor.length" class="d-flex flex-column">
          <p class="h-2">Reagent for</p>
          <div class="d-flex block flex-wrap">
            <div class="d-flex" v-for="recipeId in selectedItem.reagentFor">
              <item-preview class="m-1" :style="'preview-sm'"
                v-for="output in this.recipes.find(recipe => { return recipe.id == recipeId }).results"
                :item="this.items.find(item => { return item.id == output.itemId })" :button="true" />
            </div>
          </div>
        </div>
      </div>
      <!-- Salvageables -->
      <div class="d-flex flex-column">
        <div v-if="this.selectedItem.salvageables.length">
          <p class="h-2 mb-auto">Salvageable For</p>
          <div class="d-flex block ">
            <item-preview class="m-1" :style="'preview-sm'"
              v-for="output in this.salvageables.find(salvageable => { return salvageable.id === this.selectedItem.salvageables[0] }).results"
              :item="this.items.find(item => { return item.id === output.itemId })" :text="output.amount"
              :button="true" />
          </div>
        </div>
        <div v-if="this.selectedItem.salvageableOf.length">
          <p class="h-2 mb-auto">Salvageable From</p>
          <div class="d-flex block">
            <item-preview class="m-1" :style="'preview-sm'" v-for="input in this.selectedItem.salvageableOf	"
              :item="this.items.find(item => { return item.id === this.salvageables.find(salvageable => { return salvageable.id === input }).itemId })"
              :button="true" />
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
  },
  computed: {
    ...mapState({
      selectedItem: (state) => state.items.selectedItem,
      searchQuery: (state) => state.items.searchQuery,
      items: (state) => state.items.items,
      recipes: (state) => state.items.recipes,
      salvageables: (state) => state.items.salvageables,
      matchingFloor: (state) => state.items.matchingFloor,
      confinedRoom: (state) => state.items.confinedRoom,
    }),
  },

};
</script>

<style scoped lang="scss">
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';

@import '@/assets/styles/va_styles.scss';


.tag {
  margin-top: 5px;
  font-size: 0.8rem;
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
  padding-top: 5px;
  padding-bottom: 5px;
}

.form-check-input {
  background-color: #14131b;
  border: 1px solid rgb(77, 77, 77);
}

.form-check-input:focus {
  box-shadow: none;
}
</style>