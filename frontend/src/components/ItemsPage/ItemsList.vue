<template>
  <div
    class="items"
    v-if="!searchQuery"
    v-for="set in sets"
    :key="set.id"
  >
    <p class="items__set">{{ set.name }}</p>
    <div class="items__list">
      <MyItem
        v-for="item in set.items"
        :key="item.id"
        :item="item"
        @itemClick="itemClick"
      />
    </div>
  </div>
  <div
    class="items__list"
    v-else
  >
    <MyItem
      v-for="item in sortedAndSearchedItems"
      :key="item.id"
      :item="item"
      @itemClick="itemClick"
    />
  </div>
</template>

<script>
import MyItem from "@/components/ItemsPage/MyItem.vue";

export default {
  components: {
    MyItem,
  },
};
</script>

<script setup>
import { computed } from "vue";
import { useStore } from 'vuex';

const props = defineProps({
  sets: Object,
});

const emits = defineEmits(["itemClick"]);

const store = useStore();

const sortedAndSearchedItems = computed(() => store.getters["items/sortedAndSearchedItems"]);

const selectItem = () => store.dispatch("items/selectItem");

const searchQuery = computed(() => store.state.items.searchQuery);

function itemClick(itemInfo) {
  emits("itemClick", itemInfo);
}
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