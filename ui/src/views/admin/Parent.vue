<template>
    <div>
        <div class="sidebar-page">
            <b-sidebar
                position="static"
                :fullheight="true"
                type="is-light" open>
                <div class="p-1">
                    <b-menu>
                        <b-menu-list label="Menu">
                            <b-menu-item label="Home" @click="$router.push({name: 'AdminDashboard'})"></b-menu-item>
                            <b-menu-item expanded label="Models" v-if="schema !== null">
                                <b-menu-item :active="$route.name === 'AdminModelView' && $route.params.model === model.id" :label="model.name" @click="openModel(id)" v-for="(model, id) in schema.models" :key="id"></b-menu-item>
                            </b-menu-item>
                        </b-menu-list>
                    </b-menu>
                </div>
            </b-sidebar>
            <div class="content">
                <router-view v-if="!loadingSchema"></router-view>
            </div>
        </div>

    </div>
</template>

<script>
import {getAdminSchema} from "../../api";

export default {
    name: "Parent",

    mounted() {
        this.loadSchema();
    },

    data() {
       return {
           loadingSchema: true,
           schema: null,
       }
    },

    methods: {
        async loadSchema() {
            this.loadingSchema = true;

            try {
                const r = await getAdminSchema();
                this.schema = r.data;
                window._adminSchema = this.schema;
            } catch (e) {
                console.log(e);
            }

            this.loadingSchema = false;
        },

        openModel(m) {
            if (this.$route.name && this.$route.name === "AdminModelView" && this.$route.params.model === m) {
                return
            }

            this.$router.push({name: 'AdminModelView', params: {model: m}});
        }
    }
};
</script>

<style scoped>
.sidebar-page {
    display: flex;
    width: 100%;
    min-height: 100vh;
}

.content {
    width: 100%;
    margin-left: 1em;
    margin-top: 1em;
    margin-right: 1em;
    margin-bottom: 3em;
}
</style>