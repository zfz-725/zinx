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
	err := request.GetConnection().SendMsg(0, []byte("before ping...\n"))
	if err != nil {
		fmt.Printf("PreHandle SendMsg failed, err: %v\n", err)
	}
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	// 先读取客户端数据，再回写
	msgID := request.GetMsgID()
	data := request.GetData()
	fmt.Printf("MsgID: %d, Data: %s\n", msgID, string(data))
	err := request.GetConnection().SendMsg(1, []byte("ping... ping... ping\n"))
	if err != nil {
		fmt.Printf("Handle SendMsg failed, err: %v\n", err)
	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call PingRouter PostHandle")
	err := request.GetConnection().SendMsg(2, []byte("after ping...\n"))
	if err != nil {
		fmt.Printf("PostHandle SendMsg failed, err: %v\n", err)
	}
}

// hello test 自定义路由
type HelloRouter struct {
	znet.BaseRouter
}

// Test PreHandle
func (this *HelloRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call HelloRouter PreHandle")
	err := request.GetConnection().SendMsg(0, []byte("before hello...\n"))
	if err != nil {
		fmt.Printf("PreHandle SendMsg failed, err: %v\n", err)
	}
}

// Test Handle
func (this *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloRouter Handle")
	// 先读取客户端数据，再回写
	msgID := request.GetMsgID()
	data := request.GetData()
	fmt.Printf("MsgID: %d, Data: %s\n", msgID, string(data))
	err := request.GetConnection().SendMsg(1, []byte("hello... hello... hello\n"))
	if err != nil {
		fmt.Printf("Handle SendMsg failed, err: %v\n", err)
	}
}

// Test PostHandle
func (this *HelloRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call HelloRouter PostHandle")
	err := request.GetConnection().SendMsg(2, []byte("after hello...\n"))
	if err != nil {
		fmt.Printf("PostHandle SendMsg failed, err: %v\n", err)
	}
}

func main() {
	// 创建一个Zinx服务器句柄
	s := znet.NewServer()

	// 添加一个自定义路由
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

	// 设置服务器创建连接之后的钩子函数
	s.SetOnConnStart(func(conn ziface.IConnection) {
		fmt.Printf("OnConnStart, connID: %d\n", conn.GetConnID())
		if err := conn.SendMsg(202, []byte("DoConnection BEGIN")); err != nil {
			fmt.Printf("OnConnStart SendMsg failed, err: %v\n", err)
		}
	})

	// 设置服务器关闭连接之前的钩子函数
	s.SetOnConnStop(func(conn ziface.IConnection) {
		fmt.Printf("OnConnStop, connID: %d\n", conn.GetConnID())
		if err := conn.SendMsg(203, []byte("DoConnection END")); err != nil {
			fmt.Printf("OnConnStop SendMsg failed, err: %v\n", err)
		}
	})

	// 启动服务器
	s.Serve()
}
