import { walkIdentifiers } from "@vue/compiler-core";
import axios from "axios";

export const itemsModule = {
    state: () => ({
        weapons: [],
        armour: [],
        consumables: [],
        reagents: [],
        items: [],
        sets: [],
        recipes: [],
        isItemsLoading: false,
        searchQuery: "",
        selectedItem: null,
        loadedItems: 0,
        host: "https://vrising-academy.info/api/"
        // host: "http://localhost:8087/api/"
    }),
    getters: {
        sortedItems: (state) => (type) => {
            let items;
            switch (type) {
                case 1:
                    items = [...state.weapons];
                    break;
                case 2:
                    items = [...state.armour];
                    break;
                case 3:
                    items = [...state.consumables];
                    break;
                case 4:
                    items = [...state.reagents];
                    break;
                default:
                    items = [...state.items];
                    break;
            }
            return items.sort((item1, item2) => {
                return item1.tier - item2.tier;
            });
        },
        sortedAndSearchedItems: (state, getters) => (type) => {
            if (!state.searchQuery) {
                return getters.sortedItems(type);
            } else {
                return getters.sortedItems(0).filter(item =>
                    item.name.toLowerCase().includes(state.searchQuery.toLowerCase()) || item.tags.some(tag => tag.toLowerCase().includes(state.searchQuery.toLowerCase())));
            }
        },
    },
    mutations: {
        setWeapons(state, weapons) {
            state.weapons = weapons
        },
        setArmour(state, armour) {
            state.armour = armour
        },
        setConsumables(state, consumables) {
            state.consumables = consumables
        },
        setReagents(state, reagents) {
            state.reagents = reagents
        },
        setItems(state, items) {
            state.items = items
        },
        setSets(state, sets) {
            state.sets = sets
        },
        setRecipes(state, recipes) {
            state.recipes = recipes
        },
        setLoadedItems(state, count) {
            state.loadedItems = count
        },
        setLoading(state, bool) {
            state.isItemsLoading = bool
        },
        setSearchQuery(state, searchQuery) {
            state.searchQuery = searchQuery
        },
        setSelectedItem(state, selectedItem) {
            state.selectedItem = selectedItem
        }
    },
    actions: {
        async getItems({ state, commit }) {
            try {
                commit('setLoading', true);

                let response;
                response = await axios.get(
                    state.host + "weapon/list",
                );
                commit('setWeapons', response.data);
                response = await axios.get(
                    state.host + "armour/list",
                );
                commit('setArmour', response.data);
                response = await axios.get(
                    state.host + "consumable/list",
                );
                commit('setConsumables', response.data);
                response = await axios.get(
                    state.host + "reagent/list",
                );
                commit('setReagents', response.data);

            } catch (e) {
                alert("Error: " + e);
            } finally {
                const items = state.weapons.concat(state.armour, state.consumables, state.reagents);
                commit('setItems', items);
                commit('setLoading', false);
            }
        },
        async getSets({ state, commit }) {
            try {
                const response = await axios.get(
                    state.host + "set/list",
                );

                commit('setSets', response.data);
            } catch (e) {
                alert("Error: " + e);
            } finally {

            }
        },
        async getRecipes({ state, commit }) {
            try {
                const response = await axios.get(
                    state.host + "recipe/list",
                );

                commit('setRecipes', response.data);
            } catch (e) {
                alert("Error: " + e);
            } finally {

            }
        },
        selectItem({ state, dispatch, commit }, id) {
            const selectedItem = state.items.find(item => item.id === id)
            if (selectedItem.setId) {
                const set = [...state.sets].find(set => set.id === selectedItem.setId)
                if (!set.description.toLowerCase().includes("no bonus")) {
                    selectedItem.set = set;
                }
            }
            if (selectedItem.recipes) {
                const result = [];
                selectedItem.recipes.forEach(x => result.push([...state.recipes].find(r => r.id == x)));
                selectedItem.reagentFor.forEach(x => result.push([...state.recipes].find(r => r.id == x)));
                selectedItem.recipesInfo = result;
            }
            commit('setSelectedItem', selectedItem)

        },
        updateSearchQuery({ state, commit }, query) {
            commit('setSearchQuery', query)
        }
    },
    namespaced: true
}