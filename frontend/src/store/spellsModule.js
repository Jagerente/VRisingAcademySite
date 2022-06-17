import axios from "axios";
import router from '@/router/router';

export const spellsModule = {
    state: () => ({
        schools: [],
        types: [],
        spells: [],
        spellsGrouped: [],
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

            await axios
                .get(state.host + "spell/grouplist")
                .then(response => commit('setSpellsGrouped', response.data))
                .catch(error => alert("Error: " + error))

            if (!router.currentRoute._value.query.id) {
                router.replace({
                    query: {
                        ...router.currentRoute.params,
                        id: 1
                    }
                });
                commit('setSelectedSpell', state.spells[0]);
            }
            else if (router.currentRoute._value.query.id > 0 && router.currentRoute._value.query.id <= state.spells.length) {
                commit('setSelectedSpell', state.spells[router.currentRoute._value.query.id - 1]);
            }
            else {
                commit('setSelectedSpell', null)
            }

            commit('setSpellsLoading', false)
        },

        selectSchool({ commit }, id) {
            commit('setSelectedSchool', id);
        },
        selectSpell({ commit }, spell) {
            router.replace({
                query: {
                    ...router.currentRoute.params,
                    id: spell.id
                }
            });
            commit('setSelectedSpell', spell);
        },
    },
    namespaced: true
}