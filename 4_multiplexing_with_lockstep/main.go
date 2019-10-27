package main

import (
	"fmt"
	"patterns/common"
)

func fanIn(input1, input2 <-chan common.Message) <-chan common.Message {
	c := make(chan common.Message)

	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func main() {
	c := fanIn(common.BoringLock("a"), common.BoringLock("b"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.Str)
		msg2 := <-c
		fmt.Println(msg2.Str)
		msg1.Wait <- true
		msg2.Wait <- true
	}
	fmt.Println("You're both boring; I'm leaving.")
}
