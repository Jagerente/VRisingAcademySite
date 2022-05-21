<template>
  <input type="image" class="item rounded m-0 p-0" :title="getItemById.name" :src="
    require('@/assets/images/items/' + getPath(getItemById.type) + '/' + getItemById.name + '.webp')
  " />
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
    getPath(type) {
      switch (type) {
        case 1:
          return "weapons";
        case 2:
          return "armour";
        case 3:
          return "consumables";
        case 4:
          return "reagents";
        default:
          console.error("Wrong Item type:", type);
          return null;
      }
    },
  },
  computed: {
    ...mapState({
      items: (state) => state.items.items
    }),
    getItemById() {
      return [...this.items].find(item => item.id == this.itemId);
    },
  },
};
</script>


<style scoped>
.item {
  background: rgba(0, 0, 0, 0.5);
  /* background-size: 100%; */
  /* background-position: 50%; */

  /* --item-size: 85px; */
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