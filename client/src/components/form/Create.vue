<template>
  <v-form @submit.prevent="submitForm" fast-fail>
    <Address />

    <v-card title="Address details" class="p-4">
      <v-text-field
          v-model="state.streetNumber"
          :error-messages="v$.streetNumber.$errors.map(e => e.$message)"
          label="Street Number"
          required
          @blur="v$.streetNumber.$touch"
          @input="v$.streetNumber.$touch"
      ></v-text-field>

      <v-text-field
          v-model="state.street"
          :error-messages="v$.street.$errors.map(e => e.$message)"
          label="Street"
          required
          @blur="v$.street.$touch"
          @input="v$.street.$touch"
      ></v-text-field>

      <v-text-field
          v-model="state.city"
          :error-messages="v$.city.$errors.map(e => e.$message)"
          label="City"
          required
          @blur="v$.city.$touch"
          @input="v$.city.$touch"
      ></v-text-field>

      <v-text-field
          v-model="state.country"
          :error-messages="v$.country.$errors.map(e => e.$message)"
          label="Country"
          required
          @blur="v$.country.$touch"
          @input="v$.country.$touch"
      ></v-text-field>

      <v-text-field
          v-model="state.fullAddress"
          :error-messages="v$.fullAddress.$errors.map(e => e.$message)"
          label="Full Address"
          required
          @blur="v$.fullAddress.$touch"
          @input="v$.fullAddress.$touch"
      ></v-text-field>
    </v-card>

    <!-- Image Upload Section -->
    <v-card title="Upload Images" class="p-4 mt-4">
      <v-file-input
        model-value="state.images"
        label="Add up to 3 images"
        accept="image/*"
        multiple
        :counter=true
        :rules="[fileLimit]"
        required
        @change="handleImageUpload"
      ></v-file-input>
    </v-card>

    <v-checkbox
        v-model="state.checkbox"
        :error-messages="v$.checkbox.$errors.map(e => e.$message)"
        label="Do you agree?"
        required
        @blur="v$.checkbox.$touch"
        @change="v$.checkbox.$touch"
    ></v-checkbox>

    <v-btn type="submit" class="me-4">
      Submit
    </v-btn>
    <v-btn @click="clear">
      Clear
    </v-btn>
  </v-form>
</template>

<script setup>
import { reactive, ref, provide } from 'vue';
import { useVuelidate } from '@vuelidate/core';
import { required } from '@vuelidate/validators';
import Address from "@/components/form/Address.vue";

// Initial form state
const initialState = {
  checkbox: null,
  streetNumber: '',
  street: '',
  city: '',
  country: '',
  fullAddress: '',
  images: [],
};

const state = reactive({
  ...initialState,
});

const rules = {
  checkbox: { },
  streetNumber: {  },
  street: {  },
  city: {  },
  country: {  },
  fullAddress: {  },
  images: {  },
};

// Validation instance
const v$ = useVuelidate(rules, state);

// Custom rule for image limit
const fileLimit = (files) => {
  return files && files.length <= 3 || 'You can only upload up to 3 images';
};

// Handle image upload and convert FileList to an array
function handleImageUpload(event) {
  console.log(event.target.value)
  const files = Array.from(event.target.value); // Convert FileList to array
  state.images = files.slice(0, 3); // Limit to 3 images
}

// Clear form function
function clear() {
  v$.value.$reset();

  for (const [key, value] of Object.entries(initialState)) {
    state[key] = value;
  }
}

// Capture form submission and log values
function submitForm() {
  console.log('Form Values:', JSON.stringify(state, null, 2));
  v$.value.$validate().then((isValid) => {
    if (isValid) {
      console.log('Form Values:', JSON.stringify(state, null, 2));
    }
  });
}

// Marker state and update function
const marker = ref([36.8065, 10.181667]);
function updateMarker(val) {
  marker.value = val;
}

// Provide marker and updateMarker for other components
provide('location', {
  marker,
  updateMarker,
});

// Provide a function to update info section fields from Address component
function updateInfo(newInfo) {
  state.streetNumber = newInfo.streetNumber || '';
  state.street = newInfo.street || '';
  state.city = newInfo.city || '';
  state.country = newInfo.country || '';
  state.fullAddress = newInfo.fullAddress || '';
}

// Provide the updateInfo function for Address component to access
provide('updateInfo', updateInfo);
</script>

<style scoped>
</style>
