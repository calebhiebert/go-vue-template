import Vue from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";

import Buefy from "buefy";
import "buefy/dist/buefy.css";
import './styles.css';

// Setup fontawesome icons
import {library} from "@fortawesome/fontawesome-svg-core";
import {
    faAngleLeft,
    faAngleRight,
    faArrowUp,
    faEnvelope,
    faExclamationCircle,
    faLock,
    faPlus,
    faSyncAlt,
    faTimesCircle,
    faTrash,
} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";

// Setup vee validate
import {confirmed, email, length, required} from "vee-validate/dist/rules";
import {extend, ValidationProvider} from "vee-validate";

library.add(faLock, faEnvelope, faExclamationCircle, faArrowUp, faAngleLeft, faAngleRight, faSyncAlt, faTrash, faPlus, faTimesCircle);

Vue.component("font-awesome-icon", FontAwesomeIcon);
Vue.component("ValidationProvider", ValidationProvider);

Vue.use(Buefy, {
    defaultIconPack: "fas",
    defaultIconComponent: "font-awesome-icon",
});

extend("required", {
    ...required,
    message: "This field is required",
});

extend("email", {
    ...email,
    message: "This field must be a valid email",
});

extend("confirmed", {
    ...confirmed,
    message: "This field confirmation does not match",
});

extend("length", {
    ...length,
    message: "This field must have 2 options",
});

Vue.config.productionTip = false;

new Vue({
    router,
    render: h => h(App),
}).$mount("#app");
