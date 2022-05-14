<template>
  <my-card title="Information">
    <div v-if="selectedItem === null">
      <h3>Select item</h3>
    </div>
    <div v-else>
      <h3>{{ selectedItem.name }}</h3>
      <h4>{{ selectedItem.description }}</h4>
      <h5>{{ selectedItem.tier }}</h5>
    </div>
  </my-card>
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