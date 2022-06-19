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
    class="catalogue d-flex flex-column flex-lg-row mx-2"
  >
    <!-- LEFT TOP -->
    <div
      class="w-100 d-flex mb-2 mb-lg-0 filter__card"
      :class="!right ? 'h-100' : ''"
    >
      <!-- LEFT -->
      <div>
        <my-card
          title="Filter"
          :custom="false"
        >
          <template #left>
            <icon-filter />
          </template>
          <div
            class="nav d-flex flex-column"
            id="v-pills-tab"
            role="tablist"
            aria-orientation="vertical"
          >
            <button
              v-for="(tab, i) in this.tabs"
              @click="this.selector(i);"
              class="tab d-flex"
              :class="i == 0 ? 'active' : ''"
              data-bs-toggle="pill"
              :data-bs-target="`#v-pills-${tab.name.toLowerCase()}`"
              type="button"
              role="tab"
              :aria-controls="`v-pills-${tab.name.toLowerCase()}`"
              aria-selected="false"
            >
              <div
                class="d-flex my-auto"
                v-if="this.tabLogo"
              >
                <img
                  style="width:24px;"
                  :src="require('@/assets/images/blood_types/' + tab.name.toLowerCase() + '.webp')"
                />
              </div>
              <div class="d-flex my-auto mx-auto">
                {{ tab.name }}
              </div>
            </button>
          </div>
        </my-card>
        <!-- <slot name="left"></slot> -->
      </div>
      <!-- RIGHT -->
      <div class="flex-fill mx-0 ms-1 mx-lg-2">
        <slot name="center"></slot>
      </div>
    </div>
    <!-- RIGHT TOP -->
    <div
      v-if="right"
      class="p-0 m-0 d-lg-flex d-none"
      style="height: auto; width: 800px"
    >
      <slot name="right"></slot>
    </div>
    <!-- RIGHT TO BOTTOM -->
    <div
      v-if="right"
      class="p-0 m-0 d-flex d-lg-none flex-fill"
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
  }
}
</script>

<!--<style lang="scss">-->
<!--@import 'bootstrap/scss/_functions.scss';-->
<!--@import 'bootstrap/scss/_variables.scss';-->
<!--@import 'bootstrap/scss/_mixins.scss';-->
<!--@import '@/assets/styles/va_variables.scss';-->

<!--@include media-breakpoint-down(sm) {-->
<!--  $header-size: 13px;-->

<!--  .filter__card {-->
<!--    min-height: 220px;-->
<!--  }-->

<!--  .tab {-->
<!--    font-size: $header-size;-->
<!--    width: 100px;-->
<!--    height: 35px;-->
<!--  }-->

<!--  .logo {-->
<!--    $size: 12px;-->
<!--    width: $size;-->
<!--    height: $size;-->
<!--  }-->
<!--}-->

<!--@include media-breakpoint-up(sm) {-->
<!--  .filter__card {-->
<!--    min-height: 300px;-->
<!--  }-->

<!--  .tab {-->
<!--    height: 55px;-->
<!--  }-->

<!--  .logo {-->
<!--    $size: 12px;-->
<!--    width: size;-->
<!--    height: size;-->
<!--  }-->
<!--}-->

<!--.catalogue {-->
<!--  padding-top: 15px;-->
<!--}-->

<!--.catalogue {-->
<!--  height: calc(100vh - ($header-height + $footer-height) + 5px);-->
<!--  min-height: calc(100vh - ($header-height + 25px));-->
<!--}-->

<!--.tab {-->
<!--  margin-bottom: 5px;-->
<!--  background-color: #14141e;-->
<!--  border: 1px solid #14141e;-->
<!--  color: $text-color;-->
<!--  text-transform: capitalize;-->
<!--}-->

<!--.tab.active {-->
<!--  background: url("@/assets/images/ui/circle.webp"), rgba(0, 0, 0, 0.5);-->
<!--  background-size: 110%;-->
<!--  background-position: 50%;-->
<!--  background-repeat: no-repeat;-->
<!--  border: 1px solid white;-->
<!--  color: white;-->
<!--}-->
<!--</style>-->