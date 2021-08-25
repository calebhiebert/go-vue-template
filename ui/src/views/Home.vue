<template>
    <div class="container mt-4">
        <div class="card">
            <!-- Content -->
            <div class="card-content" v-if="user">
                Welcome, {{ user.name }}

                <p class="mt-4">
                    <b-button @click="logout">Logout</b-button>
                </p>
            </div>

            <!-- Skeleton Loader -->
            <div class="card-content" v-else>
                <b-skeleton width="20%"></b-skeleton>
            </div>
        </div>
    </div>
</template>

<script>
import {extractError, logout, usersMe} from "../api";

export default {
    name: "Home",

    data() {
        return {
            user: null,
            error: null,
            loading: false,
        }
    },

    mounted() {
        this.loadUserData();
    },

    methods: {
        async loadUserData() {
            this.loading = true;

            try {
                const result = await usersMe();

                this.user = result.data;
            } catch (e) {
                this.error = extractError(e);
            }

            this.loading = false;
        },

        logout() {
            logout();

            this.$router.push({name: "Login"});
        }
    }
};
</script>

<style scoped>

</style>