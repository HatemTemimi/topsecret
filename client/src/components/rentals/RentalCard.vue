<script setup lang="ts">
import {ref, defineProps, onMounted} from "vue";
import { useRouter } from "vue-router";

// Props
const props = defineProps({
  rental: {
    type: Object,
    required: true,
  },
});

const show = ref(false);
const router = useRouter(); // Access Vue Router

// Navigate to the rental details page
const goToDetails = () => {
  if (props.rental?.id) {
    router.push(`/rental/details/${props.rental.id}`);
  } else {
    console.error("Rental ID is missing");
  }
};

</script>

<template>
  <v-card class="mx-auto" :hover="true" color="surface" max-width="300">
    <!-- Rental Image -->
    <v-img
      height="200px"
      :src="rental.images?.[0]"
      cover
    ></v-img>

    <!-- Rental Title -->
    <v-card-title>
      {{ rental.name }}
    </v-card-title>

    <!-- Rental Subtitle -->
    <v-card-subtitle>
      {{ rental.fullAddress }}
    </v-card-subtitle>

    <v-card-actions>
      <v-btn variant="flat" size="small" color="primary-darken-1" @click="goToDetails">
        <!--<v-icon>mdi-eye-arrow-right</v-icon>-->
        Details
      </v-btn>

      <v-spacer></v-spacer>

      <!-- Expand Button -->
      <v-btn
        :icon="show ? 'mdi-chevron-up' : 'mdi-chevron-down'"
        @click="show = !show"
      ></v-btn>
    </v-card-actions>

    <!-- Expandable Rental Details -->
    <v-expand-transition>
      <div v-show="show">
        <v-divider></v-divider>

        <v-card-text>
          <!-- Display rental details dynamically -->
          <p><strong>Description:</strong> {{ rental.description || 'No description available.' }}</p>
          <p><strong>Price:</strong> ${{ rental.price }}</p>
          <p><strong>Bedrooms:</strong> {{ rental.bedrooms }}</p>
          <p><strong>Bathrooms:</strong> {{ rental.bathrooms }}</p>
          <p><strong>Availability:</strong> {{ rental.available ? 'Available' : 'Unavailable' }}</p>
        </v-card-text>
      </div>
    </v-expand-transition>
  </v-card>
</template>
