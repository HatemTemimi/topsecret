import axios from "axios";
import router from "@/router"; // Import the router directly

// Create Axios instance
const api = axios.create({
  baseURL: "http://localhost:3001/", // Replace with your API base URL
});

// Add token to all requests except login/register
api.interceptors.request.use(
  (config) => {
    if (
      config.url?.includes("/authenticate") || 
      config.url?.includes("/register")
    ) {
      // Skip token for login and register endpoints
      return config;
    }

    const token = localStorage.getItem("token");

    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    } else {
      console.error("No token found, redirecting to login.");
      router.push("/login"); // Redirect to login if no token
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Handle 401 Unauthorized responses
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      console.error("Unauthorized, redirecting to login.");
      router.push("/user/login"); // Redirect to login if unauthorized
    }
    return Promise.reject(error);
  }
);

export default api;
