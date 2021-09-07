<template>
    <div class="container">
        <h2 class="ml-4 title is-2">Dashboard</h2>

        <div v-if="stats">
            <nav class="level">
                <div class="level-item has-text-centered">
                    <div>
                        <p class="heading">Total Users</p>
                        <p class="title" v-if="stats">{{ stats.total_users }}</p>
                        <p class="title" v-else><b-skeleton position="is-centered" height="30" width="75%"></b-skeleton></p>
                    </div>
                </div>
                <div class="level-item has-text-centered">
                    <div>
                        <p class="heading">New Users (Week)</p>
                        <p class="title" v-if="stats">{{ stats.weekly_new_users }}</p>
                        <p class="title" v-else><b-skeleton position="is-centered" height="30" width="75%"></b-skeleton></p>
                    </div>
                </div>
                <div class="level-item has-text-centered">
                    <div>
                        <p class="heading">New Users (Month)</p>
                        <p class="title" v-if="stats">{{ stats.monthly_new_users }}</p>
                        <p class="title" v-else><b-skeleton position="is-centered" height="30" width="75%"></b-skeleton></p>
                    </div>
                </div>
            </nav>
        </div>

        <b-loading v-if="loading"></b-loading>
    </div>
</template>

<script>
import {getAdminDashboardStats} from "../../api";

export default {
    name: "Dashboard",

    data() {
        return {
            loading: false,
            stats: null,
        }
    },

    created() {
        this.getStats();
    },

    methods: {
        async getStats() {
            this.loading = true;

            try {
                const stats = await getAdminDashboardStats();
                this.stats = stats.data;
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