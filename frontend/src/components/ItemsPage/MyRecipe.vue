<template>
    <div class="recipe">
        <div class="header recipe__header">
            <div class="header__stations">
                <span
                    v-for="(station, i) in recipe.stations"
                    class="station"
                >
                    {{ (i > 0 ? '&nbsp;/&nbsp;' : '') + station }}
                </span>
            </div>
            <Popper
                v-if="this.recipe.time"
                hover
                placement="left"
                :content="this.confinedRoom && this.confined ? 'Time in a confined castle room.' : 'Time without bonuses.'"
            >
                <span
                    v-if="this.confined"
                    class="header__time"
                    :class="{ 'confined': this.confinedRoom && this.confined }"
                    @click="this.updateConfinedRoom"
                >
                    <span v-if="this.confinedTime.minutes">
                        {{ this.confinedRoom && this.confined ?
                                `${this.confinedTime.minutes}m${this.confinedTime.seconds ? '&nbsp;' : ''}` :
                                `${this.time.minutes}m${this.time.seconds ? '&nbsp;' : ''}`
                        }}
                    </span>
                    <span v-if="this.confinedTime.seconds">
                        {{ this.confinedRoom && this.confined ? `${this.confinedTime.seconds}s` :
                                `${this.time.seconds}s`
                        }}
                    </span>
                </span>
            </Popper>
        </div>
        <div class="content recipe__content">
            <div
                class="item content__input"
                v-for="ingridient in this.recipe.ingredients"
            >
                <ItemPreview
                    class="item__image"
                    :item="this.items.find(item => { return item.id === ingridient.itemId })"
                    :button="true"
                    @itemClick="selectItem"
                />
                <div class="item__text">
                    {{ this.matchingFloor && this.confined ?
                            Math.ceil(ingridient.amount * 0.75) : ingridient.amount
                    }}
                </div>
            </div>
            <span class="content__divider">
            </span>
            <div
                class="item content__output"
                v-for="result in this.recipe.results"
            >
                <ItemPreview
                    class="item__image output__item"
                    :item="this.items.find(item => { return item.id === result.itemId })"
                    :button="true"
                    @itemClick="selectItem"
                />
                <div class="item__text">
                    {{ result.amount }}
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import ItemPreview from "@/components/ItemsPage/ItemPreview.vue";

export default {
    components: {
        ItemPreview,
    },
    name: "my-recipe",
    props: {
        recipe: Object,
    },
    methods: {
        ...mapActions({
            updateConfinedRoom: "items/updateConfinedRoom",
            selectItem: "items/selectItem",
        }),
    },
    computed: {
        ...mapState({
            items: (state) => state.items.items,
            matchingFloor: (state) => state.items.matchingFloor,
            confinedRoom: (state) => state.items.confinedRoom,
        }),
        confined() {
            return this.recipe.stations.some(station1 => 'simple workbench,sawmill,furnace,grinder,tannery,blood press,woodworking bench,vermin nest,alchemy table,tailoring bench,smithy,loom,jewelcrafting table,gem cutting table,paper press,anvil'.split(',').some(station2 => station1.toLowerCase() === station2.toLowerCase()))
        },
        time() {
            var multiplier = 1;
            var minutes = Math.floor(this.recipe.time * multiplier / 60)
            var seconds = this.recipe.time * multiplier - (minutes * 60)
            var time = {
                seconds: seconds,
                minutes: minutes
            }
            return time;
        },
        confinedTime() {
            var multiplier = 0.8;
            var minutes = Math.floor(this.recipe.time * multiplier / 60)
            var seconds = this.recipe.time * multiplier - (minutes * 60)
            var time = {
                seconds: seconds,
                minutes: minutes
            }
            return time;
        }
    },
};
</script>

<style scoped lang="scss">
.recipe {
    &__header {
        display: flex;
        justify-content: space-between;
        margin-bottom: $m1;
    }

    &__content {
        display: flex;
    }
}

.header {
    &__stations {
        user-select: none;
    }

    &__time {
        border-radius: 100px;
        text-transform: capitalize;
        margin-right: 5px;
        padding: 2px 10px;
        background-color: #282737;
        color: $text-color;
        user-select: none;
        box-shadow: none;
        transition: box-shadow .1s ease-in-out, background-color .25s ease-in-out, color .25s ease-in-out;

        &:hover {
            box-shadow: 0 0 5px black;
        }

        &.confined {
            background-color: $primary;
            color: white;
        }
    }
}

.content {

    &__input,
    &__output {
        display: flex;
    }

    &__divider {
        $size: 1rem;
        content: "";
        display: flex;
        flex-direction: column;
        margin: auto 1rem;
        border-top: $size solid transparent;
        border-bottom: $size solid transparent;
        border-right: 0;
        border-left: $size solid;
    }
}

.item {
    position: relative;

    &__image {
        position: relative;
        $item-size: 2.5rem;
        width: $item-size;
        height: $item-size;
        background: black;
        border: 1px solid black;
    }

    &__text {
        position: absolute;
        color: white;
        font-family: sans-serif;
        pointer-events: none;
        bottom: 0px;
        right: 0px;
        font-size: 16px;
        user-select: none;
    }
}
</style>