<template>
  <div>
    <l-marker :lat-lng="latlng">
      <l-popup>
        <PopupContent :rental="props.rental"/>
      </l-popup :rental="props.rental">
    </l-marker>
  </div>
</template>

<script setup>
import { LMarker, LPopup } from '@vue-leaflet/vue-leaflet'
import PopupContent from '@/components/map/marker/PopupContent.vue'
import { onMounted, computed } from 'vue';

const props = defineProps(['rental'])


// Compute latlng from the rental prop
const latlng = computed(() => [props.rental.lat, props.rental.lng])

</script>