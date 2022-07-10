<template>
  <!-- Loader -->
  <div
    v-if="this.isLoading"
    class="loader"
  >
    <span class="loader__text">Loading...</span>
  </div>
  <div
    v-else
    class="catalogue"
  >
    <!-- TOP -->
    <div class="catalogue__wrapper">
      <!-- LEFT -->
      <my-card
        class="catalogue__filters"
        title="Filter"
        :custom="false"
      >
        <template #left>
          <icon-filter />
        </template>
        <div class="tabs">
          <button
            v-for="(tab, i) in this.tabs"
            @click="onTabSelect(i)"
            class="tabs__button button"
            :class="{ 'active': i === this.selectedTab }"
          >
            <img
              v-if="this.tabLogo"
              class="button__image"
              :src="getImageUrl('blood_types/' + tab.name.toLowerCase() + '.webp')"
            />
            <div class="button__label">
              {{ tab.name }}
            </div>
          </button>
        </div>
      </my-card>
      <!-- RIGHT -->
      <div class="catalogue__content">
        <slot name="center"></slot>
      </div>
    </div>
    <!-- BOTTOM -->
    <div
      v-if="right"
      class="catalogue__aside"
    >
      <slot name="right"></slot>
    </div>
		<div
			v-if="this.showModal && windowWidth < 992 && this.right"
			class="modal"
			@click.self="this.updateShowModal(false)"
		>
			<div
				v-show="this.showModal"
				v-drag="{ axis: 'y', handle: '#drag' }"
				@v-drag-end="onDragEnd"
				ref="modal"
				class="modal__content"
			>
				<div
					id="drag"
					class="modal__drag"
				>
				</div>
				<slot name="right"></slot>
			</div>
		</div>
  </div>
</template>

<script>
import { useWindowSize } from 'vue-window-size';
import { ref, onMounted } from "vue";

export default {
  props: {
    isLoading: Boolean,
    tabs: Array,
    selector: Function,
    left: {
      type: Boolean,
      default: true,
    },
    tabLogo: {
      type: Boolean,
      default: false,
    },
    center: {
      type: Boolean,
      default: true,
    },
    right: {
      type: Boolean,
      default: true,
    },
    showModal: {
      type: Boolean,
      default: false,
    },
    updateShowModal: Function
  },
  methods: {
    onDragEnd(event) {
      const modal = this.$refs.modal;
      const windowHeight = window.innerHeight;

      if (!modal) return;

      if (windowHeight - modal.offsetHeight - modal.offsetTop > 0) {
        modal.style.top = '0px';
      }
      else if (modal.offsetTop * 100 / windowHeight > 0 && windowHeight - modal.offsetHeight - modal.offsetTop < -50) {
        this.updateShowModal(false);
      }
    }
  },
  setup(props) {
    const selectedTab = ref(0);

    const { width, height } = useWindowSize();
    const windowWidth = width;
    const windowHeight = height;
    const getImageUrl = (name) => {
      return `images/${name}`;
    };

    function onTabSelect(i) {
      props.selector(i);
      selectedTab.value = i;
    }

    return { windowHeight, windowWidth, selectedTab, onTabSelect, getImageUrl };
  }
};
</script>

<style lang="scss" scoped>
.loader {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);

  &__text {
    font-size: 2rem;
  }
}

.catalogue {
  padding-top: 15px;
  padding-left: 10px;
  padding-right: 10px;
  height: calc(100vh - ($header-height + $footer-height) + 5px);
  min-height: calc(100vh - ($header-height + 25px));
  display: flex;
  flex-direction: column;


  @media (min-width: $lg) {
    flex-direction: row;
  }

  &__wrapper {
    display: flex;
    width: 100%;
    min-height: 220px;
    margin-bottom: $m1;

    .catalogue__filters {
      .tabs {
        display: flex;
        flex-direction: column;

        .tabs__button {
          display: flex;
          background-color: $dark;
          border: 1px solid $dark;
          color: $text-color;
          text-transform: capitalize;
          width: 100%;
          height: 40px;
          transition-property: color, border;
          transition-duration: .25s;
          transition-timing-function: ease;

          &:not(:last-child) {
            margin-bottom: 5px;
          }

          &.active {
            background: url("@/assets/images/ui/circle.webp"), $dark;
            background-size: 110%;
            background-position: 50%;
            background-repeat: no-repeat;
            border: 1px solid white;
            color: white;
          }

          .button {
            &__image {
              width: 24px;
              margin-top: auto;
              margin-bottom: auto;
            }

            &__label {
              pointer-events: none;
              display: flex;
              margin-top: auto;
              margin-bottom: auto;
              margin-left: auto;
              margin-right: auto;
            }
          }
        }
      }
    }

    .catalogue__content {
      flex: 1 1 auto;
      margin-left: $m1;

      @media (min-width: $lg) {
        margin-right: $m1;
      }
    }
  }

  &__aside {
    flex: 1 1 auto;
    width: 100%;
    display: none;

    @media (min-width: $lg) {
      margin-bottom: $m1;
      display: flex;
      height: auto;
      width: 800px;
    }
  }

  &__bottom {
    flex: 1 1 auto;

    @media (min-width: $lg) {
      display: none;
    }
  }
}

.modal {
  display: flex;
  // display: none;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  // background: rgba(0, 0, 0, 0.4);
  opacity: 1;
  align-items: flex-end;
  justify-content: center;
	background-color: rgba(0,0,0,0.6);
	animation: fade 0.25s ease-in;
	
	&__content {
		animation: translate-up 0.3s ease-in;
	}
	
	&__drag {
		position: absolute;
		top: -35px;
		left: 50%;
		transform: translateX(-50%);
		padding-top: 20px;
		padding-bottom: 55px;
		padding-left: 100%;
		padding-right: 100%;
		z-index: 11;
		
		&::after {
			content: '';
			display: block;
			height: 2px;
			width: 100px;
			background-color: #ffffff;
			border-radius: 5px;
		}
	}
	
	@keyframes fade {
		from {
			opacity: 0;
		}
		to {
			opacity: 1
		}
	}
	
	@keyframes translate-up {
		from {
			transform: translateY(100%);
		}
		to {
			transform: translateY(0%);
		}
	}

  @media (max-width: $lg) {
    &__content {
			position: relative;
      z-index: 10;
      background-color: $background;
      border-radius: 20px;
      // max-height: 75%;
      border-bottom-left-radius: 0;
      border-bottom-right-radius: 0;
      width: 100%;
			height: 75%;
    }
  }
	
}
</style>