<template>
  <img
    class="item"
    :disabled="!this.button"
    draggable="false"
    :title="item.name"
    v-lazy="{ src: getImageUrl(imagePath) }"
    @click="selectItem(itemInfo)"
    role="button"
  >
</template>

<script setup>
import { computed } from "vue";
import { useStore } from 'vuex';
const props = defineProps({
  item: Object,
  button: Boolean,
});

const store = useStore();
const items = computed(() => store.state.items.items);

const itemInfo = items.value.find(item => { return item.id == props.item.id; });

const imagePath = itemInfo.type.name.toLowerCase() + '/' + (itemInfo.type.name.toLowerCase() !== 'blueprints' ? itemInfo.name : (itemInfo.name === 'The General\'s Soul Reaper Orb' ? 'The General\'s Soul Reaper Orb' : itemInfo.tags[0])) + '.webp';

const getImageUrl = (name) => {
  return `images/items/${name}`;
};

function selectItem(item) {
  store.dispatch('items/selectItem', item);
}
</script>


<style scoped lang="scss">
.item {
  border-radius: 3px;
  -webkit-user-drag: none;
  user-select: none;
  color: white;
  text-align: right;
}
</style>