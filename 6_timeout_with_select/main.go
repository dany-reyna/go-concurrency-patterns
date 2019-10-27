package main

import (
	"fmt"
	"patterns/common"
	"time"
)

func main() {
	c := common.Boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow.")
			return
		}
	}
}
