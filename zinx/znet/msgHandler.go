package znet

import (
	"fmt"

	"github.com/zfz-725/zinx/utils"
	"github.com/zfz-725/zinx/ziface"
)

/*
	消息处理模块的实现
*/

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
	// 负责worker取任务的消息队列
	TaskQueue []chan ziface.IRequest
	// worker数量
	WorkerPoolSize uint32
}

func NewMsgHandler() ziface.IMsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

// 处理消息
func (mh *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	router, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Printf("Router not found, msgID: %d\n", request.GetMsgID())
		return
	}
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

// 注册路由
func (mh *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	if _, ok := mh.Apis[msgID]; ok {
		fmt.Printf("Router already exists, msgID: %d\n", msgID)
		return
	}
	mh.Apis[msgID] = router
	fmt.Printf("AddRouter success, msgID: %d, router: %v\n", msgID, router)
}

// 启动一个Worker工作池
func (mh *MsgHandler) StartWorkerPool() {
	// 遍历创建worker
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		// 启动worker
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskSize)
		// 启动worker
		go mh.startOneWorker(i, mh.TaskQueue[i])
	}
}

// 启动一个Worker工作流程
func (mh *MsgHandler) startOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID = ", workerID, "is started...")
	// 从任务队列中取任务
	for request := range taskQueue {
		// 调用DoMsgHandler处理消息
		mh.DoMsgHandler(request)
	}
}

// 将消息发送到任务队列
func (mh *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	// 轮询将消息发送到不同的worker的任务队列
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Printf("SendMsgToTaskQueue, workerID: %d, request: %v\n", workerID, request)
	mh.TaskQueue[workerID] <- request
}
