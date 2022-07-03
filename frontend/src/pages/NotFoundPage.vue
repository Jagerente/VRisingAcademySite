<template>
    <div class="content">
        <div class="content__image">
            <img
                class="image"
                :src="require('@/assets/images/ui/404_page/Militia - Crier.webp')"
                draggable="false"
            >
        </div>
        <div class="content__text">
            <p class="text--1">
                404
            </p>
            <p class="text--2">
                Not Found
            </p>
        </div>
    </div>
</template>

<style scoped lang="scss">
.content {
    display: flex;
    justify-content: center;
    align-items: center;

    &__image {
        display: flex;
        justify-content: flex-end;
        width: 50%;
    }

    &__text {
        width: 50%;
    }
}

.image {
    width: 35%;
}

.text {
    &--1 {
        font-size: 5rem;
    }

    &--2 {
        font-size: 1.8rem;
    }
}
</style>