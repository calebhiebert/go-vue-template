<template>
    <div>
        <b-field>
            <b-select v-model="filter" size="is-small" placeholder="🔎">
                <option value="cont" v-if="canCol('cont')">🔎</option>
                <option selected value="eq" v-if="canCol('eq')">==</option>
                <option value="gt" v-if="canCol('gt')">&gt;</option>
                <option value="gte" v-if="canCol('gte')">&gt;=</option>
                <option value="lt" v-if="canCol('lt')">&lt;</option>
                <option value="lte" v-if="canCol('lte')">&lt;=</option>
                <option value="in" v-if="canCol('in')">in</option>
            </b-select>
            <b-datetimepicker v-if="columnSchema.type === 'time'"
                              :datetime-formatter="formatTime"
                              position="is-bottom-left" size="is-small"
                              v-model="query"></b-datetimepicker>
            <b-input v-else v-model="query" size="is-small"></b-input>
        </b-field>
    </div>
</template>

<script>
export default {
    name: "ColumnFilter",

    props: {
        value: {
            type: Object,
        },

        columnSchema: {
            type: Object,
            required: true,
        },
    },

    data() {
        return {
            filter: "eq",
            query: this.columnSchema.type === 'time' ? null : "",
        };
    },

    methods: {
        canCol(col) {
            return this.columnSchema.filter_operations.includes(col);
        },

        formatTime(time) {
            return time.toISOString();
        },
    },

    watch: {
        filter() {
            this.$emit("input", {
                filter: this.filter,
                query: this.columnSchema.type === 'time' ? this.query.toISOString() : this.query,
            });
        },

        query() {
            this.$emit("input", {
                filter: this.filter,
                query: this.columnSchema.type === 'time' ? this.query.toISOString() : this.query,
            });
        },
    },
};
</script>

<style scoped>

</style>