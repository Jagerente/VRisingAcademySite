<template>
  <div class="d-flex flex-column">
    <div
      class="d-flex justify-content-evenly my-3 flex-wrap"
      v-for="stype in this.types"
    >
      <my-spell
        v-for="spell in stype.spells"
        :spell="spell"
        :class="this.selectedSpell && this.selectedSpell.id === spell.id ? 'active' : ''"
      ></my-spell>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import MySpell from "@/components/SpellsPage/MySpell";

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