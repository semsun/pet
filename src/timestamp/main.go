package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println;

	p("Calc Example!");
	p(time.Millisecond);
	p(time.Second);

	pass := time.Now().UTC().UnixNano();
	//pass := time.Now().Unix();
	p(pass);
	p(pass / 1000000);
	time.Sleep(12 * time.Millisecond);
	//time.Sleep(12 * time.Second);
	now := time.Now().UTC().UnixNano();
	//now := time.Now().Unix();
	p(now);
	p(now / 1000000);

	diff := (now - pass) / 1000000;
	p(diff);
}