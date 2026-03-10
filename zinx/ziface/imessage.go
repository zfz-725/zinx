package ziface

/*
	将请求的消息封装到Message中，定义抽象层
*/

type IMessage interface {
	// 获取消息ID
	GetMsgID() uint32
	// 获取消息数据长度
	GetMsgLen() uint32
	// 获取消息数据
	GetData() []byte

	// 设置消息ID
	SetMsgID(uint32)
	// 设置消息数据长度
	SetMsgLen(uint32)
	// 设置消息数据
	SetData([]byte)
}
