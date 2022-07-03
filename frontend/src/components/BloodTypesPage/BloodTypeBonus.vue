<template>
    <div>
        <div class="bonus">
            <div class="bonus__tier">
                <p
                    :class="{ 'active': this.isActive }"
                    class="tier"
                >
                    {{ this.getTier.name }}
                </p>
            </div>
            <div class="bonus__text">
                <p
                    :class="{ 'active': this.isActive }"
                    class="text"
                    v-html="this.getBonus"
                >
                </p>
            </div>
        </div>
    </div>
</template>

<script>
import { mapState } from "vuex";
import Markdown from 'vue3-markdown-it';

export default {
    props: {
        bonus: String,
        tier: Number,
    },
    components: {
        Markdown,
    },
    computed: {
        ...mapState({
            bloodQuality: (state) => state.bloodTypes.bloodQuality,
        }),
        getTier() {
            switch (this.tier) {
                case 0:
                    return { name: 'I', start: 0, };
                case 1:
                    return { name: 'II', start: 30, };
                case 2:
                    return { name: 'III', start: 60 };
                case 3:
                    return { name: 'IV', start: 90 };
                case 4:
                    return { name: 'V', start: 100 };
            }
        },
        isActive() {
            return this.bloodQuality >= this.getTier.start && this.bloodQuality > 0;
        },
        getBonus() {
            let bonus = this.bonus;
            let regex = /([0-9,.]+-[0-9,.]+|[0-9,.]+%-[0-9,.]+%|[0-9,.]+%)/g;
            let values = this.bonus.match(regex);
            if (values) {
                values.forEach(value => {
                    bonus = `<span>${bonus.replace(value, `<span style="color: teal;">${value}</span>`)}</span>`;
                });
            }
            if (bonus.match(/([0-9,.]+)%?-([0-9,.]+)%/) && this.bloodQuality > 0 && this.bloodQuality >= this.getTier.start) {

                let regex = /([0-9,.]+)%?-([0-9,.]+)%/;
                const bonuses = bonus.match(regex);
                const min = parseInt(bonuses[1]);
                const max = parseInt(bonuses[2]);
                const bonusCalculated = Math.round(min + (max - min) * (this.bloodQuality - this.getTier.start) / (100 - this.getTier.start));
                return bonus.replace(regex, bonusCalculated + '%');
            }
            else if (bonus.match(/([0-9,.]+)?-([0-9,.]+)/) && this.bloodQuality > 0 && this.bloodQuality >= this.getTier.start) {
                const regex = /([0-9,.]+)?-([0-9,.]+)/;
                const bonuses = bonus.match(regex);
                const min = parseInt(bonuses[1]);
                const max = parseInt(bonuses[2]);
                let bonusCalculated = Math.round(min + (max - min) * (this.bloodQuality - this.getTier.start) / (100 - this.getTier.start));
                return bonus.replace(regex, bonusCalculated);
            }
            else {
                return bonus;
            }
        },
    },
}
</script>

<style lang="scss" scoped>
@import '@/assets/styles/utility/vars.scss';

.bonus {
    display: flex;
    background-color: $dark;
    border-radius: 5px;
    padding: 15px;

    &__tier {
        width: 5%;
        margin-right: 15px;
    }

    &__text {
        width: 95%;
    }
}

.tier {
    text-align: center;
    transition-property: color;
    transition-duration: .2s;
    transition-timing-function: ease;

    &.active {
        color: yellow
    }
}

.text {
    font-size: 1.25rem;
    transition-property: font-size, color;
    transition-duration: .2s;
    transition-timing-function: ease;

    &.active {
        color: white
    }

    @media (max-width: $lg) {
        font-size: 1rem;
    }

    @media (max-width: $md) {
        font-size: .8rem;
    }
}
</style>