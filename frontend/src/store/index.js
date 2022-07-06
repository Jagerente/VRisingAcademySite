import { createStore } from "vuex";
import { itemsModule } from "@/store/itemsModule.js";
import { spellsModule } from "@/store/spellsModule.js";
import { bloodTypesModule } from "@/store/bloodTypesModule.js";

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