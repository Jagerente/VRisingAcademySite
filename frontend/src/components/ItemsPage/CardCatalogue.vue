<template>
  <my-card :custom="true">
    <template #header>
      <filter-input></filter-input>
    </template>
    <div class="tab-content d-flex justify-content-center flex-wrap" id="v-pills-tabContent">
      <div v-for="type in this.itemsGrouped" class="tab-pane fade show w-100" :class="type.id === 1 ? 'active' : ''"
        :id="`v-pills-${type.name.toLowerCase()}`" role="tabpanel"
        :aria-labelledby="`v-pills-${type.name.toLowerCase()}-tab`">
        <items-list :sets="type.sets"></items-list>
      </div>
    </div>
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
    }),
  },
};
</script>

<style scoped>
</style>