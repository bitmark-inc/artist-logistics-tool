<template>
  <div class="wrapper">
    <header>
      <a class="brand" href="home-signed-in.html">
        <img src="img/au.svg" />
        <h2>Autonomy<!--<span> Market<span>--></h2>
      </a>
    </header>

    <main>
      <h2 class="text-center">Collect Refik Anadol's Signed Prints</h2>

      <ol>
        <li>Select {number} desired items in the following artwork grid.</li>
        <li>Fill in the shipping information.</li>
        <li>Confirm and click "Send".</li>
      </ol>

      <div class="two-col">
        <div class="left">
          <h3>I. Select Artworks</h3>
          <strong style="display: block; margin-bottom: 0.5rem"
            >* Please select {number} items from your collected digital editions
            for signed prints.</strong
          >
          <div class="grids">
            <div class="card">
              <div
                class="image"
                style="background-image: url(img/01.jpg)"
              ></div>
              <div class="info">
                <p>Unsupervised — Machine Hallucinations — MoMA</p>
              </div>
            </div>

            <div class="card">
              <div
                class="image"
                style="background-image: url(img/02.jpg)"
              ></div>
              <div class="info">
                <p>
                  Unsupervised — Machine Hallucinations — MoMA Generative Study
                  1
                </p>
              </div>
            </div>

            <div class="card">
              <div
                class="image"
                style="background-image: url(img/03.jpg)"
              ></div>
              <div class="info">
                <p>Unsupervised — Machine Hallucinations — MoMA Fluid Dreams</p>
              </div>
            </div>

            <div class="card">
              <div
                class="image"
                style="background-image: url(img/04.jpg)"
              ></div>
              <div class="info">
                <p>Unsupervised — Machine Hallucinations — MoMA Dreams — A</p>
              </div>
            </div>

            <div class="card">
              <div
                class="image"
                style="background-image: url(img/05.jpg)"
              ></div>
              <div class="info">
                <p>Unsupervised — Machine Hallucinations — MoMA Dreams — B</p>
              </div>
            </div>

            <div class="card">
              <div
                class="image"
                style="background-image: url(img/06.jpg)"
              ></div>
              <div class="info">
                <p>Unsupervised — Machine Hallucinations — MoMA Dreams — C</p>
              </div>
            </div>

            <div class="card">
              <div
                class="image"
                style="background-image: url(img/07.jpg)"
              ></div>
              <div class="info">
                <p>Unsupervised — Machine Hallucinations — MoMA Dreams — D</p>
              </div>
            </div>

            <div class="card">
              <div
                class="image"
                style="background-image: url(img/08.jpg)"
              ></div>
              <div class="info">
                <p>Unsupervised — Machine Hallucinations — MoMA Dreams — E</p>
              </div>
            </div>

            <div class="card">
              <div
                class="image"
                style="background-image: url(img/09.jpg)"
              ></div>
              <div class="info">
                <p>Unsupervised — Machine Hallucinations — MoMA Dreams — F</p>
              </div>
            </div>

            <div class="card">
              <div
                class="image"
                style="background-image: url(img/10.jpg)"
              ></div>
              <div class="info">
                <p>Unsupervised — Machine Hallucinations — MoMA Dreams — G</p>
              </div>
            </div>

            <div class="card">
              <div
                class="image"
                style="background-image: url(img/11.jpg)"
              ></div>
              <div class="info">
                <p>Unsupervised — Machine Hallucinations — MoMA Dreams — H</p>
              </div>
            </div>

            <div class="card">
              <div
                class="image"
                style="background-image: url(img/12.jpg)"
              ></div>
              <div class="info">
                <p>Unsupervised — Machine Hallucinations — MoMA Dreams — I</p>
              </div>
            </div>

            <div class="card">
              <div
                class="image"
                style="background-image: url(img/13.jpg)"
              ></div>
              <div class="info">
                <p>Unsupervised — Data Universe — MoMA</p>
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
                  value=""
                />
                <input
                  type="text"
                  id="last"
                  name="last"
                  placeholder="Last Name"
                  value=""
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
                value=""
              />
            </div>
            <div class="row">
              <label>Shipping Address *</label>
              <input
                type="text"
                id="street1"
                name="street1"
                placeholder="Address Line 1"
                value=""
              />
            </div>
            <div class="row">
              <input
                type="text"
                id="street2"
                name="street2"
                placeholder="Address Line 2"
                value=""
              />
            </div>
            <div class="row">
              <input
                type="text"
                id="city"
                name="city"
                placeholder="City"
                value=""
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
                  value=""
                />
                <input
                  type="tel"
                  id="postcode"
                  name="postcode"
                  placeholder="Postcode"
                  value=""
                />
              </div>
            </div>
            <div class="row" style="position: relative">
              <input
                type="text"
                id="country"
                name="country"
                placeholder="Country"
                value=""
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
                  value=""
                />
                <input
                  type="tel"
                  id="phone"
                  name="phone"
                  placeholder="Phone Number"
                  value=""
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
            <button class="primary" id="send">Send</button>
            <button>Reset</button>
          </div>
        </div>
      </div>
    </main>

    <footer></footer>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";

@Options({
  components: {},
})
export default class Home extends Vue {}
</script>
