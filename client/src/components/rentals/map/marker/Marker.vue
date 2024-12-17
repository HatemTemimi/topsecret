<template>
  <div>
    <l-marker :lat-lng="latlng">
      <l-popup class="p-0 m-0 min-w-44">
        <PopupContent :rental="props.rental"/>
      </l-popup :rental="props.rental">
    </l-marker>
  </div>
</template>

<script setup>
import { LMarker, LPopup } from '@vue-leaflet/vue-leaflet'
import PopupContent from '@/components/rentals/map/marker/PopupContent.vue'
import { computed } from 'vue';

const props = defineProps(['rental'])

const latlng = computed(() => [props.rental.geometry.lat, props.rental.geometry.lng])

</script>