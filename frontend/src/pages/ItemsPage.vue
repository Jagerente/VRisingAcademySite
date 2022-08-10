<template>
  <MyLayout
    :isLoading="this.isItemsLoading"
    :tabs="this.types"
    :selector="this.selectType"
    :showModal="this.showModal"
    :updateShowModal="this.updateShowModal"
  >
    <template #center>
      <card-catalogue />
    </template>
    <template #right>
      <card-information />
    </template>
  </MyLayout>
</template>

<script>
import MyLayout from "@/components/UI/MyLayout.vue";
import CardCatalogue from "@/components/ItemsPage/CardCatalogue.vue";
import CardInformation from "@/components/ItemsPage/CardInformation.vue";

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
      updateShowModal: "items/updateShowModal",
    })
  },
  computed: {
    ...mapState({
      isItemsLoading: (state) => state.items.isItemsLoading,
      types: (state) => state.items.types,
      showModal: (state) => state.items.showModal,
    })
  },
  mounted() {
    this.fetchItems();
  }
};
</script>