<template>
  <Nothing>
    <p class="des">Thank you for submitting your information. We will be in touch soon.</p>
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
export default class ThankYou extends Vue {}
</script>

