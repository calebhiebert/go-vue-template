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
	{
		path: "/admin",
		name: "Admin",
		component: () => import("../views/admin/Parent.vue"),
		beforeEnter: authGuard,
		children: [
			{
				path: "dashboard",
				name: "AdminDashboard",
				component: () => import("../views/admin/Dashboard.vue"),
			},
			{
				path: "model/:model",
				name: "AdminModelView",
				component: () => import("../views/admin/ModelView.vue"),
			},
			{
				path: "edit/:model/:id",
				name: "AdminModelEdit",
				component: () => import("../views/admin/ModelEdit.vue"),
			},
		],
	},
];

const router = new VueRouter({
	mode: "history",
	base: process.env.BASE_URL,
	routes,
});

export default router;
