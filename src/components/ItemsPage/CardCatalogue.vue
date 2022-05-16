<template>
  <my-card :custom="true">
    <template #header><filter-input></filter-input></template>
    <div
      class="tab-content d-flex justify-content-center flex-wrap"
      id="v-pills-tabContent"
    >
      <div
        class="tab-pane fade show active"
        id="v-pills-weapons"
        role="tabpanel"
        aria-labelledby="v-pills-weapons-tab"
      >
        <div class="d-flex justify-content-center flex-wrap">
          <transition-group name="items-list">
            <my-item
              v-for="item in sortedAndSearchedItems"
              :key="item.name"
              :item="item"
              @click="selectItem(item.name), setSelected()"
            />
          </transition-group>
        </div>
      </div>
      <div
        class="tab-pane fade"
        id="v-pills-armour"
        role="tabpanel"
        aria-labelledby="v-pills-armour-tab"
      >
        ...
      </div>
      <div
        class="tab-pane fade"
        id="v-pills-consumables"
        role="tabpanel"
        aria-labelledby="v-pills-consumables-tab"
      >
        ...
      </div>
      <div
        class="tab-pane fade"
        id="v-pills-reagents"
        role="tabpanel"
        aria-labelledby="v-pills-reagents-tab"
      >
        ...
      </div>
    </div>
    <div style="height: 15px"></div>
  </my-card>
</template>

<script>
import { mapState, mapActions, mapGetters } from "vuex";
import FilterInput from "../FilterInput.vue";

export default {
  components: { FilterInput },
  methods: {
    ...mapActions({
      getItems: "items/getItems",
      selectItem: "items/selectItem",
      test: "items/test",
    }),
    setSelected() {
      console.log(this.$refs.itemEl);
    },
  },
  mounted() {
    this.getItems();
  },
  computed: {
    ...mapState({
      isItemsLoading: (state) => state.items.isItemsLoading,
    }),
    ...mapGetters({
      sortedAndSearchedItems: "items/sortedAndSearchedItems",
    }),
    // bgImage() {
    //   return require("@/assets/images/items" + this.backgroundImage);
    // },
    // inlineStyle() {
    //   return {
    //     backgroundImage: `url(${this.bgImage})`,
    //   };
    // },
  },
};
</script>

<style scoped>
.items-list-enter-active,
.items-list-leave-active {
  transition: all 0.45s ease;
}
.items-list-enter-from,
.items-list-leave-to {
  opacity: 0;
  transform: translateY(30px);
}
.items-list-move {
  transition: transform 0.35s ease;
}
/* .items-list-leave-active {
  position: absolute;
} */

.ps {
  height: 100%;
}
</style>