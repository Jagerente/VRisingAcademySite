<template>
    <my-card
        title="Information"
        class="information"
    >
        <!-- If none is selected -->
        <div
            v-if="this.selectedSpell === null"
            class="information__block--column"
        >
            <div class="block--unknown">
                <p class="stat--unknown">Select spell.</p>
                <p class="stat--unknown">This module is still W.I.P.</p>
            </div>
        </div>
        <div
            v-else
            class="information__block--column"
        >
            <!-- Title -->
            <div class="header information__header">
                <p class="header__title">
                    {{ this.selectedSpell.name }}
                </p>
            </div>
            <div class="block information__block">
                <div class="block__left">
                    <!-- Type -->
                    <p class="stat--md">{{ this.selectedSpell.type.name.replace(/[0-9]/g, '') }}</p>
                    <!-- Cooldown -->
                    <p
                        class="stat"
                        v-if="this.selectedSpell.cooldown"
                    >Cooldown: <span class="stat__value">{{
                            this.selectedSpell.cooldown
                    }}s</span></p>
                    <!-- Cast Time -->
                    <p
                        class="stat"
                        v-if="this.selectedSpell.castTime"
                    >Cast Time: <span class="stat__value">{{
                            this.selectedSpell.castTime
                    }}s</span></p>
                    <!-- Charges -->
                    <p
                        class="stat"
                        v-if="this.selectedSpell.charges > 1"
                    >Charges: <span class="stat__value">{{
                            this.selectedSpell.charges
                    }}</span></p>
                </div>
                <!-- Preview -->
                <div class="block__right">
                    <img
                        class="image__preview"
                        draggable="false"
                        :title="selectedSpell.name"
                        :src="
                            getImageUrl(`${selectedSpell.school.name.toLowerCase()}/${selectedSpell.name}.webp`)
                        "
                    />
                </div>
            </div>
            <!-- Description -->
            <Markdown
                :source="this.selectedSpell.description"
                class="block__card"
                html
                xhtmlOut
                v-if="this.selectedSpell.description"
            />
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
    setup() {
        const getImageUrl = (name) => {
            return `images/spells/${name}`;
        };
        return { getImageUrl };
    }
}
</script>

<style lang="scss" scoped>
.image__preview {
    user-select: none;
    --img-size: 200px;
    width: var(--img-size);
    height: var(--img-size);

    @media (max-width: $sm) {
        --img-size: 100px;
        width: var(--img-size);
        height: var(--img-size);
    }
}

.information {
    width: 100%;

    &__block {
        display: flex;
        justify-content: space-between;

        &:not(:last-child) {
            margin-bottom: $m1;
        }

        &--column {
            flex-direction: column;

            &:not(:last-child) {
                margin-bottom: $m1;
            }
        }
    }

    &__header {
        margin-bottom: $m1;
    }

}

.header {
    &__title {
        font-size: 2rem;
        color: white;
    }
}

.stat {
    font-size: 1rem;
    margin-bottom: 5px;

    &--md {
        font-size: 1.25rem;
        margin-bottom: 5px;

    }

    &--unknown {
        text-align: center;
        font-size: 2rem;
        margin-bottom: 5px;
    }

    &__value {
        color: white;
    }
}

.block {
    &__card {
        background: $dark;
        border-radius: 10px;
        margin-bottom: $m1;
        padding-left: 15px;
        padding-right: 15px;
        padding-top: 15px;
        padding-bottom: 15px;
    }

    &--unknown {
        padding-top: 15px;
    }
}
</style>