package main

import (
	"fmt"
	"reflect"
)

type TestStruct2 struct {
	Name	string
	Value	uint32
}

type Feed struct {
	Power	uint16
}
	
type TestStruct1 struct {
	Name	string
	Value	uint32
	Obj		TestStruct2
	Award	Feed
	Type	reflect.Type
	TypeValue	reflect.Value
}

func main() {
	fmt.Println("Hello World!");

	test := TestStruct2{
		Name: "test1",
		Value: 2,
	};

	test1 := TestStruct1{
		Name: "test2",
		Value: 3,
		Type: reflect.TypeOf(test),
		TypeValue: reflect.ValueOf(test),
	}

	fmt.Printf("%s: %d\n", test.Name, test.Value);
	//fmt.Printf("%s: %d {%s: %d}\n", test1.Name, test1.Value, test1.Obj.Name, test1.Obj.Value);
	fmt.Printf("%s %s %d", test1.Type.Field(1).Name, test1.TypeValue.Field(1).Type(), test1.TypeValue.Field(1).Interface());
}