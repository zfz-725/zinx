package ziface

/*
	消息管理抽象层
*/

type IMsgHandler interface {
	// 处理消息
	DoMsgHandler(request IRequest)
	// 注册路由
	AddRouter(msgID uint32, router IRouter)
}
