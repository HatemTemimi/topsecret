<template>
    <v-form @submit.prevent="submitForm" fast-fail>
    <Address />
      <!-- Address Section -->
      <v-card title="Address Details" class="p-4">
        <v-text-field
          v-model="state.address.streetNumber"
          :error-messages="v$.address.streetNumber.$errors.map(e => e.$message)"
          label="Street Number"
          required
          @blur="v$.address.streetNumber.$touch"
          @input="v$.address.streetNumber.$touch"
        ></v-text-field>
  
        <v-text-field
          v-model="state.address.street"
          :error-messages="v$.address.street.$errors.map(e => e.$message)"
          label="Street"
          required
          @blur="v$.address.street.$touch"
          @input="v$.address.street.$touch"
        ></v-text-field>
  
        <v-text-field
          v-model="state.address.city"
          :error-messages="v$.address.city.$errors.map(e => e.$message)"
          label="City"
          required
          @blur="v$.address.city.$touch"
          @input="v$.address.city.$touch"
        ></v-text-field>
  
        <v-text-field
          v-model="state.address.country"
          :error-messages="v$.address.country.$errors.map(e => e.$message)"
          label="Country"
          required
          @blur="v$.address.country.$touch"
          @input="v$.address.country.$touch"
        ></v-text-field>
        <v-text-field
          v-model="state.address.fullAddress"
          :error-messages="v$.address.country.$errors.map(e => e.$message)"
          label="Full Address"
          required
          @blur="v$.address.country.$touch"
          @input="v$.address.country.$touch"
        ></v-text-field>
      </v-card>
  
      <!-- Rental Details -->
      <v-card title="Rental Details" class="p-4 mt-4">
        <v-text-field
          v-model="state.name"
          :error-messages="v$.name.$errors.map(e => e.$message)"
          label="Rental Name"
          required
        ></v-text-field>
  
        <v-text-field
          v-model="state.price"
          type="number"
          :error-messages="v$.price.$errors.map(e => e.$message)"
          label="Price"
          required
        ></v-text-field>
  
        <v-text-field
          v-model="state.bedrooms"
          type="number"
          :error-messages="v$.bedrooms.$errors.map(e => e.$message)"
          label="Bedrooms"
          required
        ></v-text-field>
  
        <v-text-field
          v-model="state.bathrooms"
          type="number"
          :error-messages="v$.bathrooms.$errors.map(e => e.$message)"
          label="Bathrooms"
          required
        ></v-text-field>
  
        <v-text-field
          v-model="state.areaSize"
          type="number"
          :error-messages="v$.areaSize.$errors.map(e => e.$message)"
          label="Area Size (sq. meters)"
          required
        ></v-text-field>
  
        <v-select
          v-model="state.currency"
          :items="currencyOptions"
          label="Currency"
          item-title="text"
          item-value="value"
          :error-messages="v$.currency.$errors.map(e => e.$message)"
          required
        ></v-select>
  
        <v-select
          v-model="state.standing"
          :items="standingOptions"
          item-title="text"
          item-value="value"
          label="Standing"
          :error-messages="v$.standing.$errors.map(e => e.$message)"
          required
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
  
      <!-- Amenities -->
      <v-card title="Amenities" class="p-4 mt-4">
        <v-checkbox
          v-model="state.amenities.airConditioning"
          label="Air Conditioning"
          required
        ></v-checkbox>
        <v-checkbox
          v-model="state.amenities.heating"
          label="Heating"
          required
        ></v-checkbox>
        <v-checkbox
          v-model="state.amenities.refrigerator"
          label="Refrigerator"
        ></v-checkbox>
        <v-checkbox
          v-model="state.amenities.parking"
          label="Parking"
        ></v-checkbox>
      </v-card>
  
      <!-- Rules -->
      <v-card title="Rules" class="p-4 mt-4">
        <v-checkbox
          v-model="state.rules.petsAllowed"
          label="Pets Allowed"
        ></v-checkbox>
        <v-checkbox
          v-model="state.rules.partiesAllowed"
          label="Parties Allowed"
        ></v-checkbox>
        <v-checkbox
          v-model="state.rules.smokingAllowed"
          label="Smoking Allowed"
        ></v-checkbox>
      </v-card>
  
      <!-- Image Upload Section -->
      <v-card title="Upload Images" class="p-4 mt-4">
        <v-file-input
          v-model="state.images"
          label="Add up to 3 images"
          accept="image/*"
          multiple
          :rules="[fileLimit]"
          required
        ></v-file-input>
      </v-card>

      <div class="mt-4">
  
      <v-btn type="submit" class="me-4">Submit</v-btn>
      <v-btn @click="clear">Clear</v-btn>
    </div>
    </v-form>
  </template>
  
  <script setup>
  import { reactive, ref, provide, watch, onMounted } from 'vue';
import { useVuelidate } from '@vuelidate/core';
import { required, numeric } from '@vuelidate/validators';
import { useAuthStore } from "@/stores/authStore"; // Import the authStore
import Address from "@/components/rentals/form/Address.vue";
import { useRoute } from 'vue-router';
import { getRentalById, updateRental, addRental } from '@/api/rentals';


// Access auth store to get the logged-in user's data
const authStore = useAuthStore();
const loggedInUserId = authStore.user?.id; // Get the user ID from the store


const marker = ref([36.8065, 10.181667]);

function updateMarker(val) {
  marker.value = val;
}

provide('location', { marker, updateMarker });

const route = useRoute()

  const currencyOptions = [
    { text: 'Tunisian Dinar', value: 'TND' },
    { text: 'US Dollar', value: 'USD' },
    { text: 'Euro', value: 'EUR' },
  ];
  
  const standingOptions = [
    { text: 'Economy', value: 'economy' },
    { text: 'Standard', value: 'standard' },
    { text: 'Luxury', value: 'luxury' },
  ];

// Toast Notification State
const toast = reactive({
  show: false,
  message: "",
  color: "",
});

const availabilityOptions = [
  { text: "Available", value: true },
  { text: "Unavailable", value: false },
];
  
const initialState = reactive({
    address: { streetNumber: '', street: '', city: '', country: '' },
    geometry: { lat: marker.value[0], lng: marker.value[1] },
    name: 'aaa',
    price: 22,
    bedrooms: 22,
    bathrooms: 22,
    areaSize: 22,
    currency: 'TND',
    standing: 'standard',
    amenities: { airConditioning: true, heating: true, refrigerator: false, parking: false },
    rules: { petsAllowed: false, partiesAllowed: true, smokingAllowed: false },
    images: [],
    description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
    createdBy: loggedInUserId || null,
    updatedBy: loggedInUserId || null,

  });
  

  const state = reactive({ ...initialState, id:''});

  const rules = {
    name: { required },
    price: { required, numeric },
    bedrooms: { required, numeric },
    bathrooms: { required, numeric },
    areaSize: { required, numeric },
    currency: { required },
    standing: { required },
    amenities: {
      airConditioning: { required },
      heating: { required },
    },
    address: {
      streetNumber: { required },
      street: { required },
      city: { required },
      country: { required },
    },
  description: { required },
  };
  
  const v$ = useVuelidate(rules, state);


watch(marker, (newVal) => {
  state.geometry.lng = newVal[1];
  state.geometry.lat = newVal[0];
}, { deep: true });

const fileLimit = (files) => {
  return files && files.length <= 3 || 'You can only upload up to 3 images';
};
  
  function clear() {
    Object.assign(state, {
      address: { streetNumber: '', street: '', city: '', country: '', fullAddress: '' },
      geometry: { lat: '', lng: '' },
      name: '',
      price: '',
      bedrooms: '',
      bathrooms: '',
      areaSize: '',
      currency: 'TND',
      standing: 'standard',
      amenities: { airConditioning: false, heating: false, refrigerator: false, parking: false },
      rules: { petsAllowed: false, partiesAllowed: false, smokingAllowed: false },
      images: [],
    });
  }

  const isEditMode = ref(false)

  // Check if editing a rental
onMounted(async () => {
  const rentalId = route.params.id;
  if (rentalId && rentalId !== "new") {
    isEditMode.value = true;
    await loadRental(rentalId);
  } else {
    clear()
  }
});

const loadRental = async (id) => {
  try {
    const data = await getRentalById(id)
    for (const key in data) {
      if (key in state) {
        state[key] = data[key];
      }
    }

    //state.images = []
  } catch (error) {
    console.error("Failed to load rental:", error.response?.data || error.message);
    toast.message = "Could not load rental details";
    toast.color = "red";
    toast.show = true;
  }
};

async function submitForm() {
  const isValid = await v$.value.$validate();
  if (!isValid) return;

  const formData = new FormData();

  // Add flat fields
  formData.append("name", state.name);
  formData.append("price", state.price.toString());
  formData.append("bedrooms", state.bedrooms.toString());
  formData.append("bathrooms", state.bathrooms.toString());
  formData.append("areaSize", state.areaSize.toString());
  formData.append("currency", state.currency);
  formData.append("standing", state.standing);
  formData.append("description", state.description || "");
  formData.append("createdBy",  loggedInUserId);

  // Add nested fields for address
  for (const key in state.address) {
    formData.append(`address.${key}`, state.address[key]);
  }

  // Add nested fields for geometry
  for (const key in state.geometry) {
    formData.append(`geometry.${key}`, state.geometry[key]);
  }

  // Add nested fields for amenities
  for (const key in state.amenities) {
    formData.append(`amenities.${key}`, state.amenities[key] ? "true" : "false");
  }

  // Add nested fields for rules
  for (const key in state.rules) {
    formData.append(`rules.${key}`, state.rules[key] ? "true" : "false");
  }

  // Add images
  state.images.forEach((image) => {
    formData.append("images", image);
  });

  try {
    if (!isEditMode.value) {
      // Add Rental
      console.log(formData)
      const response = await addRental(formData);
      toast.message = "Rental added successfully!";
      console.log("Rental added successfully:", response.data);
    } else {
      // Update Rental
      const response = await updateRental(state.id, formData);
      toast.message = "Rental updated successfully!";
      console.log("Rental updated successfully:", response.data);
    }

    toast.color = "success";
    toast.show = true;

    clear();
  } catch (error) {
    console.error("Failed to save rental:", error.response?.data || error.message);
    toast.message = "Could not save rental";
    toast.color = "error";
    toast.show = true;
  }
}


provide('updateInfo', (newInfo) => {
  state.address.streetNumber = newInfo.streetNumber || '';
  state.address.street = newInfo.street || '';
  state.address.city = newInfo.city || '';
  state.address.country = newInfo.country || '';
  state.address.fullAddress = newInfo.fullAddress || '';
});
  </script>
  