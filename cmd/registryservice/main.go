package main

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// https://nanxiao.me/golang-http-handle-handlefunc/
	// 注意 http.Handle 与 http.Handlefunc的区别
	http.Handle("/services", &registry.RegistrationService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		// 程序报错时自动取消
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		// 手动按键取消
		fmt.Println("Registry service started. Press any key to stop.")

		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	// 等待信号
	<-ctx.Done()

	fmt.Println("shutting down registry service")
}
