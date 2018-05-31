package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	"testing"
	"fmt"
)

func TestMain(t *testing.T) {
	scc := new(PetChaincode);
	stub := shim.NewMockStub("pet", scc);

	checkTask(stub);
	checkPet(stub);
}

func checkTask(stub *shim.MockStub) {
	ci := stub.MockInvoke("invoke", [][]byte{[]byte("createTask"), []byte("TestTask"), []byte("Test Task Describtion"), []byte("awardId"), []byte("1000")});
	fmt.Println(string(ci.Payload));

	var dat map[string]string;
	json.Unmarshal(ci.Payload, &dat);
	fmt.Printf("%s %s\n", dat["task_name"], dat["id"]);

	qtask := stub.MockInvoke("invoke", [][]byte{[]byte("queryTask"), []byte(dat["task_name"]), []byte(dat["id"])});
	fmt.Println(string(qtask.Payload));

	tTask := stub.MockInvoke("invoke", [][]byte{[]byte("takeTask"), []byte(dat["task_name"]), []byte(dat["id"]), []byte("test1")});
	fmt.Println(string(tTask.Payload));

	qtask = stub.MockInvoke("invoke", [][]byte{[]byte("queryTask"), []byte(dat["task_name"]), []byte(dat["id"])});
	fmt.Println(string(qtask.Payload));

	player := stub.MockInvoke("invoke", [][]byte{[]byte("completeTask"), []byte(dat["task_name"]), []byte(dat["id"]), []byte("id")});
	fmt.Println(player.Message);

	player = stub.MockInvoke("invoke", [][]byte{[]byte("completeTask"), []byte(dat["task_name"]), []byte(dat["id"]), []byte("test1")});
	fmt.Println(string(player.Payload));
}

func checkPet(stub *shim.MockStub) {
	pet := stub.MockInvoke("invoke", [][]byte{[]byte("createPet"), []byte("TestPet"), []byte("Test Pet Describtion"), []byte("10")});
	fmt.Println(string(pet.Payload));
}
