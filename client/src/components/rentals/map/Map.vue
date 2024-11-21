<template>
  <div>
    <div class="h-full w-[65vw]">
      <l-map
          :use-global-leaflet="false"
          ref="map"
          v-model:zoom="zoom"
          :center="currentCenter"
      >
        <l-tile-layer
            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            layer-type="base"
            name="OpenStreetMap"
        ></l-tile-layer>
        <div v-for="(item, index) in props.rentals" :key="index">
          <Marker :rental="item"/>
        </div>
      </l-map>
    </div>
  </div>
</template>

<script setup>
import "leaflet/dist/leaflet.css";
import {LMap, LTileLayer} from '@vue-leaflet/vue-leaflet'
import {onMounted, ref} from "vue";
import axios from 'axios';
import Marker from '@/components/rentals/map/marker/Marker.vue'

// Define props and refs
const props = defineProps(['center', 'rentals'])
const currentCenter = ref([])
const rentals = ref([]) // Holds coordinates for markers
const zoom = ref(12)

// Set initial center of the map
if (props.center?.lat && props.center?.lng) {
  currentCenter.value = [props.center.lat, props.center.lng]
} else {
  currentCenter.value = [36.8065, 10.181667] // Default center
}

// Fetch rentals data on component mount
onMounted(() => {
  rentals.value = props.rentals
})
</script>

<style></style>
