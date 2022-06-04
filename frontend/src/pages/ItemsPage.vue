<template>
  <div class="catalogue d-flex flex-column flex-lg-row mx-2">
    <!-- LEFT -->
    <div class="w-100 d-flex mb-2 mb-lg-0 filter__card">
      <!-- FILTER -->
      <div>
        <card-filter />
      </div>
      <!-- CATALOGUE -->
      <div class="flex-fill mx-0 ms-1 mx-lg-2">
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
    <div class="p-0 m-0 d-flex d-lg-none flex-fill">
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
@import '@/assets/styles/va_variables.scss';

@include media-breakpoint-down(sm) {
  .filter__card {
    min-height: 220px;
  }
}

@include media-breakpoint-up(sm) {
  .filter__card {
    min-height: 300px;
  }
}

.catalogue {
  padding-top: 15px;
}

.catalogue {
  height: calc(100vh - ($header-height + $footer-height) + 5px);
  min-height: calc(100vh - ($header-height + 25px));
}
</style>