<script setup lang="ts">
import { defineProps } from 'vue'
import { useRouter } from 'vue-router';

const router = useRouter();

const props = defineProps<{ rental: { id: string; name: string; fullAddress: string } }>();

const navigateToRental = () => {
  router.push(`/rental/details/${props.rental.id}`);
};
</script>

<template>
    <v-card :title="props.rental.name">
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
          Edit
          <v-icon>mdi-pencil</v-icon>
        </v-btn>
        <v-btn size="small" variant="outlined" @click="navigateToRental">
        View
        <v-icon>mdi-arrow-right</v-icon>
      </v-btn>
      </v-card-actions>
    </v-card>
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