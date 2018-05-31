package main

import (
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"./utils"
	"./player"
	"./pet"
//	"./feed"
	"./task"
)

type PetChaincode struct {
}

func main() {
	err := shim.Start(new(PetChaincode));
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s\n", err);
	}
}

func (t *PetChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil);
}

func (t *PetChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters();
	fmt.Printf("Function: %s, args: %s\n", function, args);

	switch {
		case function == "createTask" :
			return createTask(stub, args);
		case function == "queryTask" :
			return queryTask(stub, args);
		case function == "takeTask" :
			return takeTask(stub, args);
		case function == "completeTask" :
			return completeTask(stub, args);
		case function == "createPet" :
			return createPet(stub, args);
		default:
			return shim.Success(nil);
	}
	//createPet(args);
}

func getPlayer(stub shim.ChaincodeStubInterface, name string) (*player.Player, error) {
	//creatorByte, err := stub.GetCreator();
	//user_name := utils.GetCertificateOwner( creatorByte );

	user_name := name;
	playerAsBytes, err := stub.GetState(user_name);
	
	if err != nil {
		return nil, fmt.Errorf("Failed get player [%s]", user_name);
	} else if playerAsBytes == nil {
		new_player := player.NewPlayer(user_name);

		return new_player, nil;
	} else {
		new_player := &player.Player{};
		err := json.Unmarshal(playerAsBytes, &new_player);
		if err != nil {
			return nil, fmt.Errorf(err.Error());
		}

		return new_player, nil;
	}
}

func _createPet(args []string) {
	new_pet := pet.NewPet("testPet1", "Test Pet One", 10);
	jsonAsBytes, _ := json.Marshal(new_pet);
	fmt.Printf("Pet: %s %d\n", new_pet.Name, new_pet.Quality);
	fmt.Printf("Pet Json: %s\n", string(jsonAsBytes) );
}

func createTask(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		fmt.Println("No enough args");
		return shim.Error("No enough args, 4 require");
	}

	name := args[0];
	des := args[1];
	awardObjId := args[2];
	awardType, err := strconv.ParseInt((args[3]), 10, 16);
	if err != nil {
		return shim.Error( "Conver AwardType Error" );
	}

	task := task.NewTask(name, des, awardObjId, uint16(awardType));
	taskJsonBytes, _ := json.Marshal(task);
	
	stub.PutState(task.Name + task.Id, taskJsonBytes);

	return shim.Success(taskJsonBytes);
}

func queryTask(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Need 2 args");
	}

	name := args[0];
	id := args[1];

	taskJsonBytes, err := stub.GetState(name + id);
	if err != nil {
		return shim.Error("Task not found");
	}

	return shim.Success(taskJsonBytes);
}

func takeTask(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//if len(args) != 2 {
	if len(args) != 3 {
		return shim.Error("Need 2 args");
	}

	name := args[0];
	id := args[1];

	taskJsonBytes, err := stub.GetState(name + id);
	if err != nil {
		return shim.Error("Task not found");
	}

	task := task.Task{};
	json.Unmarshal(taskJsonBytes, &task);

	playerByte, err := stub.GetCreator();
	player := utils.GetCertificateOwner( playerByte );
	player = args[2];

	task.Owner = player;
	taskJsonBytes, err = json.Marshal(task);

	stub.PutState(name + id, taskJsonBytes);

	return shim.Success(taskJsonBytes);
}

func completeTask(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	//if len(args) < 2 {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2");
	}

	taskName := args[0];
	taskId := args[1];

	taskBytes, err := stub.GetState(taskName + taskId);
	if err != nil {
		return shim.Error("Faild to get task: " + err.Error());
	} else if taskBytes == nil {
		return shim.Error("Task does not exits");
	}

	cur_task := task.Task{};
	err = json.Unmarshal(taskBytes, &cur_task);
	if err != nil {
		return shim.Error(err.Error());
	}

	owner := args[2];
	player, err := getPlayer(stub, owner);
	if err != nil {
		fmt.Printf("Faild to get user [%s]", err.Error());
		return shim.Error("Faild to get user:" + err.Error());
	}

	if cur_task.Owner != owner {
		return shim.Error("Not task owner!");
	}

	cur_task.State = task.COMPLETE_TASK;
	aId, aObjId := cur_task.GetAward();
	if aId == task.AWARD_TYPE_PET {
		player.Pets = append(player.Pets, aObjId);
	} else if aId == task.AWARD_TYPE_FOOD {
		player.Feeds = append(player.Feeds, aObjId);
	}

	playerJsonByte, _ := json.Marshal(player);

	return shim.Success(playerJsonByte);
}

func createFeed(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	if len(args) != 3 {
		fmt.Println("Three args require");
		return shim.Error("Three args require");
	}

	name := args[0];
	des := args[1];
	quality, err := strconv.ParseInt(args[2], 10, 32);
	if err != nil {
		fmt.Println("Parse Int Error!");
		return shim.Error("Parse Int Error!");
	}

	new_pet := pet.NewPet(name, des, uint32(quality));

	petJsonBytes, err := json.Marshal(new_pet);
	
	stub.PutState(new_pet.Name + new_pet.Id, petJsonBytes);

	return shim.Success(petJsonBytes);
}

func transferFeed(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	return shim.Success(nil);
}

func feedFeed(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	return shim.Success(nil);
}

func queryFeed(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	return shim.Success(nil);
}

func createPet(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	if len(args) != 3 {
		fmt.Println("Three args require");
		return shim.Error("Three args require");
	}

	name := args[0];
	des := args[1];
	quality, err := strconv.ParseInt(args[2], 10, 32);
	if err != nil {
		fmt.Println("Parse Int Error!");
		return shim.Error("Parse Int Error!");
	}

	new_pet := pet.NewPet(name, des, uint32(quality));

	petJsonBytes, err := json.Marshal(new_pet);
	
	stub.PutState(new_pet.Name + new_pet.Id, petJsonBytes);

	return shim.Success(petJsonBytes);
}

func transferPet(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	return shim.Success(nil);
}

func feedPet(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	return shim.Success(nil);
}

func queryPet(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	return shim.Success(nil);
}
