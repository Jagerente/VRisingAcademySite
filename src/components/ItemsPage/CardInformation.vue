<template>
  <div class="d-flex flex-column h-100 card__background">
    <div class="card__header">
      <a>Information</a>
    </div>
    <div v-if="selectedItem === null">
      <h3>Select item</h3>
    </div>
    <div v-else class="h-100 d-inline-block card__content">
      <h3>{{ selectedItem.name }}</h3>
      <h4>{{ selectedItem.description }}</h4>
      <h5>{{ selectedItem.tier }}</h5>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions, mapGetters } from "vuex";

export default {
  computed: {
    ...mapState({
      selectedItem: (state) => state.items.selectedItem,
    }),
  },
};
</script>

<style scoped>
.block {
  display: block;
  width: 100%;
}
</style>