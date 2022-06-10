<template>
  <my-card title="Filter" :custom="false">
    <template #left>
      <icon-filter />
    </template>
    <div class="nav d-flex flex-column" id="v-pills-tab" role="tablist" aria-orientation="vertical">
      <button v-for="(school, i) in this.schools" @click="selectSchool(i)" class="spells__tab"
        :class="i == 0 ? 'active' : ''" data-bs-toggle="pill" :data-bs-target="`#v-pills-${school.name.toLowerCase()}`"
        type="button" role="tab" :aria-controls="`v-pills-${school.name.toLowerCase()}`" aria-selected="false">
        {{ school.name }}
      </button>
    </div>
  </my-card>
</template>

<script>
import { mapActions, mapGetters, mapState } from "vuex";

export default {
  methods: {
    ...mapActions({
      selectSchool: "spells/selectSchool",
    })
  },
  computed: {
    ...mapState({
      schools: (state) => state.spells.schools
    }),
  },
}
</script>

<style scoped lang="scss">
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';

@import '@/assets/styles/va_variables.scss';

@include media-breakpoint-down(sm) {
  $header-size: 13px;

  .spells__tab {
    font-size: $header-size;
    width: 100px;
    height: 35px;
  }
}

@include media-breakpoint-up(sm) {
  .spells__tab {
    height: 55px;
  }
}

.spells__tab {
  margin-bottom: 5px;
  background-color: #14141e;
  border: 1px solid #14141e;
  color: $text-color;
  text-transform: capitalize;
}

.spells__tab.active {
  background: url("@/assets/images/ui/circle.webp"), rgba(0, 0, 0, 0.5);
  background-size: 110%;
  background-position: 50%;
  background-repeat: no-repeat;
  border: 1px solid white;
  color: white;
}
</style>