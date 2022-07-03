<template>
    <my-card title="Spells">
        <spells-list :types="this.spellsGrouped[this.selectedSchool].types"></spells-list>
    </my-card>
</template>

<script>
import { mapState } from "vuex";
import SpellsList from "@/components/SpellsPage/SpellsList"

export default {
    components: {
        SpellsList
    },
    computed: {
        ...mapState({
            spellsGrouped: (state) => state.spells.spellsGrouped,
            selectedSchool: (state) => state.spells.selectedSchool
        }),
    },
}
</script>