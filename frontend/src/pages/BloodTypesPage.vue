<template>
    <MyLayout
        :isLoading="this.isBloodTypesLoading"
        :selector="this.selectBloodType"
        :tabs="this.bloodTypes.slice(0, -3)"
        :right="false"
        :tabLogo="true"
    >
        <template #center>
            <BloodTypes />
        </template>
    </MyLayout>
</template>

<script>
import MyLayout from "@/components/UI/MyLayout.vue";
import BloodTypes from '@/components/BloodTypesPage/BloodTypes.vue'

import { mapState, mapActions } from 'vuex'

export default {
    components: {
        MyLayout,
        BloodTypes,
    },
    computed: {
        ...mapState({
            isBloodTypesLoading: (state) => state.bloodTypes.isBloodTypesLoading,
            bloodTypes: (state) => state.bloodTypes.bloodTypes
        }),
    },
    methods: {
        ...mapActions({
            fetchBloodTypes: "bloodTypes/fetchBloodTypes",
            selectBloodType: "bloodTypes/selectBloodType",
        })
    },
    mounted() {
        this.fetchBloodTypes();
    }

}
</script>