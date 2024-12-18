<template>
  <v-card class="h-full flex items-center justify-center">
     <v-progress-circular
     class="absolute"
      v-if="loading || !mapReady"
      size="100"
      width="50"
      color="success"
      indeterminate
    ></v-progress-circular>
    <l-map
      v-if="!loading"
      v-on:ready="setMapReady"
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
</template>


<script setup>
import "leaflet/dist/leaflet.css";
import { LMap, LTileLayer } from '@vue-leaflet/vue-leaflet';
import { inject, ref } from "vue";
import Marker from '@/components/rentals/map/marker/Marker.vue';

// Define props and refs
const props = defineProps(['rentals']);
const currentCenter = ref([36.8065, 10.181667]);
const zoom = ref(12);

const mapReady = ref(false)

const setMapReady = ()=>{
    mapReady.value = true;
}

const loading = inject("loading")

</script>
