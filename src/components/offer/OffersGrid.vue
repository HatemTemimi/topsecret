<script setup lang="ts">

import {ref} from "vue";
import CardTify from "@/components/offer/OfferCard.vue";

const cards = ref(
    [
      {title: 'Favorite road trips', src: 'https://cdn.vuetifyjs.com/images/cards/hotel.jpg', flex: 6},
      {title: 'Best airlines', src: 'https://cdn.vuetifyjs.com/images/cards/hotel.jpg', flex: 6},
      {title: 'Best airlines', src: 'https://cdn.vuetifyjs.com/images/cards/hotel.jpg', flex: 6},
      {title: 'Best airlines', src: 'https://cdn.vuetifyjs.com/images/cards/hotel.jpg', flex: 6},
      {title: 'Best airlines', src: 'https://cdn.vuetifyjs.com/images/cards/hotel.jpg', flex: 6},
      {title: 'Best airlines', src: 'https://cdn.vuetifyjs.com/images/cards/hotel.jpg', flex: 6},
      {title: 'Best airlines', src: 'https://cdn.vuetifyjs.com/images/cards/hotel.jpg', flex: 6},
      {title: 'Best airlines', src: 'https://cdn.vuetifyjs.com/images/cards/hotel.jpg', flex: 6},
      {title: 'Best airlines', src: 'https://cdn.vuetifyjs.com/images/cards/hotel.jpg', flex: 6},
    ]
)

</script>

<template>
  <v-card
      class="mx-auto overflow-auto w-full"
  >
    <v-container fluid>
      <v-row>
        <v-col
            v-for="card in cards"
            :key="card.title"
            :cols="card.flex"
        >
          <CardTify/>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>


<style scoped>

</style>