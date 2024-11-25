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

    <!-- Rental Details -->
    <v-card title="Rental Details" class="p-4 mt-4">
      <v-text-field
        v-model="state.name"
        :error-messages="v$.name.$errors.map(e => e.$message)"
        label="Rental Name"
        required
        @blur="v$.name.$touch"
        @input="v$.name.$touch"
      ></v-text-field>

      <v-text-field
        v-model="state.price"
        type="number"
        :error-messages="v$.price.$errors.map(e => e.$message)"
        label="Price"
        required
        @blur="v$.price.$touch"
        @input="v$.price.$touch"
      ></v-text-field>

      <v-text-field
        v-model="state.bedrooms"
        type="number"
        :error-messages="v$.bedrooms.$errors.map(e => e.$message)"
        label="Bedrooms"
        required
        @blur="v$.bedrooms.$touch"
        @input="v$.bedrooms.$touch"
      ></v-text-field>

      <v-text-field
        v-model="state.bathrooms"
        type="number"
        :error-messages="v$.bathrooms.$errors.map(e => e.$message)"
        label="Bathrooms"
        required
        @blur="v$.bathrooms.$touch"
        @input="v$.bathrooms.$touch"
      ></v-text-field>

      <v-text-field
        v-model="state.areaSize"
        type="number"
        :error-messages="v$.areaSize.$errors.map(e => e.$message)"
        label="Area Size (sq. meters)"
        required
        @blur="v$.areaSize.$touch"
        @input="v$.areaSize.$touch"
      ></v-text-field>

      <v-select
        v-model="state.available"
        :items="[true, false]"
        :error-messages="v$.available.$errors.map(e => e.$message)"
        label="Available"
        required
        @blur="v$.available.$touch"
        @input="v$.available.$touch"
      ></v-select>

      <v-textarea
        v-model="state.description"
        :error-messages="v$.description.$errors.map(e => e.$message)"
        label="Description"
        required
        @blur="v$.description.$touch"
        @input="v$.description.$touch"
      ></v-textarea>
    </v-card>

    <!-- Image Upload Section -->
    <v-card title="Upload Images" class="p-4 mt-4">
      <v-file-input
        v-model="state.images"
        label="Add up to 3 images"
        accept="image/*"
        multiple
        :counter="true"
        :rules="[fileLimit]"
        required
      ></v-file-input>
    </v-card>

    <v-checkbox
        v-model="state.agree"
        :error-messages="v$.agree.$errors.map(e => e.$message)"
        label="Do you agree?"
        required
        @blur="v$.agree.$touch"
        @change="v$.agree.$touch"
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
import { reactive, ref, provide, watch } from 'vue';
import { useVuelidate } from '@vuelidate/core';
import { required, numeric } from '@vuelidate/validators';
import Address from "@/components/rentals/form/Address.vue";
import axios from 'axios';

const marker = ref([36.8065, 10.181667]);

function updateMarker(val) {
  marker.value = val;
}

// Form state
const initialState = {
  name: '',
  price: '',
  bedrooms: '',
  bathrooms: '',
  areaSize: '',
  available: null,
  description: '',
  agree: null,
  geometry: {
    lat: marker.value[0],
    lng: marker.value[1],
  },
  streetNumber: '',
  street: '',
  city: '',
  country: '',
  fullAddress: '',
  images: [],
};

const state = reactive({ ...initialState });

const rules = {
  name: { required },
  price: { required, numeric },
  bedrooms: { required, numeric },
  bathrooms: { required, numeric },
  areaSize: { required, numeric },
  available: { required },
  description: { required },
  agree: { required },
  streetNumber: { required },
  street: { required },
  city: { required },
  country: { required },
  fullAddress: { required },
};

const v$ = useVuelidate(rules, state);

watch(marker, (newVal) => {
  state.geometry.lng = newVal[0];
  state.geometry.lat = newVal[1];
}, { deep: true });

const fileLimit = (files) => {
  return files && files.length <= 3 || 'You can only upload up to 3 images';
};

function clear() {
  v$.value.$reset();
  Object.assign(state, { ...initialState });
}

async function submitForm() {
  const isValid = await v$.value.$validate();
  if (!isValid) return;

  const formData = new FormData();
  for (const key in state) {
    if (key === 'images') {
      state.images.forEach((image, index) => {
        formData.append(`images[${index}]`, image);
      });
    } else if (typeof state[key] !== 'object') {
      formData.append(key, state[key]);
    } else {
      for (const subKey in state[key]) {
        formData.append(`${key}.${subKey}`, state[key][subKey]);
      }
    }
  }

  try {
    const response = await axios.post("http://localhost:3001/api/rental/add", formData, {
      headers: { "Content-Type": "multipart/form-data" },
    });
    console.log("Rental added successfully:", response.data);
    clear();
  } catch (error) {
    console.error("Failed to add rental:", error.response?.data || error.message);
  }
}

provide('location', { marker, updateMarker });
provide('updateInfo', (newInfo) => {
  state.streetNumber = newInfo.streetNumber || '';
  state.street = newInfo.street || '';
  state.city = newInfo.city || '';
  state.country = newInfo.country || '';
  state.fullAddress = newInfo.fullAddress || '';
});
</script>

<style scoped></style>
