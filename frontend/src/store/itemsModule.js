import axios from "axios";

export const itemsModule = {
    state: () => ({
        items: [],
        itemsList: [],
        types: [],
        sets: [],
        recipes: [],
        isItemsLoading: true,
        searchQuery: "",
        selectedItem: null,
        loadedItems: 0,
        // host: "https://vrising-academy.info/api/"
        host: "http://localhost:8087/api/"
    }),
    getters: {
        imagePath: (state) => (item) => {
            return '@/assets/images/items/' + item.type.toLowerCase() + '/' + (item.type.toLowerCase() !== 'blueprints' ? item.name : (item.name === 'The General\'s Soul Reaper Orb' ? 'The General\'s Soul Reaper Orb' : item.tags[1])) + '.webp';
        },
        sortedItems: (state) => {
            return state.itemsList.sort((item1, item2) => {
                return item1.tier - item2.tier;
            });
        },
        sortedAndSearchedItems: (state, getters) => {
            return getters.sortedItems.filter(item =>
                item.name.toLowerCase().includes(state.searchQuery.toLowerCase()) || item.tags.some(tag => tag.toLowerCase().includes(state.searchQuery.toLowerCase())));
        },
    },
    mutations: {
        setItems(state, items) {
            state.items = items
        },
        setItemsList(state, items) {
            state.itemsList = items
        },
        setTypes(state, types) {
            state.types = types
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
            commit('setLoading', true);
            await axios
                .get(
                    state.host + "item/list",
                )
                .then(response => {
                    commit('setItems', response.data)
                    let itemsList = [];
                    Object.keys(response.data).forEach(type => {
                        Object.keys(response.data[type]).forEach(set => {
                            itemsList = itemsList.concat(response.data[type][set])
                        })
                    })
                    commit('setItemsList', itemsList)
                })
                .catch(error => alert("Error: " + error))
                .finally(() => {
                    commit('setLoading', false)
                });
            await axios
                .get(
                    state.host + "item/types",
                )
                .then(response => {
                    commit('setTypes', response.data)
                })
                .catch(error => alert("Error: " + error));
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
        selectItem({ state, dispatch, commit }, item) {
            const selectedItem = item;
            if (selectedItem.setId) {
                const set = [...state.sets].find(set => set.id === selectedItem.setId)
                if (!set.description.toLowerCase().includes("no bonus")) {
                    selectedItem.setInfo = set;
                }
            }
            if (selectedItem.recipes) {
                const result = [];
                selectedItem.recipes.forEach(x => result.push([...state.recipes].find(r => r.id == x)));
                selectedItem.recipesInfo = result;
            }
            if (selectedItem.reagentFor) {
                const result = [];
                selectedItem.reagentFor.forEach(x => result.push([...state.recipes].find(r => r.id == x)));
                selectedItem.reagentForInfo = result;
            }
            commit('setSelectedItem', selectedItem)

        },
        updateSearchQuery({ state, commit }, query) {
            commit('setSearchQuery', query)
        }
    },
    namespaced: true
}