<template>
    <my-card :title="this.bloodTypes[this.selectedBloodType].name">
        <template #left>
            <img
                class="logo"
                :src="require('@/assets/images/blood_types/' + this.bloodTypes[this.selectedBloodType].name.toLowerCase() + '.webp')"
            />
        </template>
        <div
            class="d-flex flex-column h-100 w-100 px-1"
            style="padding-top: 30px"
        >
            <div class="d-flex flex-column">
                <BloodQualitySlider
                    class="col-11 col-lg-6 mx-auto px-2"
                    style="margin-bottom: 40px"
                />
                <BloodTypeBonus
                    :bonus="bonus"
                    :tier="j"
                    class="my-1 col-12 col-lg-6 mx-auto"
                    v-for="(bonus, j) in this.bloodTypes[this.selectedBloodType].bonuses"
                />
            </div>
        </div>
    </my-card>
</template>

<script>
import { mapState } from "vuex";
import SpellsList from '@/components/SpellsPage/SpellsList'
import BloodQualitySlider from '@/components/BloodTypesPage/BloodQualitySlider.vue'
import BloodTypeBonus from '@/components/BloodTypesPage/BloodTypeBonus.vue'

export default {
    components: {
        SpellsList,
        BloodQualitySlider,
        BloodTypeBonus
    },
    computed: {
        ...mapState({
            bloodTypes: (state) => state.bloodTypes.bloodTypes,
            selectedBloodType: (state) => state.bloodTypes.selectedBloodType,
        }),
    },
}
</script>

<style lang="scss" scoped>
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';

@include media-breakpoint-down(sm) {
    .logo {
        $size: 24px;
        width: $size;
        height: $size;
    }
}

@include media-breakpoint-up(sm) {
    .logo {
        $size: 32px;
        width: $size;
        height: $size;
    }
}
</style>