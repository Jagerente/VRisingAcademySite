<template>
  <div class="d-flex flex-column h-100">
    <my-navbar />
    <router-view />
    <my-footer v-show="true" />
  </div>
<!--	<div class="modal">-->
<!--		<div class="modal__overlay">-->
<!--			<div class="modal__content">-->
<!--				<button class="modal__button">Add</button>-->
<!--				<p class="modal__text">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus excepturi maiores mollitia, necessitatibus nemo quaerat tempore? Corporis deleniti est excepturi impedit iste iusto odio voluptates. A amet molestiae saepe sunt.</p>-->
<!--				<p class="modal__text">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus excepturi maiores mollitia, necessitatibus nemo quaerat tempore? Corporis deleniti est excepturi impedit iste iusto odio voluptates. A amet molestiae saepe sunt.</p>-->
<!--				<p class="modal__text">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus excepturi maiores mollitia, necessitatibus nemo quaerat tempore? Corporis deleniti est excepturi impedit iste iusto odio voluptates. A amet molestiae saepe sunt.</p>-->
<!--				<p class="modal__text">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus excepturi maiores mollitia, necessitatibus nemo quaerat tempore? Corporis deleniti est excepturi impedit iste iusto odio voluptates. A amet molestiae saepe sunt.</p>-->
<!--				<p class="modal__text">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus excepturi maiores mollitia, necessitatibus nemo quaerat tempore? Corporis deleniti est excepturi impedit iste iusto odio voluptates. A amet molestiae saepe sunt.</p>-->
<!--				<p class="modal__text">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus excepturi maiores mollitia, necessitatibus nemo quaerat tempore? Corporis deleniti est excepturi impedit iste iusto odio voluptates. A amet molestiae saepe sunt.</p>-->
<!--				<p class="modal__text">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus excepturi maiores mollitia, necessitatibus nemo quaerat tempore? Corporis deleniti est excepturi impedit iste iusto odio voluptates. A amet molestiae saepe sunt.</p>-->
<!--				<p class="modal__text">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus excepturi maiores mollitia, necessitatibus nemo quaerat tempore? Corporis deleniti est excepturi impedit iste iusto odio voluptates. A amet molestiae saepe sunt.</p>-->
<!--				<p class="modal__text">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus excepturi maiores mollitia, necessitatibus nemo quaerat tempore? Corporis deleniti est excepturi impedit iste iusto odio voluptates. A amet molestiae saepe sunt.</p>-->
<!--			</div>-->
<!--		</div>-->
<!--	</div>-->
</template>

<script>
import MyNavbar from "@/components/MyNavbar.vue";
import MyFooter from "./components/MyFooter.vue";

export default {
  components: {
    MyNavbar,
    MyFooter
  },
};
</script>

<style lang="scss">
@import "@/assets/styles/utility/vars.scss";
@import 'bootstrap/scss/bootstrap-grid';
@import 'bootstrap/scss/bootstrap-utilities';
@import 'assets/styles/main.scss';

body {
	font-family: $mainFontFamily;
	background-color: $background;
	color: $text-color;
}

//.modal {
//	//display: none;
//	opacity: 0;
//	transition: opacity .2s ease;
//	&__overlay {
//		position: fixed;
//		top: 0;
//		left: 0;
//		right: 0;
//		bottom: 0;
//		background: rgba(0,0,0,0.4);
//		display: flex;
//		align-items: center;
//		justify-content: center;
//	}
//	&__content {
//		border-radius: 20px;
//		background-color: #ffffff;
//		padding: 30px;
//	}
//
//	@media (max-width: $lg) {
//		//display: block;
//		max-height: 50%;
//		opacity: 1;
//		&__overlay {
//			display: flex;
//			align-items: flex-end;
//		}
//		&__content {
//			border-bottom-left-radius: 0;
//			border-bottom-right-radius: 0;
//			width: 100%;
//		}
//	}
//}
</style>