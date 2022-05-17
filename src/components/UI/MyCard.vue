<template>
  <div class="d-flex flex-column h-100 card__background">
    <div class="card__header">
      <slot name="left"></slot>
      <slot v-if="custom" name="header"></slot>
      <span v-else class="mx-2">
        {{ title }}
      </span>
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

<style scoped>
.card__background {
  border-radius: 10px;
  background-color: #282737;
  /* height: 100%; */
}

.card__header {
  background-color: #14141e;
  border-radius: 10px;
  padding-left: 15px;
  padding-right: 15px;
  /* margin-left: -5px;
  margin-right: -5px; */
  font-size: 26px;
  text-align: center;
  user-select: none;
}

.card__content {
  padding: 10px;
  overflow-y: auto;
}
</style>