<template>
  <my-card title="Information" :custom="false">
    <template #left>
      <icon-bar :left="true" />
    </template>
    <template #right>
      <icon-bar :left="false" />
    </template>

    <div v-if="selectedItem === null">
      <div></div>
      <h1 class="text-center">Select an item.</h1>
      <h5 class="text-center">This module is still W.I.P.</h5>
    </div>
    <div class="d-flex flex-column overflow p-2" v-else>
      <div class="d-flex justify-content-between">
        <div class="d-flex flex-column">
          <h1>
            {{ selectedItem.name }}
          </h1>
          <div class="d-flex flex-wrap tags mb-2">
            <input v-for="(tag, i) in selectedItem.tags" type="button" class="tag px-2" :value="tag" :key="i"
              @click="updateSearchQuery(tag)" />
          </div>
          <div v-if="selectedItem.mainStat" class="d-flex stats justify-content-between">
            <div class="d-flex flex-column w-100">
              <h3 style="" v-if="selectedItem.gearLevel">
                <span class="text-white">{{ selectedItem.gearLevel }}</span> Gear Level
              </h3>
              <h3><span class="text-white">+{{ selectedItem.mainStat }}</span> {{ getStat(selectedItem) }}</h3>
              <div v-if="selectedItem.set" class="mb-2">
                <h4 class="text-white my-0">
                  {{ selectedItem.set.name }}
                </h4>
                <h6 class="my-0" v-for="str in selectedItem.set.description.split('\\n')">
                  {{ str }}
                </h6>
              </div>
              <div v-if="selectedItem.bonusStats.length">
                <h5 v-for="(stat, i) in selectedItem.bonusStats" :key="i">
                  <span class="text-white">{{ stat.split(' ')[0] }}</span>
                  <span>
                    {{ ' ' + stat.split(' ').slice(1).join(' ') }}
                  </span>
                </h5>
              </div>
              <h5>
                <span class="text-white">Durability: {{ selectedItem.durability }}</span>
              </h5>
            </div>
          </div>
        </div>
        <div class="d-flex justify-content-center flex-fill mx-2">
          <img type="image" class="item__preview rounded" draggable="false" :title="selectedItem.name" :src="
            require('@/assets/images/items/' +
              getPath(selectedItem.type) +
              '/' +
              selectedItem.name +
              '.png')
          " />
        </div>
      </div>
      <div class="description py-2">
        <h4 class="">{{ selectedItem.description }}</h4>
      </div>
      <div class="d-flex flex-column">
        <div v-if="selectedItem.recipes.length" class="d-flex flex-column">
          <h2>Recipe</h2>
          <div class="d-flex recipes" v-for="recipes in selectedItem.recipesInfo">
            <div class="d-flex">
              <my-recipe v-for="input in recipes.ingredients" :item="input" />
            </div>
            <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" fill="#a8a9ae" class="m-2"
              viewBox="0 0 16 16">
              <path
                d="m12.14 8.753-5.482 4.796c-.646.566-1.658.106-1.658-.753V3.204a1 1 0 0 1 1.659-.753l5.48 4.796a1 1 0 0 1 0 1.506z" />
            </svg>
            <div class="d-flex flex-fill">
              <my-recipe v-for="input in recipes.results" :item="input" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </my-card>
</template>

<script>
import { mapState, mapActions, mapMutations } from "vuex";

export default {
  methods: {
    ...mapActions({
      updateSearchQuery: "items/updateSearchQuery",
    }),
    getStat(item) {
      switch (item.type) {
        case 1:
          return "Physical Power";
        case 2:
          switch (item.slotId) {
            case 7:
              return "Spell Power"
            default:
              return "Max Health"
          }
        default:
          console.error("Wrong Item type:", selectedItem.type, item);
          return '';
      }
    },
    getPath(type) {
      switch (type) {
        case 1:
          return "weapons";
        case 2:
          return "armour";
        case 3:
          return "consumables";
        case 4:
          return "reagents";
        default:
          console.error("[CardInformation] Wrong Item type:", type);
          return "all";
      }
    },
  },
  computed: {
    ...mapState({
      selectedItem: (state) => state.items.selectedItem,
      searchQuery: (state) => state.items.searchQuery,
      items: (state) => state.items.items
    }),
  },

};
</script>

<style scoped>
.tag {
  background: #ae1d1d;
  border-radius: 100px;
  text-transform: capitalize;
  margin-right: 5px;
  margin-top: 5px;
  border: none;
  color: silver;
}

.item__preview {
  --img-size: 200px;
  width: var(--img-size);
  height: var(--img-size);
  user-select: none;
}

.description {
  background: #14131b;
  border-radius: 10px;
  margin-left: -10px;
  margin-right: -10px;
  padding-left: 15px;
  padding-right: 15px;
}

.recipes {
  background: #14131b;
  border-radius: 15px;
  margin-left: -10px;
  margin-right: -10px;
  padding-left: 15px;
  padding-right: 15px;
}
</style>