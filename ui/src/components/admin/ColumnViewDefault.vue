<template>
    <div>
        <FieldView :field="field" :model="row"></FieldView>
    </div>
</template>

<script>
import FieldView from "./FieldView";
export default {
    name: "ColumnViewDefault",
    components: {FieldView},
    props: {
        row: {
            type: Object,
            required: true,
        },

        field: {
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
        },
    },
};
</script>

<style scoped>

</style>