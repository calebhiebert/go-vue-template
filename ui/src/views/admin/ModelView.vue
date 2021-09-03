<template>
    <div>
        <b-table checkable
                 :loading="loading"
                 v-if="tableColumns !== null"
                 :data="tableData === null ? [] : tableData"
                 :total="data !== null ? data.total : 0"
                 :per-page="limit"

                 detailed
                 detail-key="id"
                 :show-detail-icon="true"

                 paginated
                 backend-pagination
                 @page-change="onPageChange"
                 aria-next-label="Next page"
                 aria-previous-label="Previous page"
                 aria-page-label="Page"
                 aria-current-label="Current page"

                 :default-sort-direction="defaultSortOrder"
                 :default-sort="[sortField, sortOrder]"
                 backend-sorting
                 @sort="onSort"
        >
            <b-table-column
                v-for="col in this.tableColumns"
                :key="col.field"
                :field="col.field"
                :label="col.label"
                sortable
                v-slot="props">
                {{ props.row.original_title }}
                <component :is="getCustomColumnComponent(col.field)" :col="col" :field="col.schemaField" :p="props" :row="props.row"></component>
            </b-table-column>

            <template #detail="props">
                <ModelvViewSingle :model="props.row" :schema="modelSchema"></ModelvViewSingle>
            </template>
        </b-table>
    </div>
</template>

<script>
import axios from "axios";
import {API_BASE_URL, extractError, getToken} from "../../api";
import ColumnViewDefault from "../../components/admin/ColumnViewDefault";
import ColumnViewHTTPResponseCode from "../../components/admin/ColumnViewHTTPResponseCode";
import ModelvViewSingle from "../../components/admin/ModelViewSingle";

export default {
    name: "ModelView",
    components: {ModelvViewSingle},
    data() {
        return {
            loading: false,
            data: null,
            limit: 25,
            page: 0,
            modelSchema: null,
            tableData: null,
            tableColumns: null,
            customColumnComponents: {},

            sortField: null,
            sortOrder: 'desc',
            defaultSortOrder: 'desc',
        }
    },

    created() {
        this.registerCustomColumnComponent("*", ColumnViewDefault);
        this.registerCustomColumnComponent("access_logs.response_code", ColumnViewHTTPResponseCode);
    },

    mounted() {
        this.load();
    },

    methods: {
        async load() {
            this.loading = true;

            this.modelSchema = window._adminSchema.models[this.modelId];

            this.tableColumns = this.constructTableColumns(this.modelSchema);

            try {
                const r = await axios.get(`${API_BASE_URL}/crud/${this.modelSchema.url_name}?limit=${this.limit}&offset=${this.limit * this.page}&sort.${this.sortField}=${this.sortOrder}`, {
                    headers: {
                        'Authorization': `Bearer ${getToken()}`
                    }
                });

                this.data = r.data;
                this.tableData = this.constructTableData(this.modelSchema, r.data);
            } catch (e) {
                console.log(extractError(e));
            }

            this.loading = false;
        },

        onPageChange(page) {
            this.page = page;
            this.load();
        },

        onSort(field, order) {
            this.sortField = field;
            this.sortOrder = order;

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
                    schemaField: field,
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
                if (selector.startsWith("*.") && `*.${fieldName}` === selector) {
                    return this.customColumnComponents[selector];
                }

                if (selector === `${this.modelId}.${fieldName}`) {
                    return this.customColumnComponents[selector];
                }
            }

            return this.customColumnComponents["*"];
        }
    },

    watch: {
        modelId() {
            this.page = 0;
            this.data = null;
            this.tableData = null;
            this.tableColumns = null;
            this.sortField = null;
            this.sortOrder = "desc";
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