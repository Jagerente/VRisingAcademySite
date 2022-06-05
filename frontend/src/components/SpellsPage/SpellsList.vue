<template>
  <div class="d-flex flex-column">
    <div class="d-flex justify-content-evenly my-3 flex-wrap" v-for="spellType in clearTypes">
      <my-spell v-for="spell in spellType.spells" :spell="spell"></my-spell>
    </div>
  </div>
</template>

<script>
import MySpell from "@/components/SpellsPage/MySpell"

export default {
  name: "spells-list",
  components: {
    MySpell
  },
  props: {
    types: {
      type: Object,
      default: []
    },
  },
  computed: {
    clearTypes() {
      return this.types.filter(type => type.spells.length > 0)
    }
  },
};
</script>

<style scoped>
.items-list-enter-active,
.items-list-leave-active {
  transition: all 0.45s ease;
}

.items-list-enter-from,
.items-list-leave-to {
  opacity: 0;
  transform: translateY(30px);
}

.items-list-move {
  transition: transform 0.35s ease;
}

/* .items-list-leave-active {
  position: absolute;
} */

.input-group {
  height: 39px;
}
</style>