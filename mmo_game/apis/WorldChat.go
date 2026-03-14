package apis

import (
	"fmt"
	"github.com/zfz-725/mmo_game/core"
	"github.com/zfz-725/mmo_game/pb"
	"github.com/zfz-725/zinx/ziface"
	"github.com/zfz-725/zinx/znet"
	"google.golang.org/protobuf/proto"
)

type WorldChat struct {
	znet.BaseRouter
}

func (w *WorldChat) Handle(request ziface.IRequest) {
	// 解析聊天协议
	talk := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), talk)
	if err != nil {
		fmt.Println("Talk unmarshal err: ", err)
		return
	}
	// 获取玩家Pid
	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("get pid from conn property err: ", err)
		return
	}
	// 获取玩家
	player := core.WorldMgr.Players[pid.(int32)]
	// 发送聊天内容
	player.Talk(talk.Content)
}
