<template>
  <div class="d-flex flex-wrap">
    <div class="d-flex flex-column w-100 ">
      <transition-group name="items-list">
        <div v-if="!searchQuery" class="d-flex flex-column" v-for="(set, i) in Object.keys(this.items)" :key="i">
          <h1>{{ set }}</h1>
          <div class="d-flex flex-wrap">
            <my-item v-for="item in this.items[set]" :key="item.id" :item="item" @click="selectItem(item)" />
          </div>
        </div>
        <div v-else>
          <transition-group name="items-list">

            <my-item v-for="item in this.sortedAndSearchedItems" :key="item.id" :item="item"
              @click="selectItem(item)" />
          </transition-group>

        </div>
      </transition-group>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions, mapGetters } from "vuex";

export default {
  name: "items-list",
  props: {
    items: Object
  },
  methods: {
    ...mapActions({
      selectItem: "items/selectItem",
    }),
  },
  computed: {
    ...mapState({
      searchQuery: (state) => state.items.searchQuery,
    }),
    ...mapGetters({
      sortedAndSearchedItems: "items/sortedAndSearchedItems",
    }),
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

.items-list-leave-active {
  position: absolute;
}

.input-group {
  height: 39px;
}
</style>