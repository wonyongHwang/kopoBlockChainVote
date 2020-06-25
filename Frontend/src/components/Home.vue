<template>
  <div class="posts">
    <h1>찍먹? 부먹? 승자는?</h1>
    <h3>아래의 입력란에 계정 정보를 입력 후 투표하세요.</h3>
    <!--span><b>{{ response }}</b></span><br /-->
    <form v-on:submit="initCall">
      <input type="hidden" v-model="loginData.voterId" placeholder="Enter VoterId">
      <br>

      <input type="submit" value="투표 입장">
      <br>
      <br>
      <span v-if="loginReponse">
        <b>{{ loginReponse.data }}</b>
      </span>
      <br>
    </form>

    <br>
    <h3>계정 인증</h3>
    <form v-on:submit="registerVoter">
      <!-- <input type="text" v-model="registerData.voterId" placeholder="Enter Drivers License">
      <br>
      <input type="text" v-model="registerData.registrarId" placeholder="Enter Registrar ID">
      <br> -->
      <input type="text" v-model="registerData.firstName" placeholder="이름">
      <br>
      <input type="text" v-model="registerData.org" placeholder="조직">
      <br>
      <input type="submit" value="사용자 확인">
    </form>
    <br>
    <span v-if="registerReponse">
      <b>{{ registerReponse.data.token }}</b>
    </span>
    <br>
    <vue-instant-loading-spinner id='loader' ref="Spinner"></vue-instant-loading-spinner>
  </div>
</template>

<script>
import PostsService from "@/services/apiService";
//import { EventBus } from '../event-bus';
import VueInstantLoadingSpinner from "vue-instant-loading-spinner/src/components/VueInstantLoadingSpinner.vue";

export default {
  name: "response",
  data() {
    return {
      loginData: {},
      registerData: {},
      registerReponse: {
        data: ""
      },
      loginReponse: {
        data: ""
      }
    };
  },
  components: {
    VueInstantLoadingSpinner
  },
  methods: {
    async registerVoter() {

      await this.runSpinner();
      const apiResponse = await PostsService.registerVoter(
        this.registerData.firstName,
        this.registerData.org
      );

      console.log(apiResponse);
      this.registerReponse = apiResponse;
      console.log(this.registerReponse.data.token)
      //this.registerReponse = apiResponse.data.token
      await this.hideSpinner();
    },

    async initCall() {
      if(!this.registerReponse.data.token){
           console.log("empty token");
           this.loginReponse.data = "계정이 인증되지 않았습니다.";
      }else{
        this.$router.push({name: "CastBallot", params: {tokenId : this.registerReponse.data.token }});
      }

      //await this.runSpinner();

      // if (!this.loginData.voterId) {
      //   console.log("!thislogin");
      //   let response = 'Please enter a VoterId';
      //   this.loginReponse.data = response;
      //   await this.hideSpinner();
      // } else 
      /*
      {
        const apiResponse = await PostsService.initCall(
          //this.loginData.voterId
          this.registerReponse.data.token
        );
        console.log("apiResponse");
        console.log(apiResponse.data);

        if (apiResponse.data.error) {
          // console.log(apiResponse);
          console.log(apiResponse.data.error);
          this.loginReponse = apiResponse.data.error;
        } else {
         //  EventBus.$emit('Token',this.registerReponse.data.token);
         //$EventBus.$emit('Token',this.registerReponse.data);
          this.$router.push({name: "CastBallot", params: {tokenId : this.registerReponse.data.token }});
        }

        console.log(apiResponse);
        this.loginReponse = apiResponse;
        // this.$router.push('castBallot')
        // await this.hideSpinner();
      }
      */
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
