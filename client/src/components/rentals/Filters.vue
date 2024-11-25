<template>
  <v-container fluid class="py-4" max-height="10vh">
    <div class="flex flex-row justify-between items-center">
 

      <!-- Filters Block -->
      <div name="filters" class="flex flex-row gap-4 items-center">

        <div class="min-w-[300px]">

          <LocationAutocomplete/>
        
        </div>
        <!--
        <div class="min-w-[200px]">
          <v-slider
            v-model="filters.priceRange"
            :max="priceMax"
            :min="priceMin"
            :step="100"
            range
            label="Price Range"
            thumb-label="always"
          ></v-slider>
        </div>
-->

        <!-- Bedrooms -->
        <div class="min-w-[150px]">
          <v-select
            v-model="filters.bedrooms"
            :items="bedroomOptions"
            label="Bedrooms"
            outlined
            dense
            clearable
          ></v-select>
        </div>

        <!-- Bathrooms -->
        <div class="min-w-[150px]">
          <v-select
            v-model="filters.bathrooms"
            :items="bathroomOptions"
            label="Bathrooms"
            outlined
            dense
            clearable
          ></v-select>
        </div>

        <!-- Availability -->
        <div class="min-w-[150px]">
          <v-select
            v-model="filters.available"
            :items="availabilityOptions"
            label="Availability"
            outlined
            dense
            clearable
          ></v-select>
        </div>
      </div>
     <!-- Filter & Reset Buttons -->
      <div name="buttons" class="flex flex-row gap-2">
        <v-btn size="small" color="primary" @click="applyFilters">Apply Filters</v-btn>
        <v-btn size="small" outlined color="secondary" @click="resetFilters">Reset</v-btn>
      </div>
    </div>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from "vue";
import LocationAutocomplete from "./map/LocationAutocomplete.vue";

// Reactive state for filters
const filters = ref({
  priceRange: [0, 5000], // Default price range
  bedrooms: null,
  bathrooms: null,
  available: null,
});

// Options for filter dropdowns
const priceMin = 0;
const priceMax = 10000;

const bedroomOptions = [1, 2, 3, 4, 5];
const bathroomOptions = [1, 2, 3, 4];
const availabilityOptions = [
  { text: "Available", value: true },
  { text: "Unavailable", value: false },
];

// Emit filter event to parent
const applyFilters = () => {
  // Emit the updated filters to the parent component
  emit("filter", { ...filters.value });
};

// Reset filters to defaults
const resetFilters = () => {
  filters.value = {
    priceRange: [priceMin, priceMax],
    bedrooms: null,
    bathrooms: null,
    available: null,
  };

  // Emit reset event to the parent component
  emit("filter", { ...filters.value });
};
</script>

<style scoped>
.text-muted {
  color: #6c757d;
}
</style>
