import {createRouter, createWebHistory} from 'vue-router'
import HomeView from '../views/Home.vue'
import Search from "@/views/Search.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView
        },
        {
            path: '/search',
            name: 'search',
            component: Search
        },
    ]
})

export default router
