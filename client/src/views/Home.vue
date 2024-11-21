<script setup lang="ts">
import Map from "@/components/rentals/map/Map.vue";
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import LocationsGrid from "@/components/rentals/RentalGrid.vue";
import axios from "axios";

const router = useRouter();
const route = useRoute();
const currentCenter = ref({})
const rentals = ref([])

onMounted(async () => {
    try {
    // Fetch rentals from the API
    const response = await axios.get('http://localhost:3001/api/rental/list') // Update with your actual API endpoint
    const data = response.data
    // Extract lat and lng from each rental and add to latlng array
    rentals.value = data

  } catch (error) {
    console.error("Failed to fetch rentals:", error)
  }
  await router.isReady();
  currentCenter.value = route.query
});


</script>

<template>
  <div class="flex flex-row gap-4 w-full max-h-[91vh] overflow-y-hidden">
    <Map :rentals="rentals" :center="route.query"/>
    <LocationsGrid/>
  </div>
</template>
