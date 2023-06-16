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
}
.login h2 {
  text-align: left;
  margin-bottom: 2.5rem;
}
.login hr {
  width: 1.75rem;
  border: none;
  border-top: 3px solid #000;
  margin-bottom: 2rem;
}
.login p.des {
  width: 100%;
  margin: 0 0 1rem;
}
.login ul {
  width: 100%;
  margin: 0 0 2.5rem;
  padding-left: 1.25rem;
  box-sizing: border-box;
}
.login button {
  width: 100%;
}
.brand {
  display: flex;
  margin-top: 4rem;
}
.brand img {
  width: 2.5rem;
}
.brand h2 {
  position: relative;
  top: 0.075rem;
  font-family: "Domaine Sans";
  text-transform: uppercase;
  font-weight: 400;
  font-size: 1.5rem;
  margin: 0;
}
@media screen and (min-width: 768px) {
  .login {
    width: fit-content;
  }
  .login button {
    width: 20rem;
  }
}
</style>

<template>
  <main>
    <div class="login">

      <h2>Physical Print Redemption</h2>
      <p class="des">Connect your crypto wallet to verify ownership:</p>
      <ul>
        <li>You will be asked to share your wallet address to verify ownership.</li>
        <li>You will be asked for shipping information so that we can fulfill your request.</li>
      </ul>
      <button class="primary" @click="web3Connect">Get Started</button>
      <div class="brand">
        <img src="img/au.svg" alt="">
        <h2>Autonomy</h2>
      </div>
    </div>
  </main>
</template>


<script lang="ts">
import { Options, Vue } from "vue-class-component";

@Options({
  async created() {
    if (this.$web3Modal.cachedProvider) {
      this.$router.push({ name: "Home" });
    }
  },

  methods: {
    async web3Connect() {
      const provider = await this.$web3Modal.connect();
      this.$router.push({ name: "Home" });
    },
  },
})
export default class Welcome extends Vue {}
</script>
