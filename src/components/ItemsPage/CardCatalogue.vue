<template>
  <my-card title="Catalogue">
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

export default {
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
.items {
  background: rgba(0, 0, 0, 0.5);
  /* background-size: 100%; */
  /* background-position: 50%; */

  /* --item-size: 85px; */
  --item-size: 85px;
  font-size: 10px;
  font-family: sans-serif;
  width: var(--item-size);
  height: var(--item-size);
  color: black;
  box-shadow: 0 0 8px black;
  margin-bottom: 5px;
  -webkit-user-drag: none;
  border: 0;
  margin: 10px;
  user-select: none;
  transition: box-shadow 0.15s ease-in-out;
}

.items:hover {
  box-shadow: 0 0 13px var(--primary);
}

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
.btn-check:focus + .btn-primary,
.btn-primary:focus {
  /* background: rgba(0, 0, 0, 0.5); */
  background-color: black;
  background-size: 90%;
  background-position: 50%;
  background-repeat: no-repeat;
  border: 1px solid white;
  box-shadow: 0;
}
.btn-primary:hover {
  background-color: black;
  background-repeat: no-repeat;
  background-size: 90%;
}

.btn-check:focus + .btn-primary,
.btn-primary:focus {
  box-shadow: none;
}
.ps {
  height: 100%;
}
</style>