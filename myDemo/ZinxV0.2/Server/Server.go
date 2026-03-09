package main

import (
	"github.com/zfz-725/zinx/znet"
)

// 基于Zinx框架来开发的，服务器端应用程序

func main() {
	// 创建一个Zinx服务器句柄
	s := znet.NewServer("[zinx V0.2]")

	// 启动服务器
	s.Serve()
}
