<template>
  <v-container
    class="h-[90vh] lg:h-[100%] lg:w-full border border-gray-200 rounded-lg"
  >
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
  </v-container>
</template>


<script setup>
import "leaflet/dist/leaflet.css";
import { LMap, LTileLayer } from '@vue-leaflet/vue-leaflet';
import { inject, ref } from "vue";
import Marker from '@/components/rentals/map/marker/Marker.vue';

// Define props and refs
const props = defineProps(['center', 'rentals']);
const currentCenter = ref([36.8065, 10.181667]);
const zoom = ref(12);

const mapReady = ref(false)

const setMapReady = ()=>{
    mapReady.value = true;
}

/*
const loading = ref(true)
const updateLoading = ()=>{
  console.log('map ready..')
  loading.value = false
}
  */
const loading = inject("loading")

/*
if (props.center?.lat && props.center?.lng) {
  currentCenter.value = [props.center.lat, props.center.lng];
} else {
  currentCenter.value = ; 
}
  */

</script>
