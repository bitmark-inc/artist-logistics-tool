<template>
  <div class="wrapper">
    <header>
      <a class="brand" href="home-signed-in.html">
        <img src="img/au.svg" />
        <h2>Autonomy<!--<span> Market<span>--></h2>
      </a>
      <p v-if="web3 == null">
        <button @click="this.web3Login" :disabled="web3">Connect</button>
      </p>
      <p v-else>
        <span>{{ accountNumber }}</span>
        <button @click="this.web3Logout" :disabled="!web3">Disconnect</button>
      </p>
    </header>

    <main>
      <h2 class="text-center">Collect Refik Anadol's Signed Prints</h2>
      <ol>
        <li>
          Select {{ maxSelectableNumber }} desired items in the following
          artwork grid.
        </li>
        <li>Fill in the shipping information.</li>
        <li>Confirm and click "Send".</li>
      </ol>
      <div class="two-col">
        <div class="left">
          <h3>I. Select Artworks</h3>
          <strong style="display: block; margin-bottom: 0.5rem">
            * Please select {{ maxSelectableNumber }} items from your collected
            digital editions for signed prints. Remaining:
            {{ maxSelectableNumber - totalSelected }}
          </strong>
          <div class="grids">
            <div
              class="card"
              :class="{ selected: token.selected }"
              v-for="token in tokens"
              :key="token.id"
              @click="selectToken(token)"
            >
              <div
                class="image"
                :style="{ 'background-image': 'url(' + token.imageURL + ')' }"
              ></div>
              <div class="info">
                <p>{{ token.name }}</p>
              </div>
            </div>
          </div>
        </div>
        <div class="right">
          <h3>II. Fill Out the Following Information</h3>
          <form>
            <div class="row">
              <label>Name *</label>
              <div
                style="
                  display: grid;
                  grid-template-columns: 1fr 1fr;
                  column-gap: 0.5rem;
                "
              >
                <input
                  type="text"
                  id="first"
                  name="first"
                  placeholder="First Name"
                  v-model="form.firstName"
                />
                <input
                  type="text"
                  id="last"
                  name="last"
                  placeholder="Last Name"
                  v-model="form.lastName"
                />
              </div>
            </div>
            <div class="row">
              <label for="email">Email *</label>
              <input
                type="email"
                id="email"
                name="email"
                placeholder="eg. john@bitmark.com"
                v-model="form.email"
              />
            </div>
            <div class="row">
              <label>Shipping Address *</label>
              <input
                type="text"
                id="street1"
                name="street1"
                placeholder="Address Line 1"
                v-model="form.address1"
              />
            </div>
            <div class="row">
              <input
                type="text"
                id="street2"
                name="street2"
                placeholder="Address Line 2"
                v-model="form.address2"
              />
            </div>
            <div class="row">
              <input
                type="text"
                id="city"
                name="city"
                placeholder="City"
                v-model="form.city"
              />
            </div>
            <div class="row">
              <div
                style="
                  display: grid;
                  grid-template-columns: 2fr 1fr;
                  column-gap: 0.5rem;
                "
              >
                <input
                  type="text"
                  id="state"
                  name="state"
                  placeholder="State or Province"
                  v-model="form.state"
                />
                <input
                  type="tel"
                  id="postcode"
                  name="postcode"
                  placeholder="Postcode"
                  v-model="form.postcode"
                />
              </div>
            </div>
            <div class="row" style="position: relative">
              <input
                type="text"
                id="country"
                name="country"
                placeholder="Country"
                v-model="form.country"
              />
            </div>
            <div class="row">
              <label for="phone">Phone Number</label>
              <div
                style="
                  display: grid;
                  grid-template-columns: 1fr 2fr;
                  column-gap: 0.5rem;
                "
              >
                <input
                  type="text"
                  id="countrycode"
                  name="countrycode"
                  placeholder="Country Code"
                  v-model="form.phoneCountryCode"
                />
                <input
                  type="tel"
                  id="phone"
                  name="phone"
                  placeholder="Phone Number"
                  v-model="form.phoneNumber"
                />
              </div>
            </div>
          </form>
          <h3>III. Ready to Send</h3>
          <p>
            Please double check your desired artworks and make sure the above
            shipping information is valid before you click "send".
          </p>
          <div class="btn-set">
            <button type="button" class="primary" id="send" @click="submit">
              Send
            </button>
            <button type="button" @click="reset">Reset</button>
          </div>
        </div>
      </div>
    </main>

    <footer></footer>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import Web3 from "web3";
import { AbiItem } from "web3-utils";
import Web3Modal from "web3modal";
import WalletConnectProvider from "@walletconnect/web3-provider";

import axios from "axios";

import exhibitionABI from "./FeralFileExhibitionV2.json";

const providerOptions = {
  walletconnect: {
    package: WalletConnectProvider, // required
    options: {
      infuraId: "INFURA_ID", // required
    },
  },
};

// [Feral File — Social Codes (FF001)](https://etherscan.io/address/0x28b51ba8b990c48cb22cb6ef0ad5415fdba5210c)
// [Feral File — The Bardo: Unpacking the Real (FF002)](https://etherscan.io/address/0xaa02cc02f4531ee75d1b78cb5a155d4f3b54f830)
// [Feral File — Fragments of a Hologram Rose (FF003)](https://etherscan.io/address/0x979316f5b3f3d8db956af519553c853525a5b1af)
// [Feral File — The Long Cut (FF004)](https://etherscan.io/address/0x6e906b2e355294a6aecd6b4f75816eda9f703dda)
// [Feral File — Instructions Follow (FF005)](https://etherscan.io/address/0x1d5bdc75918600541c115b74b81a404c9e4af7d4)
// [Feral File — P1x3L (FF006)](https://etherscan.io/address/0xe5163c74ffe6563d75d750e5d767122500a1c337)
// [Feral File — Reflections in the Water (FF007)](https://etherscan.io/address/0x29c9e04e05c5d261836e458bc5b779a7de3c58d6)
// [Feral File — Polyarrythmia (FF008)](https://etherscan.io/address/0x63c8282c8705e7873b3302bd623b2bc8ebcdddd3)
// [Feral File — Unsupervised (FF009)](https://etherscan.io/address/0x7a15b36cb834aea88553de69077d3777460d73ac)
// [Feral File — –GRAPH (FF010)](https://etherscan.io/address/0xdb5f1adcffa1869b9711cbfbe3bf46cc5d5319e5)

const web3Modal = new Web3Modal({
  network: "mainnet", // optional
  cacheProvider: true, // optional
  providerOptions, // required
});

@Options({
  components: {},

  async created() {
    let query = this.$route.query;

    if (query.apiRoot) {
      this.apiRoot = query.apiRoot;
    }
    if (query.testAccountNumber) {
      this.testAccountNumber = query.testAccountNumber;
    }

    if (query.code) {
      this.web3Logout();
      this.code = query.code;
    } else {
      if (web3Modal.cachedProvider) {
        this.web3Login();
      }
    }
  },

  methods: {
    async web3Login() {
      const provider = await web3Modal.connect();
      console.log(provider);

      this.web3 = new Web3(provider);
      window.ethereum.on("accountsChanged", (accounts: Array<string>) => {
        console.log("accountsChanged");
        this.accountNumber = accounts[0];
      });

      this.contract = new this.web3.eth.Contract(
        exhibitionABI.abi as AbiItem[],
        "0x7a15b36cb834aea88553de69077d3777460d73ac"
      );

      let accounts = await this.web3.eth.getAccounts();
      this.accountNumber = accounts[0];

      let resp = await axios.get(
        this.apiRoot + "/api/owned/" + this.queryAccount
      );

      resp.data.artworks.forEach((ownedToken: any) => {
        this.pushToken(ownedToken.token_id);
      });
    },

    async pushToken(tokenID: string) {
      let tokenURI = await this.contract.methods.tokenURI(tokenID).call();
      let metadata: any = await axios.get(tokenURI);

      this.tokens.push({
        id: tokenID,
        name: metadata.data.name,
        imageURL: metadata.data.image,
        selected: false,
      });
    },

    web3Logout() {
      if (web3Modal.cachedProvider == "wallet connect") {
        this.web3.eth.currentProvider.disconnect();
      }
      web3Modal.clearCachedProvider();
      this.provider = null;
      this.web3 = null;
      this.accountNumber = "";
      this.balance = 0;
      this.tokens = [];
      this.reset();
    },

    selectToken(item: any) {
      if (item.selected) {
        item.selected = false;
        this.totalSelected -= 1;
      } else {
        if (this.totalSelected < this.maxSelectableNumber) {
          item.selected = true;
          this.totalSelected += 1;
        }
      }
    },

    async submit() {
      let selectedTokens: string[] = [];
      this.tokens.forEach((token: any) => {
        if (token.selected) {
          selectedTokens.push(token.id);
        }
      });

      let data = {
        information: this.form,
        tokens: selectedTokens,
      };

      let now = (+new Date()).toString();
      var hash = this.web3.utils.keccak256(now);

      let signature: string = await this.web3.eth.personal.sign(
        hash,
        this.accountNumber
      );

      try {
        let resp = await axios.post(this.apiRoot + "/api/claim", data, {
          headers: {
            Authorization:
              "Bearer " + signature.slice(2) + Buffer.from(now).toString("hex"),
            Requester: this.accountNumber,
          },
        });
        console.log(resp);
      } catch (error) {
        alert(error.response.data.error);
      }
    },

    reset() {
      this.tokens.forEach((token: any) => {
        token.selected = false;
      });
      this.form = {};
      this.totalSelected = 0;
    },
  },

  computed: {
    maxSelectableNumber() {
      return 3 * Math.floor(this.tokens.length / 3);
    },

    queryAccount() {
      return this.testAccountNumber || this.accountNumber;
    },
  },

  data() {
    return {
      web3: null,
      accountNumber: "",

      // form data
      form: {},
      tokens: [],
      totalSelected: 0,

      apiRoot: "",
    };
  },
})
export default class Home extends Vue {}
</script>
