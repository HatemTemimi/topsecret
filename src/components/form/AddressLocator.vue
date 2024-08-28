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
        <div v-for="item in latlng">
          <l-marker @ready="()=> console.log('marker ready')"
                    @update:latLng="(val)=> console.log('update happened', val)"
                    :draggable=true :lat-lng="item"></l-marker>
        </div>
      </l-map>
    </div>
  </div>
</template>

<script setup>
import "leaflet/dist/leaflet.css";
import {LMap, LTileLayer, LMarker} from '@vue-leaflet/vue-leaflet'
import {onBeforeMount, ref} from "vue";

const currentCenter = ref([])

onBeforeMount(() => {
  currentCenter.value = [36.8065, 10.181667]
})


const latlng = ref([
  [36.8065, 10.181667],
])

const zoom = ref(20)

</script>

<style></style>