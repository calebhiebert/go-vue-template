<template>
  <div>
    <div class="columns mb-2">
      <!--<div class="column"></div>-->
      <div class="column is-narrow">
        <b-button icon-left="sync-alt" :loading="loading" @click="load">Refresh</b-button>
        <b-button class="ml-2" type="is-primary" icon-left="plus" @click="$router.push({name: 'AdminModelCreate', params: {model: modelId}})">Create New</b-button>
        <b-button type="is-danger" icon-left="trash" class="ml-2" @click="deleteMultiple"
                  :disabled="selected.length === 0">Delete
          {{ selected.length === 0 ? "" : selected.length }}
        </b-button>
      </div>
    </div>

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
             :current-page="page"
             aria-next-label="Next page"
             aria-previous-label="Previous page"
             aria-page-label="Page"
             aria-current-label="Current page"

             :debounce-search="1000"
             backend-filtering
             @filters-change="onSearch"

             :checked-rows.sync="selected"

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
          searchable>
        <template v-slot="props">
          {{ props.row.original_title }}
          <component :is="getCustomColumnComponent(col.field)" :col="col" :field="col.schemaField" :p="props"
                     :row="props.row"></component>
        </template>

        <template #searchable="props">
          <ColumnFilter :column-schema="col.schemaField"
                        v-model="props.filters[props.column.field]"></ColumnFilter>
        </template>

      </b-table-column>

      <template #detail="props">
        <ModelvViewSingle :model="props.row" :schema="modelSchema"></ModelvViewSingle>
        <b-button class="mt-2" type="is-primary" @click="editItem(props.row)">Edit</b-button>
        <b-button class="mt-2 ml-2" type="is-danger" @click="deleteItem(props.row)">Delete</b-button>
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
import ColumnViewUUID from "../../components/admin/ColumnViewUUID";
import ColumnFilter from "../../components/admin/ColumnFilter";

export default {
  name: "ModelView",
  components: {ColumnFilter, ModelvViewSingle},
  data() {
    return {
      loading: false,
      data: null,
      limit: 10,
      page: this.$route.query.page ? parseInt(this.$route.query.page) : 1,
      modelSchema: null,
      tableData: null,
      tableColumns: null,
      customColumnComponents: {},
      selected: [],

      sortField: null,
      sortOrder: "desc",
      defaultSortOrder: "desc",
      filterQuery: "",
    };
  },

  created() {
    this.registerCustomColumnComponent("*", ColumnViewDefault);
    this.registerCustomColumnComponent("access_logs.response_code", ColumnViewHTTPResponseCode);
    this.registerCustomColumnComponent("*.id", ColumnViewUUID);
    this.registerCustomColumnComponent("*.user_id", ColumnViewUUID);
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
        let sort = "";

        if (this.sortField) {
          sort = `&sort.${this.sortField}=${this.sortOrder}`;
        }

        const r = await axios.get(`${API_BASE_URL}/crud/${this.modelSchema.url_name}?limit=${this.limit}&offset=${this.limit * (this.page - 1)}${sort}${this.filterQuery}`, {
          headers: {
            "Authorization": `Bearer ${getToken()}`,
          },
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

      this.$router.replace({
        name: this.$route.name, params: this.$route.params, query: {
          page: this.page,
        },
      });

      this.load();
    },

    onSearch(filters) {
      const filterClauses = [];

      for (let k of Object.keys(filters)) {
        if (Object.hasOwn(filters, k) && filters[k].query !== "" && filters[k].query) {
          filterClauses.push(`${k}.${filters[k].filter}=${encodeURIComponent(filters[k].query)}`);
        }
      }

      const f = filterClauses.join("&");

      if (f !== "") {
        this.filterQuery = "&" + f;
      } else {
        this.filterQuery = "";
      }

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
    },

    editItem(item) {
      this.$router.push({name: "AdminModelEdit", params: {model: this.modelId, id: item.id}});
    },

    deleteMultiple() {
      this.$buefy.dialog.confirm({
        title: "Are you sure?",
        message: `Are you sure you would like to delete ${this.selected.length} item(s)? You cannot undo this action`,
        confirmText: `Delete ${this.selected.length}`,
        cancelText: "Cancel",
        type: "is-danger",
        onConfirm: async () => {
          await axios.delete(`${API_BASE_URL}/crud/${this.modelSchema.url_name}`, {
            data: {
              ids: this.selected.map(s => s.id),
            },
            headers: {
              "Authorization": `Bearer ${getToken()}`,
            },
          });

          this.selected = [];

          this.load();
        },
      });
    },

    deleteItem(item) {
      this.$buefy.dialog.confirm({
        title: "Are you sure?",
        message: `Are you sure you would like to delete this item?${this.modelSchema.can_soft_delete ? "" : " This action cannot be undone"}`,
        confirmText: "Delete",
        cancelText: "Cancel",
        type: "is-danger",
        onConfirm: async () => {
          await axios.delete(`${API_BASE_URL}/crud/${this.modelSchema.url_name}/${item.id}`, {
            headers: {
              "Authorization": `Bearer ${getToken()}`,
            },
          });

          if (this.modelSchema.can_soft_delete) {
            this.$buefy.snackbar.open({
              duration: 5000,
              message: `${this.modelSchema.name} deleted`,
              type: "is-danger",
              position: "is-bottom-right",
              actionText: "Undo",
              onAction: async () => {
                await axios.post(`${API_BASE_URL}/crud/${this.modelSchema.url_name}/${item.id}/unDelete`, null, {
                  headers: {
                    "Authorization": `Bearer ${getToken()}`,
                  },
                });

                this.load();
              },
            });
          }

          this.load();
        },
      });
    },
  },

  watch: {
    modelId() {
      this.page = 0;
      this.data = null;
      this.tableData = null;
      this.tableColumns = null;
      this.sortField = null;
      this.sortOrder = "desc";
      this.filterQuery = null;
      this.load();
    },
  },

  computed: {
    modelId() {
      return this.$route.params.model;
    },
  },
};
</script>

<style scoped>

</style>