<script setup lang="ts">
import {VAutocomplete} from 'vuetify/components'
import {MapBoxProvider} from 'leaflet-geosearch';
import {ref} from "vue";
import {useRouter} from "vue-router";
import axios from "axios";


const model = ref(null)
const results = ref([])
const loading = ref(false)

const router = useRouter()

const provider = new MapBoxProvider({
  params: {
    access_token: import.meta.env.VITE_MAPBOX_TOKEN,
    country: 'tn',
    autofill: true,
  },
});

//const url = 'https://api.mapbox.com/autofill/v1/suggest/aut?types=address&access_token=pk.eyJ1IjoiZXhhbXBsZXMiLCJhIjoiY2p0MG01MXRqMW45cjQzb2R6b2ptc3J4MSJ9.zA2W0IkI0c6KaAhJfk9bWg&streets=true&language=en&session_token=0fb40925-45fc-4971-8919-4126168062b6&proximity=ip'

axios.get('https://api.mapbox.com/autofill/v1/suggest/aut?types=address&access_token=pk.eyJ1IjoiZXhhbXBsZXMiLCJhIjoiY2p0MG01MXRqMW45cjQzb2R6b2ptc3J4MSJ9.zA2W0IkI0c6KaAhJfk9bWg&streets=true&language=en&session_token=0fb40925-45fc-4971-8919-4126168062b6&proximity=ip')
// search
const searchMap = async (val) => {
  loading.value = true
  results.value = await provider.search({query: val});
  console.log(results.value)
}

const validateLocation = () => {
  const loc = model.value
  loading.value = false
  console.log(loc.x, loc.y)
}
</script>

<template>
  <v-card class="w-full h-1/3 flex justify-center items-center p-8"
          density="comfortable"
          :loading="loading"
  >
    <v-autocomplete
        autofocus
        clearable
        focused
        hide-no-data
        label="Choose location for results.."
        :items="results"
        v-model="model"
        item-title="label"
        no-filter
        return-object
        @update:search="(val)=> searchMap(val)"
        validate-on="blur"
        @update:modelValue="validateLocation()"
    ></v-autocomplete>
  </v-card>
</template>

<style scoped>

</style>