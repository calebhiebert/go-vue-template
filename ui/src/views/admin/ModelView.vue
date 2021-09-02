<template>
    <div>
        <b-table :loading="loading" v-if="tableColumns !== null" :data="tableData === null ? [] : tableData" :columns="tableColumns"></b-table>
    </div>
</template>

<script>
import axios from "axios";
import {API_BASE_URL, extractError, getToken} from "../../api";

export default {
    name: "ModelView",

    data() {
        return {
            loading: false,
            data: null,
            limit: 25,
            page: 0,

            tableData: null,
            tableColumns: null,
        }
    },

    mounted() {
        this.load();
    },

    methods: {
        async load() {
            this.loading = true;

            console.log(window._adminSchema.models[this.modelId]);
            const modelSchema = window._adminSchema.models[this.modelId];

            this.tableColumns = this.constructTableColumns(modelSchema);

            try {
                const r = await axios.get(`${API_BASE_URL}/crud/${modelSchema.url_name}`, {
                    headers: {
                        'Authorization': `Bearer ${getToken()}`
                    }
                });

                console.log(r.data);

                this.tableData = this.constructTableData(modelSchema, r.data);
            } catch (e) {
                console.log(extractError(e));
            }

            this.loading = false;
        },

        constructTableColumns(schema) {
            const columns = [];

            for (const field of schema.fields) {
                if (!field.config.show_on_graph) {
                    continue;
                }

                columns.push({
                    field: field.id,
                    label: field.name,
                    numeric: field.type === "int",
                });
            }

            return columns;
        },

        constructTableData(schema, d) {
            return d[schema.data_name];
        }
    },

    watch: {
        modelId() {
            this.load();
        },

        page() {
            this.load();
        }
    },

    computed: {
        modelId() {
            return this.$route.params.model;
        }
    }
}
</script>

<style scoped>

</style>