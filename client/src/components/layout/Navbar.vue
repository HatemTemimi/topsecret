<script setup lang="ts">

import {VBtn} from "vuetify/components";
</script>

<template>
  <v-app-bar rounded>
    <v-app-bar-title>App Title</v-app-bar-title>

    <v-spacer></v-spacer>

    <router-link to="/">
      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>
    </router-link>

    <router-link to="/create">
      <v-btn icon>
        <v-icon>mdi-plus</v-icon>
      </v-btn>
    </router-link>

    <router-link to="/home">
      <v-btn icon>
        <v-icon>mdi-home</v-icon>
      </v-btn>
    </router-link>

  </v-app-bar>
</template>

<style scoped>

</style>