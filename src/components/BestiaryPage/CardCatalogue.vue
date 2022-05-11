<template>
  <div class="d-flex flex-column rounded">
    <h3 class="card-content">Catalogue</h3>
    <perfect-scrollbar class="test">
      <div
        class="d-flex justify-content-center flex-wrap"
        style="max-height: 350px"
      >
        <transition-group name="items-list">
          <button
            v-for="item in sortedAndSearchedItems"
            :key="item.id"
            class="btn-primary items rounded"
            :title="item.id"
          >
            {{ item.id }}
          </button>
        </transition-group>
      </div>
    </perfect-scrollbar>
  </div>
</template>

<script>
// Compostion API
// import { useItems } from "@/hooks/useItems";
// import { useSortedItems } from "@/hooks/useSortedItems";
// import { useSortedAndSearchedItems } from "@/hooks/useSortedAndSearchedItems";

import { mapState, mapActions, mapGetters } from "vuex";

export default {
  methods: {
    ...mapActions({
      getItems: "items/getItems",
    }),
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
  },

  // Compostion API
  // setup(props) {
  //   const store = useStore();
  //   const { items, isItemsLoading } = useItems(5);
  //   const { sortedItems } = useSortedItems(items);
  //   const { sortedAndSearchedItems } = useSortedAndSearchedItems(
  //     sortedItems,
  //     store.state.items.searchQuery
  //   );
  //   // const { sortedItems } = useSortedItems(items);
  //   // const { searchQuery, sortedAndSearchedItems } =
  //   //   useSortedAndSearchedItems(sortedItems);

  //   // console.log(store.state.items.searchQuery)

  //   return {
  //     items,
  //     isItemsLoading,
  //     sortedItems,
  //     // searchQuery,
  //     sortedAndSearchedItems,
  //   };
  // },
};
</script>

<style scoped>
.items_lits {
  overflow: scroll;
  width: 300px;
  height: 150px;
  padding: 5px;
  border: solid 1px black;
}
.catalogue {
  background-color: rgba(255, 238, 0, 0.5);
  padding-top: 10px;
  overflow: hidden;
  /* max-height: 475px; */

  display: flex;
  justify-content: space-between;
  flex-direction: column;
  /* height: 100%; */
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

.ps {
  height: 400px;
}
</style>