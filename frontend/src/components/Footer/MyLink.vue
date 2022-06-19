<template>
    <Popper placement="top" arrow hover :content="this.content">
        <a :href="link" target="blank">
            <img class="image" :src="
                require('@/assets/images/ui/main_page/' + this.icon + '.webp')
            ">
        </a>
    </Popper>
</template>

<script>
export default {
    props: {
        content: String,
        link: String,
        icon: String
    },
    name: "my-link",
};
</script>

<!--<style scoped>-->
<!--.image {-->
<!--    width: 30px;-->
<!--    margin-right: 5px;-->
<!--    margin-left: 5px;-->
<!--}-->
<!--</style>-->