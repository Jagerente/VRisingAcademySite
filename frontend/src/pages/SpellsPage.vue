<template>
    <MyLayout :isLoading="this.isSpellsLoading">
        <template #left>
            <card-filter />
        </template>
        <template #center>
            <card-catalogue />
        </template>
        <template #right>
            <card-information />
        </template>
    </MyLayout>
</template>

<script>
import MyLayout from "@/components/UI/MyLayout";
import CardFilter from '@/components/SpellsPage/CardFilter.vue'
import CardInformation from '@/components/SpellsPage/CardInformation.vue'
import CardCatalogue from '@/components/SpellsPage/CardCatalogue.vue'

import { mapState, mapActions } from 'vuex'

export default {
    components: {
        MyLayout,
        CardFilter,
        CardInformation,
        CardCatalogue,
    },
    computed: {
        ...mapState({
            isSpellsLoading: (state) => state.spells.isSpellsLoading,
        }),
    },
    methods: {
        ...mapActions({
            fetchSpells: "spells/fetchSpells",
        })
    },
    mounted() {
        this.fetchSpells();
    }

}
</script>

<style lang="scss" scoped>
</style>