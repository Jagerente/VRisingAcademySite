import { walkIdentifiers } from "@vue/compiler-core";
import axios from "axios";

export const itemsModule = {
    state: () => ({
        items: [],
        weapons: [],
        armour: [],
        consumables: [],
        reagents: [],
        isItemsLoading: false,
        searchQuery: "",
        selectedItem: null,
        loadedItems: 0
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
            return getters.sortedItems(type).filter(item =>
                item.name.toLowerCase().includes(state.searchQuery.toLowerCase()) || item.tags.some(tag => tag.toLowerCase().includes(state.searchQuery.toLowerCase()))
            );
        },
    },
    mutations: {
        setItems(state, items) {
            state.items = items
        },
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
                    "http://localhost:8081/api/weapon/list",
                );
                commit('setWeapons', response.data);
                response = await axios.get(
                    "http://localhost:8081/api/armour/list",
                );
                commit('setArmour', response.data);
                response = await axios.get(
                    "http://localhost:8081/api/consumable/list",
                );
                commit('setConsumables', response.data);
                response = await axios.get(
                    "http://localhost:8081/api/reagent/list",
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
        selectItem({ state, commit }, id) {
            const selectedItem = state.items.find(item => item.id === id)
            commit('setSelectedItem', selectedItem)
        },
        updateSearchQuery({ state, commit }, query) {
            commit('setSearchQuery', query)
        }
    },
    namespaced: true
}