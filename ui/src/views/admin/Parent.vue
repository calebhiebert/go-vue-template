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
                            <b-menu-item icon="information-outline" label="Info"></b-menu-item>
                            <b-menu-item active expanded label="Models" v-if="schema !== null">
                                <b-menu-item :label="model.name" @click="openModel(id)" v-for="(model, id) in schema.models" :key="id"></b-menu-item>
                            </b-menu-item>
                            <b-menu-item icon="account" label="My Account">
                                <b-menu-item icon="account-box" label="Account data"></b-menu-item>
                                <b-menu-item icon="home-account" label="Addresses"></b-menu-item>
                            </b-menu-item>
                        </b-menu-list>
                        <b-menu-list>
                            <b-menu-item label="Expo" icon="link"></b-menu-item>
                        </b-menu-list>
                        <b-menu-list label="Actions">
                            <b-menu-item icon="logout" label="Logout"></b-menu-item>
                        </b-menu-list>
                    </b-menu>
                </div>
            </b-sidebar>
            <div class="content">
                <router-view></router-view>
            </div>
        </div>

    </div>
</template>

<script>
import {getAdminSchema} from "../../api";
import axios from 'axios';

export default {
    name: "Parent",

    mounted() {
        this.loadSchema();
    },

    data() {
       return {
           loadingSchema: false,
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
}
</style>