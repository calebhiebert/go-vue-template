<template>
    <div>{{ row }}</div>
</template>

<script>
export default {
    name: "ColumnViewDefault",

    props: {
        row: {
            type: Object,
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