<template>
  <div class="card">
    <div class="header card__header">
      <slot
        v-if="this.custom"
        name="header"
      ></slot>
      <div
        v-if="!custom"
        class="header__icon"
      >
        <slot name="left"></slot>
      </div>
      <p
        v-if="!custom"
        class="header__label"
      >
        {{ title }}
      </p>
      <slot
        v-if="!custom"
        name="right"
      ></slot>
    </div>
    <div class="card__content">
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
.card {
  display: flex;
  flex-direction: column;
  height: 100%;
  border-radius: 10px;
  background-color: $background-light;
}

.card__header {
  display: flex;
  justify-content: center;
  height: 45px;
  font-size: 26px;
  background-color: #14141e;
  border-radius: 10px;
  padding-left: $m1;
  padding-right: $m1;
  user-select: none;

  .header__icon {
    margin-top: auto;
    margin-bottom: auto;

    &:first-child {
      margin-right: 5px;
    }

    &:last-child {
      margin-left: 5px;
    }
  }

  .header__label {
    margin-top: auto;
    margin-bottom: auto;
  }
}

.card__content {
  height: 100%;
  padding: 10px;
  overflow-y: auto;
}
</style>