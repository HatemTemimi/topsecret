<template>
    <div class="hidden md:block mt-2 p-4 border h-24 border-gray-200 rounded-lg">
      <!-- Filters Block -->
      <div name="filters" class="flex flex-row gap-4 items-center">
        <!-- Region Selection -->
        <div class="sm:hidden md:block md:w-48 ">
          <v-select
            v-model="selectedRegion"
            :items="regionOptions"
            label="region"
            outlined
            dense
            clearable
            @change="updateRegion"
          ></v-select>
        </div>

        <!-- Bedrooms -->
        <div class="md:w-48">
          <v-select
            v-model="filters.bedrooms"
            :items="bedroomOptions"
            label="bedrooms"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>

        <!-- Bathrooms -->
        <div class="md:w-48">
          <v-select
            v-model="filters.bathrooms"
            :items="bathroomOptions"
            label="bathrooms"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>

        <!-- Availability -->
        <div class="md:w-48">
          <v-select
            v-model="filters.available"
            :items="availabilityOptions"
            label="availability"
            item-title="text"
            item-value="value"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>

        <div class="md:w-48">
          <v-select
            v-model="filters.available"
            :items="availabilityOptions"
            label="type"
            item-title="text"
            item-value="value"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>
        <div class="md:w-48">
          <v-select
            v-model="filters.available"
            :items="availabilityOptions"
            label="standing"
            item-title="text"
            item-value="value"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>
        <div class="md:w-48">
          <v-select
            v-model="filters.available"
            :items="availabilityOptions"
            label="price"
            item-title="text"
            item-value="value"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>
        <div class="md:w-48">
          <v-select
            v-model="filters.available"
            :items="availabilityOptions"
            label="rules"
            item-title="text"
            item-value="value"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>
        <div class="md:w-48">
          <v-select
            v-model="filters.available"
            :items="availabilityOptions"
            label="amenities"
            item-title="text"
            item-value="value"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>
        <v-btn class="w-32" size="small" variant="elevated" color="secondary" @click="resetFilters">Reset</v-btn>
      </div>

      <!-- Filter & Reset Buttons -->
    </div>
</template>

<script setup lang="ts">
import { inject, ref } from "vue";

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
