package main

import (
	"fmt"
	"time"
)

func main()  {
	//c := make( chan int, 3)
	go func() {
		for {
			fmt.Println("A")
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		i := 1
		for {
			i++
			fmt.Println("B")
			time.Sleep(1 * time.Second)
			//if i > 10 {
			//	c<- 1
			//}
		}
	}()
	//<-c
}

