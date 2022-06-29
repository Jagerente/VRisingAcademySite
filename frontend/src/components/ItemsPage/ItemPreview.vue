<template>
  <div
    class="preview"
    :class="text != null ? '' : ''"
  >
    <img
      class="preview__item"
      :class="this.style"
      :disabled="button"
      draggable="false"
      :title="item.name"
      :src="require('@/' + this.imagePath)"
      @click="selectItem(this.items.find(item => { return item.id == this.item.id }))"
      role="button"
    >
    <div
      v-if="this.text"
      class="preview__text"
    >{{ this.text }}</div>
  </div>
</template>

<script>
import { mapState, mapActions, mapMutations } from "vuex";

export default {
  name: "item-preview",
  props: {
    item: Object,
    button: Boolean,
    style: String,
    text: {
      Type: String,
      Default: null
    }
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

.preview {
  position: relative;
  text-align: center;
  color: white;

  &__text {
    pointer-events: none;
    position: absolute;
    bottom: 0px;
    right: 0px;
    font-size: 16px;
  }

  &__item {
    background: $dark;
    font-family: sans-serif;
    border: 1px solid $dark;
    border-radius: 3px;
    -webkit-user-drag: none;
    user-select: none;
    transition: border 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
    color: white;
    text-align: right;

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

  .item-sm {
    $item-size: 2.5rem;
    width: $item-size;
    height: $item-size;
    background: black;
    border: 1px solid black;
  }

  .item-md {
    $item-size: 3rem;
    width: $item-size;
    height: $item-size;
    margin: 2px;

    @media (min-width: $sm) {
      $item-size: 5rem;
      width: $item-size;
      height: $item-size;
      margin: 5px;
    }
  }

  .item-lg {
    background: none;
    $item-size: 10rem;
    width: $item-size;
    height: $item-size;
    border: none;
    pointer-events: none;

    &:hover {
      box-shadow: none;
    }

    @media (min-width: $sm) {
      $item-size: 13rem;
      width: $item-size;
      height: $item-size;
    }
  }
}
</style>