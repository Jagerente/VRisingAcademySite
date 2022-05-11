<template>
  <input
    :value="modelValue"
    @input="updateInput"
    type="text"
    class="form-control bg-dark border-primary"
  />
</template>

<script>
export default {
  name: "my-input",
  props: {
    modelValue: [String, Number],
  },
  methods: {
    updateInput(event) {
      this.$emit("update:modelValue", event.target.value);
    },
  },
};
</script>

<style>
</style>