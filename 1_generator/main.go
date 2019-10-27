package main

import (
	"fmt"
	"patterns/common"
)

func main() {
	c := common.Boring("generator!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}
