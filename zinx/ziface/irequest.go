package ziface

/*
	IRequest接口：
	实际上是把客户端请求的连接信息和请求数据包装到了一个Request中
*/

type IRequest interface {
	// 获取当前连接
	GetConnection() IConnection
	// 获取请求数据
	GetData() []byte
	// 获取请求消息ID
	GetMsgID() uint32
}
