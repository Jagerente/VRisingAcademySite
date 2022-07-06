<template>
    <my-card :title="this.bloodTypes[this.selectedBloodType].name">
        <template #left>
            <img
                class="icon"
                :src="getImageUrl(this.bloodTypes[this.selectedBloodType].name.toLowerCase() + '.webp')"
            />
        </template>
        <div
            class="content"
            style="padding-top: 30px"
        >
            <!-- <div class="d-flex flex-column"> -->
            <BloodQualitySlider
                class="content__slider"
                style="margin-bottom: 40px"
            />
            <BloodTypeBonus
                :bonus="bonus"
                :tier="j"
                class="content__bonus"
                v-for="(bonus, j) in this.bloodTypes[this.selectedBloodType].bonuses"
            />
        </div>
        <!-- </div> -->
    </my-card>
</template>

<script>
import { mapState } from "vuex";
import SpellsList from '@/components/SpellsPage/SpellsList.vue'
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
    setup() {
        const getImageUrl = (name) => {
            return new URL(`../../assets/images/blood_types/${name}`, import.meta.url).href
        }

        return {
            getImageUrl
        };

    }
}
</script>

<style lang="scss" scoped>
.content {
    display: flex;
    flex-direction: column;
    width: 100%;
    align-items: center;

    &__slider {
        width: 50%;
        margin-top: $m1;
        transition-property: width;
        transition-duration: .2s;
        transition-timing-function: ease;

        @media (max-width: $lg) {
            width: 60%;
        }

        @media (max-width: $md) {
            width: 80%;
        }

        @media (max-width: $sm) {
            width: 80%;
        }
    }

    &__bonus {
        width: 50%;
        margin-top: $m1;
        transition-property: width;
        transition-duration: .2s;
        transition-timing-function: ease;

        @media (max-width: $lg) {
            width: 60%;
        }

        @media (max-width: $md) {
            width: 80%;
        }

        @media (max-width: $sm) {
            width: 90%;
        }
    }
}

.icon {
    $size: 32px;
    width: $size;
    height: $size;
    transition-property: width, height;
    transition-duration: .2s;
    transition-timing-function: ease;

    @media (max-width: $lg) {
        $size: 24px;
        width: $size;
        height: $size;
    }
}
</style>