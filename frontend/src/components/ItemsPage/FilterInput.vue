<template>
  <div class="input">
    <div class="input__icon">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="16"
        height="16"
        fill="currentColor"
        viewBox="0 0 16 16"
      >
        <path
          d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001c.03.04.062.078.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1.007 1.007 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0z"
        >
        </path>
      </svg>
    </div>
    <my-input
      v-if="this.searchType !== 2"
      @click="updateSearchQuery('');"
      :model-value="searchQuery"
      @update:model-value="setSearchQuery"
      type="text"
      class="input__field"
      placeholder="Search..."
    />
    <button
      v-else
      class="input__tag"
      @click="updateSearchQuery('');"
    >
      {{ this.searchQuery }}
    </button>
    <button
      v-if="this.searchQuery && this.searchType !== 2"
      class="input__clear"
      @click="updateSearchQuery('');"
    >
    </button>
  </div>
</template>

<script>
import { mapState, mapMutations, mapActions } from "vuex";

export default {
  methods: {
    ...mapActions({
      updateSearchQuery: "items/updateSearchQuery",
    }),
    ...mapMutations({
      setSearchQuery: "items/setSearchQuery",
    }),
  },
  computed: {
    ...mapState({
      searchQuery: (state) => state.items.searchQuery,
      searchType: (state) => state.items.searchType,
    }),
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/styles/utility/vars.scss";
$cross_size: 16px;
$cross_weight: 1.5px;
$cross_size-tag: $cross_size * .63;
$cross_weight-tag: 1.25px;

.header__input {
  display: flex;
  align-items: center;
  width: 100%;
}

.input {
  padding: 0 15px;

  &__icon {
    margin-top: auto;
    margin-bottom: auto;
    margin-right: 15px;
  }

  &__field {
    background-color: $dark;
    border: none;
    color: white;
    width: 100%;

    &:focus {
      outline: 0;
      outline-offset: 0;
    }
  }

  &__clear {
    box-sizing: border-box;
    position: relative;
    background: transparent;
    border: none;
    height: 100%;
    width: $cross_size;
    padding: 0;

    &::after {
      content: "";
      position: absolute;
      height: $cross_size;
      width: $cross_weight;
      top: 50%;
      background-color: white;
      transform: translateY(-50%) rotate(-45deg);
    }

    &::before {
      content: "";
      position: absolute;
      height: $cross_size;
      width: $cross_weight;
      top: 50%;
      background-color: white;
      transform: translateY(-50%) rotate(45deg);
    }
  }

  &__tag {

    position: relative;
    font-size: 0.8rem;
    background: $primary;
    border-radius: 15px;
    text-transform: capitalize;
    border: none;
    color: white;
    transition: box-shadow 0.1s ease-in-out;
    padding: 4px 8px;
    padding-right: $cross_size-tag*2;

    &:hover {
      box-shadow: 0 0 8px black;
    }


    &::after {
      content: "";
      position: absolute;
      height: $cross_size-tag;
      width: $cross_weight-tag;
      top: 50%;
      right: $cross_size-tag;
      background-color: white;
      transform: translateY(-50%) rotate(-45deg);
    }

    &::before {
      content: "";
      position: absolute;
      height: $cross_size-tag;
      width: $cross_weight-tag;
      top: 50%;
      right: $cross_size-tag;
      background-color: white;
      transform: translateY(-50%) rotate(45deg);
    }
  }
}
</style>