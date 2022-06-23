<template>
    <div class="d-flex flex-column block ">
        <div class="d-flex justify-content-between">
            <div class="d-flex">
                <span
                    v-for="(station, i) in recipe.stations"
                    class="h-3 "
                >
                    {{ (i > 0 ? '&nbsp;/&nbsp;' : '') + station }}
                </span>
            </div>
            <div
                v-if="this.recipe.time"
                class="d-flex"
            >
                <Popper
                    v-if="this.confinedRoom && this.confined"
                    hover
                    placement="left"
                    content="Time in a confined castle room."
                >
                    <span
                        v-if="this.confined"
                        class="h-3 tag tag__time-confined"
                        @click="this.updateConfinedRoom"
                    >
                        <span v-if="this.confinedTime.minutes">
                            {{ `${this.confinedTime.minutes}m${this.confinedTime.seconds ? '&nbsp;' : ''}` }}
                        </span>
                        <span v-if="this.confinedTime.seconds">
                            {{ `${this.confinedTime.seconds}s` }}
                        </span>
                    </span>
                </Popper>
                <Popper
                    v-else
                    interactive
                    hover
                    placement="left"
                    content="Time without bonuses."
                >
                    <span
                        class="h-3 tag tag__time-normal"
                        @click="this.updateConfinedRoom"
                    >
                        <span v-if="this.time.minutes">
                            {{ `${this.time.minutes}m${this.time.seconds ? '&nbsp;' : ''}` }}
                        </span>
                        <span v-if="this.time.seconds">
                            {{ `${this.time.seconds}s` }}
                        </span>
                    </span>
                </Popper>
            </div>
        </div>
        <div class="d-flex mt-1">
            <div class="d-flex block ">
                <div
                    v-for="ingridient in this.recipe.ingredients"
                    class="recipe__group px-1"
                >
                    <item-preview
                        :style="'preview-sm'"
                        :item="this.items.find(item => { return item.id === ingridient.itemId })"
                        :text="this.matchingFloor && this.confined ?
                        Math.ceil(ingridient.amount * 0.75) : ingridient.amount"
                        :button="true"
                    />
                </div>
            </div>

            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="40"
                height="40"
                fill="#a8a9ae"
                class=""
                viewBox="0 0 16 16"
            >
                <path
                    d="m12.14 8.753-5.482 4.796c-.646.566-1.658.106-1.658-.753V3.204a1 1 0 0 1 1.659-.753l5.48 4.796a1 1 0 0 1 0 1.506z"
                />
            </svg>

            <div class="d-flex flex-fill">
                <div
                    v-for="result in this.recipe.results"
                    class="recipe__group px-1"
                >
                    <item-preview
                        :style="'preview-sm'"
                        :item="this.items.find(item => { return item.id === result.itemId })"
                        :text="result.amount"
                        :button="true"
                    />
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import ItemPreview from "@/components/ItemsPage/ItemPreview";

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
@import '@/assets/styles/utility/vars.scss';

* {
    user-select: none;
}

.tag__time-confined {
    background: #176744;
}

.tag__time-normal {
    background: #ae1d1d;
}

.tag {
    border-radius: 100px;
    text-transform: capitalize;
    margin-right: 5px;
    border: none;
    padding-left: 10px;
    padding-right: 10px;
    color: silver;
}

.tag:hover {
    box-shadow: 0 0 5px black;
    transition: box-shadow 0.05s ease-in-out;
}
</style>