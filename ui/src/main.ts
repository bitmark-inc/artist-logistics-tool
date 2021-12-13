import { createApp } from 'vue'
import App from './App.vue'
import router from './router'


import Web3Modal from "web3modal";
import WalletConnectProvider from "@walletconnect/web3-provider";

const providerOptions = {
  walletconnect: {
    package: WalletConnectProvider, // required
    options: {
      infuraId: "INFURA_ID", // required
    },
  },
};

const web3Modal = new Web3Modal({
  network: "mainnet", // optional
  cacheProvider: true, // optional
  providerOptions, // required
});

const app = createApp(App)
app.use(router).mount('#app')
app.config.globalProperties.$web3Modal = web3Modal
