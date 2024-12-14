<script setup lang="ts">
import { reactive, ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useVuelidate } from "@vuelidate/core";
import { required, email } from "@vuelidate/validators";
import { useAuthStore } from "@/stores/authStore";

// Pinia store
const authStore = useAuthStore();

// Reactive state for the form
const state = reactive({
  email: "",
  password: "",
});

// Validation rules
const rules = {
  email: { required, email },
  password: { required },
};

// Initialize Vuelidate
const v$ = useVuelidate(rules, state);

// Router
const router = useRouter();

// Toast Notification State
const toast = reactive({
  show: false,
  message: "",
  color: "",
});

// Error handling
const error = ref<string | null>(null);

// Submit form handler
const submitForm = async () => {
  error.value = null;

  // Validate the form
  const isValid = await v$.value.$validate();
  if (!isValid) {
    error.value = "Please complete all fields correctly.";
    return;
  }

  try {
    // Call the login action in authStore
    await authStore.login(state);

    // Show success toast
    toast.message = "Login successful! Redirecting...";
    toast.color = "success";
    toast.show = true;

    // Redirect to the dashboard or home page after 2 seconds
    setTimeout(() => {
      router.push("/rentals");
    }, 2000);
  } catch (err: any) {
    error.value = "Login failed. Please try again.";
    toast.message = error.value;
    toast.color = "error";
    toast.show = true;
  }
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

<template>
  <v-container class="py-5">
    <v-card class="mx-auto" max-width="500">
      <v-card-title class="text-center mb-4">
        <span class="text-h5">Sign in to Darwin</span>
      </v-card-title>
      <v-card-text>
        <!-- Error Alert -->
        <!-- Google Login Button -->
        <div id="googleLoginButton" class="mb-4"></div>
      <v-divider thickness="4px" opacity="100">Or</v-divider>
        <!-- Login Form -->

        <v-alert
          v-if="error"
          type="error"
          border="bottom"
          class="my-2"
        >
          {{ error }}
        </v-alert>
        <v-form class="mt-4" @submit.prevent="submitForm" fast-fail>
          <!-- Email Field -->
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

          <!-- Password Field -->
          <v-text-field
            v-model="state.password"
            :error-messages="v$.password.$errors.map((e) => e.$message)"
            label="Password"
            required
            outlined
            dense
            type="password"
            @blur="v$.password.$touch"
          ></v-text-field>
        </v-form>

        <v-btn class="w-full" color="primary" variant="flat" @click="submitForm">Login</v-btn>
        <div class="text-center mt-4">
          no account ?<v-btn to="/user/register" variant="plain">register</v-btn>
        </div>

      </v-card-text>

    </v-card>

    <!-- Toast Notification -->
    <v-snackbar
      v-model="toast.show"
      :color="toast.color"
      :timeout="10000"
    >
      {{ toast.message }}
    </v-snackbar>
  </v-container>
</template>

<style scoped>
.text-h5 {
  font-weight: bold;
}
</style>
