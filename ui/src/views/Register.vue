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

                                    <ValidationProvider v-slot="{ errors }">
                                        <BField label="Email" :message="errors">
                                            <b-input icon="envelope" v-model="email" required type="email"
                                                     placeholder="Email"></b-input>
                                        </BField>
                                    </ValidationProvider>

                                    <ValidationProvider v-slot="{ errors }">
                                        <BField label="Name" :message="errors">
                                            <b-input v-model="name" required placeholder="Name"
                                                     maxlength="30"></b-input>
                                        </BField>
                                    </ValidationProvider>


                                    <ValidationProvider rules="confirmed:confirmPassword" v-slot="{errors}">
                                        <BField label="Password" :message="errors">
                                            <b-input icon="lock" type="password" v-model="password"
                                                     placeholder="********" required></b-input>
                                        </BField>
                                    </ValidationProvider>

                                    <ValidationProvider vid="confirmPassword">
                                        <BField label="Confirm Password">
                                            <b-input icon="lock" type="password" v-model="confirmPassword"
                                                     placeholder="Confirm Password" required></b-input>
                                        </BField>
                                    </ValidationProvider>

                                    <b-notification :closable="false" type="is-danger" v-if="error">
                                        {{ error }}
                                    </b-notification>


                                    <div class="columns mt-4">
                                        <div class="column">
                                            <b-button native-type="submit" :loading="loading">
                                                Register
                                            </b-button>
                                        </div>
                                        <div class="column has-text-centered">
                                            <router-link to="/login">Log in instead</router-link>
                                        </div>
                                    </div>

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