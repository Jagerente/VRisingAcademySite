<template>
  <my-card :custom="true">
    <template #header>
      <filter-input class="header__input"></filter-input>
    </template>
    <items-list
      :sets="itemsGrouped[selectedType].sets"
      @itemClick="itemClick"
    ></items-list>
  </my-card>
</template>

<script>
import FilterInput from "@/components/ItemsPage/FilterInput.vue";
import ItemsList from "@/components/ItemsPage/ItemsList.vue";

export default {
  components: { FilterInput, ItemsList },
};
</script>

<script setup>
import { computed } from "vue";
import { useStore } from 'vuex';
const store = useStore();
const selectItem = (itemInfo) => store.dispatch("items/selectItem", itemInfo);

function itemClick(itemInfo) {
  selectItem(itemInfo);
}

const itemsGrouped = computed(() => store.state.items.itemsGrouped);
const selectedType = computed(() => store.state.items.selectedType);

</script>