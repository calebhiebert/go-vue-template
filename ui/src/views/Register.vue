<template>
    <div>
        <section class="hero is-primary is-fullheight">
            <div class="hero-body">
                <div class="container">
                    <div class="columns is-centered">
                        <div class="column is-5-tablet is-4-desktop is-3-widescreen">
                            <form @submit.prevent="doRegister" class="box">
                                <div class="title is-4 has-text-dark">Register</div>

                                <fieldset :disabled="loading">
                                    <b-field label="Name" label-position="inside">
                                        <b-input v-model="name" required placeholder="Name"></b-input>
                                    </b-field>
                                    <b-field label="Email" label-position="inside">
                                        <b-input icon="envelope" v-model="email" required type="email"
                                                 placeholder="Email"></b-input>
                                    </b-field>
                                    <b-field label="Password" label-position="inside">
                                        <b-input icon="lock" type="password" v-model="password"
                                                 placeholder="********" required></b-input>
                                    </b-field>
                                    <b-field label="Confirm Password" label-position="inside">
                                        <b-input icon="lock" type="password" v-model="confirmPassword"
                                                 placeholder="Confirm Password" required></b-input>
                                    </b-field>
                                    <b-notification :closable="false" type="is-danger" v-if="error">
                                        {{ error }}
                                    </b-notification>

                                    <p class="mb-4"><router-link to="/login">Log in instead</router-link></p>

                                    <b-button native-type="submit" :loading="loading">
                                        Register
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
import {extractError, register} from "../api";

export default {
    name: "Register",

    data() {
        return {
            email: "",
            name: "",
            password: "",
            confirmPassword: "",
            loading: false,
            error: null,
        };
    },

    methods: {
        async doRegister() {
            this.loading = true;
            this.error = null;

            try {
                await register({
                    name: this.name,
                    email: this.email,
                    login: this.email,
                    password: this.password,
                });

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