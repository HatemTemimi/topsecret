<script setup lang="ts">
import {MapBoxProvider} from 'leaflet-geosearch';
import {ref} from "vue";
import {useRouter} from "vue-router";

const search = ref('')
const results = ref([])
const isOpen = ref(false)

const router = useRouter()

const provider = new MapBoxProvider({
  params: {
    access_token: 'pk.eyJ1IjoiaGF0ZW0tdGVtaW1pIiwiYSI6ImNtMDE2ZWxkYzFmMHQycnIyZmVwNHBucXEifQ.1Ca9bPDHbgtbsoytvEvWJw',
    country: 'tn',
  },
});

/*
const searchControl = new GeoSearchControl({
  provider: provider,
  autoComplete: true, // optional: true|false  - default true
  autoCompleteDelay: 250
}) ;
*/

// search
const searchMap = async () => {
  results.value = await provider.search({query: search.value});
  isOpen.value = true;
}

const validateLocation = (loc) => {
  search.value = location.label;
  isOpen.value = false;
  loc.x && loc.y ?
      router.push({path: '/', query: {lat: loc.raw.center[0], lng: loc.raw.center[1]}})
      : router.push(`/`)
  console.log(results.value)
}
</script>

<template>
  <div class="flex flex-col w-full items-center">
    <input @keyup="searchMap" v-model="search" type="text" placeholder="search by location.."
           class="rounded text-center w-1/4 h-10"/>
    <ul v-if="isOpen && results.length > 0"
        class="w-1/4 mt-1 border-2 border-slate-50 overflow-auto shadow-lg rounded list-none"
    >
      <li :class="[
  'hover:bg-blue-100 hover:text-blue-800',
  'w-full list-none text-left py-2 px-3 cursor-pointer']"
          v-for="(location,i) in results" :key="i"
          @click="validateLocation(location)"
      >
        {{ location.label }}
      </li>
    </ul>
  </div>
</template>
