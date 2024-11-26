<script setup lang="ts">
import { useAuthStore } from "@/stores/authStore"; // Import the authStore
import { VBtn, VIcon } from "vuetify/components";
import { useRouter } from "vue-router";

const authStore = useAuthStore(); // Access Pinia authStore
const router = useRouter();

// Logout function
const logout = () => {
  authStore.logout(); // Clear auth state
  router.push("/user/login"); // Redirect to login page
};
</script>

<template>
  <v-app-bar rounded>
    <v-app-bar-title>
      Darmap
    </v-app-bar-title>

    <v-spacer></v-spacer>

    <!-- Navbar for Authenticated Users -->
    <template v-if="authStore.isAuthenticated">
      <!-- Greeting -->
      <span>Hello, {{ authStore.user?.firstName }}</span>

      <!-- Links available to authenticated users -->
      <router-link to="/rentals/search">
        <v-btn icon>
          <v-icon>mdi-magnify</v-icon>
        </v-btn>
      </router-link>

      <router-link to="/rental/create">
        <v-btn icon>
          <v-icon>mdi-plus</v-icon>
        </v-btn>
      </router-link>

      <router-link to="/rentals">
        <v-btn icon>
          <v-icon>mdi-home</v-icon>
        </v-btn>
      </router-link>

      <!-- Logout Button -->
      <v-btn @click="logout" icon>
        <v-icon>mdi-logout</v-icon>
      </v-btn>
    </template>

    <!-- Navbar for Unauthenticated Users -->
    <template v-else>
      <router-link to="/user/login">
        <v-btn icon>
          <v-icon>mdi-login</v-icon>
        </v-btn>
      </router-link>

      <router-link to="/user/register">
        <v-btn icon>
          <v-icon>mdi-account-plus</v-icon>
        </v-btn>
      </router-link>
    </template>
  </v-app-bar>
</template>

<style scoped>
</style>
