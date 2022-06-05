import axios from "axios";

export const spellsModule = {
    state: () => ({
        schools: [],
        types: [],
        spells: [],
        isSpellsLoading: true,
        selectedSpell: null,
        selectedSchool: 1,
        host: "https://vrising-academy.info/api/"
        // host: "http://localhost:8087/api/"
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
        async getSchools({ state, commit }) {
            await axios
                .get(state.host + "spell/schools")
                .then(response => commit('setSchools', response.data))
                .catch(error => alert("Error: " + error));

        },
        async getTypes({ state, commit }) {
            await axios
                .get(state.host + "spell/types")
                .then(response => commit('setTypes', response.data))
                .catch(error => alert("Error: " + error));

        },
        async getSpells({ state, commit }) {
            await axios
                .get(state.host + "spell/list")
                .then(response => {
                    let spells = [];
                    state.schools.forEach((school, i) => {
                        spells.push({ name: school.name, types: [] });
                        state.types.forEach((type, j) => {
                            spells[i].types.push({ title: type.title, spells: [] })
                            response.data.forEach(spell => {
                                if (spell.school === school.name && spell.type === type.title) {
                                    spells[i].types[j].spells.push(spell);
                                }
                            });
                        });
                    });
                    commit('setSpells', spells);
                })
                .catch(error => alert("Error: " + error))
                .finally(commit('setSpellsLoading', false));
        },
        selectSchool({ commit }, id) {
            commit('setSelectedSchool', id);
        },
        selectSpell({ state, commit }, spell) {
            commit('setSelectedSpell', spell);
        },
    },
    namespaced: true
}