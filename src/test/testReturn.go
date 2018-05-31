package main

import (
	"fmt"
)

func return1() (int, string) {
	return 1, "test";
}

func main() {
	vi, vs := return1();
	fmt.Printf("VInt:%d, VString:%s!\n", vi, vs);
}