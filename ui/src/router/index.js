import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/Login";
import {isAuthenticated} from "../api";
import Register from "../views/Register";

Vue.use(VueRouter);

const authGuard = (to, from, next) => {
	if (!isAuthenticated()) {
		next({name: "Login"});
	} else {
		next();
	}
};

const routes = [
	{
		path: "/",
		name: "Home",
		component: Home,
		beforeEnter: authGuard,
	},
	{
		path: "/login",
		name: "Login",
		component: Login,
	},
	{
		path: "/register",
		name: "Register",
		component: Register,
	},
];

const router = new VueRouter({
	mode: "history",
	base: process.env.BASE_URL,
	routes,
});

export default router;
