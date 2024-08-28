<script setup lang="ts">
import {RouterView} from 'vue-router'
import Navbar from "@/components/layout/Navbar.vue";
</script>

<template>
  <v-layout class="rounded-md p-2">
    <Navbar/>
    <v-main class="d-flex align-center justify-center" style="min-height: 300px;">
      <RouterView/>
    </v-main>
  </v-layout>
</template>


