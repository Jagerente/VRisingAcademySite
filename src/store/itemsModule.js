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
        selectedItem: null
    }),
    getters: {
        // sortedItems(state) {
        //     return [...state.items].sort((item1, item2) => {
        //         return item1[state.selectedSort]?.localeCompare(
        //             item2[state.selectedSort]
        //         );
        //     });
        // },
        sortedAndSearchedItems(state, getters) {
            return [...state.items].filter(item =>
                item.name.toLowerCase().includes(state.searchQuery.toLowerCase())
            );
        },
    },
    mutations: {
        setItems(state, items) {
            state.items = items
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
                const response = await axios.get(
                    "http://localhost:8081/api/items/list",
                    // {
                    //     params: {
                    //         _page: state.page,
                    //         _limit: state.limit,
                    //     },
                    // }
                );
                commit('setItems', response.data);
            } catch (e) {
                alert("Error: " + e);
            } finally {
                commit('setLoading', false);
            }
        },
        selectItem({ state, commit }, name) {
            const selectedItem = state.items.find(item => item.name === name)
            commit('setSelectedItem', selectedItem)
        },
    },
    namespaced: true
}