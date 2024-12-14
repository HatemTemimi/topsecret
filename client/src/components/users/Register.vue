<template>
  <v-container class="py-5">
    <v-card class="mx-auto" max-width="800">
      <v-card-title class="text-center mb-4">
        <span class="text-h5">Register to Darwin</span>
      </v-card-title>
      <v-card-text>

        <div id="googleLoginButton" class="mb-4"></div>
          <v-divider thickness="4px" opacity="100">Or</v-divider>
        <!-- Login Form -->
        <v-form class="mt-4" @submit.prevent="submitForm" fast-fail>
          <v-text-field
            v-model="state.firstName"
            :error-messages="v$.firstName.$errors.map((e) => e.$message)"
            label="First Name"
            required
            outlined
            dense
            @blur="v$.firstName.$touch"
          ></v-text-field>

          <v-text-field
            v-model="state.lastName"
            :error-messages="v$.lastName.$errors.map((e) => e.$message)"
            label="Last Name"
            required
            outlined
            dense
            @blur="v$.lastName.$touch"
          ></v-text-field>

          <v-text-field
            v-model="state.email"
            :error-messages="v$.email.$errors.map((e) => e.$message)"
            label="Email"
            required
            outlined
            dense
            type="email"
            @blur="v$.email.$touch"
          ></v-text-field>

          <v-text-field
            v-model="state.password"
            :error-messages="v$.password.$errors.map((e) => e.$message)"
            label="Password"
            required
            outlined
            dense
            :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
            @click:append="show1 = !show1"
            :type="show1 ? 'text' : 'password'"
            @blur="v$.password.$touch"
          ></v-text-field>

          <v-text-field
            v-model="state.confirmPassword"
            :error-messages="v$.confirmPassword.$errors.map((e) => e.$message)"
            :append-icon="show2 ? 'mdi-eye' : 'mdi-eye-off'"
            :type="show2 ? 'text' : 'password'"
             @click:append="show2 = !show2"
            label="Confirm Password"
            required
            outlined
            dense
            @blur="v$.confirmPassword.$touch"
          ></v-text-field>

          <v-text-field
            v-model="state.phone"
            :error-messages="v$.phone.$errors.map((e) => e.$message)"
            label="Phone"
            outlined
            dense
            @blur="v$.phone.$touch"
          ></v-text-field>

          <v-textarea
            v-model="state.address"
            :error-messages="v$.address.$errors.map((e) => e.$message)"
            label="Address"
            outlined
            dense
            @blur="v$.address.$touch"
          ></v-textarea>

        <!--<v-btn outlined color="secondary" @click="clearForm">Clear</v-btn>-->

        </v-form>

        <v-btn class="w-full" color="primary" @click="submitForm">Register</v-btn>
        <div class="text-center mt-4">
          already have an account ?<v-btn to="/user/login" variant="plain">Sign in</v-btn>
        </div>
      </v-card-text>

    </v-card>

    <!-- Snackbar for Success or Error Notifications -->
    <v-snackbar
      v-model="snackbar.show"
      :color="snackbar.color"
      :timeout="2000"
      bottom
    >
      {{ snackbar.message }}
    </v-snackbar>
  </v-container>
</template>

<script setup lang="ts">
import { reactive, ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useVuelidate } from "@vuelidate/core";
import { required, email, minLength, sameAs } from "@vuelidate/validators";
import axios from "axios";

// Custom function to validate that passwords match
const passwordsMatch = (value: string, state: any) => {
  return value === state.password || "Passwords do not match.";
};

// Router
const router = useRouter();

// Reactive state for the form
const state = reactive({
  firstName: "",
  lastName: "",
  email: "",
  password: "",
  confirmPassword: "",
  phone: "",
  address: "",
  role: "user", // Default role
});

// Show/Hide Password
const show1 = ref(false);
// Show/Hide Password
const show2 = ref(false);

// Validation rules
const rules = {
  firstName: { required },
  lastName: { required },
  email: { required, email },
  password: { required, minLength: minLength(6) },
  confirmPassword: {
    required,
    sameAsPassword: sameAs(computed(()=> state.password)),
  },
  phone: { required },
  address: { required },
};

// Initialize Vuelidate
const v$ = useVuelidate(rules, state);

// Error and Snackbar State
const error = ref<string | null>(null);
const snackbar = reactive({
  show: false,
  message: "",
  color: "",
});

// Submit form handler
const submitForm = async () => {
  error.value = null;

  // Validate the form
  const isValid = await v$.value.$validate();
  if (!isValid) {
    error.value = "Please complete all fields correctly.";
    snackbar.message = error.value
    snackbar.color = "error";
    snackbar.show = true;
    return;
  }

  try {
    // Send registration data to the backend
    await axios.post("http://localhost:3001/users/create", state);

    // Show success snackbar
    snackbar.message = "Registration successful! Redirecting to login...";
    snackbar.color = "success";
    snackbar.show = true;

    // Redirect to login page after 2 seconds
    setTimeout(() => {
      router.push("/user/login");
    }, 2000);
  } catch (err: any) {
    error.value = err.response?.data?.error || "Registration failed.";

    // Show error snackbar
    snackbar.message = error.value;
    snackbar.color = "error";
    snackbar.show = true;
  }
};

// Clear the form
const clearForm = () => {
  v$.value.$reset();
  for (const key in state) {
    state[key] = "";
  }
  state.role = "user"; // Reset role to default
  error.value = null;
};


// Google Login handler
const handleGoogleLogin = async (response: any) => {
  try {
    const idToken = response.credential; // Get the ID token from Google

    // Send ID token to the backend for verification and login
    await authStore.googleLogin(idToken);

    // Show success toast
    toast.message = "Google Login successful! Redirecting...";
    toast.color = "success";
    toast.show = true;

    // Redirect to the dashboard or home page after 2 seconds
    setTimeout(() => {
      router.push("/rentals");
    }, 2000);
  } catch (err: any) {
    console.log(err)
    error.value = "Google Login failed. Please try again.";
    toast.message = error.value;
    toast.color = "error";
    toast.show = true;
  }
};

// Initialize Google Login
onMounted(() => {
  google.accounts.id.initialize({
    client_id: "428909233068-npvor85d7nu27i0ulm5tp278h5ajagtf.apps.googleusercontent.com", // Replace with your Google Client ID
    callback: handleGoogleLogin,
  });
  google.accounts.id.renderButton(
    document.getElementById("googleLoginButton"),
    { theme: "outline", size: "large" } // Customize the button appearance
  );
});
</script>

<style scoped>
.text-h5 {
  font-weight: bold;
}
</style>
