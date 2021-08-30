<template>
    <b-navbar>
        <template #brand>
            <b-navbar-item tag="router-link" :to="{ name: 'Home' }">
                <img
                    src="https://raw.githubusercontent.com/buefy/buefy/dev/static/img/buefy-logo.png"
                    alt="Lightweight UI components for Vue.js based on Bulma"
                >
            </b-navbar-item>
        </template>
        <template #start>
            <b-navbar-item tag="router-link" :to="{ name: 'Home' }">
                Home
            </b-navbar-item>
        </template>

        <template #end>
            <b-navbar-item tag="div">
                <figure style="width: 28px; height: 36px; padding-top: 4px;" class="mr-2">
                    <b-image :responsive="true" ratio="1by1" v-if="authenticated && user !== null" :rounded="true"
                             :src="`${apiUrl}/avatar/${user.id}`">
                    </b-image>
                </figure>

                <div class="buttons">
                    <router-link v-if="!authenticated" :to="{name: 'Register'}" class="button is-primary">
                        <strong>Register</strong>
                    </router-link>
                    <router-link v-if="!authenticated" :to="{name: 'Login'}" class="button is-light">
                        Log in
                    </router-link>

                    <button @click="doLogout" v-if="authenticated" class="button is-light">
                        Log out
                    </button>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>

<script>
import {API_BASE_URL, logout} from "../api";

export default {
    name: "NavBar",

    props: {
        authenticated: {
            type: Boolean,
            required: true,
        },

        user: {
            type: Object,
            required: false,
        },
    },

    data() {
        return {
            apiUrl: API_BASE_URL,
        };
    },

    methods: {
        doLogout() {
            logout();
            this.$router.push({name: "Login"});
        },
    },
};
</script>

<style scoped>

</style>