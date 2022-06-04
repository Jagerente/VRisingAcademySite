<template>
    <div class="recipe__group px-1 py-2">
        <input class="recipe__item rounded" type="image" :title="getItemById.name" @click="selectItem(getItemById)"
            :src="
            require('@/assets/images/items/' + getItemById.type.toLowerCase() + '/' + (getItemById.type.toLowerCase() !== 'blueprints' ? getItemById.name : (getItemById.name === 'The General\'s Soul Reaper Orb' ? 'The General\'s Soul Reaper Orb' : getItemById.tags[1])) + '.webp')" />
        <div class="recipe__count">{{ item.amount }}</div>
    </div>
</template>

<script>
import { mapState, mapActions, mapMutations } from "vuex";

export default {
    name: "my-recipe",
    props: {
        item: Object,
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
                return item.id == this.item.itemId;
            })
        },
    },
};
</script>

<style scoped>
.recipe__group {
    position: relative;
    text-align: center;
    color: white;
}

.recipe__count {
    pointer-events: none;
    position: absolute;
    bottom: 8px;
    right: 16px;
    font-size: 16px;
}

.recipe__item {
    background: rgba(0, 0, 0, 0.5);
    /* background-size: 100%; */
    /* background-position: 50%; */

    /* --item-size: 85px; */
    --item-size: 32px;
    font-family: sans-serif;
    width: var(--item-size);
    height: var(--item-size);
    -webkit-user-drag: none;
    border: 0;
    user-select: none;
    margin-right: 5px;
    transition: box-shadow 0.15s ease-in-out;
    color: white;
    text-align: right;
}

.recipe__item:hover {
    box-shadow: 0 0 8px black;
}
</style>