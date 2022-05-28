<template>
    <div>
        <my-card title="Spells">
            <div class="d-flex flex-column h-100 test justify-content-between">
                <div class="d-flex justify-content-center">
                    <div class="btn-primary mx-5" v-for="spell in spells" :key="spell.id">{{ spell.name }}</div>
                </div>
                <div class="d-flex justify-content-center">
                    <div class="btn-primary mx-5" v-for="spell in advancedSpells" :key="spell.id">{{ spell.name }}</div>
                </div>
                <div class="d-flex justify-content-center">
                    <div class="btn-primary mx-5" v-for="spell in spells" :key="spell.id">{{ spell.name }}</div>
                </div>
            </div>
        </my-card>
    </div>
</template>

<script>
export default {
    name: 'card-information',
    data() {
        return {
            spells: [
                {
                    id: 1,
                    name: 'Spell Name',
                    type: 'Basic Spell',
                    description: 'Spell Description'
                },
                {
                    id: 2,
                    name: 'Spell Name',
                    type: 'Basic Spell',
                    description: 'Spell Description'
                },
            ],
            advancedSpells: [
                {
                    id: 1,
                    name: 'Spell Name',
                    type: 'Advanced Spell',
                    description: 'Spell Description'
                },
                {
                    id: 2,
                    name: 'Spell Name',
                    type: 'Advanced Spell',
                    description: 'Spell Description'
                },
                {
                    id: 3,
                    name: 'Spell Name',
                    type: 'Advanced Spell',
                    description: 'Spell Description'
                },
            ]
        }
    },
}
</script>

<style lang="scss" scoped>
</style>