import Vue from 'vue'
import Router from 'vue-router'

// Lazy loading components below
const AppMain = () => System.import('@/components/AppMain')
const FormComponent = () => System.import('@/components/FormComponent')
const PingPongComponent = () => System.import('@/components/PingPongComponent')

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'AppMain',
            component: AppMain,                
        },
        {
            path: '/attribute-form',
            name: 'AttributeForm',
            component: FormComponent,                
        },
        {
            path: '/ping-pong',
            name: 'PingPong',
            component: PingPongComponent,                
        },
    ]
})
