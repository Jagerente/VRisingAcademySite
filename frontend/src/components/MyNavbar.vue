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
							href=""
							class="nav__link"
							@click="$router.push('/')"
						>Home</a>
					</li>
					<li class="nav__item">
						<div class="nav__dropdown dropdown">
							<div
								class="dropdown__title"
								@click="this.toggleDropdown"
							>Catalogue</div>
							<ul
								v-show="this.catalogueDropdown"
								class="dropdown__list"
								name="catalogue"
							>
								<li class="dropdown__item">
									<a
										href="#"
										class="dropdown__link"
										@click="$router.push('/items')"
									>Items</a>
								</li>
								<li class="dropdown__item">
									<a
										href="#"
										class="dropdown__link"
										@click="$router.push('/spells')"
									>Spells</a>
								</li>
								<li class="dropdown__item">
									<a
										href="#"
										class="dropdown__link"
										@click="$router.push('/bloodtypes')"
									>Blood types</a>
								</li>
								<li class="dropdown__item">
									<a
										href="#"
										class="dropdown__link disabled"
									>Hunts</a>
								</li>
								<li class="dropdown__item">
									<a
										href="#"
										class="dropdown__link disabled"
									>Bestiary</a>
								</li>
							</ul>
						</div>
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
export default {
	name: "my-navbar",
	data() {
		return {
			dropdowns: [],
		}
	},
	methods: {
		toggleDropdown(el) {
			el.target.nextSibling.style.display = el.target.nextSibling.style.display === "none" ? "block" : "none";
		},
		closeDropdowns(el) {
			if (el.target.className.includes("dropdown") && !el.target.className.includes("dropdown__link")) return;
			this.dropdowns.forEach(dropdown => {
				if (dropdown.style.display === "none") return;
				dropdown.style.display = "none";
			});
		}
	},
	mounted() {
		var dropdowns = document.getElementsByClassName("dropdown__list");
		[...dropdowns].forEach(dropdown => {
			this.dropdowns.push(dropdown);
		});
	},
	created: function () {
		document.body.addEventListener('click', this.closeDropdowns);
	},
	destroyed: function () {
		document.body.removeEventListener('click', this.closeDropdowns);
	}
};
</script>

<style lang="scss">
@import "@/assets/styles/utility/vars.scss";

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

.dropdown {
	position: relative;

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