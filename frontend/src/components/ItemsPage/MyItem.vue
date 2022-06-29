<template>
  <ItemPreview
    :style="`item-md ${this.selectedItem !== null && this.item.id === this.selectedItem.id ? 'active' : ''}`"
    :item="item"
    @click="selectItem(item)"
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