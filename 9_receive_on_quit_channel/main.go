package main

import (
	"fmt"
	"math/rand"
	"patterns/common"
)

func main() {
	quit := make(chan string)
	c := common.BoringQuitReceive("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n", <-quit)
}
