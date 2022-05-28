<template>
    <footer class="d-flex flex-column sticky-bottom h-100">
        <div class="d-flex mb-1 justify-content-center mt-auto">
            <my-link content="Join VK!" link="https://vk.com/vrisingacademy" icon="vk" />
            <my-link content="Join Telegram!" link="https://t.me/vrisingacademy" icon="tg" />
            <my-link content="Join Discord!" link="https://discord.gg/7jmczsbrCH" icon="discord" />
            <my-link content="Support us!" link="https://boosty.to/vrisingacademy" icon="boosty" />
        </div>
        <p class="text-center">© 2022 V Rising Academy · Genshin Academy Team</p>
    </footer>
</template>

<script>
import MyLink from "@/components/Footer/MyLink";

export default {
    components: {
        MyLink
    },
    name: "my-footer",
};
</script>

<style>
</style>