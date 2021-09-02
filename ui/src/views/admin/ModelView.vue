<template>
    <div>{{ $route.params }}</div>
</template>

<script>
export default {
    name: "ModelView",

    data() {
        return {
            loading: false,
            data: null,
            limit: 25,
            page: 0,
        }
    },

    methods: {
        async load() {
            this.loading = true;

            console.log(window._adminSchema.models[this.modelId]);

            this.loading = false;
        }
    },

    watch: {
        modelId() {
            this.load();
        },

        page() {
            this.load();
        }
    },

    computed: {
        modelId() {
            return this.$route.params.model;
        }
    }
}
</script>

<style scoped>

</style>