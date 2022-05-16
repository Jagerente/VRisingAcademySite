<template>
  <div class="d-flex mx-5 my-4 h-75">
    <div class="w-50 d-flex flex-column" style="margin-right: 15px">
      <!-- SEARCH -->
      <!-- <div class="d-flex">
        <card-search />
      </div> -->
      <!-- OTHER -->
      <!-- Why 90%? -->
      <div class="d-flex justify-content-between h-100">
        <!-- FILTER -->
        <div class="" style="margin-right: 15px">
          <card-filter />
        </div>
        <!-- CATALOGUE -->
        <div class="flex-fill">
          <card-catalogue />
        </div>
      </div>
    </div>
    <!-- RIGHT -->
    <div class="w-50">
      <card-information />
    </div>
  </div>
</template>

<script>
import CardFilter from "@/components/ItemsPage/CardFilter";
import CardCatalogue from "@/components/ItemsPage/CardCatalogue";
import CardInformation from "@/components/ItemsPage/CardInformation";

export default {
  components: {
    CardFilter,
    CardCatalogue,
    CardInformation,
  },
};
</script>

<style>
.card__background {
  border-radius: 10px;
  background-color: #282737;
  /* height: 100%; */
}

.card__header {
  background-color: #14141e;
  border-radius: 10px;
  padding-left: 15px;
  padding-right: 15px;
  /* margin-left: -5px;
  margin-right: -5px; */
  font-size: 26px;
  text-align: center;
}

.card__content {
  padding-top: 10px;
  padding-left: 10px;
  padding-right: 10px;
}
</style>