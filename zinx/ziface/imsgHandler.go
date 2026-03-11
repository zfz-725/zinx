package ziface

/*
	消息管理抽象层
*/

type IMsgHandler interface {
	// 处理消息
	DoMsgHandler(request IRequest)
	// 注册路由
	AddRouter(msgID uint32, router IRouter)
	// 启动一个Worker工作池
	StartWorkerPool()
	// 将消息发送到任务队列
	SendMsgToTaskQueue(request IRequest)
}
