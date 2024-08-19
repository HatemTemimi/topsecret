<script setup lang="ts">

import { RouterView } from 'vue-router'
import NavBar from "@/components/layout/Navbar.vue";

</script>

<template>
  <div class="flex flex-col">
    <NavBar/>
    <RouterView />
  </div>
</template>

