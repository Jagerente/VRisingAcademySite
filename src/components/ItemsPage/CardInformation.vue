<template>
  <my-card title="Information" :custom="false">
    <template #left><icon-bar :left="true" /></template>
    <template #right><icon-bar :left="false" /></template>

    <div v-if="selectedItem === null">
      <div></div>
      <h1 class="text-center">Select an item.</h1>
      <h5 class="text-center">This module is still W.I.P.</h5>
    </div>
    <div class="d-flex flex-column overflow p-2" v-else>
      <!-- <img
        type="image"
        class="btn-primary rounded"
        :title="selectedItem.name"
        :src="
          require('@/assets/images/items/' +
            getPath() +
            '/' +
            selectedItem.name +
            '.png')
        "
      /> -->
      <h1>
        {{ selectedItem.name }}
      </h1>
      <div class="description py-2 mb-3">
        <h4 class="">{{ selectedItem.description }}</h4>
      </div>
      <div
        v-if="selectedItem.mainStat"
        class="d-flex stats justify-content-between"
      >
        <div class="d-flex flex-column w-100">
          <h4>Main Stat:</h4>
          <h5>{{ selectedItem.mainStat }}</h5>
          <div v-if="selectedItem.bonusStats.length">
            <h4>Bonus Stats:</h4>
            <h5 v-for="(stat, i) in selectedItem.bonusStats" :key="i">
              {{ stat }}
            </h5>
          </div>
        </div>
        <div class="d-flex flex-column w-50">
          <h4>
            Durability: <span>{{ selectedItem.durability }}</span>
          </h4>
          <h4>
            Gear Level: <span>{{ selectedItem.gearLevel }}</span>
          </h4>
        </div>
      </div>
      <div class="d-flex tags">
        <input
          v-for="(tag, i) in selectedItem.tags"
          type="button"
          class="tag px-2"
          :value="tag"
          :key="i"
          @click="updateSearchQuery(tag)"
        />
      </div>
      <!-- <h1>{{ selectedItem }}</h1> -->
    </div>
  </my-card>
</template>

<script>
import { mapState, mapActions, mapMutations } from "vuex";

export default {
  methods: {
    ...mapActions({
      selectItem: "items/selectItem",
      updateSearchQuery: "items/updateSearchQuery",
    }),
  },
  computed: {
    ...mapState({
      selectedItem: (state) => state.items.selectedItem,
      searchQuery: (state) => state.items.searchQuery,
    }),
  },
  // methods: {
  //   getPath() {
  //     switch (this.type) {
  //       case 1:
  //         return "weapons";
  //       case 2:
  //         return "armour";
  //       case 3:
  //         return "consumables";
  //       case 4:
  //         return "reagents";
  //       default:
  //         console.error("Wrong Item type:", selectedItem.type, item);
  //         return "all";
  //     }
  //   },
  // },
};
</script>

<style scoped>
.description {
  background: #14131b;
  border-radius: 10px;
  margin-left: -10px;
  margin-right: -10px;
  padding-left: 15px;
  padding-right: 15px;
}
.tags {
  margin-top: 15px;
}
.tag {
  background: #ae1d1d;
  border-radius: 100px;
  text-transform: capitalize;
  margin-right: 5px;
  border: none;
  color: silver;
}
</style>