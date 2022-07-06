<template>
    <label class="switch">
        <input
            @click="this.update"
            class="switch__input"
            :checked="this.checked"
            type="checkbox"
        >
        <span class="switch__slider"></span>
        <span class="switch__label">{{ this.label }}</span>
    </label>
</template>

<script>
export default {
    name: 'MySwitch',
    props: {
        label: String,
        update: Function,
        checked: Boolean
    },
}
</script>

<style lang="scss" scoped>
$slider_heigth: 20px;
$slider_width: 40px;
$slider_translate: $slider_width - 20px;
$dot_size: 16px;
$dot_translate: 2px;

.switch {
    position: relative;
    display: inline-block;
    height: 100%;

    &:not(:last-child) {
        margin-bottom: 2px;
    }

    &__input {
        opacity: 0;
        width: 0;
        height: 0;

        &:checked+.switch__slider {
            background-color: $primary;
        }

        &:not(:checked)+.switch__slider {
            background-color: $dark;
        }

        &:checked+.switch__slider::before {
            -webkit-transform: translateX($slider_translate);
            -ms-transform: translateX($slider_translate);
            transform: translateX($slider_translate);
        }

        &:focus+.switch__slider {
            box-shadow: none;
        }
    }

    &__slider {
        position: relative;
        display: inline-block;
        vertical-align: middle;
        margin-right: $m1;
        height: $slider_heigth;
        width: $slider_width;
        border-radius: 34px;
        background-color: #ccc;
        -webkit-transition: .4s;
        transition: .4s;

        &::before {
            position: absolute;
            content: "";
            height: $dot_size;
            width: $dot_size;
            left: $dot_translate;
            bottom: $dot_translate;
            border-radius: 50%;
            background-color: white;
            -webkit-transition: .4s;
            transition: .4s;
        }
    }

    &__label {
        display: inline-block;
        vertical-align: middle;
        color: white;
        user-select: none;
    }
}
</style>