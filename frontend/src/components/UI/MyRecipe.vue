<template>
    <div class="recipe__group px-1 py-2">
        <input class="recipe__item rounded" type="image" :title="getItemById.name" @click="selectItem(item.itemId)"
            :src="
            require('@/assets/images/items/' + getPath(getItemById) + '/' + getItemById.name + '.webp')" />
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
        getPath(item) {
            switch (item.type) {
                case 1:
                    return "weapons";
                case 2:
                    return "armour";
                case 3:
                    return "consumables";
                case 4:
                    return "reagents";
                default:
                    console.error("Wrong Item type:", this.item.type, item);
                    return "all";
            }
        },
    },
    computed: {
        ...mapState({
            items: (state) => state.items.items
        }),
        getItemById() {
            return [...this.items].find(item => item.id == this.item.itemId)
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