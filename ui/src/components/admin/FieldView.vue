<template>
    <p v-if="field.type === 'time'">
        <span v-if="model[field.id]">{{ model[field.id] | localTime }}</span>
        <span class="has-text-grey is-italic" v-else>NULL</span>
    </p>
    <p v-else-if="field.type === 'json'">
        <json-viewer :value="model[field.id]"></json-viewer>
    </p>
    <b-taglist v-else-if="field.type === 'array'">
        <b-tag v-for="t of model[field.id]" :key="t">{{ t }}</b-tag>
    </b-taglist>
    <p class="has-text-grey is-italic" v-else-if="model[field.id] === null || model[field.id] === undefined">
        NULL
    </p>
    <p v-else>{{ model[field.id] }}</p>
</template>

<script>
import JsonViewer from 'vue-json-viewer';

export default {
    name: "FieldView",

    components: {JsonViewer},

    props: {
        field: {
            type: Object,
            required: true,
        },

        model: {
            type: Object,
            required: true,
        }
    },


    filters: {
        localTime(value) {
            return new Date(value).toLocaleString();
        }
    }
};
</script>

<style scoped>

</style>