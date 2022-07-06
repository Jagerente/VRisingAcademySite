<template>
  <input
    @click="selectSpell(spell)"
    :class="spell.type.name === 'Ultimate' ? 'spell-ultimate' : spell.type.name === 'Travel Skill' ? 'spell-travel' : 'spell-basic'"
    class="spell"
    type="image"
    :title="spell.name"
    :src="
      getImageUrl(`${spell.school.name.toLowerCase()}/${spell.name}.webp`)
    "
  />
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
  setup() {
    const getImageUrl = (name) => {
      return new URL(`../../assets/images/spells/${name}`, import.meta.url).href
    }
    return { getImageUrl }
  }
};
</script>

<style scoped lang="scss">
$item-size: 5em;

.spell {
  background: none;
  -webkit-user-drag: none;
  user-select: none;
  transition: opacity 0.15s ease, box-shadow 0.15s ease, border 0.15s ease;
  border-radius: 3px;
  width: $item-size;
  height: $item-size;

  @media (max-width: $sm) {
    $item-size: 3em;
    width: $item-size;
    height: $item-size;
  }

  &:hover {
    opacity: 0.85;
  }

  &.spell-basic {
    border: 3px solid #534837;

    &:hover {
      box-shadow: 0 0 8px black;
    }

    &.active {
      border: 3px solid #917141;
      box-shadow: 0;
    }
  }

  &-ultimate {
    border-image: url("@/assets/images/spells/ui/UltiFrame.webp") 27 / 25px / 1rem;
    -webkit-border-image: url("@/assets/images/spells/ui/UltiFrame.webp") 27 / 25px / 1rem;

    &:hover {
      border-image: url("@/assets/images/spells/ui/UltiFrame.webp") 27 / 25px / 1rem;
      -webkit-border-image: url("@/assets/images/spells/ui/UltiFrame.webp") 27 / 25px / 1rem;
      box-shadow: 0 0 28px black;
    }

    &.active {
      border-image: url("@/assets/images/spells/ui/UltiFrame_Active.webp") 27 / 25px / 1rem;
      -webkit-border-image: url("@/assets/images/spells/ui/UltiFrame_Active.webp") 27 / 25px / 1rem;
    }
  }

  &-travel {
    border-image: url("@/assets/images/spells/ui/TravelFrame.webp") 27 / 25px / 1rem;
    -webkit-border-image: url("@/assets/images/spells/ui/TravelFrame.webp") 27 / 25px / 1rem;

    &:hover {
      border-image: url("@/assets/images/spells/ui/TravelFrame.webp") 27 / 25px / 1rem;
      -webkit-border-image: url("@/assets/images/spells/ui/TravelFrame.webp") 27 / 25px / 1rem;
      box-shadow: 0 0 28px black;
    }

    &.active {
      border-image: url("@/assets/images/spells/ui/TravelFrame_Active.webp") 27 / 25px / 1rem;
      -webkit-border-image: url("@/assets/images/spells/ui/TravelFrame_Active.webp") 27 / 25px / 1rem;
    }
  }
}
</style>