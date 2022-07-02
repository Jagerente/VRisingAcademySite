<template>
  <img
    class="item"
    :disabled="button"
    draggable="false"
    :title="item.name"
    :src="require('@/' + this.imagePath)"
    @click="selectItem(this.items.find(item => { return item.id == this.item.id }))"
    role="button"
  >
</template>

<script>
import { mapState, mapActions, mapMutations } from "vuex";

export default {
  name: "item-preview",
  props: {
    item: Object,
    button: Boolean,
  },
  methods: {
    ...mapActions({
      selectItem: "items/selectItem",
    }),
  },
  computed: {
    imagePath() {
      return 'assets/images/items/' + this.item.type.name.toLowerCase() + '/' + (this.item.type.name.toLowerCase() !== 'blueprints' ? this.item.name : (this.item.name === 'The General\'s Soul Reaper Orb' ? 'The General\'s Soul Reaper Orb' : this.item.tags[0])) + '.webp';
    },
    ...mapState({
      items: (state) => state.items.items,
      selectedItem: (state) => state.items.selectedItem,
    }),
  },
};
</script>


<style scoped lang="scss">
@import "@/assets/styles/utility/vars.scss";

.item {
  border-radius: 3px;
  -webkit-user-drag: none;
  user-select: none;
  color: white;
  text-align: right;
}
</style>