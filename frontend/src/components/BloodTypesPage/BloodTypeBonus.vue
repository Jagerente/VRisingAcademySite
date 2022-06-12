<template>
    <div>
        <div class="d-flex block rounded py-2 py-lg-3 w-100">
            <div class="col-2 d-flex">
                <div
                    class="d-flex flex-column mx-auto justify-content-center tier"
                    :class="this.isActive ? 'tier-active' : ''"
                >
                    {{ this.getTier.name }}
                </div>
            </div>
            <div class="col-10 d-flex">
                <div
                    class="d-flex flex-column justify-content-center flex-wrap bonus"
                    :class="this.isActive ? 'description-active' : ''"
                    v-html="this.getBonus"
                >
                </div>
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
            if (bonus.match(/([0-9,.]+)%?-([0-9,.]+)%/) && this.bloodQuality > 0 && this.bloodQuality > this.getTier.start) {

                let regex = /([0-9,.]+)%?-([0-9,.]+)%/;
                const bonuses = bonus.match(regex);
                const min = parseInt(bonuses[1]);
                const max = parseInt(bonuses[2]);
                const bonusCalculated = Math.round(min + (max - min) * (this.bloodQuality - this.getTier.start) / (100 - this.getTier.start));
                return bonus.replace(regex, bonusCalculated + '%');
            }
            else if (bonus.match(/([0-9,.]+)?-([0-9,.]+)/) && this.bloodQuality > 0 && this.bloodQuality > this.getTier.start) {
                const regex = /([0-9,.]+)?-([0-9,.]+)/;
                const bonuses = bonus.match(regex);
                const min = parseInt(bonuses[1]);
                const max = parseInt(bonuses[2]);
                let bonusCalculated = Math.round(min + (max - min) * (this.bloodQuality - this.getTier.start) / (100 - this.getTier.start));
                return bonus.replace(regex, bonusCalculated);
            }
            else {
                // let bonus = this.bonus;
                // let regex = /([0-9,.]+-[0-9,.]+ |[0-9,.]+%-[0-9,.]+% |[0-9,.]+%)/g;
                // let values = bonus.match(regex);
                // if (values) {
                //     values.forEach(value => {
                //         bonus = bonus.replace(value, `<span style="color: teal; opacity: 0.5">${value}</span>`);
                //     });
                // }
                return bonus;
            }
        },
    },
}
</script>

<style lang="scss" scoped>
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';
@import '@/assets/styles/va_variables.scss';

p {
    margin: 0;
}

@include media-breakpoint-down(sm) {
    .bonus {
        font-size: 13px;
    }

    .tier {
        font-size: 18px;
    }
}

@include media-breakpoint-up(sm) {
    .tier {
        font-size: 22px;
    }
}

.tier-active {
    color: yellow;
}

.description-active {
    color: white;
    opacity: 1;
}

.block {
    background-color: #14141e;
}
</style>