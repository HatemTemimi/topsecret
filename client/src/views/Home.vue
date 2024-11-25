<script setup lang="ts">
import Map from "@/components/rentals/map/Map.vue";
import { useRoute, useRouter } from "vue-router";
import { onMounted, ref, watch } from "vue";
import LocationsGrid from "@/components/rentals/RentalGrid.vue";
import { getRentals } from "@/api/rentals";

const router = useRouter();
const route = useRoute();
const currentCenter = ref({ lat: 36.8065, lng: 10.181667 }); // Default center (Tunis)
const rentals = ref([]);

onMounted(async () => {
  try {
    // Fetch rentals from the API
    const data = await getRentals();
    rentals.value = data; // Assign fetched rentals
    console.log(rentals.value);
  } catch (error) {
    console.error("Failed to fetch rentals:", error);
  }

  await router.isReady();

  // Update the map's center using route query or fallback to default
  currentCenter.value = {
    lat: route.query.lat ? parseFloat(route.query.lat) : 36.8065,
    lng: route.query.lng ? parseFloat(route.query.lng) : 10.181667,
  };
});

// Watch for route changes and update center dynamically
watch(
  () => route.query,
  (query) => {
    currentCenter.value = {
      lat: query.lat ? parseFloat(query.lat) : 36.8065,
      lng: query.lng ? parseFloat(query.lng) : 10.181667,
    };
  }
);
</script>

<template>
  <div class="flex flex-row gap-4 w-full h-[90vh] overflow-y-hidden">
    <Map :rentals="rentals" :center="currentCenter" />
    <LocationsGrid :rentals="rentals" />
  </div>
</template>
