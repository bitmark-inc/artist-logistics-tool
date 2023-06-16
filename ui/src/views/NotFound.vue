<template>
  <Nothing>
    <h2>404 Not Found</h2>
    <p class="des">Sorry! We couldn't find what you were requested. Please contact <a href="mailto:support@feralfile.com">support@feralfile.com</a> with your Feral File alias or email.</p>
  </Nothing>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import Nothing from "./Nothing.vue";

@Options({
  components: {
    Nothing,
  },
})
export default class NotFound extends Vue {}
</script>

