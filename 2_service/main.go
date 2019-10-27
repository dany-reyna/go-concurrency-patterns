package main

import (
	"fmt"
	"patterns/common"
)

func main() {
	joe := common.Boring("Joe")
	ann := common.Boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving.")
}
