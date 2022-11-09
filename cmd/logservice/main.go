package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stdLog "log"
)

func main() {
	// 项目中需要在配置文件、或环境变量中读取的
	log.Run("./distributed.log")

	host, port := "localhost", "4000"

	reg := registry.Registration{
		ServiceName: "Log service",
		ServiceURL:  fmt.Sprintf("http://%s:%v", host, port),
	}

	// 启动服务
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		reg,
		log.RegisterHandlers,
	)
	if err != nil {
		// 因为自定义的日志库没有启动成功，使用标准库的日志
		stdLog.Fatalln(err)
	}

	// 等待接收到退出信号，startService中两个协程出错或主动退出时会发送信号，调用 cancel() 时会发送信息
	<-ctx.Done()

	// 卡在上面那行，如果接收到信息，则往下执行
	fmt.Println("Shutting down log service.")
}
