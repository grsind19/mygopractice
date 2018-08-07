package main

import (
	"fmt"
	"time"
)

func PrintCount(c chan int) {
	num := 0
	for num >= 0 {
		num := <-c
		fmt.Println(num, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	go PrintCount(c)
	for _, v := range a {
		c <- v
	}
	time.Sleep(time.Millisecond * 5)
	fmt.Println("ENd of main")
}
