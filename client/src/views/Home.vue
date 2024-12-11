<script setup lang="ts">
import Map from "@/components/rentals/map/Map.vue";
import { useRoute, useRouter } from "vue-router";
import { onMounted, ref, watch, computed, provide } from "vue";
import LocationsGrid from "@/components/rentals/RentalGrid.vue";
import { getRentals } from "@/api/rentals";

// Vue Router
const router = useRouter();
const route = useRoute();

// Center and Rentals State
const currentCenter = ref({ lat: 36.8065, lng: 10.181667 }); // Default center (Tunis)
const rentals = ref([]);

// Filters State
const filters = ref({
  bedrooms: null,
  bathrooms: null,
  available: null,
});

// Provide Filter State
provide("filters", filters);
provide("updateFilters", (newFilters: Partial<typeof filters.value>) => {
  filters.value = { ...filters.value, ...newFilters };
});

// Fetch Rentals
onMounted(async () => {
  try {
   
    rentals.value  =  await getRentals();

  } catch (error) {
    console.error("Failed to fetch rentals:", error);
  }

  await router.isReady();

  currentCenter.value = {
    lat: route.query.lat ? parseFloat(route.query.lat) : 36.8065,
    lng: route.query.lng ? parseFloat(route.query.lng) : 10.181667,
  };
});

// Watch for Route Changes
watch(
  () => route.query,
  (query) => {
    currentCenter.value = {
      lat: query.lat ? parseFloat(query.lat) : 36.8065,
      lng: query.lng ? parseFloat(query.lng) : 10.181667,
    };
  }
);

// Computed: Filtered Rentals
const filteredRentals = computed(() => {
  return rentals.value.filter((rental) => {
    const matchesBedrooms =
      filters.value.bedrooms === null || rental.bedrooms === filters.value.bedrooms;
    const matchesBathrooms =
      filters.value.bathrooms === null || rental.bathrooms === filters.value.bathrooms;
    const matchesAvailability =
      filters.value.available === null || rental.available === filters.value.available;

    return matchesBedrooms && matchesBathrooms && matchesAvailability;
  });
});
</script>

<template>
  <div class="flex flex-row gap-4 w-full h-[90vh] overflow-y-hidden">
    <!-- Pass Filtered Rentals to Child Components -->
    <Map :rentals="filteredRentals" :center="currentCenter" />
    <LocationsGrid :rentals="filteredRentals" />
  </div>
</template>
