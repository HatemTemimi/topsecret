import { defineStore } from "pinia";
import { useCookies } from "@vueuse/integrations/useCookies";
import axios from "axios";
import router from "@/router"; 


export const useAuthStore = defineStore("auth", {
  state: () => ({
    isAuthenticated: false, // Authentication status
    user: null as { email: string; firstName: string; lastName: string; id: string; role: string } | null, // User details
  }),
  actions: {
    async login(credentials: { email: string; password: string }) {

      try {
        const cookies = useCookies(); // Access cookies

        await axios.post("http://localhost:3001/api/auth/login", credentials, {
          withCredentials: true, // Include cookies in the request
        });

        // Load user information from cookies
        const encodedUserInfo = cookies.get("user_info");
        if (encodedUserInfo) {
          this.isAuthenticated = true;
          this.user = encodedUserInfo;
        } else {
          throw new Error("User information not found in cookies.");
        }
      } catch (error) {
        console.error("Login failed:", error);
        throw error;
      }
    },

    async loadSession() {

      try {
        const cookies = useCookies(); // Access cookies

        // Check for user info in cookies
        const encodedUserInfo = cookies.get("user_info");
        if (encodedUserInfo) {
          this.isAuthenticated = true;
          this.user = encodedUserInfo;
        } else {
          this.isAuthenticated = false;
          this.user = null;
          router.push("/user/login"); // Redirect to login page
        }
      } catch (error) {
        console.error("Auth check failed:", error);
        this.isAuthenticated = false;
        this.user = null;
      }
    },

    async googleLogin(idToken: string) {
      try {
        const cookies = useCookies(); // Access cookies

        const resp = await axios.post(
          "http://localhost:3001/api/auth/google/login",
          { idToken },
          { withCredentials: true }
        );
        console.log(resp.data.url)
        //router.push(resp.data.url)
        window.location.href = resp.data.url;
        /*
        const encodedUserInfo = cookies.get("user_info");
        if (encodedUserInfo) {
          this.isAuthenticated = true;
          this.user = encodedUserInfo
        } else {
          throw new Error("User information not found in cookies.");
        }
        */
      } catch (error) {
        console.error("Google Login failed:", error.message);
        throw error;
      }
    },

    async logout() {
      try {
        const cookies = useCookies(); // Access cookies

        await axios.post("http://localhost:3001/api/auth/logout", {}, { withCredentials: true });
        cookies.remove("user_info");
        this.isAuthenticated = false;
        this.user = null;
        router.push("/user/login");
      } catch (error) {
        console.error("Logout failed:", error);
        throw error;
      }
    },
  },
});
