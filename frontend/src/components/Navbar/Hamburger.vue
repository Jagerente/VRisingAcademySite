<template>
    <div
        :class="{ 'open': isOpen }"
        class="hamburger"
        role="button"
        @click="change"
    >
        <div class="icon-left"></div>
        <div class="icon-right"></div>
    </div>
</template>

<script setup>
import { ref } from "vue";

const isOpen = ref(false);

const emit = defineEmits();

function change() {
    isOpen.value = !isOpen.value;
    emit('onChange', isOpen.value);
};
</script>

<style scoped lang="scss">
$color: $text-color;
$burger-top: 20px;
$burger-right: 25px;
$icon-thickness: 4px;
$icon-width: 20px;
$icon-gap: 15px;
$icon-top: $burger-top - calc($icon-gap / 2) - calc($icon-thickness / 3) + 8px;
$icon-degree: 45deg;
$transition-duration: .25s;

@mixin burger($right, $top) {
    position: absolute;
    width: 40px;
    height: calc($icon-gap*2) + calc($icon-thickness*3);
    top: $top;
    right: $right;
    transition-duration: $transition-duration;
}

@mixin icon($height, $width, $top) {
    transition-duration: $transition-duration;
    position: absolute;
    top: $top;
    height: $height;
    width: $width;
    background-color: $color;
}

@mixin icon-before($height, $width, $top) {
    transition-duration: $transition-duration;
    position: absolute;
    width: $width;
    height: $height;
    background-color: $color;
    content: "";
    top: $top;
}

@mixin icon-after($height, $width, $top) {
    transition-duration: $transition-duration;
    position: absolute;
    width: $width;
    height: $height;
    background-color: $color;
    content: "";
    top: $top;
}

.hamburger {
    @include burger($burger-right, $burger-top);

    @media (min-width: $lg) {
        display: none;
    }

    .icon-left {
        @include icon($icon-thickness, $icon-width, $icon-top);
        left: 0px;

        &:before {
            @include icon-before($icon-thickness, $icon-width, -$icon-gap);
        }

        &:after {
            @include icon-after($icon-thickness, $icon-width, $icon-gap);
        }
    }

    .icon-right {
        @include icon($icon-thickness, $icon-width, $icon-top);
        left: $icon-width;

        &:before {
            @include icon-before($icon-thickness, $icon-width, -$icon-gap);
        }

        &:after {
            @include icon-after($icon-thickness, $icon-width, $icon-gap);
        }
    }

    &.open {
        .icon-left {
            transition-duration: $transition-duration;
            background: transparent;

            &:before {
                transform: rotateZ($icon-degree) scaleX(1.4) translate(4px, 4px);
            }

            &:after {
                transform: rotateZ(-$icon-degree) scaleX(1.4) translate(4px, -4px);
            }
        }

        .icon-right {
            transition-duration: $transition-duration;
            background: transparent;

            &:before {
                transform: rotateZ(-$icon-degree) scaleX(1.4) translate(-4px, 4px);
            }

            &:after {
                transform: rotateZ($icon-degree) scaleX(1.4) translate(-4px, -4px);
            }
        }
    }
}
</style> 