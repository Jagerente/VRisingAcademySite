<template>
  <li class="nav-item">
    <slot />
    <img class="selector" src="@/assets/images/ui/blood_drop.png" width="10" />
  </li>
</template>

<script>
export default {
  name: "blood-selector",
};
</script>

<style scoped>
.selector {
  display: block;
  margin-left: auto;
  margin-right: auto;
}

.nav-pills {
  background-color: 0;
}
</style>