<template>
  <!-- Loader -->
  <div v-if="this.isLoading">
    <div
      class="position-absolute top-50 start-50 translate-middle"
      role="status"
    >
      <span class="h1">Loading...</span>
    </div>
  </div>
  <div
    v-else
    class="catalogue"
  >
    <!-- LEFT TOP -->
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
            @click="this.selector(i); this.selectedTab = i;"
            class="tabs__button button"
            :class="{ 'active': i === this.selectedTab }"
          >
            <img
              v-if="this.tabLogo"
              class="button__image"
              :src="require('@/assets/images/blood_types/' + tab.name.toLowerCase() + '.webp')"
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
    <!-- RIGHT TOP -->
    <div
      v-if="right"
      class="catalogue__aside"
    >
      <slot name="right"></slot>
    </div>
    <div
      class="modal"
      :class="{ 'active': this.showModal }"
      v-if="this.showModal"
    >
      <div
        class="modal__overlay"
        @click.self="this.updateShowModal(false)"
      >
        <div class="modal__content">
          <slot name="right"></slot>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
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
  data() {
    return {
      selectedTab: 0,
    }
  },
}
</script>

<style lang="scss" scoped>
@import '@/assets/styles/utility/vars.scss';

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
  //display: none;
  opacity: 0;
  transition: opacity .2s ease;

  &__overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.4);
    display: flex;
    align-items: center;
    justify-content: center;
  }

  &__content {
    border-radius: 20px;
    background-color: $background;
    // padding: 15px 5px 0 5px;
    translate: 0 100%;
    transition: translate .2s ease, display .2s ease;
  }

  @media (max-width: $lg) {

      display: flex;
      max-height: 50%;
      opacity: 1;

      &__overlay {
        display: flex;
        align-items: flex-end;
      }

      &__content {
        height: 75%;
        translate: 0 0%;
        border-bottom-left-radius: 0;
        border-bottom-right-radius: 0;
        width: 100%;
      }
  }
}
</style>