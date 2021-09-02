<template>
    <div>{{ value }}</div>
</template>

<script>
export default {
    name: "ColumnViewDateTime",

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
        value() {
            return new Date(this.row[this.col.field]).toLocaleString();
        }
    }
}
</script>

<style scoped>

</style>