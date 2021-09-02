<template>
    <div>
        <b-tag :type="tagType"><span class="has-text-weight-bold">{{ row[col.field] }}</span></b-tag>
    </div>
</template>

<script>
export default {
    name: "ColumnViewHTTPResponseCode",

    props: {
        row: {
            type: Object,
            required: true,
        },

        col: {
            type: Object,
            required: true,
        },

        p: {
            type: Object,
            required: true,
        }
    },

    computed: {
        tagType() {
            const data = this.row[this.col.field];

            if (data >= 200 && data < 300) {
                return "is-success";
            } else if (data >= 400 && data < 500) {
                return "is-warning";
            } else if (data >= 500 && data < 600) {
                return "is-danger";
            }

            return "is-primary";
        }
    }
}
</script>

<style scoped>

</style>