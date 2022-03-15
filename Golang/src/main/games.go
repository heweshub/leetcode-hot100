package main

import (
	"fmt"
	"time"
)

var pool = 100

func odd(c chan int) {
	for i := 1; i <= pool; i++ {
		c <- i
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}
func even(c chan int) {
	for i := 1; i <= pool; i++ {
		<-c
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	c := make(chan int)
	go odd(c)
	go even(c)
	time.Sleep(time.Second * 1)
}
