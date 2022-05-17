<template>
    <div class="recipe__group">
        <input class="recipe__item rounded" type="image" :title="getItemById(item.itemId).name"
            @click="selectItem(item.itemId)"
            :src="
            require('@/assets/images/items/' + getPath(getItemById(item.itemId)) + '/' + getItemById(item.itemId).name + '.png')" />
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
        getItemById(id) {
            return [...this.items].find(item => item.id == id)
        },
    },
    computed: {
        ...mapState({
            items: (state) => state.items.items
        }),
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
    margin: 10px;
    user-select: none;
    transition: box-shadow 0.15s ease-in-out;
    color: white;
    text-align: right;
}

.recipe__item:hover {
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