<template>
  <my-card title="Filter" :custom="false">
    <template #left>
      <icon-filter />
    </template>
    <div class="nav d-flex flex-column nav-pills" id="v-pills-tab" role="tablist" aria-orientation="vertical">
      <input class="nav-link active" id="v-pills-weapons-tab" data-bs-toggle="pill" data-bs-target="#v-pills-weapons"
        type="button" role="tab" aria-controls="v-pills-weapons" aria-selected="true" value="Weapons"
        style="cursor: pointer" />
      <button class="nav-link" id="v-pills-armour-tab" data-bs-toggle="pill" data-bs-target="#v-pills-armour"
        type="button" role="tab" aria-controls="v-pills-armour" aria-selected="false">
        Armour
      </button>
      <button class="nav-link" id="v-pills-consumables-tab" data-bs-toggle="pill" data-bs-target="#v-pills-consumables"
        type="button" role="tab" aria-controls="v-pills-consumables" aria-selected="false">
        Consumables
      </button>
      <button class="nav-link" id="v-pills-reagents-tab" data-bs-toggle="pill" data-bs-target="#v-pills-reagents"
        type="button" role="tab" aria-controls="v-pills-reagents" aria-selected="false">
        Reagents
      </button>
    </div>
  </my-card>
</template>

<script>
</script>

<style scoped lang="scss">
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';

@include media-breakpoint-down(sm) {
  $header-size: 13px;

  .nav-link {
    font-size: $header-size;
    width: 100px;
    height: 35px;

  }
}

@include media-breakpoint-up(sm) {
  .nav-link {
    height: 55px;
  }
}

.nav-pills .nav-link.active,
.nav-pills .show>.nav-link {
  background: url("@/assets/images/ui/circle.webp"), rgba(0, 0, 0, 0.5);
  background-size: 110%;
  background-position: 50%;
  background-repeat: no-repeat;
  border: 1px solid white;
}

.nav-pills .nav-link {
  margin-bottom: 5px;
  background-color: #14141e;
  border: 1px solid #14141e;
}
</style>