<template>
  <input @click="selectSpell(spell)" type="image" class="spell rounded"
    :class="spell.type === 'Ultimate' ? 'spell-ultimate' : spell.type === 'Travel Skill' ? 'spell-travel' : 'spell-basic'"
    :title="spell.name" :src="
      require(`@/assets/images/spells/${spell.school.toLowerCase()}/${spell.name}.webp`)
    " />
</template>

<script>
import { mapActions } from "vuex"

export default {
  name: "my-spell",
  props: {
    spell: Object,
  },
  methods: {
    ...mapActions({
      selectSpell: "spells/selectSpell"
    })
  },
};
</script>

<style scoped lang="scss">
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';


@include media-breakpoint-down(sm) {
  $item-size: 3em;
  $item-margin: 2px;

  .spell {
    width: $item-size;
    height: $item-size;
  }
}

@include media-breakpoint-up(sm) {
  $item-size: 5em;
  $item-margin: 5px;

  .spell {
    width: $item-size;
    height: $item-size;
  }
}

.spell {
  background: rgba(0, 0, 0, 0.5);
  -webkit-user-drag: none;
  user-select: none;
  transition: opacity 0.15s ease-in-out;
}

.spell:hover {
  opacity: 0.85;
}

.spell-basic {
  border: 3px solid #534837
}

.spell-basic:hover {
  box-shadow: 0 0 8px black;
}

.spell-basic:focus {
  border: 3px solid #917141;
  box-shadow: 0;
}

.spell-ultimate {
  border-image: url("@/assets/images/spells/ui/UltiFrame.webp") 27 / 25px / 1rem;
}

.spell-ultimate:hover {
  border-image: url("@/assets/images/spells/ui/UltiFrame.webp") 27 / 25px / 1rem;
  box-shadow: 0 0 28px black;
}

.spell-ultimate:focus {
  border-image: url("@/assets/images/spells/ui/UltiFrame_Active.webp") 27 / 25px / 1rem;
}

.spell-travel {
  border-image: url("@/assets/images/spells/ui/TravelFrame.webp") 27 / 25px / 1rem;
}

.spell-travel:hover {
  border-image: url("@/assets/images/spells/ui/TravelFrame.webp") 27 / 25px / 1rem;
  box-shadow: 0 0 28px black;
}

.spell-travel:focus {
  border-image: url("@/assets/images/spells/ui/TravelFrame_Active.webp") 27 / 25px / 1rem;
}
</style>