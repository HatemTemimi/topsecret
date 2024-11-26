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



const vuetify = createVuetify({
    components,
    directives,
    theme: {
        defaultTheme: 'dark'
    }
})

const pinia = createPinia();

createApp(App)
    .use(vuetify)
    .use(router)
    .use(pinia)
    .mount('#app')

