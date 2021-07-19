package main

import "fmt"

func main(){
	//数组
	var arr [2]int
	fmt.Printf("var初始化数组,%p\n",&arr)
	arr1 := new([2]int)
	fmt.Printf("var初始化数组指针,%p\n",arr1)
	//切片
	var sli []int
	fmt.Printf("var初始化切片,%p\n",sli)
	sli = append(sli,1)
	fmt.Printf("var赋值切片,%p\n",sli)
	makeSli := make([]int,0)
	fmt.Printf("make切片,%p\n",makeSli)
}
