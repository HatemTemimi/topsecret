<template>
  <v-container fluid class="py-4" style="max-height: 10vh;">
    <div class="flex flex-row justify-between items-center">
      <!-- Filters Block -->
      <div name="filters" class="flex flex-row gap-4 items-center">
        <!-- Location Autocomplete -->
        <div class="min-w-[250px]">
          <LocationAutocomplete />
        </div>

        <!-- Bedrooms -->
        <div class="min-w-[150px]">
          <v-select
            v-model="filters.bedrooms"
            :items="bedroomOptions"
            label="Bedrooms"
            outlined
            dense
            clearable
            @change="updateFiltersState"
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
            @change="updateFiltersState"
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
            @change="updateFiltersState"
          ></v-select>
        </div>
      </div>

      <!-- Filter & Reset Buttons -->
      <div name="buttons" class="flex flex-row gap-2 ml-4">
        <v-btn size="small" color="primary" @click="applyFilters">Apply Filters</v-btn>
        <v-btn size="small" outlined color="secondary" @click="resetFilters">Reset</v-btn>
      </div>
    </div>
  </v-container>
</template>

<script setup lang="ts">
import { inject } from "vue";
import LocationAutocomplete from "./map/LocationAutocomplete.vue";

// Inject shared state and update method from the parent
const filters = inject("filters");
const updateFilters = inject("updateFilters");

// Options for filter dropdowns
const bedroomOptions = [1, 2, 3, 4, 5];
const bathroomOptions = [1, 2, 3, 4];
const availabilityOptions = [
  { text: "Available", value: true },
  { text: "Unavailable", value: false },
];

// Update filters state
const updateFiltersState = () => {
  updateFilters(filters);
};

// Apply filters explicitly
const applyFilters = () => {
  updateFilters(filters);
};

// Reset filters to defaults
const resetFilters = () => {
  updateFilters({
    bedrooms: null,
    bathrooms: null,
    available: null,
  });
};
</script>

<style scoped>
.text-muted {
  color: #6c757d;
}
</style>
