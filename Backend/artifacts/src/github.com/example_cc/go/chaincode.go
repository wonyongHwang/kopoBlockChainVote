/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("kopoVote")

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Init is somthing...["A","0","B","0","C,"0"]
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### kopoVote Init ###########")

	_, args := stub.GetFunctionAndParameters()
	var A, B, C string       // Entities
	var Aval, Bval, Cval int // Asset holdings
	var err error

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	C = args[4]
	Cval, err = strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	logger.Info("Aval = %d, Bval = %d, Cval = %d\n ", Aval, Bval, Cval)

	Avalbytes, err := stub.GetState(A)
	if Avalbytes == nil { // 체인코드 업그레이드에 대비한 방어로직, 업그레이드시 Init호출될때 기존에 값이 존재하는 경우 초기화 하지 않도록 함
		// Write the state to the ledger
		err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	Bvalbytes, err := stub.GetState(B)
	if Bvalbytes == nil {
		err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	Cvalbytes, err := stub.GetState(C)
	if Cvalbytes == nil {
		err = stub.PutState(C, []byte(strconv.Itoa(Cval)))
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	return shim.Success(nil)

}

// Invoke is somthing ... Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### kopoVote Invoke ###########")

	function, args := stub.GetFunctionAndParameters()

	if function == "query" {
		// queries an entity state
		return t.query(stub, args)
	}
	if function == "vote" {
		// Deletes an entity from its state
		return t.vote(stub, args)
	}

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}

func (t *SimpleChaincode) vote(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// must be an invoke
	var A string // Entities
	var Aval int // Asset holdings
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A = args[0]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))
	logger.Infof("Before> Aval = %d\n", Aval)

	Aval = Aval + 1
	logger.Infof("After> Aval = %d\n", Aval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var A, B, C string // Entities
	var err error

	// if len(args) != 1 {
	// 	return shim.Error("Incorrect number of arguments")
	// }

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil qty for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Qty\":\"" + string(Avalbytes) + "\"}"
	logger.Infof("Query Response:%s\n", jsonResp)

	B = args[1]

	// Get the state from the ledger
	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + B + "\"}"
		return shim.Error(jsonResp)
	}

	if Bvalbytes == nil {
		jsonResp := "{\"Error\":\"Nil qty for " + B + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp = "{\"Name\":\"" + B + "\",\"Qty\":\"" + string(Bvalbytes) + "\"}"
	logger.Infof("Query Response:%s\n", jsonResp)

	C = args[2]

	// Get the state from the ledger
	Cvalbytes, err := stub.GetState(C)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + C + "\"}"
		return shim.Error(jsonResp)
	}

	if Cvalbytes == nil {
		jsonResp := "{\"Error\":\"Nil qty for " + C + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp = "{\"Name\":\"" + C + "\",\"Qty\":\"" + string(Cvalbytes) + "\"}"
	logger.Infof("Query Response:%s\n", jsonResp)

	retStr := []string{string(Avalbytes), string(Bvalbytes), string(Cvalbytes)}
	//return shim.Success(Avalbytes)
	//var str = []string{"str1","str2"}
	var x = []byte{}
	x = append(x, '{')
	for i := 0; i < len(retStr); i++ {
		b := []byte(retStr[i])
		for j := 0; j < len(b); j++ {
			x = append(x, b[j])
		}
		x = append(x, ',')
	}
	x = append(x, '}')
	return shim.Success(x)
}
