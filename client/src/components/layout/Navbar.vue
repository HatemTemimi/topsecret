<script setup lang="ts">
import { useAuthStore } from "@/stores/authStore"; // Import the authStore
import { VBtn, VIcon } from "vuetify/components";

const authStore = useAuthStore(); 

const logout = async  () => {
  authStore.logout(); 
};

</script>

<template>
  <v-app-bar rounded color="surface">
    <v-app-bar-title>
      <router-link to="/rentals">
        <v-btn>
          Dar/win
        </v-btn>
        </router-link>
    </v-app-bar-title>

    <template v-if="authStore.isAuthenticated">

      <!--
      <router-link to="/rentals/search">
        <v-btn color="success" icon>
          <v-icon>mdi-magnify</v-icon>
        </v-btn>
      </router-link>
-->

      <router-link to="/rental/new">
        <v-btn variant="elevated" class="mr-2"  color="primary">
          Create your rental
          <v-icon>mdi-plus</v-icon>
        </v-btn>
      </router-link>
      <div class="ml-2">

        <v-menu
      open-on-hover
    >
      <template v-slot:activator="{ props }">
        <v-btn
          color="success"
          v-bind="props"
          class="mr-2"
        >
      Hello, {{ authStore.user?.firstName }}
        </v-btn>
      </template>

      <v-list>
        <v-list-item>
        <router-link  to="/rentals/user">
          <v-btn>
            My Rentals
          </v-btn>
        </router-link>
        </v-list-item>

        <v-list-item>
          <v-btn @click="logout">
            Logout
          </v-btn>
        </v-list-item>

      </v-list>
    </v-menu>
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
