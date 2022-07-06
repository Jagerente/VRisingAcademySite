<template>
	<footer class="footer">
		<div class="footer__container">
			<div class="footer__social">
				<my-link content="Join VK!" link="https://vk.com/vrisingacademy" icon="vk" />
				<my-link content="Join Telegram!" link="https://t.me/vrisingacademy" icon="tg" />
				<my-link content="Join Discord!" link="https://discord.gg/7jmczsbrCH" icon="discord" />
				<my-link content="Support us!" link="https://boosty.to/vrisingacademy" icon="boosty" />
			</div>
			<p class="footer__copy">© 2022 V Rising Academy · Genshin Academy Team</p>
		</div>
	</footer>
</template>

<script>
import MyLink from "@/components/Footer/MyLink.vue";

export default {
    components: {
        MyLink
    },
    name: "my-footer",
};
</script>

<style lang="scss">
.footer {
	padding: 20px 0;
	&__container {
		margin: 0 auto;
	}
	&__social,
	&__copy {
		display: flex;
		justify-content: center;
		align-items: center;
	}
	&__social {
		margin-bottom: 5px;
	}
}
</style>