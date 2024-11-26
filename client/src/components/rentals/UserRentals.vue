<script setup lang="ts">
import { getRentalsByUserId } from "@/api/rentals";
import CardTify from "@/components/rentals/RentalCard.vue";
import type { Rental } from "@/models/rental";
import { onMounted, ref } from "vue";
import { useAuthStore } from "@/stores/authStore"; // Import the Pinia store

const rentals = ref<Rental[] | null>(null);

// Access the auth store to get the userId
const authStore = useAuthStore();

onMounted(async () => {
  if (authStore.user?.id) {
    try {
      rentals.value = await getRentalsByUserId(authStore.user.id);
    } catch (error) {
      console.error("Failed to fetch rentals:", error);
    }
  } else {
    console.warn("No user ID found in the auth store.");
  }
});
</script>


<template>
  <v-card class="mx-auto overflow-auto w-full">
    <v-container fluid>
      <v-row v-if="rentals && rentals.length > 0">
        <!-- Iterate over rentals to display each as a card -->
        <v-col
          v-for="rental in rentals"
          :key="rental.name"
          cols="12"
          md="6"
          lg="4"
        >
          <CardTify
            :rental="rental"
          />
        </v-col>
      </v-row>
      <v-row v-else>
            You did not add any rentals yet
      </v-row>
    </v-container>
  </v-card>
</template>

<style scoped>
</style>