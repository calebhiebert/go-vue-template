<template>
    <div>
        <b-table checkable
                 :loading="loading"
                 v-if="tableColumns !== null"
                 :data="tableData === null ? [] : tableData"
                 :total="data !== null ? data.total : 0"
                 :per-page="limit"
                 paginated
                 backend-pagination
                 @page-change="onPageChange"
                 aria-next-label="Next page"
                 aria-previous-label="Previous page"
                 aria-page-label="Page"
                 aria-current-label="Current page"
        >
            <b-table-column
                v-for="col in this.tableColumns"
                :key="col.field"
                :field="col.field"
                :label="col.label"
                v-slot="props">
                {{ props.row.original_title }}
                <component :is="getCustomColumnComponent(col.field)" :col="col" :p="props" :row="props.row"></component>
            </b-table-column>

        </b-table>
    </div>
</template>

<script>
import axios from "axios";
import {API_BASE_URL, extractError, getToken} from "../../api";
import ColumnViewDefault from "../../components/admin/ColumnViewDefault";

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
            customColumnComponents: {}
        }
    },

    created() {
        this.registerCustomColumnComponent("*", ColumnViewDefault);
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
                const r = await axios.get(`${API_BASE_URL}/crud/${modelSchema.url_name}?limit=${this.limit}&offset=${this.limit * this.page}`, {
                    headers: {
                        'Authorization': `Bearer ${getToken()}`
                    }
                });

                console.log(r.data);
                this.data = r.data;
                this.tableData = this.constructTableData(modelSchema, r.data);
            } catch (e) {
                console.log(extractError(e));
            }

            this.loading = false;
        },

        onPageChange(page) {
            this.page = page;
            this.load();
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
        },

        registerCustomColumnComponent(selector, component) {
            this.customColumnComponents[selector] = component;
        },

        getCustomColumnComponent(fieldName) {
            for (const selector in this.customColumnComponents) {
                if (selector === `${this.modelId}.${fieldName}`) {
                    return this.customColumnComponents[fieldName];
                }
            }

            return this.customColumnComponents["*"];
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