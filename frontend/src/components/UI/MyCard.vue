<template>
  <div class="d-flex flex-column h-100 card__background">
    <div class="d-flex card__header justify-content-center">
      <slot name="left"></slot>
      <slot v-if="custom" name="header"></slot>
      <p v-else class="mx-2 my-auto">
        {{ title }}
      </p>
      <slot name="right"></slot>
    </div>
    <div class="card__content h-100">
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: "my-card",
  props: {
    title: String,
    custom: Boolean,
  },
};
</script>

<style scoped lang="scss">
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';

.card__background {
  border-radius: 10px;
  background-color: #282737;
}

@include media-breakpoint-down(sm) {
  $header-size: 18px;

  .card__header {
    font-size: $header-size;
    height: 39px;
  }
}

@include media-breakpoint-up(sm) {
  $header-size: 26px;

  .card__header {
    font-size: $header-size;
  }
}

.card__header {
  background-color: #14141e;
  border-radius: 10px;
  padding-left: 15px;
  padding-right: 15px;
  text-align: center;
  user-select: none;
}

.card__content {
  padding: 10px;
  overflow-y: auto;
}
</style>