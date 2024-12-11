<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';
import { ref, onMounted } from 'vue';
import type { Rental } from '@/models/rental';
import { getRentalById } from '@/api/rentals';

// Access the route parameters to fetch the rental ID
const route = useRoute();
const router = useRouter();
const rentalId = route.params.id;

// Reactive variable for rental details
const rental = ref<Rental | null>(null);

// Reactive variable for error handling
const error = ref<string | null>(null);

// Fetch rental details from the API
const fetchRental = async () => {
  try {
    rental.value = await getRentalById(rentalId);
    console.log(rental.value)
  } catch (err) {
    console.error('Failed to fetch rental:', err);
    error.value = 'Failed to fetch rental details.';
  }
};

// Lifecycle hook to fetch data on component mount
onMounted(() => {
  if (!rentalId) {
    error.value = 'Invalid rental ID.';
    return;
  }
  fetchRental();
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
        <!-- Error Handling -->
        <v-alert v-if="error" type="error" border="start">
          {{ error }}
        </v-alert>

        <!-- Rental Details -->
        <template v-else-if="rental">
          <!-- Address -->
          <v-row class="mb-3">
            <v-col cols="12">
              <v-icon class="mr-2" icon="mdi-map-marker-outline"></v-icon>
              <span>{{ rental.fullAddress || 'Address not available' }}</span>
            </v-col>
          </v-row>

          <v-spacer />

          <!-- City -->
          <v-row class="mb-3">
            <v-col cols="12">
              <v-icon class="mr-2" icon="mdi-city"></v-icon>
              <span>{{ rental.city || 'City not available' }}</span>
            </v-col>
          </v-row>

          <v-spacer />

          <!-- Country -->
          <v-row class="mb-3">
            <v-col cols="12">
              <v-icon class="mr-2" icon="mdi-earth"></v-icon>
              <span>{{ rental.country || 'Country not available' }}</span>
            </v-col>
          </v-row>

          <v-spacer />

          <!-- Street -->
          <v-row class="mb-3">
            <v-col cols="12">
              <v-icon class="mr-2" icon="mdi-road-variant"></v-icon>
              <span>{{ rental.street || 'Street not available' }}</span>
            </v-col>
          </v-row>

          <v-spacer />

          <!-- Status -->
          <v-row class="mb-3">
            <v-col cols="12">
              <v-icon class="mr-2" icon="mdi-check-circle-outline"></v-icon>
              <span>{{ rental.status ? 'Available' : 'Unavailable' }}</span>
            </v-col>
          </v-row>

          <v-spacer />

          <!-- Agreement -->
          <v-row class="mb-3">
            <v-col cols="12">
              <v-icon class="mr-2" icon="mdi-handshake-outline"></v-icon>
              <span>{{ rental.agree ? 'Agreed to terms' : 'Not agreed' }}</span>
            </v-col>
          </v-row>
        </template>

        <!-- Loading State -->
        <template v-else>
          <v-alert type="info" border="start">
            Loading rental details...
          </v-alert>
        </template>
      </v-card-text>

      <v-divider></v-divider>

      <!-- Images Carousel -->
      <v-card-text v-if="rental?.images && rental?.images?.length > 0">
        <v-row>
          <v-col cols="12">
            <h4>Images</h4>
            <v-carousel height="200" show-arrows>
              <v-carousel-item
                v-for="(image, index) in rental?.images"
                :key="index"
              >
                <v-img :src="image" :alt="'Image ' + index"></v-img>
              </v-carousel-item>
            </v-carousel>
          </v-col>
        </v-row>
      </v-card-text>

      <v-card-actions>
        <v-btn variant="outlined" color="primary" @click="router.back()">
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
