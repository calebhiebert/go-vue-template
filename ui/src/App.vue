<template>
  <div id="app">
      <NavBar v-if="showNavbar" :authenticated="authenticated" :user="currentUser" />
    <router-view/>
  </div>
</template>

<script>
    import NavBar from "./components/NavBar";
    import {extractError, isAuthenticated, usersMe} from "./api";
    export default {
        name: "App",
        components: {NavBar},
        computed: {
            showNavbar() {
                return !['Login', 'Register', 'Admin', 'AdminDashboard'].includes(this.$route.name);
            }
        },

        created() {
            this.revalidateAuthenticationStatus();
        },

        data() {
            return {
                authenticated: false,
                currentUser: null,
            }
        },

        methods: {
            revalidateAuthenticationStatus() {
                this.authenticated = isAuthenticated();

                if (!this.authenticated && this.$route.name !== "Login") {
                    this.$router.push({name: "Login"});
                }

                if (this.authenticated && this.currentUser === null) {
                    this.fetchCurrentUser();
                }
            },

            async fetchCurrentUser() {
                try {
                    const result = await usersMe();

                    this.currentUser = result.data;
                } catch (e) {
                    if (e.response.statusCode === 401) {
                        this.$router.push({name: "Login"});
                    }

                    this.error = extractError(e);
                }
            }
        },

        watch: {
            '$route.name'() {
                this.revalidateAuthenticationStatus();
            }
        }
    }
</script>

<style>
</style>
