<script setup lang="ts">
import RentalCard from "@/components/rentals/RentalCard.vue";
import type { Rental } from "@/models/rental";

// Define props to accept rentals
const props = defineProps<{
  rentals: Array<Rental>;
}>();
</script>

<template>
  <v-card class="bg-white mx-auto border-thin overflow-auto sm:py-4 sm:w-full">
    <v-container fluid>
      <v-row>
        <!-- Iterate over rentals to display each as a card -->
        <v-col
          v-for="rental in rentals"
          :key="rental.id"
          cols="12"
          sm="2"
          md="3"
          lg="4"
        >
          <RentalCard
            :rental="rental"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<style scoped>
</style>
