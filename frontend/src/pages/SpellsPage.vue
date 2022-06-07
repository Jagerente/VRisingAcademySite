<template>
    <div v-if="this.isSpellsLoading">
        <div class="position-absolute top-50 start-50 translate-middle">
            <span class="h1">Loading...</span>
        </div>
    </div>
    <div v-else class="catalogue d-flex flex-column flex-lg-row w-100 py-3 px-2">
        <!-- LEFT -->
        <!-- UP -->
        <div class="d-flex w-100 mb-2 mb-lg-0">
            <!-- FILTER -->
            <card-filter class=""></card-filter>
            <!-- CATALOGUE -->
            <card-catalogue class="flex-fill mx-2"></card-catalogue>
        </div>
        <!-- RIGHT -->
        <!-- INFORMATION -->
        <div class="d-lg-flex d-none">
            <card-information style="width: 600px;"></card-information>
        </div>
        <!-- BOTTOM -->
        <!-- INFORMATION -->
        <div class="d-flex d-lg-none flex-fill">
            <card-information style="width: 100%;"></card-information>
        </div>
    </div>
</template>

<script>
import CardFilter from '@/components/SpellsPage/CardFilter.vue'
import CardInformation from '@/components/SpellsPage/CardInformation.vue'
import CardCatalogue from '@/components/SpellsPage/CardCatalogue.vue'

import { mapState, mapActions } from 'vuex'

export default {
    components: {
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
@import '@/assets/styles/va_variables.scss';

.catalogue {
    padding-top: 15px;
}

.catalogue {
    height: calc(100vh - ($header-height + $footer-height) + 5px);
    min-height: calc(100vh - ($header-height + 25px));
}
</style>