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

                if (!this.authenticated && !['Login', 'Register'].includes(this.$route.name)) {
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

        computed: {
            showNavbar() {
                if (this.$route.name && this.$route.name.startsWith("Admin")) {
                    return false;
                }

                return !['Login', 'Register', 'Admin', 'AdminDashboard'].includes(this.$route.name);
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
