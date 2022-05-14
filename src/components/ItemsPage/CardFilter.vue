<template>
  <my-card title="Filter">
    <div
      class="nav d-flex flex-column nav-pills"
      id="v-pills-tab"
      role="tablist"
      aria-orientation="vertical"
    >
      <button
        class="nav-link active"
        id="v-pills-weapons-tab"
        data-bs-toggle="pill"
        data-bs-target="#v-pills-weapons"
        type="button"
        role="tab"
        aria-controls="v-pills-weapons"
        aria-selected="true"
      >
        Weapons
      </button>
      <button
        class="nav-link"
        id="v-pills-armour-tab"
        data-bs-toggle="pill"
        data-bs-target="#v-pills-armour"
        type="button"
        role="tab"
        aria-controls="v-pills-armour"
        aria-selected="false"
      >
        Armour
      </button>
      <button
        class="nav-link"
        id="v-pills-consumables-tab"
        data-bs-toggle="pill"
        data-bs-target="#v-pills-consumables"
        type="button"
        role="tab"
        aria-controls="v-pills-consumables"
        aria-selected="false"
      >
        Consumables
      </button>
      <button
        class="nav-link"
        id="v-pills-reagents-tab"
        data-bs-toggle="pill"
        data-bs-target="#v-pills-reagents"
        type="button"
        role="tab"
        aria-controls="v-pills-reagents"
        aria-selected="false"
      >
        Reagents
      </button>
    </div>
  </my-card>
</template>

<script>
</script>

<style scoped>
.nav-pills .nav-link.active,
.nav-pills .show > .nav-link {
  background: url("@/assets/images/ui/circle.png"), rgba(0, 0, 0, 0.5);
  background-size: 110%;
  background-position: 50%;
  background-repeat: no-repeat;
  border: 1px solid white;
}

.nav-pills .nav-link {
  margin-bottom: 5px;
  height: 55px;
  background-color: #14141e;
}
</style>