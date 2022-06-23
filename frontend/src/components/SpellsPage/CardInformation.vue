<template>
    <my-card
        title="Information"
        class="w-100"
    >
        <div class="d-flex flex-column p-2">
            <div
                v-if="this.selectedSpell"
                class="d-flex flex-column"
            >
                <div class="d-flex">
                    <div class="d-flex flex-column w-50">
                        <!-- Title -->
                        <p class="h-1">{{ this.selectedSpell.name }}</p>
                        <!-- Type -->
                        <p class="h-2">{{ this.selectedSpell.type.name.replace(/[0-9]/g, '') }}</p>
                        <!-- Cooldown -->
                        <p
                            class="h-3"
                            v-if="this.selectedSpell.cooldown"
                        >Cooldown: <span class="text-white">{{
                                this.selectedSpell.cooldown
                        }}s</span></p>
                        <!-- Cast Time -->
                        <p
                            class="h-3"
                            v-if="this.selectedSpell.castTime"
                        >Cast Time: <span class="text-white">{{
                                this.selectedSpell.castTime
                        }}s</span></p>
                        <!-- Charges -->
                        <p
                            class="h-3"
                            v-if="this.selectedSpell.charges > 1"
                        >Charges: <span class="text-white">{{
                                this.selectedSpell.charges
                        }}</span></p>
                    </div>
                    <!-- Preview -->
                    <div class="d-flex justify-content-center flex-fill mx-2 w-50">
                        <img
                            class="image__preview rounded"
                            draggable="false"
                            :title="selectedSpell.name"
                            :src="
                                require(`@/assets/images/spells/${selectedSpell.school.name.toLowerCase()}/${selectedSpell.name}.webp`)
                            "
                        />
                    </div>
                </div>
                <!-- Description -->
                <Markdown
                    :source="this.selectedSpell.description"
                    class="description py-2 my-3"
                    html
                    xhtmlOut
                    v-if="this.selectedSpell.description"
                />
            </div>
            <!-- If none is selected -->
            <div v-else>
                <h1 class="text-center">Select spell.</h1>
                <h5 class="text-center">This module is still W.I.P.</h5>
            </div>
        </div>
    </my-card>
</template>

<script>
import { mapState } from 'vuex'
import Markdown from 'vue3-markdown-it';

export default {
    components: {
        Markdown,
    },
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

@import '@/assets/styles/utility/vars.scss';


@include media-breakpoint-down(sm) {
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

p {
    margin: 0;
    padding: 0;
}
</style>