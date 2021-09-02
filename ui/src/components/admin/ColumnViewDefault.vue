<template>
    <div>{{ row[col.field] }}</div>
</template>

<script>
export default {
    name: "ColumnViewDefault",

    props: {
        row: {
            type: Object,
            required: true,
        },

        col: {
            type: String,
            required: true,
        },

        p: {
            type: Object,
            required: true,
        }
    }
}
</script>

<style scoped>

</style>