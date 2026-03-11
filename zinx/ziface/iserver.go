package ziface

// 定义一个服务器接口
type IServer interface {
	// 启动服务器
	Start()
	// 停止服务器
	Stop()
	// 运行服务器
	Serve()
	// 添加路由
	AddRouter(msgID uint32, router IRouter)
	// 获取当前Server的连接管理器
	GetConnManager() IConnectionManager
	// 设置服务器创建连接之后的钩子函数
	SetOnConnStart(hookFunc func(conn IConnection))
	// 设置服务器关闭连接之前的钩子函数
	SetOnConnStop(hookFunc func(conn IConnection))
	// 调用服务器创建连接之后的钩子函数
	CallOnConnStart(conn IConnection)
	// 调用服务器关闭连接之前的钩子函数
	CallOnConnStop(conn IConnection)
}
