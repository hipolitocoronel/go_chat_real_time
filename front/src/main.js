import { createApp } from "vue";
import { createPinia } from "pinia";
import router from "./router";
import App from "./App.vue";
import "./assets/main.css";
import vue3GoogleLogin from "vue3-google-login";

const GOOGLE_CLIENT_ID =
  "417404445638-sdnhtmpkf4f1l4ncpdj8isdjma6gvlau.apps.googleusercontent.com";

const pinia = createPinia();
const app = createApp(App);

app.use(router);
app.use(pinia);
app.use(vue3GoogleLogin, { clientId: GOOGLE_CLIENT_ID });

app.mount("#app");
