<style scoped>
.names {
  display: grid;
  grid-template-columns: 1fr;
  grid-gap: 0.5rem;
}
p.note {
  color: #555;
  font-size: .75rem;
}
@media screen and (min-width: 1024px) {
  .names {
    grid-template-columns: 1fr 1fr;
    row-gap: 0;
  }
}
</style>

<template>
  <div v-if="loggedIn && hasToken">
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
              <input type="email" id="email" name="email" placeholder="eg. john@bitmark.com" v-model="form.email" />
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
              <div style=" display: grid; grid-template-columns: 1fr 2fr; column-gap: 0.5rem; ">
                <input type="text" id="countrycode" name="countrycode" placeholder="Country Code" v-model="form.phoneCountryCode" />
                <input type="tel" id="phone" name="phone" placeholder="Phone Number" v-model="form.phoneNumber" />
              </div>
              <div class="input-errors" v-for="(error, index) of v$.form.phoneCountryCode.$errors" :key="index">
                <div class="error-msg">{{ error.$message }}</div>
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

import exhibitionABI from "./FeralFileExhibitionV2.json";

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

type StringMap = {
  [key: string]: string;
};

const imageSrc: StringMap = {
  "01b502397dcfdd7337b62a45694ceb219eb8513577646cb46f357b379719c06d49fe2eef9bccb98c98d7f09915d5751c30afb720aaff4fd30401e2253c2588c0":
    "https://cdn.feralfileassets.com/thumbnails/e601569d-5611-4a82-93ba-a7f55b260001/1637053732",
  "539d0f5499fb82cc3d519d68dcffc540e8278ea55e22de4229cd1f32f5b3e805a8dc1594155163b7bf47258d461b66edeff32f5ae395d984e34ec3bfbad37249":
    "https://cdn.feralfileassets.com/thumbnails/96f1663b-5216-4758-aa7f-3a70e6d9d45b/1637052722",
  a1a1c08ec9c7d62c2496b75c6ebd37ed1e2f59c83694ada29b1b2d9e3507a81f9d0be6bb34079fe6f38d37a47fefa08ed1649355b8bd8d51c400dc0f92ae55dd:
    "https://cdn.feralfileassets.com/thumbnails/61c2c2b6-4f2c-4e1b-9e45-23faf3e76594/1637046234",
  b294e029328010fd8e71954c3527d4acf978c7fc5091707a780cfc98baed45e4180d8d3ce6fc65a6fd944db9805df8c0a93e8bb874cc80e5f789fc177019c528:
    "https://cdn.feralfileassets.com/thumbnails/29b720ff-fcbf-4467-a53d-e79a90ad0d60/1637054909",
  "10781635a7ca427c9ca0831018e63b668d56bf9493f6d59fbe8e6cab9585660aa00ab14578ae8f7c391b6f6d437d1fc93418d4c70a90f4c32314c4b2e26bf4ce":
    "https://cdn.feralfileassets.com/thumbnails/d35cdb38-917f-42f3-9914-e276dc5f0ded/1637038450",
  b1f969b7704cb938d1ab5d9ef2ee5fdfb8d21b2aa5c5e016b18ee5f9286ff64787523f35a0d414b74267789230969c20ea98cb79352f7893ecb2347d77c387a7:
    "https://cdn.feralfileassets.com/thumbnails/821dd86a-8764-441f-af58-95c857a7c4d3/1637046126",
  b3b9a02ce920bbc94e72b9f64884b68a95163bb88813e9e8f25754bf0ea84c0efb32c546b54f0244278aaa049cf263ec45252ce828108cf33782f7e3dc2332db:
    "https://cdn.feralfileassets.com/thumbnails/c3b2c10b-91c6-4e2f-83bc-fcb5068e98be/1637046195",
  "19df39ff870c1229962a53b754b797a48cf6dd0b2c5b93267d10be0d3bf4da408bb7aa7a5c421a22faa38b378f1c9fd41b09caf9d1021fb72079a0fd8f5dac6b":
    "https://cdn.feralfileassets.com/thumbnails/e362fa91-c218-45df-8167-d7e7c3440052/1637046307",
  "76460ab711429437bfc0c3b46141e068e7451f5da5e0bbdbbb2d17b939f78a0b73e5d8bc6e6faa402407567d632a9f5efaa93e1111bf92ea809aa9e60b7c9fc3":
    "https://cdn.feralfileassets.com/thumbnails/29cd44b3-2ed9-4d8b-beb5-073d248a1a00/1637046395",
  c41d4e6d2e2d5f17e325b62ab2af801c850d400100480f13c7d5e9dc74b7fe0cebcf701c9f7e77e5c0b840892aa0f47fd1162bea1106a16bddf4709d53e54223:
    "https://cdn.feralfileassets.com/thumbnails/c4c9bfde-13bc-4cf5-bc4d-817280673ec3/1637046279",
  "574eafa537953a1af58467788a9a84640c6a80c072d625e2315b111e445ecaf280a1fdb4ceb1043e0d3d03bec85a5ebfb1cae1e9f9d75a9eef11c846e063139a":
    "https://cdn.feralfileassets.com/thumbnails/31bc21a0-c834-49a3-866f-d4839cbe14df/1637046336",
  "56c69488317e1c464a20b3c6d025d4aa0f4c15d373b81f453b28e5662dc0cd840e4883541f95a2528d2212faa5a00f49d691109afe96a0da04c303185ceec972":
    "https://cdn.feralfileassets.com/thumbnails/6a0781b7-93ff-4cef-9a92-6e197990be57/1637046364",
  "105520675453f8b45271044006420520186e5a948e05329d7d08e15fa21f1b916a4754b4ea7db06cff0e3cc71b777043f2173554c037090e93f9489ca53135ee":
    "https://cdn.feralfileassets.com/thumbnails/d6e1e0f1-7a9d-46f1-a691-bf0604e755b5/1637053332",
};

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

    this.loginByEmail = query.code && query.requester;
    if (this.loginByEmail) {
      this.web3Logout();
      this.accountNumber = query.requester;
      this.code = query.code;
    } else {
      if (this.$web3Modal.cachedProvider) {
        await this.web3Login();
      } else {
        this.$router.push({ name: "Welcome" });
        return;
      }
    }

    this.loggedIn = true;
    await this.checkSubmission();
    await this.loadToken();
  },

  methods: {
    async checkSubmission() {
      try {
        await axios.head(
          this.apiRoot + "/api/claim?requester=" + this.queryAccount,
          {
            headers: {
              Requester: this.queryAccount,
            },
          }
        );
        this.$router.push({ name: "ThankYou" });
      } catch (error) {
        console.log(error);
      }
    },

    async web3Login() {
      const provider = await this.$web3Modal.connect();

      this.web3 = new Web3(provider);
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

      this.contract = new this.web3.eth.Contract(
        exhibitionABI.abi as AbiItem[],
        "0x7a15b36cb834aea88553de69077d3777460d73ac"
      );

      let accounts = await this.web3.eth.getAccounts();
      this.accountNumber = accounts[0];
    },

    async loadToken() {
      let resp = await axios.get(
        this.apiRoot + "/api/owned/" + this.queryAccount
      );

      if (resp.data.artworks.length < 3) {
        this.$router.push({ name: "Nothing" });
      }

      this.hasToken = true;

      resp.data.artworks.forEach(async (ownedToken: any) => {
        if (this.loginByEmail) {
          await this.pushTokenByBitmark(ownedToken.token_id);
        } else {
          await this.pushTokenByEthereum(ownedToken.token_id);
        }
      });
    },

    async pushTokenByBitmark(tokenID: string) {
      let resp: any = await axios.get(
        "https://api.bitmark.com/v1/bitmarks/" + tokenID + "?asset=true"
      );

      let data = resp.data;
      this.tokens.push({
        id: tokenID,
        name: data.asset.metadata.title,
        imageURL: imageSrc[data.asset.id],
        selected: false,
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

    web3Logout() {
      if (this.$web3Modal.cachedProvider == "walletconnect") {
        this.web3.eth.currentProvider.disconnect();
      }
      this.$web3Modal.clearCachedProvider();

      if (!this.loginByEmail) {
        this.$router.push({ name: "Welcome" });
      }
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
      return 3 * Math.floor(this.tokens.length / 3);
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
        phoneCountryCode: {
          required: helpers.withMessage(
            "Please fill in your phone country code",
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
        phoneCountryCode: "",
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
