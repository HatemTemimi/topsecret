<template>
  <v-form fast-fail>
    <v-card title="Address" class="p-4">
      <Address />
    </v-card>

    <v-card title="Info" class="p-4 mt-4">
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

    <v-checkbox
        v-model="state.checkbox"
        :error-messages="v$.checkbox.$errors.map(e => e.$message)"
        label="Do you agree?"
        required
        @blur="v$.checkbox.$touch"
        @change="v$.checkbox.$touch"
    ></v-checkbox>

    <v-btn
        class="me-4"
        @click="v$.$validate"
    >
      submit
    </v-btn>
    <v-btn @click="clear">
      clear
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
};

const state = reactive({
  ...initialState,
});

const rules = {
  items: { required },
  checkbox: { required },
  streetNumber: { required },
  street: { required },
  city: { required },
  country: { required },
  fullAddress: { required },
};

const v$ = useVuelidate(rules, state);

function clear() {
  v$.value.$reset();

  for (const [key, value] of Object.entries(initialState)) {
    state[key] = value;
  }
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
