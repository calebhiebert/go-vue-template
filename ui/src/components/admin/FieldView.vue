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
  <p v-else-if="t === 'uuid'" @click="copyToClip(model[field.id])" class="clippy">
    <b-tooltip :label="expanded ? 'Click to Copy' : model[field.id] + ' (Click to copy)'">
      <b-tag type="is-primary" v-if="!expanded">
        {{ model[field.id] | shortenUUID }}
      </b-tag>
      <b-tag v-else>
        {{ model[field.id] }}
      </b-tag>
    </b-tooltip>
  </p>
  <p v-else>{{ model[field.id] }}</p>
</template>

<script>
import JsonViewer from "vue-json-viewer";
import {copyTextToClipboard} from "../../util";

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

    expanded: {
      type: Boolean,
      required: false,
      default: false,
    }
  },

  methods: {
    async copyToClip(id) {
      try {
        await copyTextToClipboard(id);

        this.$buefy.toast.open("Copied to clipboard!");
      } catch (e) {
        this.$buefy.toast.open("Failed to copy to clipboard");
      }
    }
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
.clippy {
  cursor: pointer;
}
</style>