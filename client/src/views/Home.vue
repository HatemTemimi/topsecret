<script setup lang="ts">
import Map from "@/components/rentals/map/Map.vue";
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref} from "vue";
import LocationsGrid from "@/components/rentals/RentalGrid.vue";

const router = useRouter();
const route = useRoute();
const currentCenter = ref({})

onMounted(async () => {
  await router.isReady();
  currentCenter.value = route.query
});


</script>

<template>
  <div class="flex flex-row gap-4 w-full max-h-[91vh] overflow-y-hidden">
    <Map :center="route.query"/>
    <LocationsGrid/>
  </div>
</template>
