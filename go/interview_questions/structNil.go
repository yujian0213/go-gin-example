package main

import "fmt"

type Peopl interface {
	Show()
}

type Student struct{}

func (stu Student) Show() {

}

func live() Peopl {
	var stu Student
	return stu
}

func main() {
	l := live()
	fmt.Printf("l:%v\n",l)
	if l == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
