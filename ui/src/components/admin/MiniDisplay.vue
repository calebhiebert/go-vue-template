<template>
  <div class="media">
    <div class="media-left" v-if="hasImage">
      <b-image :src="imageSrc"></b-image>
    </div>
    <div class="media-content">
      <div class="title is-6">{{ displayData }}</div>
    </div>
  </div>
</template>

<script>
import {API_BASE_URL} from "../../api";

export default {
  name: "MiniDisplay",
  props: {
    model: {
      type: Object,
      required: true,
    }
  },

  computed: {

    // Guess what should be displayed
    displayData() {
      if (this.model['name']) {
        return this.model['name'];
      } else if (this.model['id']) {
        return this.model['id'];
      } else {
        return JSON.stringify(this.model);
      }
    },

    hasImage() {
      return this.model['image'] !== null && this.model['image'] !== undefined;
    },

    imageSrc() {
      if (!this.model['image']) {
        return null;
      } else {
        try {
          // If the image id is a uuid, this operation will fail
          new URL(this.model['image']);
          return this.model['image'];
        } catch (e) {
          // Assume uuid image
          return `${API_BASE_URL}/api/image/${this.model['image']}`;
        }
      }
    }
  }
}
</script>

<style scoped>

</style>