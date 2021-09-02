<template>
    <div>
        <b-tag>{{ row[col.field] }}</b-tag>
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
            } else {
                return "is-light"
            }
        }
    }
}
</script>

<style scoped>

</style>