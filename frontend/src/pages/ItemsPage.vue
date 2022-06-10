<template>
  <MyLayout :isLoading="this.isItemsLoading">
    <template #left>
      <card-filter />
    </template>
    <template #center>
      <card-catalogue />
    </template>
    <template #right>
      <card-information />
    </template>
  </MyLayout>
</template>

<script>
import MyLayout from "@/components/UI/MyLayout";
import CardFilter from "@/components/ItemsPage/CardFilter";
import CardCatalogue from "@/components/ItemsPage/CardCatalogue";
import CardInformation from "@/components/ItemsPage/CardInformation";

import { mapState, mapActions } from "vuex"

export default {
  components: {
    MyLayout,
    CardFilter,
    CardCatalogue,
    CardInformation,
  },
  methods: {
    ...mapActions({
      fetchItems: "items/fetchItems",
    })
  },
  computed: {
    ...mapState({
      isItemsLoading: (state) => state.items.isItemsLoading,
    })
  },
  mounted() {
    this.fetchItems();
  }
};
</script>

<style scoped lang="scss">
</style>