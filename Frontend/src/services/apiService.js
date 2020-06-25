import Api from '@/services/api'

export default {
  castBallot(electionId, tokenId, picked) {


    const headers = {
      'authorization': 'Bearer ' + tokenId,
      'content-type': 'application/json'
    }

    return Api().post('channels/mychannel/chaincodes/mycc', {
      fcn: "vote",
      peers: ["peer0.org1.example.com","peer0.org2.example.com"],
      args :[picked]
    },{headers : headers}) 
  },
  // queryAll() {
  //   return Api().get('queryAll')
  // },
  // queryByObjectType() {
  //   return Api().get('queryByObjectType')
  // },
  // queryWithQueryString(selected) {
  //   return Api().post('queryWithQueryString', {
  //     selected: selected
  //   }) 
  // },
  registerVoter(firstName, org) {
    return Api().post('users', {
      username: firstName,
      orgName: org
    }) 
  },
  initCall(token) {
    const headers = {
      'authorization': 'Bearer ' + token,
      'content-type': 'application/json'
    }

    return Api().post('channels/mychannel/chaincodes', {
      chaincodeName: "mycc",
      chaincodeVersion: "v0",
      chaincodeType: "go",
      args :["a","0","b","0","c","0"]
    },{headers : headers}) 
  },
  queryByKey(key) {
    return Api().post('queryByKey', {
      key: key
    }) 
  },
  getCurrentStanding(token) {
    const headers = {
      'authorization': 'Bearer ' + token,
      'content-type': 'application/json'
    }
    return Api().get('channels/mychannel/chaincodes/mycc', { params:{  fcn: "query",
    peer: "peer0.org1.example.com",
    args :["a","b","c"]}, headers: { 'authorization': 'Bearer ' + token, } })
    
  }
}