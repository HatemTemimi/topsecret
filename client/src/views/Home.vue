<script setup lang="ts">
import Map from "@/components/rentals/map/Map.vue";
import { useRoute, useRouter } from "vue-router";
import { onMounted, ref, watch, computed, provide } from "vue";
import LocationsGrid from "@/components/rentals/RentalGrid.vue";
import FilterBar from "@/components/rentals/filters/Filterbar-v2.vue";
import { getRentals } from "@/api/rentals";

const route = useRoute();

// Center and Rentals State
const currentCenter = ref({ lat: 36.8065, lng: 10.181667 }); // Default center (Tunis)
const rentals = ref([]);

const loading = ref(true)

// Filters State
const filters = ref({
  bedrooms: null,
  bathrooms: null,
  available: null,
  type: null,
  standing: null,
  priceRange: null,
  rules: null,
  amenities: null,
  region: null,
});

// Provide Filter State
provide("filters", filters);
provide("updateFilters", (newFilters: Partial<typeof filters.value>) => {
  filters.value = { ...filters.value, ...newFilters };
});

provide("loading", loading)
provide("updateLoading", (val)=>{
  loading.value = val
})

const mapReady = ref(false);

// Fetch Rentals
onMounted(async () => {
  try {
    rentals.value = await getRentals();
    loading.value = false
  } catch (error) {
    console.error("Failed to fetch rentals:", error);
  } 
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

// Price range helper function
const parsePriceRange = (price: number, range: string | null) => {
  if (!range) return true;
  switch (range) {
    case "Under 500 TND":
      return price < 500;
    case "500-1000 TND":
      return price >= 500 && price <= 1000;
    case "1000-2000 TND":
      return price > 1000 && price <= 2000;
    case "Above 2000 TND":
      return price > 2000;
    default:
      return true;
  }
};

// Computed: Filtered Rentals
const filteredRentals = computed(() => {
  return rentals.value.filter((rental) => {
    const matchesBedrooms =
      filters.value.bedrooms === null || rental.bedrooms === filters.value.bedrooms;

    const matchesBathrooms =
      filters.value.bathrooms === null || rental.bathrooms === filters.value.bathrooms;

    const matchesAvailability =
      filters.value.available === null || rental.available === filters.value.available;

    const matchesType =
      filters.value.type === null || rental.type === filters.value.type;

    const matchesStanding =
      filters.value.standing === null || rental.standing === filters.value.standing;

    const matchesPriceRange = parsePriceRange(rental.price, filters.value.priceRange);

    const matchesRules =
      filters.value.rules === null ||
      (Array.isArray(filters.value.rules) && filters.value.rules.every((rule) => rental.rules[rule]));

    const matchesAmenities =
      filters.value.amenities === null ||
      (Array.isArray(filters.value.amenities) &&
        filters.value.amenities.every((amenity) => rental.amenities[amenity]));

    const matchesRegion =
      filters.value.region === null || filters.value.region.some((coord: any) => {
        return (
          parseFloat(rental.geometry.lat) === coord[0] &&
          parseFloat(rental.geometry.lng) === coord[1]
        );
      });

    return (
      matchesBedrooms &&
      matchesBathrooms &&
      matchesAvailability &&
      matchesType &&
      matchesStanding &&
      matchesPriceRange &&
      matchesRules &&
      matchesAmenities &&
      matchesRegion
    );
  });
});
</script>

<template>
  <div class="w-full">
    <!-- FilterBar Component -->
    <FilterBar />
    <div
      class="flex flex-col lg:flex-row lg:gap-4 gap-24 w-full lg:h-[80vh] h-full md:overflow-y-hidden mt-4"
    >
      <!-- Pass Filtered Rentals to Child Components -->
      <Map :rentals="filteredRentals" :center="currentCenter" />
      <LocationsGrid :rentals="filteredRentals" />
    </div>
  </div>
</template>

<style scoped>
.text-muted {
  color: #6c757d;
}
</style>
