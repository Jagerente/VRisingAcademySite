<template>
  <ItemPreview
    class="item"
    :class="{ 'active': selectedItem !== null && item.id === selectedItem.id }"
    :item="item"
    button
    @itemClick="itemClick"
  />
</template>

<script>
import ItemPreview from '@/components/ItemsPage/ItemPreview.vue';

export default {
  components: {
    ItemPreview,
  },
};
</script>

<script setup>
import { computed } from "vue";
import { useStore } from 'vuex';

const props = defineProps({
  item: Object,
});

const store = useStore();

const selectedItem = computed(() => store.state.items.selectedItem);

const emits = defineEmits(["itemClick"]);

function itemClick(itemInfo) {
  emits("itemClick", itemInfo);
}
</script>

<style scoped lang="scss">
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