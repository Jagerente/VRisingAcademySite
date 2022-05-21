<template>
  <div class="input-group">
    <span class="input-group-text bg-dark border-0">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-search"
        viewBox="0 0 16 16">
        <path
          d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001c.03.04.062.078.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1.007 1.007 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0z">
        </path>
      </svg>
    </span>
    <my-input :model-value="searchQuery" @update:model-value="setSearchQuery" type="text"
      class="form-control bg-dark border-0" placeholder="Search..." focus />
    <button v-if="searchQuery" class="clear p-0 m-0 mx-2 test" @click="updateSearchQuery('')">
      <svg xmlns="http://www.w3.org/2000/svg" style="margin-top: -5px;" width="20" height="20" fill="#a8a9ae"
        viewBox="0 0 16 16">
        <path
          d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z" />
      </svg>
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
    }),
  },
};
</script>

<style scoped lang="scss">
.clear {
  background: #14141e;
  border: none;
}

.input-group {
  height: 39px;
}
</style>