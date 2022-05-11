<template>
  <div class="content d-flex justify-content-between test">
    <div class="flex-fill flex-column test" style="width: 60%">
      <card-search v-model="searchQuery" />
      <div class="d-flex mt-auto justify-content-between">
        <div class="test">
          <card-filter />
        </div>
        <div class="flex-fill test">
          <card-catalogue :searchQuery=searchQuery />
        </div>
      </div>
    </div>
    <div class="flex-fill test" style="width: 40%">
      <card-information />
    </div>
  </div>
</template>

<script>
import CardSearch from "@/components/BestiaryPage/CardSearch";
import CardFilter from "@/components/BestiaryPage/CardFilter";
import CardCatalogue from "@/components/BestiaryPage/CardCatalogue";
import CardInformation from "@/components/BestiaryPage/CardInformation";

export default {
  data() {
    return {
      searchQuery: '',
    };
  },
  components: {
    CardSearch,
    CardFilter,
    CardCatalogue,
    CardInformation,
  },
};
</script>

<style>
.content {
  margin-top: 25px;
  margin-left: auto;
  margin-right: auto;
}

.menu {
  height: 100%;
  width: 100%;
}

.row {
  padding-top: 5px;
}

.text-vertical {
  transform: rotate(-90deg);
}

/* .card {
  padding-top: 10px;
  margin-right: 5px;
  vertical-align: top;
} */

.content {
  width: 800px;
  height: 500px;
}

.items {
  --item-size: 56px;
  font-size: 10px;
  font-family: sans-serif;
  width: var(--item-size);
  height: var(--item-size);
  color: black;
  box-shadow: 0 0 8px black;
  margin-bottom: 5px;
  -webkit-user-drag: none;
  /* background: url("../resources/images/items/purple_eye.jpg") no-repeat; */
  border: 0;
  /* background-size: var(--item-size); */
  margin-right: 5px;
  user-select: none;
  transition: box-shadow 0.15s ease-in-out;
}
</style>