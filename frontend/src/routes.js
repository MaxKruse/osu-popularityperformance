import HelloWorld from './components/HelloWorld.vue'
import Code from './components/Code.vue'

export const routes = [
    {
        path: '/',
        component: HelloWorld
    },
    {
        path: '/code',
        component: Code
    }
]
