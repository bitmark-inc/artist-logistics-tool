<style scoped>
.names {
  display: grid;
  grid-template-columns: 1fr;
  grid-gap: 0.5rem;
}
p.note {
  color: #555;
  font-size: 0.75rem;
}
@media screen and (min-width: 1024px) {
  .names {
    grid-template-columns: 1fr 1fr;
    row-gap: 0;
  }
}
</style>

<template>
  <div v-if="loggedIn">
    <header>
      <a class="brand">
        <img src="img/au.svg" />
        <h2>Autonomy
          <!--<span> Market<span>-->
        </h2>
      </a>
      <p class="wallet">
        <span>{{ accountNumber }}</span>
        <button class="secondary" @click="this.web3Logout" v-if="web3 != null" :disabled="!web3">
          Disconnect
        </button>
      </p>
    </header>
    <main>
      <div v-if="hasToken">
        <h2 class="text-center">Collect Prints</h2>
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
              <div class="card" :class="{ selected: token.selected }" v-for="token in tokens" :key="token.id" @click="selectToken(token)">
                <div class="image" :style="{ 'background-image': 'url(' + token.imageURL + ')' }"></div>
                <div class="info">
                  <p>{{ token.name }}</p>
                </div>
              </div>
            </div>
            <p class="note">Something missing? Contact <a href="mailto:support@feralfile.com">support@feralfile.com</a> with your Feral File alias or email. </p>
          </div>
          <div class="right">
            <h3>II. Fill Out the Following Information</h3>
            <form>
              <div class="row">
                <label>Name *</label>
                <div class="names">
                  <input type="text" id="first" name="first" placeholder="First Name" v-model="form.firstName" />
                  <input type="text" id="last" name="last" placeholder="Last Name" v-model="form.lastName" />
                  <div class="input-errors" v-for="(error, index) of v$.form.firstName.$errors" :key="index">
                    <div class="error-msg">{{ error.$message }}</div>
                  </div>
                  <div class="input-errors" v-for="(error, index) of v$.form.lastName.$errors" :key="index">
                    <div class="error-msg">{{ error.$message }}</div>
                  </div>
                </div>
              </div>
              <div class="row">
                <label for="email">Email *</label>
                <input type="email" id="email" name="email" placeholder="eg. john@bitmark.com" autocomplete="email" v-model="form.email" />
                <div class="input-errors" v-for="(error, index) of v$.form.email.$errors" :key="index">
                  <div class="error-msg">{{ error.$message }}</div>
                </div>
              </div>
              <div class="row">
                <label>Shipping Address *</label>
                <input type="text" id="street1" name="street1" placeholder="Address Line 1" v-model="form.address1" />
                <div class="input-errors" v-for="(error, index) of v$.form.address1.$errors" :key="index">
                  <div class="error-msg">{{ error.$message }}</div>
                </div>
              </div>
              <div class="row">
                <input type="text" id="street2" name="street2" placeholder="Address Line 2" v-model="form.address2" />
              </div>
              <div class="row">
                <input type="text" id="city" name="city" placeholder="City" v-model="form.city" />
                <div class="input-errors" v-for="(error, index) of v$.form.city.$errors" :key="index">
                  <div class="error-msg">{{ error.$message }}</div>
                </div>
              </div>
              <div class="row">
                <div style=" display: grid; grid-template-columns: 2fr 1fr; column-gap: 0.5rem;">
                  <input type="text" id="state" name="state" placeholder="State or Province" v-model="form.state" />
                  <input type="tel" id="postcode" name="postcode" placeholder="Postcode" v-model="form.postcode" />
                  <div class="input-errors" v-for="(error, index) of v$.form.state.$errors" :key="index">
                    <div class="error-msg">{{ error.$message }}</div>
                  </div>
                  <div class="input-errors" v-for="(error, index) of v$.form.postcode.$errors" :key="index">
                    <div class="error-msg">{{ error.$message }}</div>
                  </div>
                </div>
              </div>
              <div class="row" style="position: relative">
                <input type="text" id="country" name="country" placeholder="Country" v-model="form.country" />
                <div class="input-errors" v-for="(error, index) of v$.form.country.$errors" :key="index">
                  <div class="error-msg">{{ error.$message }}</div>
                </div>
              </div>
              <div class="row">
                <label for="phone">Phone Number *</label>
                <div>
                  <input type="tel" id="phone" name="phone" placeholder="Phone Number (With Country Code)" autocomplete="tel" v-model="form.phoneNumber" />
                </div>
                <div class="input-errors" v-for="(error, index) of v$.form.phoneNumber.$errors" :key="index">
                  <div class="error-msg">{{ error.$message }}</div>
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
      </div>
      <div v-else>
        <h2>No token found</h2>
      </div>
    </main>
    <div class="modal" :class="{ show: openErrorModal }">
      <div class="modal-inner">
        <div class="modal-head"></div>
        <div class="modal-body">
          <h3>Error</h3>
          <p>{{ errorModalMessage }}</p>
        </div>
        <div class="modal-foot">
          <div class="btn-set">
            <button class="btn primary" @click="closeErrorModel">Okay</button>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script lang="ts">
import { Options, Vue, setup } from "vue-class-component";
import Web3 from "web3";
import { AbiItem } from "web3-utils";

import useVuelidate from "@vuelidate/core";
import { required, email, helpers } from "@vuelidate/validators";

import axios from "axios";

declare const window: any;

const ERC721TokenURIABI = [
  {
    inputs: [
      {
        internalType: "uint256",
        name: "tokenId",
        type: "uint256",
      },
    ],
    name: "tokenURI",
    outputs: [
      {
        internalType: "string",
        name: "",
        type: "string",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
];

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

    if (this.$web3Modal.cachedProvider) {
      await this.web3Login();
    } else {
      this.$router.push({ name: "Welcome" });
      return;
    }

    this.loggedIn = true;

    if (this.$web3Modal.cachedProvider === "injected") {
      try {
        window.ethereum.on("accountsChanged", (accounts: Array<string>) => {
          console.log("ethereum accounts changed");
          this.accountNumber = accounts[0];
        });
      } catch (error) {
        console.log("can not listen to account changed events");
      }
    }

    await this.checkSubmission();
    await this.loadToken();
  },

  methods: {
    async checkSubmission() {
      try {
        await axios.head(this.apiRoot + "/api/claim/" + this.queryAccount, {
          headers: {
            LogisticID: this.logisticID,
          },
        });
        this.$router.push({ name: "ThankYou" });
      } catch (error) {
        console.log(error);
      }
    },

    async web3Login() {
      const provider = await this.$web3Modal.connect();

      this.web3 = new Web3(provider);

      this.contract = new this.web3.eth.Contract(
        ERC721TokenURIABI as AbiItem[],
        "0x7a15b36cb834aea88553de69077d3777460d73ac"
      );

      let accounts = await this.web3.eth.getAccounts();
      this.accountNumber = accounts[0];
    },

    async loadToken() {
      let resp = {
        data: {
          artworks: [
            {
              token_id:
                "0x144a21328383edb0a5a14ea5e990edf734d7d8b8cbc967c03cbab6e04ff5f116",
              owner_address: "0x480eA382A94BB8cc9f30bEbF7031019749c429d2",
            },
            {
              token_id:
                "0x78e500466ffc3d6ce4ada8dbb9efbc07bd8bc322ca8756e60294fb2f99feb21b",
              owner_address: "0x480eA382A94BB8cc9f30bEbF7031019749c429d2",
            },
          ],
        },
      };

      if (resp.data.artworks.length < 1) {
        this.$router.push({
          name: "Nothing",
          params: {
            message:
              "We were unable to find any artworks under your control that qualify for receiving prints.",
          },
        });
      }

      this.hasToken = true;

      resp.data.artworks.forEach(async (ownedToken: any) => {
        await this.pushTokenByEthereum(ownedToken.token_id);
      });
    },

    async pushTokenByEthereum(tokenID: string) {
      let tokenURI = await this.contract.methods.tokenURI(tokenID).call();
      let metadata: any = await axios.get(tokenURI);

      this.tokens.push({
        id: tokenID,
        name: metadata.data.name,
        imageURL: metadata.data.image,
        selected: false,
      });
    },

    async web3Logout() {
      if (this.$web3Modal.cachedProvider == "walletconnect") {
        if (this.web3) {
          this.web3.eth.currentProvider.disconnect();
        }
      }

      await this.$web3Modal.clearCachedProvider();
      this.$router.push({ name: "Welcome" });
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
      this.v$.$touch();
      if (this.v$.$error) {
        return;
      }

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

      let authToken: string;

      if (this.web3) {
        let now = (+new Date()).toString();
        var hash = this.web3.utils.keccak256(now);
        let signature: string = await this.web3.eth.personal.sign(
          hash,
          this.accountNumber
        );
        authToken = signature.slice(2) + Buffer.from(now).toString("hex");
      } else {
        authToken = this.code;
      }

      try {
        let resp = await axios.post(this.apiRoot + "/api/claim", data, {
          headers: {
            Authorization: "Bearer " + authToken,
            Requester: this.accountNumber,
            LogisticID: this.logisticID,
          },
        });
        this.$router.push({ name: "ThankYou" });
      } catch (error) {
        this.showErrorModal("Server error: " + error.response.data.error);
      }
    },

    showErrorModal(message: string) {
      this.errorModalMessage = message;
      this.openErrorModal = true;
    },

    closeErrorModel() {
      this.openErrorModal = false;
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
      return 1;
    },

    queryAccount() {
      return this.testAccountNumber || this.accountNumber;
    },
  },

  validations() {
    return {
      form: {
        firstName: {
          required: helpers.withMessage(
            "Please fill in your first name",
            required
          ),
        },
        lastName: {
          required: helpers.withMessage(
            "Please fill in your last name",
            required
          ),
        },
        email: {
          required: helpers.withMessage("Please fill in your email", required),
          email: helpers.withMessage("Invalid email address", email),
        },
        address1: {
          required: helpers.withMessage(
            "Please fill in your address",
            required
          ),
        },
        city: {
          required: helpers.withMessage("Please fill in your city", required),
        },
        state: {
          required: helpers.withMessage("Please fill in your state", required),
        },
        postcode: {
          required: helpers.withMessage(
            "Please fill in your postcode",
            required
          ),
        },
        country: {
          required: helpers.withMessage(
            "Please fill in your country",
            required
          ),
        },
        phoneNumber: {
          required: helpers.withMessage(
            "Please fill in your phone number",
            required
          ),
        },
      },
    };
  },

  data() {
    return {
      web3: null,
      loggedIn: false,
      hasToken: false,
      accountNumber: "",
      logisticID: "test-001",

      // form data
      form: {
        firstName: "",
        lastName: "",
        email: "",
        address1: "",
        city: "",
        state: "",
        postcode: "",
        country: "",
        phoneNumber: "",
      },
      tokens: [],
      totalSelected: 0,

      apiRoot: "",

      openErrorModal: false,
      errorModalMessage: "",
    };
  },
})
export default class Home extends Vue {
  v$ = setup(() => useVuelidate());
}
</script>
