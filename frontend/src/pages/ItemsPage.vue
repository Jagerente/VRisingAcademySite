<template>
  <MyLayout :isLoading="this.isItemsLoading" :tabs="this.types" :selector="this.selectType">
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
import CardCatalogue from "@/components/ItemsPage/CardCatalogue";
import CardInformation from "@/components/ItemsPage/CardInformation";

import { mapState, mapActions } from "vuex"

export default {
  components: {
    MyLayout,
    CardCatalogue,
    CardInformation,
  },
  methods: {
    ...mapActions({
      fetchItems: "items/fetchItems",
      selectType: "items/selectType",
    })
  },
  computed: {
    ...mapState({
      isItemsLoading: (state) => state.items.isItemsLoading,
      types: (state) => state.items.types,
    })
  },
  mounted() {
    this.fetchItems();
  }
};
</script>

<style scoped lang="scss">
</style>