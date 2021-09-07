<template>
    <div>

        <!-- Field -->
        <div v-for="field of schema.fields" :key="field.id">
            <fieldset :disabled="disabled" v-if="field.editable && field.config.editable !== false">
                <b-field :label="field.name">
                    <b-input v-if="field.type === 'string'" :type="field.config.is_email ? 'email' : undefined" v-model="value[field.id]" :required="!field.nullable"
                             :placeholder="field.name">

                    </b-input>
                    <b-taginput
                        v-else-if="field.type === 'array'"
                        v-model="value[field.id]"
                        :placeholder="field.name"
                        :required="!field.nullable"
                        ellipsis>
                    </b-taginput>
                </b-field>

            </fieldset>
        </div>

    </div>
</template>

<script>
export default {
    name: "ModelForm",

    props: {
        schema: {
            type: Object,
            required: true,
        },

        value: {
            type: Object,
            required: false,
        },

        inputData: {
            type: Object,
            required: false,
        },

        disabled: {
            type: Boolean,
            required: false,
            default: false,
        },
    },

    created() {
        this.initData();
    },

    methods: {
        initData() {
            if (this.inputData !== undefined && this.inputData !== null) {
                for (let key of Object.keys(this.inputData)) {
                    const field = this.schema.fields.find(f => f.id === key);

                    if (field && field.editable && field.config.editable) {
                        this.$set(this.value, key, this.inputData[key]);
                    }
                }
            } else {
                for (let field of this.schema.fields) {
                    if (field.nullable && field.editable) {
                        this.$set(this.value, field.id, null);
                    } else if (field.editable) {
                        let val;

                        switch (field.type) {
                            case "string":
                                val = "";
                                break;
                            default:
                                val = null;
                                break;
                        }

                        this.$set(this.value, field.id, val);
                    }
                }
            }
        }
    },

    watch: {
        inputData: {
            deep: true,
            handler() {
                this.initData();
            }
        }
    }
};
</script>

<style scoped>

</style>