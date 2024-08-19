<script setup lang="ts">

import Card from "@/components/card/Card.vue";
</script>

<template>
  <div class="container mx-auto px-4 md:px-12 overflow-auto">
    <div class="flex flex-wrap -mx-1 lg:-mx-4">
      <div class="w-full flex md:w-1/2 px-1 md:px-4"  v-for="i in 12">
        <Card />
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>