package main

import (
	"fmt"
	"strconv"
)

//生产
func producer(data chan int)  {
	for i := 0; i < 4; i++ {
		data <- i	//发送数据
	}
	close(data)		//关闭通道
}
//消费
func consumer(data chan int,done chan bool)  {
	for  x := range data {
		println("recv",x)	//接收数据,直至通道关闭
	}
	done <- true				//将消费结束的消息写入channel
}
func main()  {
	s,_ := strconv.Atoi("1")
	println(s)
	data := make(chan int,0)
	done := make(chan bool)
	fmt.Printf("data: %v\n",data)
	fmt.Printf("done: %v\n",done)
	go producer(data)
	go consumer(data,done)
	<-done
}