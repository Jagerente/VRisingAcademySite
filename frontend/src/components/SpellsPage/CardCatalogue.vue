<template>
    <my-card title="Spells">
        <div class="tab-content d-flex justify-content-center h-100 flex-wrap" id="v-pills-tabContent">
            <div v-for="(school, i) in this.spellsGrouped" class="tab-pane fade show h-100 w-100" :class="i === 0 ? 'active' : ''"
                :id="`v-pills-${school.name.toLowerCase()}`" role="tabpanel"
                :aria-labelledby="`v-pills-${school.name.toLowerCase()}-tab`">
                <spells-list :types="school.types"></spells-list>
            </div>
        </div>
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
        }),
    },
}
</script>

<style lang="scss" scoped>
</style>