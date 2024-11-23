import {createRouter, createWebHistory} from 'vue-router'
import Home from '../views/Home.vue'
import Search from "@/views/Search.vue";
import Create from "@/views/Create.vue";
import Details from "@/views/Details.vue";
import Register from "@/views/Register.vue";
import Login from '@/views/Login.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/home',
            name: 'home',
            component: Home
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
        {
            path: '/details/:id',
            name: 'details',
            component: Details
        },
        {
            path: '/user/register',
            name: 'register',
            component: Register
        },
        {
            path: '/user/login',
            name: 'login',
            component: Login
        },
    ]
})

export default router
