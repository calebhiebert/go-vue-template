<template>
  <div>
    <h2 class="title is-2">Create {{ modelSchema.name_singular }}</h2>
    <form ref="frm" @submit.prevent="create">
      <ModelForm :schema="modelSchema" :disabled="creating" v-model="formModel"></ModelForm>

      <div class="mt-4">
        <b-button :disabled="creating"
                  @click="$router.push({name: 'AdminModelView', params: {model: $route.params.model}})">Cancel
        </b-button>
        <b-button :loading="creating" class="ml-2" type="is-primary" @click="create">Create</b-button>
      </div>
    </form>
  </div>
</template>

<script>
import ModelForm from "../../components/admin/ModelForm";
import axios from "axios";
import {API_BASE_URL, extractError, getToken} from "../../api";

export default {
  name: "ModelCreate",
  components: {ModelForm},

  data() {
    return {
      modelSchema: window._adminSchema.models[this.$route.params.model],
      loading: false,
      creating: false,
      formModel: {},
    };
  },

  methods: {
    async create() {
      const valid = this.$refs.frm.checkValidity();

      if (!valid) {
        this.$buefy.dialog.alert({
          title: "Invalid input",
          message: "There are some issues with the form, please correct them and try again",
        });
        return
      }

      this.creating = true;

      try {
        await axios.post(`${API_BASE_URL}/crud/${this.modelSchema.url_name}`, this.formModel, {
          headers: {
            "Authorization": `Bearer ${getToken()}`,
          },
        });

        await this.$router.push({name: 'AdminModelView', params: {model: this.$route.params.model}});
      } catch (e) {
        const err = extractError(e);

        this.$buefy.dialog.alert({
          title: "Something went wrong",
          message: err.message,
        });
      }

      this.creating = false;
    }
  }
}
</script>

<style scoped>

</style>