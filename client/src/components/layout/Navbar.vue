<script setup lang="ts">
import { VBtn, VIcon } from "vuetify/components";
import { ref } from "vue";
import { useRouter } from "vue-router";

// State for authentication
const isAuthenticated = ref(!!localStorage.getItem("token")); // Check if token exists in sessionStorage
const router = useRouter();

console.log(isAuthenticated.value)

// Logout function
const logout = () => {
  sessionStorage.removeItem("token"); // Clear the token
  isAuthenticated.value = false; // Update authentication state
  router.push("/user/login"); // Redirect to login page
};
</script>

<template>
  <v-app-bar rounded>
    <v-app-bar-title>
      Darmap
    </v-app-bar-title>

    <v-spacer></v-spacer>


    <!-- Links available only when authenticated -->
    <template v-if="isAuthenticated">
    <!-- Links available to all users -->
    <router-link to="/rentals/search">
      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>
    </router-link>

      <router-link to="rentals/create">
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

    <!-- Login and Register links for unauthenticated users -->
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
