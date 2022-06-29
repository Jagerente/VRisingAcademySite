<template>
  <my-card :custom="true">
    <template #header>
      <filter-input></filter-input>
    </template>
    <items-list :sets="this.itemsGrouped[this.selectedType].sets"></items-list>
  </my-card>
</template>

<script>
import FilterInput from "@/components/ItemsPage/FilterInput.vue";
import ItemsList from "@/components/ItemsPage/ItemsList.vue";
import { mapState } from "vuex";

export default {
  components: { FilterInput, ItemsList },
  computed: {
    ...mapState({
      itemsGrouped: (state) => state.items.itemsGrouped,
      selectedType: (state) => state.items.selectedType,
    }),
  },
};
</script>