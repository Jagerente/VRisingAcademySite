<template>
  <!-- Loader -->
  <div v-if="this.isLoading">
    <div class="position-absolute top-50 start-50 translate-middle" role="status">
      <span class="h1">Loading...</span>
    </div>
  </div>
  <div v-else class="catalogue d-flex flex-column flex-lg-row mx-2">
    <!-- LEFT TOP -->
    <div class="w-100 d-flex mb-2 mb-lg-0 filter__card">
      <!-- LEFT -->
      <div>
        <slot name="left"></slot>
      </div>
      <!-- RIGHT -->
      <div class="flex-fill mx-0 ms-1 mx-lg-2">
        <slot name="center"></slot>
      </div>
    </div>
    <!-- RIGHT TOP -->
    <div class="p-0 m-0 d-lg-flex d-none" style="height: auto; width: 800px">
      <slot name="right"></slot>
    </div>
    <!-- RIGHT TO BOTTOM -->
    <div class="p-0 m-0 d-flex d-lg-none flex-fill">
      <slot name="right"></slot>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    isLoading: Boolean,
  }
}
</script>

<style scoped lang="scss">
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';
@import '@/assets/styles/va_variables.scss';

@include media-breakpoint-down(sm) {
  .filter__card {
    min-height: 220px;
  }
}

@include media-breakpoint-up(sm) {
  .filter__card {
    min-height: 300px;
  }
}

.catalogue {
  padding-top: 15px;
}

.catalogue {
  height: calc(100vh - ($header-height + $footer-height) + 5px);
  min-height: calc(100vh - ($header-height + 25px));
}
</style>