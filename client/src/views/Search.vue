<script setup lang="ts">
import Geosearch from "@/components/rentals/map/Geosearch.vue";
import { ref, onMounted } from "vue";

// Control the visibility of the second card
const showSecondCard = ref(false);

// Trigger animation after the component mounts
onMounted(() => {
  setTimeout(() => {
    showSecondCard.value = true;
  }, 500); // Delay before showing the second card
});
</script>

<template>
  <v-card
    class="h-[90vh] w-full flex flex-col justify-center items-center p-8"
    density="comfortable"
    image="/img.jpg"
  >
    <transition name="fade-in">

    <v-card
      class="h-[18vh] flex flex-col justify-center items-center"
      hover
      density="comfortable"
        v-if="showSecondCard"
        elevation-17
    >
      <v-card-title>
        Welcome to Dar/win
      </v-card-title>
      <v-card-subtitle>
        Discover your next rental in Tunis. Use the search below to start your journey.
      </v-card-subtitle>
      <Geosearch class="mt-4" />
  </v-card>

    </transition>
    </v-card>
</template>

<style scoped>
/* Fade-in animation */
.fade-in-enter-active {
  animation: fadeIn 0.5s ease-in;
}
.fade-in-enter-from {
  opacity: 0;
}
.fade-in-enter-to {
  opacity: 1;
}

/* Keyframes for fade-in animation */
@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>
