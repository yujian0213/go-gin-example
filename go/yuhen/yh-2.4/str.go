package main

import (
	"fmt"
	"strconv"
)

func main()  {
	strconv.Itoa(1)
	strconv.Atoi("1")
	fmt.Println( 32 << ((^uint64(0)) >> 63) )
	fmt.Println( ^5)
	//strconv.ParseInt()
}
