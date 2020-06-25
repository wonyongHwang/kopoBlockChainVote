<template>
  <div class="posts">
    <h1>Cast Ballot</h1>
    <input type="radio" id="one" value="a" v-model="picked">
    <label for="a">부먹</label>
    <br>
    <input type="radio" id="two" value="b" v-model="picked">
    <label for="b">찍먹</label>
    <br>
    <input type="radio" id="three" value="c" v-model="picked">
    <label for="c">아무 생각 없음</label>
    <br>
    <span v-if="picked">
      Picked:
      <b>{{ picked }}</b>
    </span>
    <br>
    <br>
    <!--span><b>{{ response }}</b></span><br /-->
    <form v-on:submit="castBallot">
      <input type="submit" value="투표">
      <br>
    </form>

    <br>
    <span v-if="response">
      <b>{{ response }}</b>
    </span>
    <br>
    <vue-instant-loading-spinner id="loader" ref="Spinner"></vue-instant-loading-spinner>
  </div>
</template>

<script>
import PostsService from "@/services/apiService";
import VueInstantLoadingSpinner from "vue-instant-loading-spinner/src/components/VueInstantLoadingSpinner.vue";
//import { EventBus } from '../event-bus';
// const clickHandler = function(clickCount) {
//   console.log(`Oh, that's nice. It's gotten ${clickCount} clicks! :)`)
// }
export default {
  name: "response",
  data() {
    return {
      input: {},
      picked: null,
      response: null,
      token : null
    }
  },
  created(){
    this.token = this.$route.params.tokenId
    console.log("this.$route.params.tokenId ${this.$route.params.tokenId}")
  },
  components: {
    VueInstantLoadingSpinner
  },
  methods: {
    clickHandler(token) {
    console.log(`Oh, that's nice. It's gotten ${token} got! :)`);
    this.token = token;
    },

    async castBallot() {
      await this.runSpinner();

      // TODO : Check whether a user has already been voted

      // const electionRes = await PostsService.queryWithQueryString('election');

      // let electionId = electionRes.data[0].Key;

       console.log("picked: ");
       console.log(this.picked);
      // console.log("voterId: ");
      // console.log(this.input.voterId);
      // this.response = null;

 

      //error checking for making sure to vote for a valid party
      if (this.picked === null ) {
        console.log('this.picked === null')

        let response = "You have to pick a party to vote for!";
        this.response = response;
        await this.hideSpinner();
      
      } 
      // else if (this.input.voterId === undefined) {
      //   console.log('this.voterId === undefined')

      //   let response = "You have to enter your voterId to cast a vote!";
      //   this.response = response;
      //   await this.hideSpinner();

      // }
       else {
          console.log("token in castBallot: ");
          console.log(this.token);
        const apiResponse = await PostsService.castBallot(
          null,
          this.token,
          this.picked
        );

        console.log('apiResponse: &&&&&&&&&&&&&&&&&&&&&&&');
        console.log(apiResponse);

        if (apiResponse.data.error) {
          this.response= apiResponse.data.error;
          await this.hideSpinner();
        } else if (apiResponse.data.message) {
          this.response= apiResponse.data.message // `Could not find voter with voterId ${this.input.voterId} in the state. Make sure you are entering a valid voterId`;
          console.log(this.response);
          this.$router.push({name: "GetCurrentStanding", params: {tokenId : this.token }});
          await this.hideSpinner();
        } 
        else {
          let response = `Successfully recorded vote for ${this.picked} party 
            for voter with voterId ${apiResponse.data.voterId}. Thanks for 
            doing your part and voting!`;

          this.response = response;

          console.log("cast ballot");
          console.log(this.input);
          await this.hideSpinner();

          
        }
      }
    },
    async runSpinner() {
      this.$refs.Spinner.show();
    },
    async hideSpinner() {
      this.$refs.Spinner.hide();
    }
  }
};
</script>
