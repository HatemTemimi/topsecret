<script setup lang="ts">
import {VAutocomplete} from 'vuetify/components'
import {ref, watch} from "vue";
import {getPlaceDetails, getPlacesGoogle} from "@/api/getPlaces";
import _ from 'lodash'
import AddressLocator from "@/components/form/AddressLocator.vue";
import {provide} from 'vue'

const model = ref(null)
const results = ref([])
const loading = ref(false)
const marker = ref([36.8065, 10.181667])

// Debounced search function using watch
let timeoutId: number | null = null;
const debouncedSearch = (val: string) => {
  if (timeoutId) clearTimeout(timeoutId);
  timeoutId = setTimeout(async () => {
    if (!_.isEmpty(val) && model.value === null) {
      results.value = await getPlacesGoogle(val)
    }
  }, 500);
}

watch(model, debouncedSearch, {immediate: true}); // Call search on initial render

function updateMarker(val) {
  console.log("updating marker..")
  marker.value = val
}

provide('location', {
  marker,
  updateMarker
})

const validateLocation = async (model) => {
  console.log(model)
  const details = await getPlaceDetails(model.place_id)
  console.log(details)
 /* const lat = model.value.center[1]
  model.value.center[1] = model.value.center[0]
  model.value.center[0] = lat
  */
  //marker.value.push(model.value.geometry.location)

  const tmp = []
  tmp.push(details.geometry.location.lat, details.geometry.location.lng)
  
  console.log("marker")

  marker.value = tmp

  console.log(marker.value)
  //set marker position on locator map

}
</script>

<template>
  <v-card class="w-full h-1/3  p-8"
          density="comfortable"
          :loading="loading"
  >
  <v-autocomplete
        clearable
        focused
        hide-no-data
        label="Choose location for results.."
        :items="results"
        v-model="model"
        item-title="description"
        no-filter
        return-object
        @focus="loading=true"
        @blur="loading=false"
        @update:search="debouncedSearch"
        validate-on="blur"
        @update:modelValue="(model)=>validateLocation(model)"
    ></v-autocomplete>
    <AddressLocator :marker="marker" :updateMarker="updateMarker"/>
  </v-card>
</template>

<style scoped>

</style>