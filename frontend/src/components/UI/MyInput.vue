<template>
  <input :value="modelValue" @input="debounce(() => updateInput($event))" type="text"
    class="form-control bg-dark border-primary" />
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
  setup() {
    function createDebounce() {
      let timeout = null;
      return function (fnc, delayMs) {
        clearTimeout(timeout);
        timeout = setTimeout(() => {
          fnc();
        }, delayMs || 250);
      };
    }

    return {
      debounce: createDebounce(),
    };
  }
};
</script>

<style>
</style>