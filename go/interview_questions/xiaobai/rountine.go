package main

import (
	"fmt"
	"sync"
)

func main() {
	  var wg sync.WaitGroup
	   do := make(chan int)
	   done := make(chan int)
	   for i := 0; i < 10; i++ {
		        wg.Add(1)

		        go func(i int) {
			            defer wg.Done()

			            select {
				            case <-do:
				                fmt.Printf("Work: %d\\n", i)
				            case <-done:
				                fmt.Printf("Quit: %d\\n", i)
				            }
			        }(i)
		    }

	    close(done)

	    wg.Wait()
	}

