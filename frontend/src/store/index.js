import { createStore } from "vuex";
import { itemsModule } from "@/store/itemsModule";
import { spellsModule } from "@/store/spellsModule";
import { bloodTypesModule } from "@/store/bloodTypesModule";

export default createStore({
    state: {
    },
    getters: {
    },
    mutations: {
    },
    actions: {
    },
    modules: {
        items: itemsModule,
        spells: spellsModule,
        bloodTypes: bloodTypesModule
    }
})