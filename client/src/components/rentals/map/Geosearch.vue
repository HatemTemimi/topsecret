<script setup lang="ts">
import { VAutocomplete } from 'vuetify/components';
import { ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import { getPlaceDetails, getPlacesGoogle } from '@/api/getPlaces';

const model = ref<any | null>(null);
const results = ref<any[]>([]);
const isOpen = ref(false);
const loading = ref(false);

const router = useRouter();

// Utility for debouncing
const debounce = (func: Function, delay: number) => {
  let timer: number | undefined;
  return (...args: any[]) => {
    clearTimeout(timer);
    timer = setTimeout(() => func(...args), delay);
  };
};

// Debounced search function
const performSearch = async (query: string) => {
  if (query && model.value === null) {
    loading.value = true;
    results.value = await getPlacesGoogle(query);
    isOpen.value = true;
    loading.value = false;
  }
};

const debouncedSearch = debounce(performSearch, 500);

// Watcher for search input
watch(
  model,
  (val) => {
    if (typeof val === 'string') {
      debouncedSearch(val);
    }
  }
);

// Validate and route on location selection
const validateLocation = async (selectedModel: any) => {
  if (!selectedModel?.place_id) return;

  loading.value = true;
  isOpen.value = false;

  try {
    const placeDetails = await getPlaceDetails(selectedModel.place_id);
    const { lat, lng } = placeDetails.geometry.location;

    console.log('Latitude:', lat, 'Longitude:', lng);

    router.push({
      path: '/rentals',
      query: { lat, lng },
    });
  } catch (error) {
    console.error('Failed to fetch place details:', error);
  } finally {
    loading.value = false;
  }
};
</script>

<template>
    <v-autocomplete
      class="w-2/3"
      clearable
      focused
      hide-no-data
      label="Choose location for results.."
      :items="results"
      v-model="model"
      item-title="description"
      no-filter
      return-object
      @focus="loading = true"
      @blur="loading = false"
      @update:search="debouncedSearch"
      @update:modelValue="validateLocation"
    ></v-autocomplete>
</template>
