<template>
  <img
    type="image"
    class="btn-primary items rounded"
    :title="item.name"
    :src="
      require('@/assets/images/items/' + getPath() + '/' + item.name + '.png')
    "
  />
</template>

<script>
import { mapState, mapMutations } from "vuex";

export default {
  name: "my-item",
  props: {
    item: Object,
    type: Number,
  },
  methods: {
    getPath() {
      switch (this.type) {
        case 1:
          return "weapons";
        case 2:
          return "armour";
        case 3:
          return "consumables";
        case 4:
          return "reagents";
        default:
          console.error("Wrong Item type:", type, item);
          return "all";
      }
    },
  },
};
</script>

<style scoped>
.items {
  background: rgba(0, 0, 0, 0.5);
  /* background-size: 100%; */
  /* background-position: 50%; */

  /* --item-size: 85px; */
  --item-size: 85px;
  font-size: 10px;
  font-family: sans-serif;
  width: var(--item-size);
  height: var(--item-size);
  color: black;
  box-shadow: 0 0 8px black;
  margin-bottom: 5px;
  -webkit-user-drag: none;
  border: 0;
  margin: 10px;
  user-select: none;
  transition: box-shadow 0.15s ease-in-out;
}

.items:hover {
  box-shadow: 0 0 13px var(--primary);
}

.btn-check:focus + .btn-primary,
.btn-primary:focus {
  /* background: rgba(0, 0, 0, 0.5); */
  background-color: black;
  background-size: 90%;
  background-position: 50%;
  background-repeat: no-repeat;
  border: 1px solid white;
  box-shadow: 0;
}
.btn-primary:hover {
  background-color: black;
  background-repeat: no-repeat;
  background-size: 90%;
}

.btn-check:focus + .btn-primary,
.btn-primary:focus {
  box-shadow: none;
}
</style>