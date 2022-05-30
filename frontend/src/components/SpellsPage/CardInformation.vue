<template>
    <div>
        <my-card title="Information">
            <div class="d-flex flex-column p-2">
                <div v-if="this.selectedSpell" class="d-flex flex-column">
                    <div class="d-flex">
                        <div class="d-flex flex-column w-50">
                            <!-- Title -->
                            <h1>{{ this.selectedSpell.name }}</h1>
                            <h4>{{ this.selectedSpell.type.replace(/[0-9]/g, '') }}</h4>
                            <h4>Cooldown: <span class="text-white">{{ this.selectedSpell.cooldown }}s</span></h4>
                            <h4 v-if="this.selectedSpell.castTime">Cast Time: <span class="text-white">{{
                                    this.selectedSpell.castTime
                            }}s</span></h4>
                            <h4 v-if="this.selectedSpell.charges > 1">Charges: <span class="text-white">{{
                                    this.selectedSpell.charges
                            }}</span></h4>
                        </div>
                        <div class="d-flex justify-content-center flex-fill mx-2 w-50">
                            <img class="image__preview rounded" draggable="false" :title="selectedSpell.name" :src="
                                require(`@/assets/images/spells/${selectedSpell.school.toLowerCase()}/${selectedSpell.name}.webp`)
                            " />
                        </div>
                    </div>
                    <div v-if="selectedSpell.description" class="description py-2 my-3">
                        <h4 class="">{{ selectedSpell.description }}</h4>
                    </div>
                </div>
                <div v-else>
                    <h1 class="text-center">Select spell.</h1>
                    <h5 class="text-center">This module is still W.I.P.</h5>
                </div>
            </div>
        </my-card>
    </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
    name: 'card-information',
    computed: {
        ...mapState({
            selectedSpell: (state) => state.spells.selectedSpell
        })
    },
}
</script>

<style lang="scss" scoped>
@import 'bootstrap/scss/_functions.scss';
@import 'bootstrap/scss/_variables.scss';
@import 'bootstrap/scss/_mixins.scss';

@include media-breakpoint-down(sm) {
    h1 {
        font-size: 1.5rem;
    }

    h2 {
        font-size: 1.5rem;
    }

    h3 {
        font-size: 1rem;
    }

    h4 {
        font-size: 1rem;
    }

    h5 {
        font-size: 1rem;
    }

    .tag {
        margin-top: 5px;
        font-size: 0.7rem;
    }

    .image__preview {
        --img-size: 100px;
        width: var(--img-size);
        height: var(--img-size);
    }

}

@include media-breakpoint-up(sm) {
    .image__preview {
        --img-size: 200px;
        width: var(--img-size);
        height: var(--img-size);
    }
}

.image__preview {
    user-select: none;
}

.description {
    background: #14131b;
    border-radius: 10px;
    margin-left: -10px;
    margin-right: -10px;
    padding-left: 15px;
    padding-right: 15px;
}
</style>