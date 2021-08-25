import Vue from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";

import {library} from "@fortawesome/fontawesome-svg-core";
import {faEnvelope, faExclamationCircle, faLock} from "@fortawesome/free-solid-svg-icons";
import Buefy from "buefy";
import "buefy/dist/buefy.css";
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";

library.add(faLock, faEnvelope, faExclamationCircle);

Vue.component("font-awesome-icon", FontAwesomeIcon);

Vue.use(Buefy, {
	defaultIconPack: "fas",
	defaultIconComponent: "font-awesome-icon",
});

Vue.config.productionTip = false;

new Vue({
	router,
	render: h => h(App),
}).$mount("#app");
