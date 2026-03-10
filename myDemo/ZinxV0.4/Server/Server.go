package main

import (
	"fmt"

	"github.com/zfz-725/zinx/ziface"
	"github.com/zfz-725/zinx/znet"
)

// 基于Zinx框架来开发的，服务器端应用程序

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call PingRouter PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Printf("PreHandle Write failed, err: %v\n", err)
	}
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping... ping... ping\n"))
	if err != nil {
		fmt.Printf("Handle Write failed, err: %v\n", err)
	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call PingRouter PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Printf("PostHandle Write failed, err: %v\n", err)
	}
}

func main() {
	// 创建一个Zinx服务器句柄
	s := znet.NewServer()

	// 添加一个自定义路由
	s.AddRouter(&PingRouter{})

	// 启动服务器
	s.Serve()
}
