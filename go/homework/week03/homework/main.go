package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func main()  {
	//获取group和传递上下文
	g,ctx := errgroup.WithContext(context.Background())
	//开启服务
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(w, "hello, server is started")
	})
	//关闭服务
	serverOut := make(chan struct{})//无缓冲channel，接收server退出信号
	mux.HandleFunc("/shutdown",func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(w, "hello, server is stop")
		serverOut<- struct{}{}
	})
	server := http.Server{
		Handler: mux,
		Addr: ":8080",
	}
	//1.start http server
	g.Go(func() error {
		return server.ListenAndServe()
	})
	//2.stop http server
	g.Go(func() error {
		select {
		case <-ctx.Done()://阻塞
			return ctx.Err()
		case <-serverOut:
			fmt.Println("server will stop")
			_ =server.Shutdown(ctx)
		}
		return nil
	})
	_,cancel := context.WithCancel(ctx)
	//用channel接收信号量
	sig := make(chan os.Signal,1)
	signal.Notify(sig)
	//3.监听信号量
	g.Go(func() error {
		for  {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-sig:
				cancel()
			}
		}
	})
	//等待所有goroutine退出
	if err :=g.Wait();err != nil{
		fmt.Println("goroutine err")
	}
	fmt.Println("all goroutine done")
}