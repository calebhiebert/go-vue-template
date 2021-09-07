<template>
    <div>
        <b-field>
            <b-select v-model="filter" size="is-small" placeholder="🔎">
                <option value="cont">🔎</option>
                <option selected value="eq">==</option>
                <option value="gt">&gt;</option>
                <option value="gte">&gt;=</option>
                <option value="lt">&lt;</option>
                <option value="lte">&lt;=</option>
                <option value="in">in</option>
            </b-select>
            <b-input v-model="query" size="is-small"></b-input>
        </b-field>
    </div>
</template>

<script>
export default {
    name: "ColumnFilter",

    props: {
        value: {
            type: Object,
        }
    },

    data() {
        return {
            filter: "cont",
            query: ""
        }
    },

    watch: {
        filter() {
            this.$emit("input", {
                filter: this.filter,
                query: this.query
            });
        },

        query() {
            this.$emit("input", {
                filter: this.filter,
                query: this.query
            });
        }
    }
};
</script>

<style scoped>

</style>