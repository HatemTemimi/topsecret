<script setup lang="ts">
import { defineProps } from 'vue'
const props = defineProps<{ rental: { name: string; fullAddress: string } }>()
</script>

<template>
  <div>
    <h3>{{ props.rental.name || 'just a rent' }}</h3>
    <p>{{ props.rental.fullAddress }}</p>
  </div>
</template>

<style scoped>
h3 {
  margin: 0;
  font-weight: bold;
}
p {
  margin: 0;
}
</style>