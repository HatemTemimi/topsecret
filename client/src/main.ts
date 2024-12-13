import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import './index.css'
import 'vuetify/styles'
import {createVuetify, type ThemeDefinition} from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import '@mdi/font/css/materialdesignicons.css'
import { createPinia } from "pinia";
import { useAuthStore } from './stores/authStore'

const lemonade: ThemeDefinition = {
    dark: false,
    colors: {
      background: '#ffffff', // Base-100
      surface: '#f5f5f5', // Base-200
      primary: '#f4b8c0', // Primary
      'primary-darken-1': '#f7d4c4', // Secondary (used for a slightly darker shade)
      secondary: '#f9e4c8', // Accent
      'secondary-darken-1': '#e5e5e5', // Base-300
      error: '#ff5724', // Error
      info: '#2094f3', // Info
      success: '#009485', // Success
      warning: '#ff9900', // Warning
    },
};

const vuetify = createVuetify({
    components,
    directives,
    theme: {
    defaultTheme: 'lemonade',
    themes: {
      lemonade,
    },
  },
})

const pinia = createPinia();

const app = createApp(App)
app.use(vuetify)
app.use(router)
app.use(pinia)

const authStore = useAuthStore();
authStore.loadSession();

app.mount('#app')

