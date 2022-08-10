import axios from "axios";

export const bloodTypesModule = {
  state: () => ({
    bloodTypes: [],
    isBloodTypesLoading: true,
    selectedBloodType: 0,
    bloodQuality: 0,
    host: "https://dev.vrising-academy.info/api/"
    // host: "http://localhost:8087/api/",
  }),
  getters: {},
  mutations: {
    setBloodTypes(state, bloodTypes) {
      state.bloodTypes = bloodTypes;
    },
    setBloodTypesLoading(state, bool) {
      state.isBloodTypesLoading = bool;
    },
    setSelectedBloodType(state, selectedBloodType) {
      state.selectedBloodType = selectedBloodType;
    },
    setBloodQuality(state, bloodQuality) {
      state.bloodQuality = bloodQuality;
    },
  },
  actions: {
    async fetchBloodTypes({ state, commit }) {
      if (state.bloodTypes.length > 0) {
        return;
      }
      
      commit("setBloodTypesLoading", true);
      await axios
        .get(state.host + "bloodtype/list")
        .then((response) => commit("setBloodTypes", response.data))
        .catch((error) => alert("Error: " + error));

      commit("setBloodTypesLoading", false);
    },

    selectBloodType({ commit }, id) {
      commit("setSelectedBloodType", id);
    },

    updateBloodQuality({ commit, state }, value) {
      commit("setBloodQuality", value);
    },
  },
  namespaced: true,
};
