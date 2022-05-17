<template>
  <div>
    <div>
      <input type="file" @change="uploadFile" ref="file" />
      <button @click="submitFile">Upload!</button>
    </div>
    <div v-if="files.length > 0">
      <div v-for="(file, i) in files" :key="i">
        <p>{{ file }}</p>
        <button @click="downloadFile(file)">Download!</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      files: [],
      fileName: "",
    };
  },

  methods: {
    uploadFile() {
      this.Files = this.$refs.file.files[0];
    },

    submitFile() {
      const formData = new FormData();
      formData.append("file", this.Files, this.Files.name);
      const headers = { "Content-Type": "multipart/form-data" };
      axios.post("http://fileloader-env-1.eba-bfab5ipr.us-east-1.elasticbeanstalk.com:5000/upload", formData, headers).then(res => console.log(res));
      this.files.push(this.Files.name);
    },

    downloadFile(filename) {
      const requestBody = { file_name: filename };

      axios.post("http://fileloader-env-1.eba-bfab5ipr.us-east-1.elasticbeanstalk.com:5000/download", requestBody).then((res) => {
        const byteCharacters = atob(res.data.file);
        const byteNumbers = new Array(byteCharacters.length);
        for (let i = 0; i < byteCharacters.length; i++) {
          byteNumbers[i] = byteCharacters.charCodeAt(i);
        }
        const byteArray = new Uint8Array(byteNumbers);

        const blob = new Blob([byteArray], { type: "octet-stream" });
        const href = URL.createObjectURL(blob);
        const a = Object.assign(document.createElement("a"), {
          href,
          style: "display:none",
          download: filename,
        });
        document.body.appendChild(a);

        a.click();
      });
    },
  },
};
</script>
