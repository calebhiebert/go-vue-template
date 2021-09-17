<template>
    <div class="image-display has-text-centered" @click="openUppy">
      <b-image :src="value" v-if="value" alt="Upload Image" ></b-image>
      <b-button style="margin: 0 auto" @click="openUppy">{{ value ? 'Change Image' : 'Upload Image' }}</b-button>
    </div>
</template>

<script>
import '@uppy/core/dist/style.css';
import '@uppy/dashboard/dist/style.css';
import '@uppy/image-editor/dist/style.css';

import Uppy from '@uppy/core';
import XHRUpload from '@uppy/xhr-upload';
import ImageEditor from '@uppy/image-editor';
import Dashboard from "@uppy/dashboard";
import {API_BASE_URL, getToken} from "../api";

export default {
  name: "ImageUpload",

  props: {
    value: {
      type: String,
    }
  },

  data() {
    return {
      uppy: null,
    }
  },

  methods: {
    openUppy() {
      this.uppy.getPlugin('Dashboard').openModal();
    }
  },

  mounted() {
    this.uppy = new Uppy({
      allowMultipleUploadBatches: false,
      restrictions: {
        maxFileSize: 5242880,
        maxNumberOfFiles: 1,
        allowedFileTypes: ['image/jpeg', 'image/png', 'image/webm'],
      }
    })
        .use(Dashboard, {
          proudlyDisplayPoweredByUppy: false,
          autoOpenFileEditor: true,
          closeAfterFinish: true,
          theme: 'auto',
        })
        .use(XHRUpload, {
          endpoint: `${API_BASE_URL}/image`,
          fieldName: 'image',
          headers: {
            'Authorization': `Bearer ${getToken()}`,
          }
        })
        .use(ImageEditor, {
          target: Dashboard,
        });

    this.uppy.on('complete', (result) => {
      if (result.successful.length > 0) {
        const url = `${API_BASE_URL}/image/${result.successful[0].response.body.id}`;

        this.$emit('input', url);
      }
    })
  },

  beforeDestroy() {
    this.uppy.close();
  }
}
</script>

<style scoped>
.image-display {
  width: 256px;
  cursor: pointer;
}
</style>