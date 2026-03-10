package znet

import (
	"fmt"

	"github.com/zfz-725/zinx/ziface"
)

/*
	消息处理模块的实现
*/

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() ziface.IMsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
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
