<template>
    <div class="recipe__group px-1 py-2">
        <item-preview class="preview-recipe" :item="this.items.find(item => { return item.id == this.item.itemId })"
            :button="true" @click="selectItem(this.items.find(item => { return item.id == this.item.itemId }))" />

        <div class="recipe__count">{{ item.amount }}</div>
    </div>
</template>

<script>
import { mapState, mapActions, mapGetters } from "vuex";

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
    pointer-events: none;
    position: absolute;
    bottom: 8px;
    right: 16px;
    font-size: 16px;
}
</style>