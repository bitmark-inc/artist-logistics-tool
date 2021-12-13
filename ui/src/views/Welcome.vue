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
  font-family: "Domaine Sans";
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
      <p class="des">Get started by clicking the Connect button below - you will be asked for personal information that will then be shared directly with Refik Anadol so they can fulfill your request. You can reach them on <a href="https://discord.gg/vXvpeETTD6" target="blank">Discord</a> through your designated exhibition channel if you have any followup questions.</p>
      <div class="btn-set">
        <button class="primary" @click="web3Connect">Connect</button>
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
