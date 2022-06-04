<template>
  <input type="image" class="item rounded m-0 p-0" @click="selectItem(getItemById)" :title="getItemById.name" :src="
    require('@/assets/images/items/' + getItemById.type.toLowerCase() + '/' + (getItemById.type.toLowerCase() !== 'blueprints' ? getItemById.name : (getItemById.name === 'The General\'s Soul Reaper Orb' ? 'The General\'s Soul Reaper Orb' : getItemById.tags[1])) + '.webp')" />
</template>

<script>
import { mapState, mapActions } from "vuex";

export default {
  name: "item-preview",
  props: {
    itemId: Number,
    button: Boolean
  },
  methods: {
    ...mapActions({
      selectItem: "items/selectItem",
    }),
  },
  computed: {
    ...mapState({
      itemsList: (state) => state.items.itemsList
    }),
    getItemById() {
      return this.itemsList.find(item => {
        return item.id == this.itemId;
      })
    },
  },
};
</script>


<style scoped>
.item {
  background: rgba(0, 0, 0, 0.5);
  --item-size: 48px;
  font-family: sans-serif;
  width: var(--item-size);
  height: var(--item-size);
  -webkit-user-drag: none;
  border: 0;
  user-select: none;
  transition: box-shadow 0.15s ease-in-out;
  color: white;
  text-align: right;
}

.item:hover {
  box-shadow: 0 0 8px black;
}
</style>