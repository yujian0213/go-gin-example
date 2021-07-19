package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Param map[string]interface{}
type Show struct {
	Param
}
//弟1题
//new 关键字无法初始化 Show 结构体中的 Param 属性，所以直接 对 s.Param 操作会出错
func NewParam() {
	s := new(Show)
	s.Param["RMB"] = 10000
}

type student struct { Name string
}
//弟2题
//golang中有规定， switch type 的 case T1 ，类型列表只有一个，那么 v := m.(type) 中的 v 的类型就是T1类型。
//如果是 case T1, T2 ，类型列表中有多个，那 v 的类型还是多对应接口的类型，也就 是 m 的类型。
//所以这里 msg 的类型还是 interface{} ，所以他没有 Name 这个字段，编译阶段就会 报错。
/*func zhoujielun(v interface{}) {
	switch msg := v.(type) { case *student, student:
		msg.Name }
}*/
//弟3题
//按照 golang 的语法，小写开头的方法、属性或 struct 是私有的，同样，在 json 解 码或转码的时候也无法上线私有属性的转换。
//题目中是无法正常得到 People 的 name 值的。而且，私有属性 name 也不应该加
//json 的标签
type People struct {
	Name string `json:"name"`
}
func Peopler() {
	js := `{ "name":"11"
}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return }
	fmt.Println("people: ", p)
}
//弟4题
func (p *People) String() string {
	return fmt.Sprintf("print:%v",p)
}

func Stringer()  {
	p := &People{}
	_ = p.String()
}
//第5题
func Five() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)

}