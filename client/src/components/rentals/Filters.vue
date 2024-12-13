<template>
    <v-card class="w-full p-4">
    <div class="flex flex-row justify-between items-center">
      <!-- Filters Block -->
      <div name="filters" class="flex flex-row gap-4 items-center">
        <!-- Region Selection -->
        <div class="min-w-[200px]">
          <v-select
            v-model="selectedRegion"
            :items="regionOptions"
            label="Select Region"
            outlined
            dense
            clearable
            @change="updateRegion"
          ></v-select>
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
            item-title="text"
            item-value="value"
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
</v-card>
</template>

<script setup lang="ts">
import { inject, ref } from "vue";
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

// Regions in Tunis Capital
const regionCoordinates = ref({
  Ariana: [
    [36.866785, 10.165079],
    [36.867583, 10.195508],
    [36.854747, 10.195373],
    [36.854033, 10.163997],
  ],
  "Tunis Centre": [
    [36.806112, 10.167519],
    [36.807963, 10.190217],
    [36.800972, 10.190947],
    [36.799844, 10.167839],
  ],
  "El Menzah": [
    [36.834, 10.163],
    [36.838, 10.185],
    [36.825, 10.185],
    [36.823, 10.163],
  ],
  "El Manar": [
    [36.836, 10.155],
    [36.840, 10.178],
    [36.830, 10.178],
    [36.828, 10.155],
  ],
  Mutuelleville: [
    [36.822, 10.164],
    [36.825, 10.180],
    [36.815, 10.180],
    [36.812, 10.164],
  ],
  Bardo: [
    [36.806, 10.145],
    [36.810, 10.165],
    [36.800, 10.165],
    [36.798, 10.145],
  ],
});

// Region options for dropdown
const regionOptions = Object.keys(regionCoordinates.value);

// Selected region
const selectedRegion = ref(null);

// Update filters state
const updateFiltersState = () => {
  updateFilters(filters);
};

// Update selected region
const updateRegion = () => {
  if (selectedRegion.value) {
    const coords = regionCoordinates.value[selectedRegion.value];
    updateFilters({ region: coords }); // Update filters with region's coordinates
  }
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
    region: null, // Reset region
  });
  selectedRegion.value = null; // Clear region selection
};
</script>

<style scoped>
.text-muted {
  color: #6c757d;
}
</style>
