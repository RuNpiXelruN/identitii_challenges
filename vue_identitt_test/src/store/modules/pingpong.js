const state = {
    pingPong: true,
    pingPongValue: "ping",
}

const getters = {
    getPingPong: (state) => state.pingPongValue
}

const actions = {
    setPingPong({commit}) {
        commit('updatePingPong')
    }
}

const mutations = {
    updatePingPong(state) {
        if (state.pingPong) {
            state.pingPongValue = "pong"
        } else {
            state.pingPongValue = "ping"
        }
        state.pingPong = !state.pingPong
    }
}

export default {
    state,
    getters,
    actions,
    mutations,
}