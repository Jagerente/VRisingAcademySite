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
    <div
      class="catalogue__top-left"
      :class="{ 'h-100': !right }"
    >
      <!-- LEFT -->
      <my-card
        class="top-left__left"
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
      <div class="top-left__right">
        <slot name="center"></slot>
      </div>
    </div>
    <!-- RIGHT TOP -->
    <div
      v-if="right"
      class="catalogue__top-right"
      style="height: auto; width: 800px"
    >
      <slot name="right"></slot>
    </div>
    <!-- RIGHT TO BOTTOM -->
    <div
      v-if="right"
      class="catalogue__bottom"
    >
      <slot name="right"></slot>
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
  },
  data() {
    return {
      selectedTab: 0
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

  &__top-left {
    display: flex;
    width: 100%;
    min-height: 220px;
    margin-bottom: $m1;

    .top-left__left {
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

    .top-left__right {
      flex: 1 1 auto;
      margin-left: $m1;

      @media (min-width: $lg) {
        margin-right: $m1;
      }
    }
  }

  &__top-right {
    display: none;

    @media (min-width: $lg) {
      margin-bottom: $m1;
      display: flex;
    }
  }

  &__bottom {
    flex: 1 1 auto;

    @media (min-width: $lg) {
      display: none;
    }
  }
}
</style>