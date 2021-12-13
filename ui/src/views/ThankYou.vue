<style scoped>
  .wrapper {
    width: 100%;
    height: 100%;
  }
  main {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .login {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .login img {
    width: 4rem;
  }
  .login h1 {
    font-family: 'Domaine Sans';
    text-transform: uppercase;
    font-weight: 400;
  }
  .login hr {
    width: 1.75rem;
    border: none;
    border-top: 3px solid #000;
    margin-bottom: 2rem;
  }
  .login p.des {
    margin: 0 0 2rem;
    text-align: justify;
  }
  @media screen and (min-width: 768px) {
    .login {
      width: 50%;
    }
  }
</style>

<template>
  <main>
    <div class="login">
      <img src="img/au.svg" alt="">
      <h1>Autonomy</h1>
      <hr>
      <p class="des">Thank you for submitting your information. We will be in touch soon.</p>
    </div>
  </main>
</template>


<script lang="ts">
import { Options, Vue } from "vue-class-component";

@Options({
  async created() {
    if (this.$web3Modal.cachedProvider == "wallet connect") {
      this.web3.eth.currentProvider.disconnect();
    }
    this.$web3Modal.clearCachedProvider();
  },
})
export default class ThankYou extends Vue {}
</script>

