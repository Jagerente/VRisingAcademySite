import axios from "axios";

export const spellsModule = {
    state: () => ({
        schools: [],
        types: [],
        spells: [],
        spellsGrouped: [],
        isSpellsLoading: true,
        selectedSpell: null,
        selectedSchool: 1,
        // host: "https://vrising-academy.info/api/"
        host: "http://localhost:8087/api/"
    }),
    getters: {

    },
    mutations: {
        setSchools(state, schools) {
            state.schools = schools
        },
        setTypes(state, types) {
            state.types = types
        },
        setSpells(state, spells) {
            state.spells = spells
        },
        setSpellsGrouped(state, spellsGrouped) {
            state.spellsGrouped = spellsGrouped
        },
        setSpellsLoading(state, bool) {
            state.isSpellsLoading = bool
        },
        setSelectedSchool(state, selectedSchool) {
            state.selectedSchool = selectedSchool
        },
        setSelectedSpell(state, selectedSpell) {
            state.selectedSpell = selectedSpell
        },
    },
    actions: {
        async fetchSpells({ state, commit }) {
            commit('setSpellsLoading', true)

            await axios
                .get(state.host + "spell/schools")
                .then(response => commit('setSchools', response.data))
                .catch(error => alert("Error: " + error));

            await axios
                .get(state.host + "spell/types")
                .then(response => commit('setTypes', response.data))
                .catch(error => alert("Error: " + error));

            await axios
                .get(state.host + "spell/list")
                .then(response => commit('setSpells', response.data))
                .catch(error => alert("Error: " + error))

            // await axios
            //     .get(state.host + "spell/grouplist")
            //     .then(response => commit('setSpellsGrouped', response.data))
            //     .catch(error => alert("Error: " + error))

            // --Remove after fix--
            let spells = [];
            state.schools.forEach((school, i) => {
                spells.push({ name: school.name, types: [] });
                state.types.forEach((type, j) => {
                    spells[i].types.push({ title: type.title, spells: [] })
                    state.spells.forEach(spell => {
                        if (spell.school.name === school.name && spell.type.name === type.title) {
                            spells[i].types[j].spells.push(spell);
                        }
                    });
                });
            })
            commit('setSpellsGrouped', spells)
            // --Remove after fix--
            
            commit('setSpellsLoading', false)
        },

        selectSchool({ commit }, id) {
            commit('setSelectedSchool', id);
        },
        selectSpell({ commit }, spell) {
            commit('setSelectedSpell', spell);
        },
    },
    namespaced: true
}