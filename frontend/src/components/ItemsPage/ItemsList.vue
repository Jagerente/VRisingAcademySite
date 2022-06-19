<template>
  <div class="d-flex flex-wrap">
    <div class="d-flex flex-column w-100 ">
      <!-- <transition-group name="items-list"> -->
      <div
        v-if="!searchQuery"
        class="d-flex flex-column"
        v-for="set in this.sets"
        :key="set.id"
      >
        <h1>{{ set.name }}</h1>
        <div class="d-flex flex-wrap">
          <MyItem
            v-for="item in set.items"
            :key="item.id"
            :item="item"
          />
        </div>
      </div>
      <div
        class="d-flex flex-wrap"
        v-else
      >
        <!-- <transition-group name="items-list"> -->
        <MyItem
          v-for="item in this.sortedAndSearchedItems"
          :key="item.id"
          :item="item"
        />
        <!-- </transition-group> -->
      </div>
      <!-- </transition-group> -->
    </div>
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

<!--<style scoped>-->
<!--/* .items-list-enter-active,-->
<!--.items-list-leave-active {-->
<!--  transition: all 0.45s ease;-->
<!--}-->

<!--.items-list-enter-from,-->
<!--.items-list-leave-to {-->
<!--  opacity: 0;-->
<!--  transform: translateY(30px);-->
<!--}-->

<!--.items-list-move {-->
<!--  transition: transform 0.35s ease;-->
<!--}-->

<!--.items-list-leave-active {-->
<!--  position: absolute;-->
<!--} */-->

<!--.input-group {-->
<!--  height: 39px;-->
<!--}-->
<!--</style>-->