<script setup lang="ts">
import RentalCard from "@/components/rentals/RentalCard.vue";
import type { Rental } from "@/models/rental";
import { inject } from "vue";

// Define props to accept rentals
const props = defineProps<{
  rentals: Array<Rental>;
}>();

const loading = inject("loading")

</script>

<template>
  <v-card class="bg-white md:w-[50vw] mx-auto border-thin overflow-auto sm:py-4 sm:w-full rounded-lg">
    <div class="h-full w-full flex items-center justify-center" v-if="loading">
          <v-progress-circular
     class="absolute"
      v-show="loading"
      size="100"
      width="50"
      color="success"
      indeterminate
    ></v-progress-circular>

    </div>
    <v-container v-else fluid>
      <v-row>
        <!-- Iterate over rentals to display each as a card -->
        <v-col
          v-for="rental in rentals"
          :key="rental.id"
          cols="12"
          sm="2"
          md="6"
          lg="6"
        >
          <RentalCard
          :with-expand="true"
          :rental="rental"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<style scoped>
</style>
