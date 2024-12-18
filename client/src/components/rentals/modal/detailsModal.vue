<template>
  <v-dialog
      v-model="dialog"
      width="auto"
      :persistent="false"
    >
      <v-card
        max-width="400"
        prepend-icon="mdi-update"
        text="Your application will relaunch automatically after the update is complete."
        title="Update in progress"
      >
      </v-card>
  </v-dialog>
</template>
<script setup>
import { inject } from 'vue';

const dialog = inject("dialog")


</script>