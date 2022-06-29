<template>
  <div class="card">
    <div class="card__header">
      <div class="header__icon">
        <slot name="left"></slot>
      </div>
      <slot
        v-if="custom"
        name="header"
      ></slot>
      <p
        v-else
        class="header__label"
      >
        {{ title }}
      </p>
      <slot name="right"></slot>
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
@import '@/assets/styles/utility/vars.scss';


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