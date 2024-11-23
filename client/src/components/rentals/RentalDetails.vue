<script setup lang="ts">
import { useRoute } from 'vue-router';
import { ref, onMounted } from 'vue';

// Access the route parameters to fetch the rental ID
const route = useRoute();
const rentalId = route.params.id;

// Reactive variable for rental details
const rental = ref<{ id: string; name: string; fullAddress: string } | null>(null);

onMounted(() => {
  // Replace this with an API call or state management to fetch the rental details by ID
  rental.value = {
    id: rentalId as string,
    name: 'Luxury Apartment',
    fullAddress: '456 Elm Street, Springfield, USA',
  };

});
</script>

<template>
  <v-container class="py-5">
    <v-card class="mx-auto" max-width="600">
      <v-card-title>
        <v-icon class="mr-3" icon="mdi-home-city-outline"></v-icon>
        <span class="text-h5">{{ rental?.name || 'Loading...' }}</span>
      </v-card-title>

      <v-divider></v-divider>

      <v-card-text>
        <v-row class="mb-3">
          <v-col cols="12">
            <v-icon class="mr-2" icon="mdi-map-marker-outline"></v-icon>
            <span>{{ rental?.fullAddress || 'No address available' }}</span>
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12">
            <!-- Additional rental details can go here -->
            <v-alert type="info" border="start">
              ID: {{ rental?.id }}
            </v-alert>
          </v-col>
        </v-row>
      </v-card-text>

      <v-card-actions>
        <v-btn variant="outlined" color="primary" @click="$router.back()">
          <v-icon left>mdi-arrow-left</v-icon>
          Go Back
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-container>
</template>

<style scoped>
.text-h5 {
  font-weight: bold;
}
</style>
