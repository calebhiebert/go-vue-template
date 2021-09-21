<template>
  <div>

    <!-- Field -->
    <div v-for="field of visibleFields" :key="field.id">
      <fieldset :disabled="disabled">
        <b-field :type="errorFor(field) ? 'is-danger' : null"
                 :message="errorFor(field) ? errorFor(field).message : null"
                 :label="field.name + (field.required ? ' (Required)' : '')">
          <component v-if="getCustomFieldComponent(field.id) !== null"
                     v-model="value[field.id]" :required="field.required"
                     :placeholder="field.name" :is="getCustomFieldComponent(field.id)">
          </component>
          <b-input v-else-if="field.type === 'string'" :type="field.config.is_email ? 'email' : undefined"
                   v-model="value[field.id]" :required="field.required"
                   :placeholder="field.name"></b-input>
          <b-checkbox v-else-if="field.type === 'bool'" v-model="value[field.id]"></b-checkbox>
          <b-datetimepicker
              v-model="value[field.id]"
              :placeholder="field.name"
              :datetime-formatter="formatTime"
              :required="field.required"
              v-else-if="field.type === 'time'">
          </b-datetimepicker>
          <b-taginput
              v-else-if="field.type === 'array'"
              v-model="value[field.id]"
              :placeholder="field.name"
              ellipsis>
          </b-taginput>
          <v-jsoneditor v-else-if="field.type === 'json'" v-model="value[field.id]" :options="{mode: 'code'}">

          </v-jsoneditor>

        </b-field>

      </fieldset>
    </div>

    <slot v-if="visibleFields.length === 0" name="empty">
      <div class="has-text-centered">
        <h5 class="is-5 has-text-grey is-italic">There are no editable fields</h5>
      </div>
    </slot>
  </div>
</template>

<script>
import ImageUpload from "../ImageUpload";
import VJsoneditor from 'v-jsoneditor'

export default {
  name: "ModelForm",

  components: {
    VJsoneditor
  },

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

    validationErrors: {
      type: Array,
      required: false,
      default() {
        return [];
      }
    }
  },

  created() {
    this.initData();

    this.registerCustomField("users.image", ImageUpload);
  },

  data() {
    return {
      customFieldComponents: {},
    }
  },

  methods: {
    initData() {
      if (this.inputData !== undefined && this.inputData !== null) {
        for (let key of Object.keys(this.inputData)) {
          const field = this.schema.fields.find(f => f.id === key);

          if (field && field.editable && field.config.editable) {
            if (field.type === 'time') {
              this.$set(this.value, key, new Date(this.inputData[key]));
            } else {
              this.$set(this.value, key, this.inputData[key]);
            }
          }
        }
      } else {
        for (let field of this.schema.fields) {
          if (!field.required && field.editable) {
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
    },

    errorFor(field) {
      return this.validationErrors.find(ve => ve.field === field.name || ve.field === field.id);
    },

    registerCustomField(selector, component) {
      this.customFieldComponents[selector] = component;
    },

    getCustomFieldComponent(fieldName) {
      for (const selector in this.customFieldComponents) {
        if (selector.startsWith("*.") && `*.${fieldName}` === selector) {
          return this.customFieldComponents[selector];
        }

        if (selector === `${this.schema.data_name}.${fieldName}`) {
          return this.customFieldComponents[selector];
        }
      }

      return null;
    },

    formatTime(time) {
      return time.toISOString();
    },
  },

  computed: {
    visibleFields() {
      return this.schema.fields.filter(f => (f.editable && f.config.editable !== false));
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