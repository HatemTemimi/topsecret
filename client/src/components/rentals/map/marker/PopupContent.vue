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
    <v-card :elevation="0" color="white" class="w-full" :title="props.rental.name">
    <v-img
      height="100px"
      :src="rental.images?.[0]"
      cover
    ></v-img>
    <div class="flex flex-row gap-1">
        <v-chip
          class="mt-2"
          color="success"
          label
          size="small"
        >
        <v-icon icon="mdi-currency-usd" start></v-icon>
        {{ props.rental.price }} TND
        </v-chip>
        <v-chip
          class="mt-2"
          color="success"
          label
          size="small"
        >
        <v-icon icon="mdi-currency-usd" start></v-icon>
        {{ props.rental.bedrooms }} Bedrooms
        </v-chip>
    </div>
      <v-card-actions>
        <v-btn class="w-full" size="small" variant="elevated" color="secondary" @click="navigateToRental">
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