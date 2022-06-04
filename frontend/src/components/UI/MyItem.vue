<template>
  <input type="image" class="btn-primary items rounded" :title="item.name" :src="
    require('@/assets/images/items/' + item.type.toLowerCase() + '/' + (item.type.toLowerCase() !== 'blueprints' ? item.name : (item.name === 'The General\'s Soul Reaper Orb' ? 'The General\'s Soul Reaper Orb' : item.tags[1])) + '.webp')
  " />
</template>

<script>
import { mapGetters } from "vuex";

export default {
  name: "my-item",
  props: {
    item: Object,
  },
  computed: {
    ...mapGetters({
      imagePath: "items/imagePath",
    })
  },
};
</script>

<style scoped lang="scss">
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';


@include media-breakpoint-down(sm) {
  $item-size: 3em;
  $item-margin: 2px;

  .items {
    width: $item-size;
    height: $item-size;
    margin: $item-margin;
  }
}

@include media-breakpoint-up(sm) {
  $item-size: 5em;
  $item-margin: 5px;

  .items {
    width: $item-size;
    height: $item-size;
    margin: $item-margin;
  }
}

.items {
  background: rgba(0, 0, 0, 0.5);
  font-family: sans-serif;
  color: #14131b;
  -webkit-user-drag: none;
  border: 0;
  user-select: none;
  transition: box-shadow 0.15s ease-in-out;
}


.items:hover {
  box-shadow: 0 0 8px black;
}

.btn-check:focus+.btn-primary,
.btn-primary:focus {
  /* background: rgba(0, 0, 0, 0.5); */
  background-color: #14131b;
  background-size: 90%;
  background-position: 50%;
  background-repeat: no-repeat;
  border: 1px solid white;
  box-shadow: 0;
}

.btn-primary:hover {
  box-shadow: 0 0 8px black;
  background-color: #14131b;
  background-repeat: no-repeat;
  background-size: 90%;
}

.btn-check:focus+.btn-primary,
.btn-primary:focus {
  box-shadow: none;
}
</style>