import {createApp} from 'vue'
import App from './App.vue'
import router from './router'
import './index.css'
import 'vuetify/styles'
import {createVuetify} from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import '@mdi/font/css/materialdesignicons.css'
import { createPinia } from "pinia";
import { useAuthStore } from './stores/authStore'



const vuetify = createVuetify({
    components,
    directives,
    theme: {
        defaultTheme: 'dark'
    }
})

const pinia = createPinia();

const app = createApp(App)
app.use(vuetify)
app.use(router)
app.use(pinia)

const authStore = useAuthStore();
authStore.loadSession();

app.mount('#app')

