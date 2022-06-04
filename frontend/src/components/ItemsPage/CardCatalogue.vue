<template>
  <my-card :custom="true">
    <template #header>
      <filter-input></filter-input>
    </template>
    <div class="tab-content d-flex justify-content-center flex-wrap" id="v-pills-tabContent">
      <div v-for="(type, i) in types" class="tab-pane fade show w-100" :class="i == 0 ? 'active' : ''"
        :id="`v-pills-${type.title.toLowerCase()}`" role="tabpanel"
        :aria-labelledby="`v-pills-${type.title.toLowerCase()}-tab`">
        <items-list :items="this.items[type.title]"></items-list>
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
      types: (state) => state.items.types,
      items: (state) => state.items.items,
    }),
  },
};
</script>

<style scoped>
</style>