<template>
    <div>
        <h4 class="title is-4">{{ schema.name }}</h4>

        <div class="mt-4" v-for="field of schema.fields" :key="field.id">
            <h6 class="title is-6 mb-0">{{ field.name }}</h6>

            <FieldView :field="field" :model="model"></FieldView>
        </div>
    </div>
</template>

<script>
import FieldView from "./FieldView";

export default {
    name: "ModelViewSingle",

    components: {FieldView},

    props: {
        model: {
            type: Object,
            required: true,
        },

        schema: {
            type: Object,
            required: true,
        },
    }
};
</script>

<style scoped>

</style>