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
    <Transition
      v-if="right"
      name="translate-up"
    >
      <div
        v-if="this.showModal && windowWidth < 992"
        class="modal"
        @click.self="this.updateShowModal(false)"
      >
        <div
          v-if="this.showModal"
          v-drag="{ axis: 'y' }"
          @v-drag-end="onDragEnd"
          ref="modal"
          class="modal__content"
        >
          <slot name="right"></slot>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { useWindowSize } from 'vue-window-size';
import { ref, onMounted } from "vue";

const props = defineProps({
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
});
const selectedTab = ref(0);

const { width, height } = useWindowSize();
const windowWidth = width;
const windowHeight = height;

const modal = ref(null);

const getImageUrl = (name) => {
  return `images/${name}`;
};

function onTabSelect(i) {
  props.selector(i);
  selectedTab.value = i;
}

function onDragEnd(event) {
  if (!modal.value) return;

  if (windowHeight.value - modal.value.offsetHeight - modal.value.offsetTop > 0) {
    modal.value.style.top = '0px';
  }
  else if (modal.value.offsetTop * 100 / windowHeight.value > 0 && windowHeight.value - modal.value.offsetHeight - modal.value.offsetTop < -50) {
    props.updateShowModal(false);
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

  @media (max-width: $lg) {
    &__content {
      z-index: 10;
      background-color: $background;
      border-radius: 20px;
      // max-height: 75%;
      border-bottom-left-radius: 0;
      border-bottom-right-radius: 0;
      width: 100%;
    }
  }
}
</style>