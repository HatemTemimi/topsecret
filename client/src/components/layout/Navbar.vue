<script setup lang="ts">
import { useAuthStore } from "@/stores/authStore"; // Import the authStore
import { VBtn, VIcon } from "vuetify/components";
import { useRouter } from "vue-router";

const authStore = useAuthStore(); // Access Pinia authStore
const router = useRouter();

// Logout function
const logout = async  () => {
  await router.isReady()
  authStore.logout(); // Clear auth state
  router.push("/user/login"); // Redirect to login page
};
</script>

<template>
  <v-app-bar rounded
  color="surface"
  >
    <v-app-bar-title>
      <router-link to="/rentals">
        <v-btn>
          Dar/win
        </v-btn>
        </router-link>
    </v-app-bar-title>

    <template v-if="authStore.isAuthenticated">
      <router-link to="/rentals/search">
        <v-btn color="success" icon>
          <v-icon>mdi-magnify</v-icon>
        </v-btn>
      </router-link>

      <router-link to="/rental/new">
        <v-btn color="success" icon>
          <v-icon>mdi-plus</v-icon>
        </v-btn>
      </router-link>
      <router-link  to="/rentals/user">
        <v-btn color="success" icon>
          <v-icon>mdi-arrange-bring-forward</v-icon>
        </v-btn>
      </router-link>

      <div class="ml-2">
      <span>Hello, {{ authStore.user?.firstName }}</span>

      <!-- Logout Button -->
      <v-btn @click="logout" icon>
        <v-icon>mdi-logout</v-icon>
      </v-btn>
      </div>

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
