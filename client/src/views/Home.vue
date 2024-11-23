<script setup lang="ts">
import Map from "@/components/rentals/map/Map.vue";
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import LocationsGrid from "@/components/rentals/RentalGrid.vue";
import axios from "axios";
import { getRentals } from "@/api/rentals";
import Filters from "@/components/rentals/Filters.vue"

const router = useRouter();
const route = useRoute();
const currentCenter = ref({})
const rentals = ref([])

onMounted(async () => {
    try {
    // Fetch rentals from the API
    const data = await getRentals()
    // Extract lat and lng from each rental and add to latlng array
    rentals.value = data
    console.log(rentals.value)
  } catch (error) {
    console.error("Failed to fetch rentals:", error)
  }
  await router.isReady();
  currentCenter.value = route.query
});


</script>

<template>
  <div class="flex flex-row gap-4 w-full max-h-[90vh] overflow-y-hidden">
    <Map :rentals="rentals" :center="route.query"/>
    <LocationsGrid :rentals="rentals"/>
  </div>
</template>
