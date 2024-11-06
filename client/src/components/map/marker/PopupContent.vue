<script setup lang="ts">
import { defineProps } from 'vue'

// Define rental as a prop
const props = defineProps<{ rental: { name: string; fullAddress: string } }>()
</script>

<template>
  <div>
    <h3>{{ props.rental.name }}</h3>
    <p>{{ props.rental.fullAddress }}</p>
  </div>
</template>

<style scoped>
/* Optional styling */
h3 {
  margin: 0;
  font-weight: bold;
}
p {
  margin: 0;
}
</style>