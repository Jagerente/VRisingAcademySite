<template>
    <MyLayout
        :isLoading="this.isSpellsLoading"
        :tabs="this.schools"
        :selector="this.selectSchool"
        :showModal="this.showModal"
        :updateShowModal="this.updateShowModal"
    >
        <template #center>
            <card-catalogue />
        </template>
        <template #right>
            <card-information />
        </template>
    </MyLayout>
</template>

<script>
import MyLayout from "@/components/UI/MyLayout.vue";
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
            schools: (state) => state.spells.schools,
            showModal: (state) => state.spells.showModal,
        }),
    },
    methods: {
        ...mapActions({
            fetchSpells: "spells/fetchSpells",
            selectSchool: "spells/selectSchool",
            updateShowModal: "spells/updateShowModal",
        })
    },
    mounted() {
        this.fetchSpells();
    }
}
</script>