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
      </p>
    </header>

    <main>
      <h2 class="text-center">Claim Printed Artworks from –GRAPH</h2>
      <ol>
        <li>
          Review your collected artworks from the show –GRAPH.
        </li>
        <li>Fill in the shipping information.</li>
        <li>Confirm and click "Send".</li>
      </ol>
      <div class="two-col">
        <div class="left">
          <h3>I. Review Artworks</h3>
          <strong style="display: block; margin-bottom: 0.5rem">
            * Please review the following {{ maxSelectableNumber }} items from your collected
            digital editions for prints.
          </strong>
          <div class="grids">
            <div class="card" v-for="token in tokens" :key="token.id">
              <div class="image" :style="{ 'background-image': 'url(' + token.imageURL + ')' }"></div>
              <div class="info">
                <p>{{ token.name }}</p>
              </div>
            </div>
          </div>
          <p class="note">Something missing? Contact <a href="mailto:support@feralfile.com">support@feralfile.com</a> with your Feral File alias or email.</p>
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
                <!--<input type="text" id="countrycode" name="countrycode" placeholder="Country Code" autocomplete="tel" v-model="form.phoneCountryCode" />-->
                <input type="tel" id="phone" name="phone" placeholder="Phone Number (with country code)" autocomplete="tel" v-model="form.phoneNumber" />
              </div>
              <!--<div class="input-errors" v-for="(error, index) of v$.form.phoneCountryCode.$errors" :key="index">
                <div class="error-msg">{{ error.$message }}</div>
              </div>-->
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

import useVuelidate from "@vuelidate/core";
import { required, email, helpers } from "@vuelidate/validators";

import axios from "axios";

type StringMap = {
  [key: string]: string;
};

const imageSrc: StringMap = {
  "4fd83a576ee3d7788eb58704b021f1dd4683abc7c19709ecc4284d907fb22c32bfd4efa83b6a5a5397733edba9de416628e2b7564fe8f4bad6beddc4a4e11f5b":
    "https://cdn.feralfileassets.com/thumbnails/cee5b80c-af17-44d0-9a56-c17a68f92569/1637146824",
  b11df60a3ac1f87204558587531b6acec9bc841513bfb8dffae240979edb2aa7c2e11933685a6c149cc864b4faa04861115699129841b64f233d512b80169c22:
    "https://cdn.feralfileassets.com/thumbnails/5e0a93aa-22d0-4305-8bdd-5806dab5503a/1636853639",
  e757eabdb3dba00db25ba720be40c3537c5311a70b071b1ea40ae24fff380f36325dd45e291236078426709c36bf8f03a42acac442eb5510afb6c18abd2e9a0c:
    "https://cdn.feralfileassets.com/thumbnails/1b1ffd48-5fc2-4d19-a563-364e00ddfb01/1637210985",
  e71d0b1635546929cea414e34cc8f9e7c14d913a7d1cac2dbb9d9e3764ecb503e1d65efb5c115fb8625dab8810d8fb9a112df54a5ecfe72d54685d43875fe52d:
    "https://cdn.feralfileassets.com/thumbnails/2fb649d0-3ed2-4539-8ea2-e073b85e9a21/1637249094",
  "4fa259125b48951aab7dd988cecbe20cd12b2ecf49dadb92771ba3e8ec6df57220f6186a628bd9915ba703995fcc60522d0772c069b849c62db55b2afcaccc6f":
    "https://cdn.feralfileassets.com/thumbnails/99d373e5-cc9d-45e8-8aa5-1ed47e481791/1637145106",
  "8ea9d880bd9e39675d437992c3d52f9b334d8e8281dd315df06b982b863d98847200db0813bca302addc58c5f51b3583966fc4f31cb8e8793d58643bdf447836":
    "https://cdn.feralfileassets.com/thumbnails/e7252d16-c248-4de0-aade-433e9e7d2484/1637677549",
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
      this.accountNumber = query.requester;
      this.code = query.code;
    } else {
      this.$router.push({ name: "NotFound" });
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
              LogisticID: this.logisticID,
            },
          }
        );
        this.$router.push({ name: "ThankYou" });
      } catch (error) {
        console.log(error);
      }
    },

    async loadToken() {
      let resp = await axios.get(
        this.apiRoot + "/api/owned/" + this.queryAccount
      );

      if (resp.data.artworks.length < 3) {
        this.$router.push({
          name: "Nothing",
          params: {
            message:
              'We were unable to find any artworks under your control that qualify for receiving prints. Email <a href="mailto:support@feralfile.com">support@feralfile.com</a> or visit our Discord <a href="https://discord.gg/NHMJAnjjCS" target="_blank">troubleshooting channel</a> if you have any followup questions.',
          },
        });
      }

      this.hasToken = true;

      resp.data.artworks.forEach(async (ownedToken: any) => {
        await this.pushTokenByBitmark(ownedToken.token_id);
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
      });
    },

    async submit() {
      this.v$.$touch();
      if (this.v$.$error) {
        return;
      }

      let data = {
        information: this.form,
      };

      let authToken = this.code;

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
      this.form = {};
      this.totalSelected = 0;
    },
  },

  computed: {
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
      loggedIn: false,
      hasToken: false,
      accountNumber: "",

      logisticID: "graph-002",

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
export default class Graph002 extends Vue {
  v$ = setup(() => useVuelidate());
}
</script>
