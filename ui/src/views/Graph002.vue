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
                <input type="tel" id="phone" name="phone" placeholder="Phone Number (With Country Code)" autocomplete="tel" v-model="form.phoneNumber" />
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
import Web3 from "web3";
import { AbiItem } from "web3-utils";

import useVuelidate from "@vuelidate/core";
import { required, email, helpers } from "@vuelidate/validators";

import axios from "axios";

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

      let authToken = this.code;

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
