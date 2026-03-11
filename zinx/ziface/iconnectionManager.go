package ziface

/*
	连接管理模块接口
*/

type IConnectionManager interface {
	// 添加连接
	Add(conn IConnection)
	// 删除连接
	Remove(conn IConnection)
	// 获取连接
	Get(connID uint32) (IConnection, error)
	// 获取连接数量
	Len() int
	// 清空所有连接
	Clear()
}
