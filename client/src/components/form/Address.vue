<script setup lang="ts">
import {VAutocomplete} from 'vuetify/components'
import {ref, watch} from "vue";
import {getPlacesGoogle} from "@/api/getPlaces";
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
    if (!_.isEmpty(val)) {
      results.value = await getPlaces(val)
    }
  }, 500);
}

watch(model, debouncedSearch, {immediate: true}); // Call search on initial render

function updateMarker(val) {
  marker.value = val
}

provide('location', {
  marker,
  updateMarker
})
const validateLocation = () => {
  const lat = model.value.center[1]
  model.value.center[1] = model.value.center[0]
  model.value.center[0] = lat
  marker.value = model.value.center
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
        label="Choose your area.."
        :items="results"
        v-model="model"
        item-title="place_name"
        no-filter
        return-object
        @focus="loading=true"
        @blur="loading=false"
        @update:search="debouncedSearch"
        validate-on="blur"
        @update:modelValue="validateLocation()"
    ></v-autocomplete>
    <AddressLocator :marker="marker" :updateMarker="updateMarker"/>
  </v-card>
</template>

<style scoped>

</style>