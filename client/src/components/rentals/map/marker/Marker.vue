<template>
  <div>
    <l-marker @click="dialog=true" :lat-lng="latlng">
        <l-tooltip :options="{opacity: 1}">
          <PopupContent :rental="props.rental"/>
        </l-tooltip>
    </l-marker>
  </div>
  <DetailsModal v-if="dialog"/>

</template>

<script setup>
import { LMarker, LTooltip } from '@vue-leaflet/vue-leaflet'
import PopupContent from '@/components/rentals/map/marker/PopupContent.vue'
import { computed, provide, ref } from 'vue';
import DetailsModal from '../../modal/detailsModal.vue';

const props = defineProps(['rental'])

const dialog = ref(false)

const close = (val)=>{
  dialog.value = val
}

provide("dialog", dialog)
provide("close", close)

const latlng = computed(() => [props.rental.geometry.lat, props.rental.geometry.lng])

</script>