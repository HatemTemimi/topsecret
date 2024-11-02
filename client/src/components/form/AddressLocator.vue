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
        <l-marker 
          @ready="() => console.log('marker ready')"
          @update:latLng="(val) => debouncedGetAddressFromLatLng(val.lat, val.lng)"
          :draggable="true" 
          :lat-lng="currentCenter">
        </l-marker>
      </l-map>
    </div>
  </div>
</template>

<script setup>
import "leaflet/dist/leaflet.css";
import { LMap, LTileLayer, LMarker } from '@vue-leaflet/vue-leaflet';
import { onBeforeMount, ref, inject } from "vue";
import { getAddress } from "@/api/getAddress.ts";
import { getAddressFromLatLng } from "@/api/getPlaces";

const { marker, updateMarker } = inject('location');
const updateInfo = inject('updateInfo'); // Inject updateInfo to update the state in the main component

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
const debouncedGetAddressFromLatLng = debounce(async (lat, lng) => {
  try {
    const response = await getAddressFromLatLng(lat, lng);
    console.log(response);

    // Parse the address components and map to the fields
    const addressData = {
      streetNumber: '',
      street: '',
      city: '',
      country: '',
      fullAddress: '',
    };

    // Extract data based on types
    response.address_components.forEach((component) => {
      const types = component.types;
      if (types.includes("street_number")) addressData.streetNumber = component.long_name;
      if (types.includes("route")) addressData.street = component.long_name;
      if (types.includes("locality")) addressData.city = component.long_name;
      if (types.includes("country")) addressData.country = component.long_name;
    });

    // Set full address if it's available in response
    addressData.fullAddress = response.formatted_address

    // Call updateInfo to update the main component state
    updateInfo(addressData);
  } catch (error) {
    console.error("Error fetching address:", error);
  }
}, 500);
</script>

<style></style>
