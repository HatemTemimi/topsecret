<script setup lang="ts">
import {VAutocomplete} from 'vuetify/components'
import {inject, ref, watch} from "vue";
import {getPlaceDetails, getPlacesGoogle} from "@/api/getPlaces";
import _ from 'lodash'
import AddressLocator from "@/components/form/AddressLocator.vue";
import {provide} from 'vue'

const model = ref(null)
const results = ref([])
const loading = ref(false)
//const marker = ref([36.8065, 10.181667])
const {marker, updateMarker} = inject('location');

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


const validateLocation = async (model) => {
  const details = await getPlaceDetails(model.place_id)
  const tmp = []
  tmp.push(details.geometry.location.lat, details.geometry.location.lng)
  marker.value = tmp
}
</script>

<template>
  <v-card title="Address"
   class="w-full h-1/3  p-8"
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
    <AddressLocator/>
  </v-card>
</template>

<style scoped>

</style>