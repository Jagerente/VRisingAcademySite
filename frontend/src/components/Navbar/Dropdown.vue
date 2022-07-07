<template>
    <div class="dropdown">
        <div
            class="dropdown__title"
            @click="show = !show"
        >{{ name }}</div>
        <Transition name="fade">
            <ul
                v-show="show"
                v-click-away="onClickAway"
                class="dropdown__list"
                name="catalogue"
            >
                <li
                    v-for="item in items"
                    class="dropdown__item"
                >
                    <a
                        href="#"
                        class="dropdown__link"
                        :class="{ 'disabled': item.url === '' }"
                        @click="linkClick(item.url)"
                    >
                        {{ item.name }}
                    </a>
                </li>
            </ul>
        </Transition>
    </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import type { Ref } from 'vue'

import { useRouter, useRoute } from 'vue-router';
const router = useRouter();

type DropdownItem = { name: string, url: string };

const props = defineProps({
    name: String,
    items: Array<DropdownItem>
});

const show: Ref<Boolean> = ref(false);

function onClickAway(event) {
    if (!event.target.className.includes("dropdown__title"))
        show.value = false;
};

function linkClick(url) {
    router.push(url);
    show.value = false;
};
</script>

<style scoped lang="scss">
.dropdown {
    position: relative;
    z-index: 10;

    &__link,
    &__title {
        color: $text-color;
        text-decoration: none;
        transition: color .2s ease;

        &.active {
            color: #ffffff;
        }

        &:hover {
            color: #ffffff;
        }
    }

    &__title:after {
        display: inline-block;
        margin-left: 5px;
        vertical-align: 5px;
        content: "";
        border-top: 0.3em solid;
        border-right: 0.3em solid transparent;
        border-bottom: 0;
        border-left: 0.3em solid transparent;
        transition: transform .2s ease;
    }

    &__link {
        font-size: 16px;

        &.disabled {
            opacity: 0.5;
            pointer-events: none;
            cursor: default;
        }
    }

    &__list {
        // display: none;
        position: absolute;
        top: 40px;
        left: 0;
        right: 0;
        background-color: $background;
        padding: 15px;
        border: 1px solid rgba(0, 0, 0, .15);
        border-radius: 5px;

        // &.active {
        // 	display: block;
        // }
    }

    &__item {
        &+& {
            margin-top: 5px;
        }
    }
}
</style>