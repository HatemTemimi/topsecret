<script setup lang="ts">
import LeafletMap from '@/components/map/Map.vue'
import CardContainer from "@/components/card/CardGrid.vue";
import Filters from "@/components/layout/Filters.vue";
</script>

<template>
  <Filters/>
  <div class="flex flex-row gap-4 w-full h-[91vh] overflow-y-hidden">
    <LeafletMap/>
    <card-container/>
  </div>
</template>
