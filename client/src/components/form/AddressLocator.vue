<template>
  <div class="h-96 w-full">
    <div class="h-full w-full">
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
        <l-marker @ready="()=> console.log('marker ready')"
                  @update:latLng="(val) => debouncedGetAddressFromLatLng(val.lat, val.lng)"
                  :draggable=true :lat-lng="currentCenter"></l-marker>
      </l-map>
    </div>
  </div>
</template>

<script setup>
import "leaflet/dist/leaflet.css";
import {LMap, LTileLayer, LMarker} from '@vue-leaflet/vue-leaflet';
import {onBeforeMount, ref} from "vue";
import {getAddress} from "@/api/getAddress.ts";
import {inject} from 'vue';
import {getAddressFromLatLng} from "@/api/getPlaces";

const {marker, updateMarker} = inject('location');
const currentCenter = ref(marker);

onBeforeMount(() => {
  currentCenter.value = marker.value;
  getAddress(currentCenter.value[1], currentCenter.value[0])
      .then((val) => {
        console.log(val);
      });
});

const zoom = ref(20);

// Debounce utility function
function debounce(func, delay) {
  let timeoutId;
  return function(...args) {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => func.apply(this, args), delay);
  };
}

// Wrap getAddressFromLatLng with debounce
const debouncedGetAddressFromLatLng = debounce((lat, lng) => {
  getAddressFromLatLng(lat, lng);
}, 500);

</script>

<style></style>
