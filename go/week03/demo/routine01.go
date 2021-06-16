package main

import (
	"context"
	"fmt"
	"net/http"
)

func main()  {
	done := make(chan error,2)
	stop := make(chan struct{})
	go func() {
		done <- serverDebug(stop)
	}()
	go func() {
		done <- serverApp(stop)
	}()
	var stopped bool
	for i := 0; i < cap(stop) ; i++ {
		if err:= <-done;err !=nil{
			fmt.Printf("error %v\n",err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
func serverApp(stop <-chan struct{}) error  {
	return server("127.0.0.1：80001",http.DefaultServeMux,stop)
}
func serverDebug(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(w,"Hello,QCon!")
	})
	return server("127.0.0.1：80002",mux,stop)
}
func server(addr string,handler http.Handler,stop <-chan struct{}) error  {
	s := http.Server{
		Addr : addr,
		Handler: handler,
	}
	go func() {
		<-stop	//wait for stop signal
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}
