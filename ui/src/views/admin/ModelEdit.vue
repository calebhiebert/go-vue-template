<template>
  <div>
    <h2 class="title is-2">Edit {{ modelSchema.name_singular }}</h2>

    <ModelForm :schema="modelSchema" :input-data="model" :disabled="loading || saving" v-model="formModel"></ModelForm>

    <div class="mt-4">
      <b-button :disabled="saving || loading"
                @click="$router.push({name: 'AdminModelView', params: {model: $route.params.model}})">Cancel
      </b-button>
      <b-button :loading="saving || loading" class="ml-2" type="is-primary" @click="save">Save</b-button>
    </div>
  </div>
</template>

<script>
import ModelForm from "../../components/admin/ModelForm";
import axios from "axios";
import {API_BASE_URL, extractError, getToken} from "../../api";

export default {
  name: "ModelEdit",
  components: {ModelForm},

  data() {
    return {
      modelSchema: window._adminSchema.models[this.$route.params.model],
      model: null,
      loading: false,
      saving: false,
      formModel: {},
    };
  },

  created() {
    this.load();
  },

  methods: {
    async load() {
      this.loading = true;

      try {
        const model = await axios.get(`${API_BASE_URL}/crud/${this.modelSchema.url_name}/${this.modelId}`, {
          headers: {
            "Authorization": `Bearer ${getToken()}`,
          },
        });

        this.model = model.data;

        console.log(this.model);
      } catch (e) {
        const err = extractError(e);

        this.$buefy.dialog.alert({
          title: "Something went wrong",
          message: err.message,
        });
      }

      this.loading = false;
    },

    async save() {
      this.saving = true;

      try {
        const updatedModel = await axios.put(`${API_BASE_URL}/crud/${this.modelSchema.url_name}/${this.modelId}`, this.formModel, {
          headers: {
            "Authorization": `Bearer ${getToken()}`,
          },
        });

        this.model = updatedModel.data;
      } catch (e) {
        const err = extractError(e);

        this.$buefy.dialog.alert({
          title: "Something went wrong",
          message: err.message,
        });
      }

      this.saving = false;
    },
  },

  computed: {
    modelTypeId() {
      return this.$route.params.model;
    },

    modelId() {
      return this.$route.params.id;
    },
  },
};
</script>

<style scoped>

</style>