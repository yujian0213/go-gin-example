package main

import "fmt"

var a1 string

func f1(c chan int) {
	a1 = "hello, world"
	c <- 0
}

func main() {
	//c := make(chan int, 1)
	//go f1(c)
	//<-c
	//print(a1)
	a := 1
	b := 2
	defer fmt.Println(a, b)
	a = 0
	defer fmt.Println(a, b)
	b = 1

}
