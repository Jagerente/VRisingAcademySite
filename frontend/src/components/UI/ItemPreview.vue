<template>
  <img class="rounded" :disabled="button" draggable="false" :title="item.name" :src="require('@/' + this.imagePath)"
    @click="selectItem(this.items.find(item => { return item.id == this.item.id }))" />
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
      return 'assets/images/items/' + this.item.type.name.toLowerCase() + '/' + (this.item.type.name.toLowerCase() !== 'blueprints' ? this.item.name : (this.item.name === 'The General\'s Soul Reaper Orb' ? 'The General\'s Soul Reaper Orb' : this.item.tags[1])) + '.webp';
    },
    ...mapState({
      items: (state) => state.items.items,
    }),
  },
};
</script>


<style scoped lang="scss">
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';

@include media-breakpoint-down(sm) {
  .preview-md {
    $size: 3em;
    $margin: 2px;
    width: $size;
    height: $size;
    margin: $margin;
  }

  .preview-lg {
    $size: 100px;
    width: $size;
    height: $size;
  }
}

@include media-breakpoint-up(sm) {

  .preview-md {
    $size: 5em;
    $margin: 5px;
    width: $size;
    height: $size;
    margin: $margin;
  }

  .preview-lg {
    $size: 200px;
    width: $size;
    height: $size;
  }
}

.preview-sm {
  $size: 48px;
  background: rgba(0, 0, 0, 0.5);
  font-family: sans-serif;
  width: $size;
  height: $size;
  -webkit-user-drag: none;
  border: 0;
  user-select: none;
  transition: box-shadow 0.15s ease-in-out;
  color: white;
  text-align: right;
}

.preview-sm:hover {
  box-shadow: 0 0 8px black;
}

.preview-md {
  background: rgba(0, 0, 0, 0.5);
  font-family: sans-serif;
  -webkit-user-drag: none;
  border: 0;
  user-select: none;
  transition: box-shadow 0.15s ease-in-out;
}

.preview-md:hover {
  box-shadow: 0 0 8px black;
}

.active {
  border: 1px solid white;
}

.active:hover {
  box-shadow: none
}
</style>