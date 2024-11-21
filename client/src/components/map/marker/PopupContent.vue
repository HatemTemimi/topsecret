<script setup lang="ts">
import { defineProps } from 'vue'
const props = defineProps<{ rental: { name: string; fullAddress: string } }>()
</script>

<template>
  <v-container fluid min-width="13vw">
        <v-card theme="light" :title="props.rental.name">
          <v-card-item>
    <p>{{ props.rental.fullAddress }}</p>
    <v-chip
      class="ma-2"
      color="pink"
      label
    >
      <v-icon icon="mdi-label" start></v-icon>
      some tag
    </v-chip>
</v-card-item>
<v-card-actions>
  <v-btn size="small" variant="outlined">
    View Rental
    <v-icon>mdi-arrow-right</v-icon>
  </v-btn>
</v-card-actions>

    </v-card>
  </v-container>
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