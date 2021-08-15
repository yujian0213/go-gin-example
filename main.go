package main

import (
	"fmt"
	"github.com/fvbock/endless"
	setting "go-gin-example/pkg"
	"go-gin-example/routers"
	"log"
	"syscall"
)

func main()  {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1<<20
	endPoint := fmt.Sprintf(":%d",setting.HTTPPort)
	server := endless.NewServer(endPoint,routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Println("Actual pid is %d",syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err:%s",err)
	}

	/*r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"msg":"test",
		})
	})*/


	/*
	router := routers.InitRouter()
		s := &http.Server{
		Addr: fmt.Sprintf(":%d",setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	quit := make(chan os.Signal)
	signal.Notify(quit,os.Interrupt)
	<-quit
	log.Println("shutdown server ...")
	ctx,cancle := context.WithTimeout(context.Background(),5*time.Second)
	defer cancle()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln("server shutdown:",err)
	}
	log.Println("server exiting")
	*/

	//s.ListenAndServe()
}