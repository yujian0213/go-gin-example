package main

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

func main()  {

}
//交替打印数字和字母
func AlternatePrint(){
	letter,number := make(chan bool),make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		i := 1
		for  {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			default:
				break
			}
		}
	}()
	wait.Add(1)
	go func() {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for  {
			select {
			case <-letter:
				if i >= strings.Count(str,"")-1 {
					wait.Done()
					return
				}
				fmt.Print(str[i:i+1])
				i++
				if i >= strings.Count(str,"") {
					i = 0
				}
				fmt.Print(str[i:i+1])
				i++
				number <- true
			default:
				break
			}
		}


	}()
	number <- true
	wait.Wait()
}
//翻转字符串
func ReverString(s string) (string,bool){
	str := []rune(s)
	l := len(str)
	if l >5000{
		return s,false
	}
	for i := 0; i <=l/2 ; i++ {
		str[i],str[l-i-1] = str[l-i-1],str[i]
	}
	return string(str),true
}
func AlternatePrint1(){
	var (
		i int32 = 1
		max int32 = 10
	)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i <= max {
			if i & 1 == 0 {
				fmt.Printf("g1:%d\n", i)
				atomic.AddInt32(&i, 1)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i <= max {
			if i & 1 != 0 {
				fmt.Printf("g2:%d\n", i)
				atomic.AddInt32(&i, 1)
			}
		}
	}()
	wg.Wait()
}
