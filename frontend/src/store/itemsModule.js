import axios from "axios";

export const itemsModule = {
    state: () => ({
        items: [],
        itemsGrouped: [],
        types: [],
        sets: [],
        locations: [],
        recipes: [],
        salvageables: [],
        searchType: 1,
        isItemsLoading: true,
        searchQuery: "",
        selectedType: 0,
        selectedItem: null,
        matchingFloor: true,
        confinedRoom: true,
        host: "https://vrising-academy.info/api/"
        // host: "http://localhost:8087/api/"
    }),
    getters: {
        sortedItems: (state) => {
            return state.items.sort((item1, item2) => {
                return item1.tier - item2.tier;
            });
        },
        sortedAndSearchedItems: (state, getters) => {
            switch (state.searchType) {
                case 1:
                    return getters.sortedItems.filter(item => { return item.name.toLowerCase().includes(state.searchQuery.toLowerCase()) });
                case 2:
                    return getters.sortedItems.filter(item => { return item.tags.some(tag => tag.toLowerCase() === state.searchQuery.toLowerCase()) });
                default:
                    return getters.sortedItems.filter(item => { return item.name.toLowerCase().includes(state.searchQuery.toLowerCase()) || item.tags.some(tag => tag.toLowerCase() === state.searchQuery.toLowerCase()) });
            }
        },

        mapGenieLocations: (state) => {
            const locations = state.locations
                .filter(location => state.selectedItem.locations.includes(location.id))
                .map(location => (location.mapgenieId));
            return locations.join();
        }
    },
    mutations: {
        setItems(state, items) {
            state.items = items
        },
        setItemsGrouped(state, items) {
            state.itemsGrouped = items
        },
        setTypes(state, types) {
            state.types = types
        },
        setSets(state, sets) {
            state.sets = sets
        },
        setLocations(state, locations) {
            state.locations = locations
        },
        setRecipes(state, recipes) {
            state.recipes = recipes
        },
        setSalvageables(state, salvageables) {
            state.salvageables = salvageables
        },
        setLoading(state, bool) {
            state.isItemsLoading = bool
        },
        setSearchType(state, searchType) {
            state.searchType = searchType
        },
        setSearchQuery(state, searchQuery) {
            state.searchQuery = searchQuery
        },
        setSelectedType(state, selectedType) {
            state.selectedType = selectedType
        },
        setSelectedItem(state, selectedItem) {
            state.selectedItem = selectedItem
        },
        setMatchingFloor(state, matchingFloor) {
            state.matchingFloor = matchingFloor
        },
        setConfinedRoom(state, confinedRoom) {
            state.confinedRoom = confinedRoom
        }
    },
    actions: {
        async fetchItems({ state, commit }) {
            commit('setLoading', true);

            await axios
                .get(state.host + "item/list")
                .then(response => commit('setItems', response.data))
                .catch(error => alert("Error: " + error))

            await axios
                .get(state.host + "item/grouplist")
                .then(response => commit('setItemsGrouped', response.data))
                .catch(error => alert("Error: " + error))

            await axios
                .get(state.host + "item/types")
                .then(response => commit('setTypes', response.data))
                .catch(error => alert("Error: " + error));

            await axios
                .get(state.host + "set/list")
                .then(response => commit('setSets', response.data))
                .catch(error => alert("Error: " + error));

            await axios
                .get(state.host + "location/list")
                .then(response => commit('setLocations', response.data))
                .catch(error => alert("Error: " + error));

            await axios
                .get(state.host + "recipe/list")
                .then(response => commit('setRecipes', response.data))
                .catch(error => alert("Error: " + error));

            await axios
                .get(state.host + "salvageable/list")
                .then(response => commit('setSalvageables', response.data))
                .catch(error => alert("Error: " + error));

            commit('setLoading', false)
        },

        selectType({ commit }, id) {
            commit('setSelectedType', id);
            commit('setSearchQuery', '')
        },

        selectItem({ commit }, item) {
            commit('setSelectedItem', item)
        },

        updateSearchQuery({ commit }, { query = '', type = 0 }) {
            commit('setSearchType', type)
            commit('setSearchQuery', query)
        },

        updateMatchingFloor({ state, commit }) {
            commit('setMatchingFloor', !state.matchingFloor);
            if (state.matchingFloor) commit('setConfinedRoom', true);
        },

        updateConfinedRoom({ state, commit }) {
            commit('setConfinedRoom', !state.confinedRoom);
        },
    },
    namespaced: true
}