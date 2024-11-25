<template>
    <div class="map-container">
      <Filters></Filters>
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
</template>

<script setup>
import "leaflet/dist/leaflet.css";
import {LMap, LTileLayer} from '@vue-leaflet/vue-leaflet'
import { ref } from "vue";
import Marker from '@/components/rentals/map/marker/Marker.vue'

import Filters from '@/components/rentals/Filters.vue'

// Define props and refs
const props = defineProps(['center', 'rentals'])
const currentCenter = ref([])
const zoom = ref(12)
const rentals = ref([])

// Set initial center of the map
if (props.center?.lat && props.center?.lng) {
  currentCenter.value = [props.center.lat, props.center.lng]
} else {
  currentCenter.value = [36.8065, 10.181667] // Default center
}


</script>

<style>
.map-container {
  height: 100%;
  width: 60vw;
}
</style>
