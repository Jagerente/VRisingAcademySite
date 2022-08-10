<template>
  <div class="spells d-flex flex-column">
    <div
      class="spells__row"
      v-for="stype in this.types"
    >
      <my-spell
        v-for="spell in stype.spells"
        :spell="spell"
        :class="{ 'active': this.selectedSpell && this.selectedSpell.id === spell.id }"
      ></my-spell>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import MySpell from "@/components/SpellsPage/MySpell.vue";

export default {
  name: "spells-list",
  components: {
    MySpell
  },
  props: {
    types: Array
  },
  computed: {
    ...mapState({
      selectedSpell: (state) => state.spells.selectedSpell
    })
  },
};
</script>

<style lang="scss" scoped>
.spells {
  display: flex;
  flex-direction: column;

  &__row {
    display: flex;
    justify-content: space-evenly;
    flex-wrap: wrap;
    margin: $m2 0px;
  }
}
</style>