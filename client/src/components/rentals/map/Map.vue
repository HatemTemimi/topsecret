<template>
  <v-container
    class="h-[60vh] lg:h-[100%] lg:w-[60vw] border border-gray-200 "
  >
  <Filters></Filters>
  <v-card class="h-full mt-4">
    <l-map
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
      <div v-for="(item, index) in props.rentals" :key="index">
        <Marker :rental="item" />
      </div>
    </l-map>
  </v-card>
  </v-container>
</template>


<script setup>
import "leaflet/dist/leaflet.css";
import { LMap, LTileLayer, LPolygon } from '@vue-leaflet/vue-leaflet';
import { ref, computed } from "vue";
import Marker from '@/components/rentals/map/marker/Marker.vue';
import Filters from '@/components/rentals/Filters.vue';

// Define props and refs
const props = defineProps(['center', 'rentals']);
const currentCenter = ref([]);
const zoom = ref(12);

// Set initial center of the map
if (props.center?.lat && props.center?.lng) {
  currentCenter.value = [props.center.lat, props.center.lng];
} else {
  currentCenter.value = [36.8065, 10.181667]; // Default center
}

// Polygon coordinates (example around the center point)
const polygonCoordinates = computed(() => [
  [currentCenter.value[0] + 0.01, currentCenter.value[1] - 0.01],
  [currentCenter.value[0] + 0.01, currentCenter.value[1] + 0.01],
  [currentCenter.value[0] - 0.01, currentCenter.value[1] + 0.01],
  [currentCenter.value[0] - 0.01, currentCenter.value[1] - 0.01],
]);

// Polygon style options
const polygonOptions = {
  color: 'red', // Border color
  fillColor: 'red', // Fill color
  fillOpacity: 0.5, // Fill transparency
};
</script>

<style>
/* Add custom Tailwind utility class */
.bg-custom-background {
  background-color: var(--v-theme-background);
}
</style>