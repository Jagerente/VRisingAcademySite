<template>
  <div class="d-flex flex-column h-100 card__background">
    <div class="card__header">
      <span>
        {{ title }}
      </span>
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
}

.card__content {
  padding: 10px;
  overflow-y: auto;
}
</style>