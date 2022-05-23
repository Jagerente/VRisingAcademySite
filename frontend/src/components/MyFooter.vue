<template>
    <div class="d-flex flex-column justify-content-center">
        <div class="d-flex justify-content-center redirects">
            <a href="https://vk.com/vrisingacademy" target="blank">
                <img class="image" :src="
                    require('@/assets/images/ui/main_page/VK.webp')
                ">
            </a>
            <a href="https://t.me/vrisingacademy" target="blank">
                <img class="image" :src="
                    require('@/assets/images/ui/main_page/TG.webp')
                ">
            </a>
            <a href="https://boosty.to/vrisingacademy" target="blank">
                <img class="image" type="image" :src="
                    require('@/assets/images/ui/main_page/Boosty.webp')
                ">
            </a>
        </div>
        <p class="text-center">© 2022 V Rising Academy · Genshin Academy Team</p>
    </div>
</template>

<script>
export default {
    name: "my-footer",
};
</script>

<style scoped>
.image {
    width: 30px;
    margin-right: 5px;
    margin-left: 5px;
}
</style>