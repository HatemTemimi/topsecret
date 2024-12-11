import axios from "axios";
import router from "@/router"; // Import the router
import { useAuthStore } from "@/stores/authStore"; // Import the auth store

// Create Axios instance
const api = axios.create({
  baseURL: "http://localhost:3001/", // Replace with your API base URL
  withCredentials: true, // Include cookies in requests
});

// Request Interceptor
api.interceptors.request.use(
  (config) => {
    const authStore = useAuthStore(); // Access the auth store

    if (
      config.url?.includes("/login") || 
      config.url?.includes("/register")
    ) {
      // Skip additional checks for login and register endpoints
      return config;
    }

    // If not authenticated, redirect to login
    if (!authStore.isAuthenticated) {
      console.error("User not authenticated, redirecting to login.");
      throw new axios.Cancel("Request canceled due to unauthenticated user.");
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
  async (error) => {
    const authStore = useAuthStore(); // Access the auth store

    if (error.response?.status === 401) {
      console.error("Unauthorized, logging out and redirecting to login.");

      try {
        await authStore.logout(); // Clear auth state
      } catch (logoutError) {
        console.error("Failed to log out:", logoutError);
      }
    }

    return Promise.reject(error);
  }
);

export default api;
