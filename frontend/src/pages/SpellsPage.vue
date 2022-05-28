<template>
<div class="d-flex h-100 w-100 test justify-content-center">
    <card-catalogue class="w-50 m-2"></card-catalogue>
    <card-information class="w-50 m-2"></card-information>
</div>
</template>

<script>
import CardInformation from '@/components/SpellsPage/CardInformation.vue'
import CardCatalogue from '@/components/SpellsPage/CardCatalogue.vue'
export default {
    components: {
        CardInformation,
        CardCatalogue
    },
}
</script>

<style lang="scss" scoped>
</style>