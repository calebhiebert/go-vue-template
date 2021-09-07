<template>
    <p v-if="t === 'time'">
        <span v-if="model[field.id]">{{ model[field.id] | localTime }}</span>
        <span class="has-text-grey is-italic" v-else>NULL</span>
    </p>
    <p v-else-if="t === 'json'">
        <json-viewer v-if="model[field.id] !== null" :value="model[field.id]"></json-viewer>
    </p>
    <b-taglist v-else-if="t === 'array'">
        <b-tag v-for="t of model[field.id]" :key="t">{{ t }}</b-tag>
    </b-taglist>
    <p class="has-text-grey is-italic" v-else-if="model[field.id] === null || model[field.id] === undefined">
        NULL
    </p>
    <p v-else-if="t === 'uuid'">
        <b-tooltip :label="model[field.id]">
            <b-tag type="is-primary">
                {{ model[field.id] | shortenUUID }}
            </b-tag>
        </b-tooltip>
    </p>
    <p v-else>{{ model[field.id] }}</p>
</template>

<script>
import JsonViewer from "vue-json-viewer";

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
        },

        overrideType: {
            type: String,
            required: false,
            default() {
                return null;
            },
        },
    },


    filters: {
        localTime(value) {
            return new Date(value).toLocaleString();
        },

        shortenUUID(value) {
            let split = value.split('-');

            return split[split.length - 1];
        },
    },

    computed: {
        t() {
            if (this.overrideType) {
                return this.overrideType;
            } else {
                return this.field.type;
            }
        },
    },
};
</script>

<style scoped>

</style>