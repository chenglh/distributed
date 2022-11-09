package service

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

// Start 集中启动项目中的所有服务(注册、启动、关闭service)
/*
 host: 地址
 port: 端口号
 serviceName:  注册的服务对象
 registerHandlersFunc: 注册方法
*/
func Start(ctx context.Context, host, port string, reg registry.Registration, registerHandlersFunc func()) (context.Context, error) {
	// 注册函数调用
	registerHandlersFunc()

	// 启动/关闭服务
	ctx = startService(ctx, reg.ServiceName, host, port)

	// 在注册中心中维护服务
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	// 上下文，取消函数
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = host + ":" + port

	// 该协程为监听http服务，停止服务的时候cancel
	go func() {
		// 如果启动出错，返回出错信息，这里会打印出错信息；如果不出错，则进入for阻塞协程
		log.Println(srv.ListenAndServe())

		fmt.Println("程序异常信号")

		// 注销服务
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			fmt.Println(err)
		}

		// 如果在启动服务器的时候出错了，会发送信号，调用cancel()函数，执行取消操作
		cancel()
	}()

	// 该协程为监听手动停止服务的信号
	go func() {
		fmt.Printf("%v started. Press any key to stop.\n", serviceName)

		//按任意键中止服务
		var s string
		fmt.Scan(&s)

		fmt.Println("程序捕捉到中止信号")

		// 注销服务
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			fmt.Println(err)
		}
		srv.Shutdown(ctx)

		//或者用户中止也会调用cancel()关闭当前服务
		cancel()
	}()

	return ctx
}
