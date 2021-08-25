<template>
    <div class="container mt-4">
        <div class="card">
            <div class="card-content" v-if="user">
                Welcome, {{ user.name }}
            </div>

            <b-loading v-model="loading"></b-loading>
        </div>
    </div>
</template>

<script>
import {usersMe} from "../api";

export default {
    name: "Home",

    data() {
        return {
            user: null,
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
                console.log(e);
            }

            this.loading = false;
        }
    }
};
</script>

<style scoped>

</style>