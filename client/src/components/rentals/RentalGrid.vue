<script setup lang="ts">
import CardTify from "@/components/rentals/RentalCard.vue";

// Define props to accept rentals
const props = defineProps<{
  rentals: Array<{
    name: string;
    fullAddress: string;
    images: string[];
  }>;
}>();
</script>

<template>
  <v-card class="mx-auto overflow-auto w-full">
    <v-container fluid>
      <v-row>
        <!-- Iterate over rentals to display each as a card -->
        <v-col
          v-for="rental in rentals"
          :key="rental.name"
          cols="12"
          md="6"
          lg="4"
        >
          <CardTify
            :title="rental.name"
            :src="rental.images[0]"
            :address="rental.fullAddress"
          />
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<style scoped>
</style>
