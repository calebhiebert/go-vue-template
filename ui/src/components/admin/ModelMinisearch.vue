<template>
  <section>
    <b-autocomplete
        :data="data"
        field="id"
        :loading="loading"
        :check-infinite-scroll="true"
        :placeholder="'Find '+ foreignTableSchema.name_singular"
        @typing="fetchData"
        @select="option => selected = option"
        @infinite-scroll="fetchMoreData">

      <template slot-scope="props">
        <MiniDisplay :model="props.option"></MiniDisplay>
      </template>
      <template #footer>
        <span v-show="page > totalPages" class="has-text-grey">Sorry, there's nothing else here</span>
      </template>
    </b-autocomplete>
  </section>
</template>

<script>
import axios from "axios";
import {API_BASE_URL, getToken} from "../../api";
import MiniDisplay from "./MiniDisplay";

export default {
  name: "ModelMinisearch",
  components: {MiniDisplay},
  props: {
    schema: {
      type: Object,
      required: true,
    }
  },

  data() {
    return {
      foreignTableSchema: window._adminSchema.models[this.schema.foreign_fields[0].model],
      data: null,
      selected: null,
      loading: false,
      page: 1,
      perPage: 10,
      totalPages: 1,
      q: ""
    }
  },

  methods: {
    async fetchData(q) {
      this.loading = true;

      // String update
      if (this.q !== q) {
        this.q = q
        this.data = []
        this.page = 1
        this.totalPages = 1
      }

      // String cleared
      if (!q.length) {
        this.data = []
        this.page = 1
        this.totalPages = 1
        return
      }

      // Reached last page
      if (this.page > this.totalPages) {
        return
      }

      try {
        const res = await axios.get(`${API_BASE_URL}/crud/${this.foreignTableSchema.url_name}/minisearch?q=${q}&limit=${this.perPage}&offset=${(this.page -1) * this.perPage}`, {
          headers: {
            "Authorization": `Bearer ${getToken()}`,
          },
        });

        if (res.data.items) {
          res.data.items.forEach(i => this.data.push(i));
        }

        this.page++;

        this.totalPages = Math.ceil(res.data.total / this.perPage);
      } catch (e) {
        console.log(e);
      }

      this.loading = false;
    },

    async fetchMoreData() {
      return this.fetchData(this.q)
    }
  }
}
</script>

<style scoped>

</style>