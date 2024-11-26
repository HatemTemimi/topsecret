<script setup lang="ts">
import CardTify from "@/components/rentals/RentalCard.vue";
import type { Rental } from "@/models/rental";

// Define props to accept rentals
const props = defineProps<{
  rentals: Array<Rental>;
}>();
</script>

<template>
  <v-card class="mx-auto overflow-auto w-full">
    <v-container fluid>
      <v-row>
        <!-- Iterate over rentals to display each as a card -->
        <v-col
          v-for="rental in rentals"
          :key="rental.id"
          cols="12"
          md="6"
          lg="4"
        >
          <CardTify
            :rental="rental"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<style scoped>
</style>
