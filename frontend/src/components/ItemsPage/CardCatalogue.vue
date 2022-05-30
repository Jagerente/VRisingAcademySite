<template>
  <my-card :custom="true">
    <template #header>
      <filter-input></filter-input>
    </template>
    <div class="tab-content d-flex justify-content-center flex-wrap" id="v-pills-tabContent">
      <div class="tab-pane fade show active" id="v-pills-weapons" role="tabpanel" aria-labelledby="v-pills-weapons-tab">
        <items-list :type="1"></items-list>
      </div>
      <div class="tab-pane fade" id="v-pills-armour" role="tabpanel" aria-labelledby="v-pills-armour-tab">
        <items-list :type="2"></items-list>
      </div>
      <div class="tab-pane fade" id="v-pills-consumables" role="tabpanel" aria-labelledby="v-pills-consumables-tab">
        <items-list :type="3"></items-list>
      </div>
      <div class="tab-pane fade" id="v-pills-reagents" role="tabpanel" aria-labelledby="v-pills-reagents-tab">
        <items-list :type="4"></items-list>
      </div>
    </div>
    <div style="height: 15px"></div>
  </my-card>
</template>

<script>
import FilterInput from "@/components/ItemsPage/FilterInput.vue";
import ItemsList from "@/components/ItemsPage/ItemsList.vue";

export default {
  components: { FilterInput, ItemsList },
};
</script>

<style scoped>
</style>