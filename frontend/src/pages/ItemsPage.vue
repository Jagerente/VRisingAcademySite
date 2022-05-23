<template>
  <div class="catalogue d-flex flex-column flex-lg-row h-100 mx-2">
    <!-- LEFT -->
    <div class="w-100 d-flex mb-2 mb-lg-0 filter__card">
      <!-- FILTER -->
      <div>
        <card-filter />
      </div>
      <!-- CATALOGUE -->
      <div class="flex-fill mx-0 ms-2 mx-lg-3">
        <card-catalogue />
      </div>
    </div>
    <!-- RIGHT -->
    <!-- INFORMATION -->
    <div class="p-0 m-0 d-lg-flex d-none" style="height: auto;">
      <card-information style="width: 500px;" />
    </div>
    <!-- BOTTOM -->
    <!-- INFORMATION -->
    <div class="p-0 m-0 d-flex d-lg-none">
      <card-information style="width: 100%;" />
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

<style scoped lang="scss">
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';

@include media-breakpoint-down(sm) {
  .filter__card {
    min-height: 220px;
  }

  .catalogue {
    padding-top: 100px;
    padding-bottom: 15px;
  }
}

@include media-breakpoint-up(sm) {
  .filter__card {
    min-height: 300px;
  }

  .catalogue {
    padding-top: 125px;
    padding-bottom: 50px;
  }
}
</style>