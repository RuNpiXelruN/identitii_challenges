import api from '@/sampleData.json'

const state = {
    formData: {}
}

const getters = {
    formData: (state) => state.formData
}

const actions = {
    async initFormData({ commit }) {
        try {
            let data = await api
                commit('setFormData', data)          
        } catch (err) {
            console.log("Error initialising form data")
        }
    }
}

const mutations = {
    setFormData(state, data) {
        state.formData = data
    }
}

export default {
    state,
    getters,
    actions,
    mutations
}