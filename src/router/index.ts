import {createRouter, createWebHistory} from 'vue-router'
import HomeView from '../views/Home.vue'
import Search from "@/views/Search.vue";
import Create from "@/views/Create.vue";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/home',
            name: 'home',
            component: HomeView
        },
        {
            path: '/',
            name: 'search',
            component: Search
        },
        {
            path: '/create',
            name: 'create',
            component: Create
        },
    ]
})

export default router
