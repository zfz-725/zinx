package ziface

import "net"

// 定义连接模块的抽象层
type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 获取当前连接的绑定socket conn
	GetTCPConnection() *net.TCPConn
	// 获取连接ID
	GetConnID() uint32
	// 获取连接的远程节点地址
	GetRemoteAddr() net.Addr
	// 发送数据
	SendMsg(msgID uint32, data []byte) error
	// 设置连接属性
	SetProperty(key string, value interface{})
	// 获取连接属性
	GetProperty(key string) (interface{}, error)
	// 删除连接属性
	DelProperty(key string)
}

// 定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
