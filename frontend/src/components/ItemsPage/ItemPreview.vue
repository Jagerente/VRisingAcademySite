<template>
  <img
    class="item"
    draggable="false"
    :disabled="!button"
    :title="item.name"
    v-lazy="{ src: getImageUrl(imagePath) }"
    @click="itemClick"
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
    itemClick() {
      this.$emit("itemClick", this.items.find(item => { return item.id == this.item.id; }));
    }
  },
  computed: {
    imagePath() {
      const item = this.items.find(item => { return item.id == this.item.id; });
      return item.type.name.toLowerCase() + '/' + (item.type.name.toLowerCase() !== 'blueprints' ? item.name : (item.name === 'The General\'s Soul Reaper Orb' ? 'The General\'s Soul Reaper Orb' : item.tags[0])) + '.webp';
    },
    ...mapState({
      items: (state) => state.items.items,
    }),
  },
  setup(props) {
    const getImageUrl = (name) => {
      return `images/items/${name}`;
    };
    return { getImageUrl };
  }
};
</script>

<!-- BUG Image not changes -->
<!-- <script setup>
import { computed, onMounted } from "vue";
import { useStore } from 'vuex';

const props = defineProps({
  item: Object,
  button: Boolean,
});

const emits = defineEmits(["itemClick"]);

const store = useStore();

const items = computed(() => store.state.items.items);
const selectedItem = computed(() => store.state.items.selectedItem);

const itemInfo = computed(() => items.value.find(item => { return item.id == props.item.id; })).value;

const imagePath = computed(() =>
  itemInfo.type.name.toLowerCase() + '/' +
  (itemInfo.type.name.toLowerCase() !== 'blueprints' ?
    itemInfo.name :
    (itemInfo.name === 'The General\'s Soul Reaper Orb' ?
      'The General\'s Soul Reaper Orb' :
      itemInfo.tags[0])) + '.webp'
);

const getImageUrl = (name) => {
  return `images/items/${name}`;
};

function itemClick() {
  emits("itemClick", itemInfo);
}
</script> -->


<style scoped lang="scss">
.item {
  border-radius: 3px;
  -webkit-user-drag: none;
  user-select: none;
  color: white;
  text-align: right;
}
</style>