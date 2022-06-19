<template>
    <MyLayout :isLoading="this.isSpellsLoading" :selector="this.selectSchool" :tabs="this.schools">
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
import CardInformation from '@/components/SpellsPage/CardInformation.vue'
import CardCatalogue from '@/components/SpellsPage/CardCatalogue.vue'

import { mapState, mapActions } from 'vuex'

export default {
    components: {
        MyLayout,
        CardInformation,
        CardCatalogue,
    },
    computed: {
        ...mapState({
            isSpellsLoading: (state) => state.spells.isSpellsLoading,
            schools: (state) => state.spells.schools
        }),
    },
    methods: {
        ...mapActions({
            fetchSpells: "spells/fetchSpells",
            selectSchool: "spells/selectSchool",
        })
    },
    mounted() {
        this.fetchSpells();
    }

}
</script>

<style lang="scss" scoped>
</style>