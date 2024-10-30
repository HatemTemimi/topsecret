<script setup lang="ts">
import {VAutocomplete} from 'vuetify/components'
import {ref, watch} from "vue";
import {useRouter} from "vue-router";
import { getPlacesGoogle } from "@/api/getPlaces";
import _ from 'lodash'

const model = ref(null)
const results = ref([])
const isOpen = ref(false)
const loading = ref(false)

const router = useRouter()

// Debounced search function using watch
let timeoutId: number | null = null;
const debouncedSearch = (val: string) => {
  if (timeoutId) clearTimeout(timeoutId);
  timeoutId = setTimeout(async () => {
    if (!_.isEmpty(val)) {
      results.value = await getPlacesGoogle(val)
      console.log(results.value)
      isOpen.value = true;
    }
  }, 500);
}

watch(model, debouncedSearch, {immediate: true});

const validateLocation = () => {
  const loc = model.value
  isOpen.value = false;
  loc?.center ?
      router.push({path: '/home', query: {lat: loc.center[0], lng: loc.center[1]}})
      : router.push(`/home`)
}
</script>

<template>
  <v-card class="w-1/3 h-1/3 flex justify-center items-center p-8"
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
        @update:modelValue="validateLocation()"
    ></v-autocomplete>
  </v-card>
</template>

<style scoped>

</style>