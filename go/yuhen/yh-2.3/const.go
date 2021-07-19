package main

import "fmt"

func main()  {
	const (
		_ = iota
		KB = 1<<10*iota
		MB
		GB
	)
	fmt.Println(KB,MB,GB)
	const (
		a = iota
		b
		c = 100
		d
		e = iota
		f

	)
	//
	fmt.Println(a,b,c,d,e,f)
}
