import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from '@/stores/authStore'; // Import the auth store
import Home from '@/views/Home.vue';
import Search from "@/views/Search.vue";
import Create from "@/views/Create.vue";
import Details from "@/views/Details.vue";
import Register from "@/views/Register.vue";
import Login from '@/views/Login.vue';

// Define routes
const routes = [
    {
        path: '/rentals',
        name: 'home',
        component: Home,
        meta: { requiresAuth: true }, // Protect this route
    },
    {
        path: '/rentals/search',
        name: 'search',
        component: Search,
    },
    {
        path: '/rentals/create',
        name: 'create',
        component: Create,
        meta: { requiresAuth: true }, // Protect this route
    },
    {
        path: '/rentals/details/:id',
        name: 'details',
        component: Details,
        meta: { requiresAuth: true },
    },
    {
        path: '/user/register',
        name: 'register',
        component: Register,
    },
    {
        path: '/user/login',
        name: 'login',
        component: Login,
    },
];

// Create router
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

// Middleware to check for authentication
router.beforeEach((to, from, next) => {
    const authStore = useAuthStore();

    if (to.meta.requiresAuth && !authStore.isAuthenticated) {
        // Redirect to login if the route requires authentication and user is not authenticated
        next({ name: 'login' });
    } else {
        // Proceed to the requested route
        next();
    }
});

export default router;
