package main

import (
	"fmt"
)

func main() {
	res := generatNum(1, 2)
	vals := addVals(res)
	for vl := range vals {
		fmt.Println(vl)
	}

}

func generatNum(nums ...int) <-chan int {
	sender := make(chan int)
	go func() {
		for _, n := range nums {
			sender <- n
		}
		close(sender)
	}()
	return sender
}

func addVals(n <-chan int) <-chan int {
	rec := make(chan int)
	go func() {
		for v := range n {
			rec <- v + v
		}
		close(rec)
	}()
	return rec
}
