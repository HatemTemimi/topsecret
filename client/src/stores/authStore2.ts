import { defineStore } from "pinia";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    token: sessionStorage.getItem("token") || null, // JWT token
    isAuthenticated: !!sessionStorage.getItem("token"),
    user: null as { email: string; firstName: string; lastName: string; id: string; role: string } | null, // User details
  }),
  actions: {
    login(token: string, user: any) {
      this.token = token;
      this.isAuthenticated = true;
      this.user = user;
      sessionStorage.setItem("token", token); // Store token in sessionStorage
      sessionStorage.setItem("user", JSON.stringify(user)); // Store user details in sessionStorage
    },
    loadSession() {
      // Load session from sessionStorage
      const token = sessionStorage.getItem("token");
      const user = sessionStorage.getItem("user");

      this.token = token;
      this.isAuthenticated = !!token;
      this.user = user ? JSON.parse(user) : null;
    },
    logout() {
      this.token = null;
      this.isAuthenticated = false;
      this.user = null;
      sessionStorage.removeItem("token"); // Clear token from sessionStorage
      sessionStorage.removeItem("user"); // Clear user details
    },
  },
});
