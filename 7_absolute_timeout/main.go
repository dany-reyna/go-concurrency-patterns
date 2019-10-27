package main

import (
	"fmt"
	"patterns/common"
	"time"
)

func main() {
	c := common.Boring("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}
