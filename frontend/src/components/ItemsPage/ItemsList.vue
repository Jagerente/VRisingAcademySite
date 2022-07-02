<template>
  <div
    class="items"
    v-if="!searchQuery"
    v-for="set in this.sets"
    :key="set.id"
  >
    <p class="items__set">{{ set.name }}</p>
    <div class="items__list">
      <MyItem
        v-for="item in set.items"
        :key="item.id"
        :item="item"
      />
    </div>
  </div>
  <div
    class="items__list"
    v-else
  >
    <MyItem
      v-for="item in this.sortedAndSearchedItems"
      :key="item.id"
      :item="item"
    />
  </div>
</template>

<script>
import { mapState, mapActions, mapGetters } from "vuex";
import MyItem from "@/components/ItemsPage/MyItem.vue"

export default {
  components: {
    MyItem,
  },
  name: "items-list",
  props: {
    sets: Object
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

<style lang="scss" scoped>
.items {
  &__list {
    display: flex;
    flex-wrap: wrap;
  }

  &__set {
    font-size: 26px;
    user-select: none;
  }
}
</style>