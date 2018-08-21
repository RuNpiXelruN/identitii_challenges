import Vuex from 'vuex';
import Vue from 'vue';
import pingPong from './modules/pingPong'
import form from './modules/form'

Vue.use(Vuex);

export default new Vuex.Store({
    modules: {
        form,
        pingPong,
    }
})