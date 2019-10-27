package main

import (
	"fmt"
	"math/rand"
	"patterns/common"
)

func main() {
	quit := make(chan bool)
	c := common.BoringQuit("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}
