import axios from "axios";

export const itemsModule = {
    state: () => ({
        items: [],
        isItemsLoading: false,
        searchQuery: "",
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
                item.title.toLowerCase().includes(state.searchQuery.toLowerCase())
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
    },
    actions: {
        async getItems({ state, commit }) {
            try {
                commit('setLoading', true);
                const response = await axios.get(
                    "https://jsonplaceholder.typicode.com/posts",
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
    },
    namespaced: true
}