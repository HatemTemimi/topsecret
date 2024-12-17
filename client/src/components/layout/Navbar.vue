<script setup lang="ts">
import { useAuthStore } from "@/stores/authStore"; // Import the authStore
import { ref } from "vue";
import { VBtn, VIcon } from "vuetify/components";

const authStore = useAuthStore(); 

const logout = async  () => {
  authStore.logout(); 
};

const drawer = ref(null)

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

      <v-btn @click.stop="drawer = !drawer" class="sm:block" color="success" icon>
          <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <router-link to="/rental/new">
        <v-btn variant="elevated" class="mr-2 hidden md:block"  color="primary">
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
       <v-navigation-drawer
        v-model="drawer"
        temporary
      >
        <v-list-item
          prepend-avatar="https://randomuser.me/api/portraits/men/78.jpg"
          title="John Leider"
        ></v-list-item>

        <v-divider></v-divider>

        <v-list density="compact" nav>
          <v-list-item prepend-icon="mdi-view-dashboard" title="Home" value="home"></v-list-item>
          <v-list-item prepend-icon="mdi-forum" title="About" value="about"></v-list-item>
        </v-list>
      </v-navigation-drawer>

</template>

<style scoped>
</style>
