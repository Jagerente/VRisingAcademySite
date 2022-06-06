<template>
  <div class="d-flex flex-column h-100">
    <my-navbar />
    <router-view />
    <my-footer v-if="false" />
  </div>
</template>

<script>
import MyNavbar from "@/components/MyNavbar.vue";
import MyFooter from "./components/MyFooter.vue";
import { mapGetters, mapActions } from "vuex";

export default {
  components: {
    MyNavbar,
    MyFooter
  },
  methods: {
    ...mapActions({
      fetchItems: "items/fetchItems",
      getSchools: "spells/getSchools",
      getTypes: "spells/getTypes",
      getSpells: "spells/getSpells",
    })
  },
  mounted() {
    this.fetchItems();
    this.getSchools();
    this.getTypes();
    this.getSpells();
  }
};
</script>

<style>
</style>