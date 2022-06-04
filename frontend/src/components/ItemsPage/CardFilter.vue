<template>
  <my-card title="Filter" :custom="false">
    <template #left>
      <icon-filter />
    </template>
    <div class="nav d-flex flex-column" id="v-pills-tab" role="tablist" aria-orientation="vertical">
      <button v-for="(type, i) in this.types" @click="updateSearchQuery('')" class="items__tab rounded"
        :class="i == 0 ? 'active' : ''" data-bs-toggle="pill" :data-bs-target="`#v-pills-${type.title.toLowerCase()}`"
        type="button" role="tab" :aria-controls="`v-pills-${type.title.toLowerCase()}`" aria-selected="false">
        {{ type.title }}
      </button>
    </div>
  </my-card>
</template>

<script>
import { mapActions, mapGetters, mapState } from "vuex";

export default {
  methods: {
    ...mapActions({
      updateSearchQuery: "items/updateSearchQuery",
    }),
  },
  computed: {
    ...mapState({
      types: (state) => state.items.types,
      items: (state) => state.items.items,
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

  .items__tab {
    font-size: $header-size;
    // width: 100px;
    height: 35px;
  }
}

@include media-breakpoint-up(sm) {
  .items__tab {
    height: 55px;
  }
}

.items__tab {
  margin-bottom: 5px;
  background-color: #14141e;
  border: 1px solid #14141e;
  color: $text-color;
  text-transform: capitalize;
}

.items__tab.active {
  background: url("@/assets/images/ui/circle.webp"), rgba(0, 0, 0, 0.5);
  background-size: 110%;
  background-position: 50%;
  background-repeat: no-repeat;
  border: 1px solid white;
  color: white;
}

// @import 'bootstrap/scss/_functions.scss';
// @import 'bootstrap/scss/_variables.scss';
// @import 'bootstrap/scss/_mixins.scss';

// @include media-breakpoint-down(sm) {
//   $header-size: 13px;

//   .nav-link {
//     font-size: $header-size;
//     width: 100px;
//     height: 35px;

//   }
// }

// @include media-breakpoint-up(sm) {
//   .nav-link {
//     height: 55px;
//   }
// }

// .nav-pills .nav-link.active,
// .nav-pills .show>.nav-link {
//   background: url("@/assets/images/ui/circle.webp"), rgba(0, 0, 0, 0.5);
//   background-size: 110%;
//   background-position: 50%;
//   background-repeat: no-repeat;
//   border: 1px solid white;
// }

// .nav-pills .nav-link {
//   margin-bottom: 5px;
//   background-color: #14141e;
//   border: 1px solid #14141e;
// }
</style>