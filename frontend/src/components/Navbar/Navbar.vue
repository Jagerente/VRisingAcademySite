<template>
	<header class="header">
		<div class="header__container">
			<img
				src="@/assets/images/ui/logo.webp"
				height="60"
				width="210"
				alt="logo"
				class="header__logo"
			>
			<nav class="header__nav nav">
				<ul class="nav__list">
					<li class="nav__item">
						<a
							href="#"
							class="nav__link"
							@click="$router.push('/')"
						>Home</a>
					</li>
					<li class="nav__item">
						<Dropdown
							class="nav__dropdown"
							:name="catalogueDropdown.name"
							:items="catalogueDropdown.items"
						/>
					</li>
					<li class="nav__item">
						<a
							href="https://mapgenie.io/v-rising/maps/vardoran"
							target="blank"
							class="nav__link"
						>Map</a>
					</li>
					<li class="nav__item">
						<a
							href=""
							class="nav__link disabled"
						>Guides</a>
					</li>
				</ul>
			</nav>
		</div>
	</header>
</template>

<script>
import Dropdown from "@/components/Navbar/Dropdown.vue";
export default {
	components: {
		Dropdown,
	},
	data() {
		return {
			catalogueDropdown: {
				name: 'Catalogue',
				items: [
					{
						name: 'Items',
						url: '/items'
					},
					{
						name: 'Spells',
						url: '/spells'
					},
					{
						name: 'Bestiary',
						url: ''
					},
					{
						name: 'Hunts',
						url: ''
					},
				]
			}
		}
	},
};
</script>

<style lang="scss" scoped>
.header {
	user-select: none;
	background-color: $background;
	box-shadow: 0 5px 5px #11111d;
	padding: 10px;

	&__container {
		max-width: 1320px;
		margin: 0 auto;
	}

	&__logo {
		display: inline-block;
		vertical-align: middle;
		width: auto;
		margin-right: 30px;
	}

	&__nav {
		display: inline-block;
		vertical-align: middle;
	}
}

.nav {
	&__list {
		margin: 0;
		padding: 0;
		list-style: none;
	}

	&__item {
		display: inline-block;
		font-size: 21px;
	}

	&__dropdown {
		color: $text-color;
		padding: 8px 16px;
		text-transform: uppercase;
		letter-spacing: 0.35rem;
		text-decoration: none;
		transition: color .2s ease;

		&.active {
			color: #ffffff;
		}

		&:hover {
			color: #ffffff;
		}
	}

	&__link {
		color: $text-color;
		padding: 8px 16px;
		text-transform: uppercase;
		letter-spacing: 0.35rem;
		text-decoration: none;
		transition: color .2s ease;

		&.active {
			color: #ffffff;
		}

		&:hover {
			color: #ffffff;
		}

		&.disabled {
			opacity: 0.5;
			pointer-events: none;
			cursor: default;
		}
	}
}
</style>