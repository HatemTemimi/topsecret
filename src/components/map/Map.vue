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
        <div v-for="item in latlng">
          <Marker :latlng="item"/>
        </div>
      </l-map>
    </div>
  </div>
</template>

<script setup>
import "leaflet/dist/leaflet.css";
import {LMap, LTileLayer} from '@vue-leaflet/vue-leaflet'
import {onBeforeMount, onMounted, ref} from "vue";
import Marker from '@/components/map/marker/Marker.vue'

const props = defineProps(['center'])
const currentCenter = ref([])

onBeforeMount(() => {
  if (props.center.lat && props.center.lng) {
    currentCenter.value.push(props.center.lng, props.center.lat)
  } else {
    currentCenter.value = [36.8065, 10.181667]
  }
})


const latlng = ref([
  [36.8065, 10.181667],
  [36.8065, 10.172667],
  [36.7065, 10.162667],
  [36.7045, 10.162667],
  [36.8045, 10.162667],
  [36.8075, 10.171663],
])

const zoom = ref(12)

</script>

<style></style>
