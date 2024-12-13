<script setup lang="ts">
import {RouterView} from 'vue-router'
import Navbar from "@/components/layout/Navbar.vue";
</script>

<template>
  <v-layout class="rounded-md lg:p-4">
    <Navbar/>
    <v-main class="d-flex align-center justify-center">
      <RouterView/>
    </v-main>
  </v-layout>
</template>


