<template>
    <div>
        <section class="hero is-primary is-fullheight">
            <div class="hero-body">
                <div class="container">
                    <div class="columns is-centered">
                        <div class="column is-5-tablet is-4-desktop is-3-widescreen">
                            <form @submit.prevent="doLogin" class="box">
                                <div class="title is-4 has-text-dark">Login</div>

                                <fieldset :disabled="loading">
                                    <b-field label="Username or Email" label-position="inside">
                                        <b-input icon="envelope" v-model="username" required
                                                 placeholder="Username or Email"></b-input>
                                    </b-field>
                                    <b-field label="Password" label-position="inside">
                                        <b-input icon="lock" type="password" v-model="password"
                                                 placeholder="********" required></b-input>
                                    </b-field>
                                    <b-notification :closable="false" type="is-danger" v-if="error">
                                        {{ error }}
                                    </b-notification>

                                    <p class="mb-4"><router-link to="/register">Register instead</router-link></p>

                                    <b-button native-type="submit" :loading="loading">
                                        Login
                                    </b-button>
                                </fieldset>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>

</template>

<script>
import {extractError, loginUsernamePassword} from "../api";

export default {
    name: "Login",

    data() {
        return {
            username: "",
            password: "",
            loading: false,
            error: null,
        };
    },

    methods: {
        async doLogin() {
            this.loading = true;
            this.error = null;

            try {
                await loginUsernamePassword(this.username, this.password);
                this.$router.push({name: "Home"});
            } catch (e) {
                this.error = extractError(e).message;
            }

            this.loading = false;
        },
    },
};
</script>

<style scoped>

</style>