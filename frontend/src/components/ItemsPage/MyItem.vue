<template>
  <ItemPreview
    class="item"
    :class="{ 'active': this.selectedItem !== null && this.item.id === this.selectedItem.id }"
    :item="item"
    button
  />
</template>

<script>
import { mapState, mapActions } from "vuex";
import ItemPreview from '@/components/ItemsPage/ItemPreview.vue';

export default {
  components: {
    ItemPreview,
  },
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
@import "@/assets/styles/utility/vars.scss";

.item {
  $item-size: 3rem;
  border: 1px solid $dark;
  background-color: $dark;
  width: $item-size;
  height: $item-size;
  margin: 2px;
  transition: border 0.15s ease-in-out, box-shadow 0.15s ease-in-out;

  @media (min-width: $sm) {
    $item-size: 5rem;
    width: $item-size;
    height: $item-size;
    margin: 5px;
  }

  &:hover {
    box-shadow: 0 0 8px black;
  }

  &.active {
    border: 1px solid white;

    &:hover {
      box-shadow: none
    }
  }
}
</style>