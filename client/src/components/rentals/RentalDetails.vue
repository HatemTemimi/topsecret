<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';
import { ref, onMounted } from 'vue';
import type { Rental } from '@/models/rental';
import { getRentalById } from '@/api/rentals';
import "leaflet/dist/leaflet.css";
import { LMap, LTileLayer } from '@vue-leaflet/vue-leaflet';
import Marker from '@/components/rentals/map/marker/Marker.vue';
import Map from "@/components/rentals/map/Map.vue";
// Access the route parameters to fetch the rental ID
const route = useRoute();
const rentalId = route.params.id;

// Reactive variable for rental details
const rental = ref<Rental | null>(null);

// Reactive variable for error handling
const error = ref<string | null>(null);

const loading = ref(true)

const currentCenter = ref([36.8065, 10.181667])
const zoom = ref(12)

// Fetch rental details from the API
const fetchRental = async () => {
  try {
    loading.value = true; // Start loading
    const fetchedRental = await getRentalById(rentalId);
    if (fetchedRental?.geometry) {
      currentCenter.value = [
        parseFloat(fetchedRental.geometry.lat),
        parseFloat(fetchedRental.geometry.lng),
      ];
    }
    rental.value = fetchedRental;
  } catch (err) {
    console.error('Failed to fetch rental:', err);
    error.value = 'Failed to fetch rental details.';
  } finally {
    loading.value = false; // Stop loading
  }
};


// Lifecycle hook to fetch data on component mount
onMounted(async () => {
  if (!rentalId) {
    error.value = 'Invalid rental ID.';
    return;
  }
  await fetchRental();
});

</script>

<template>
    <div class="flex flex-col gap-4">
    <v-card class="h-[45vh] md:h-full">
      <v-card-text v-if="rental?.images && rental?.images?.length > 0">
          <v-row>
            <v-col cols="12">
              <v-carousel class="md:h-full"  hide-delimiters show-arrows="hover" :cycle="true">
                <v-carousel-item
                  v-for="(image, index) in rental?.images"
                  :key="index"
                >
                  <v-img aspect-ratio="1" rounded :src="image" :alt="'Image ' + index"></v-img>
                </v-carousel-item>
              </v-carousel>
            </v-col>
          </v-row>
        </v-card-text>
    </v-card>
    <div class="flex flex-col md:flex-row gap-4">
    <div class="flex flex-col gap-4">
    <v-card class="w-full lg:w-[50vw]">
      <v-card-title>
        <span class="text-h5">{{ rental?.name || 'Loading...' }}</span>
      </v-card-title>
      <v-card-subtitle>{{ rental?.address.city }}, {{ rental?.address.country }}</v-card-subtitle>

      <v-spacer></v-spacer>

      <v-card-text>
        <!-- Error Handling -->
        <v-alert v-if="error" type="error" border="start">
          {{ error }}
        </v-alert>

        <!-- Rental Details -->
        <template v-else-if="rental">
          <v-row class="mb-3">
            <v-col cols="12">
              <v-chip variant="elevated" label color="success">
                <span>{{ rental.price || 'price not available' }} TND</span>
              </v-chip>
            </v-col>
          </v-row>

          <v-spacer />

          <v-row class="mb-3">
            <v-col cols="12">
              <v-chip v-if="rental.status" variant="outlined" label color="success">
                <span>Available</span>
              </v-chip>
              <v-chip v-else variant="outlined" label color="error">
                <span>Unavailable</span>
              </v-chip>
            </v-col>
          </v-row>

          <v-spacer />
        </template>

        <!-- Loading State -->
        <template v-else>
          <v-alert type="info" border="start">
            Loading rental details...
          </v-alert>
        </template>
      </v-card-text>

      <v-divider></v-divider>

    </v-card>
    <v-card class="lg:w-[50vw] sm:w-full">
      <v-card-title>
        <span class="font-bold">
          Decription
        </span>
      </v-card-title>
      <v-card-text>
        {{ rental?.description }}
      </v-card-text>
    </v-card>
</div>
<v-card class="w-full h-96">
    <l-map
      v-if="rental"
      :use-global-leaflet="false"
      ref="map"
      v-model:zoom="zoom"
      :center="currentCenter"
    >
      <!-- Tile Layer -->
      <l-tile-layer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        layer-type="base"
        name="OpenStreetMap"
      ></l-tile-layer>
      
      <!-- Markers for Rentals -->
      <Marker :rental="rental" />
    </l-map>
    <div v-else>
      loading..
    </div>
</v-card>
</div>
</div>
</template>

<style scoped>
.text-h5 {
  font-weight: bold;
}
</style>
