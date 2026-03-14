package ziface

// IServer 定义一个服务器接口
type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Server 运行服务器
	Serve()

	// AddRouter 给当前的服务注册一个路由方法，供客户端的连接处理使用
	AddRouter(msgID uint32, router IRouter)

	// GetConnMgr 获取当前 server 的连接管理器
	GetConnMgr() IConnManager

	// SetOnConnStart 注册 OnConnStart 钩子函数
	SetOnConnStart(func(conn IConnection))
	// CallOnConnStart 调用 CallOnConnStart 钩子函数
	CallOnConnStart(conn IConnection)
	// SetOnConnStop 注册 SetOnConnStop 钩子函数
	SetOnConnStop(func(conn IConnection))
	// CallOnConnStop 调用 CallOnConnStop 钩子函数
	CallOnConnStop(conn IConnection)
}
