<template>
  <div class="d-flex justify-content-center flex-wrap">
    <div v-if="isItemsLoading" class="spinner-border text-primary position-absolute" role="status">
      <span class="sr-only"></span>
    </div>
    <transition-group v-else name="items-list">
      <my-item v-for="item in sortedAndSearchedItems(type)" :key="item.id" :item="item" @click="selectItem(item.id)" />
    </transition-group>
  </div>
</template>

<script>
import { mapState, mapActions, mapGetters } from "vuex";

export default {
  name: "items-list",
  props: {
    type: Number,
  },
  methods: {
    ...mapActions({
      selectItem: "items/selectItem",
    }),
  },
  computed: {
    ...mapState({
      isItemsLoading: (state) => state.items.isItemsLoading,
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

/* .items-list-leave-active {
  position: absolute;
} */

.input-group {
  height: 39px;
}
</style>