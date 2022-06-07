<template>
  <item-preview class="preview-md rounded"
    :class="this.selectedItem !== null && this.item.id === this.selectedItem.id ? 'active' : ''" :item="item"
    @click="selectItem(item)" />
</template>

<script>
import { mapState, mapActions } from "vuex";

export default {
  name: "my-item",
  props: {
    item: Object,
  },
  methods: {
    ...mapActions({
      selectItem: "items/selectItem",
    }),
  },
  computed: {
    ...mapState({
      selectedItem: (state) => state.items.selectedItem,
    })
  }
};
</script>

<style scoped lang="scss">
</style>