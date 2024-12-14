<script lang="ts">
import { onMounted, reactive } from 'vue';
import { useRouter } from 'vue-router';



// Toast Notification State
const toast = reactive({
  show: false,
  message: "",
  color: "",
});

// Router
const router = useRouter();


const handleGoogleLogin = async (response: any) => {
  try {
    const idToken = response.credential; // Get the ID token from Google

    // Send ID token to the backend for verification and login
    //await authStore.googleLogin(idToken);

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
        <div id="googleLoginButton" class="mb-4"></div>
</template>