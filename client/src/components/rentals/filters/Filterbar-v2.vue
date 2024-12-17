<template>
    <div class="hidden md:block mt-2 p-4 border h-24 border-gray-200 rounded-lg">
      <!-- Filters Block -->
      <div name="filters" class="flex flex-row gap-4 items-center">
        <!-- Region Selection -->
        <div class="md:w-48">
          <v-select
            v-model="selectedRegion"
            :items="regionOptions"
            label="Region"
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
            label="Bedrooms"
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
            label="Bathrooms"
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
            label="Availability"
            item-title="text"
            item-value="value"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>
  
        <!-- Type -->
        <div class="md:w-48">
          <v-select
            v-model="filters.type"
            :items="typeOptions"
            label="Type"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>
  
        <!-- Standing -->
        <div class="md:w-48">
          <v-select
            v-model="filters.standing"
            :items="standingOptions"
            label="Standing"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>
  
        <!-- Price -->
        <div class="md:w-48">
          <v-select
            v-model="filters.priceRange"
            :items="priceOptions"
            label="Price Range"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>
  
        <!-- Rules -->
        <div class="md:w-48">
          <v-select
            v-model="filters.rules"
            :items="rulesOptions"
            label="Rules"
            outlined
            dense
            clearable
            @change="updateFiltersState"
          ></v-select>
        </div>
  
        <!-- Amenities -->
        <div class="md:w-48">
          <v-select
            v-model="filters.amenities"
            :items="amenitiesOptions"
            label="Amenities"
            outlined
            dense
            clearable
            multiple
            @change="updateFiltersState"
          ></v-select>
        </div>
  
        <!-- Reset Button -->
        <v-btn class="w-32" size="small" variant="elevated" color="secondary" @click="resetFilters">
          Reset
        </v-btn>
      </div>
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
  const typeOptions = ["shared", "independent", "sale"];
  const standingOptions = ["economy", "standard", "luxury"];
  const priceOptions = [
    "Under 500 TND",
    "500-1000 TND",
    "1000-2000 TND",
    "Above 2000 TND",
  ];
  const rulesOptions = ["Pets Allowed", "Parties Allowed", "Smoking Allowed"];
  const amenitiesOptions = ["Air Conditioning", "Heating", "Refrigerator", "Parking"];
  
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
      updateFilters({ region: coords });
    }
  };
  
  // Reset filters to defaults
  const resetFilters = () => {
    updateFilters({
      bedrooms: null,
      bathrooms: null,
      available: null,
      type: null,
      standing: null,
      priceRange: null,
      rules: null,
      amenities: null,
      region: null,
    });
    selectedRegion.value = null;
  };
  </script>
  
  <style scoped>
  .text-muted {
    color: #6c757d;
  }
  </style>
  