<template>
  <div>
    <l-marker :lat-lng="latlng">
      <l-popup>
        <PopupContent/>
      </l-popup>
    </l-marker>

  </div>
</template>

<script setup>
import { LMarker, LPopup } from '@vue-leaflet/vue-leaflet'
import PopupContent from '@/components/map/marker/PopupContent.vue'

const props = defineProps(['latlng'])

</script>