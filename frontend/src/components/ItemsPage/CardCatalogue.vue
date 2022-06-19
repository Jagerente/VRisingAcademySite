<template>
  <my-card :custom="true">
    <template #header>
      <filter-input></filter-input>
    </template>
    <!-- <div
      class="tab-content d-flex justify-content-center flex-wrap"
      id="v-pills-tabContent"
    >
      <div
        v-for="itemType in this.itemsGrouped"
        class="tab-pane fade show w-100"
        :class="itemType.id === 1 ? 'active' : ''"
        :id="`v-pills-${itemType.name.toLowerCase()}`"
        role="tabpanel"
        :aria-labelledby="`v-pills-${itemType.name.toLowerCase()}-tab`"
      > -->
    <items-list :sets="this.itemsGrouped[this.selectedType].sets"></items-list>
    <!-- </div>
    </div> -->
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