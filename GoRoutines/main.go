package main

import (
	"fmt"
	"time"
)

func main() {

	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, 10)
	go add(s[:len(s)/2], c)
	go add(s[len(s)/2:], c)

	x := <-c

	fmt.Println(x)

	y := <-c
	fmt.Println(y)

	go echo(c)

	for i := range c {
		fmt.Println(i)
	}
}

func add(s []int, c chan int) {
	aux := 0

	for _, item := range s {
		time.Sleep(1000 * time.Millisecond)
		aux += item
	}

	c <- aux
}

func echo(c chan int) {
	for {
		time.Sleep(1000 * time.Millisecond)
		c <- 1
	}
}
