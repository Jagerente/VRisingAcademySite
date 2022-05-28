<template>
    <div>
        <my-card title="Information">
            <div class="d-flex flex-column">
                <div class="d-flex">
                    <div class="btn-primary mx-5" v-for="spell  in spells" :key="spell.id">{{ spell.name }}</div>
                </div>
            </div>
        </my-card>
    </div>
</template>

<script>
export default {
    name: 'card-information'
}
</script>

<style lang="scss" scoped>
</style>