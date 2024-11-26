import axios from "axios";
import router from "@/router"; // Import the router
import { useAuthStore } from "@/stores/authStore"; // Import the auth store

// Create Axios instance
const api = axios.create({
  baseURL: "http://localhost:3001/", // Replace with your API base URL
});

// Request Interceptor
api.interceptors.request.use(
  (config) => {
    const authStore = useAuthStore(); // Access the auth store

    if (
      config.url?.includes("/authenticate") || 
      config.url?.includes("/register")
    ) {
      // Skip token for login and register endpoints
      return config;
    }

    if (authStore.token) {
      config.headers["Authorization"] = `Bearer ${authStore.token}`;
    } else {
      console.error("No token found, redirecting to login.");
      router.push("/user/login"); // Redirect to login if no token
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response Interceptor
api.interceptors.response.use(
  (response) => response,
  (error) => {
    const authStore = useAuthStore(); // Access the auth store

    if (error.response?.status === 401) {
      console.error("Unauthorized, logging out and redirecting to login.");
      authStore.logout(); // Clear auth state
      router.push("/user/login"); // Redirect to login if unauthorized
    }

    return Promise.reject(error);
  }
);

export default api;
