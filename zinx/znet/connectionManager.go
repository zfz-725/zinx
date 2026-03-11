package znet

import (
	"errors"
	"fmt"
	"sync"

	"github.com/zfz-725/zinx/ziface"
)

/*
	连接管理模块
*/

type ConnectionManager struct {
	// 存储所有连接的map
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

// 创建当前连接的方法
func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

// 添加连接
func (cm *ConnectionManager) Add(conn ziface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	cm.connections[conn.GetConnID()] = conn
	fmt.Println("Add connection, ConnID:", conn.GetConnID())
}

// 删除连接
func (cm *ConnectionManager) Remove(conn ziface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	delete(cm.connections, conn.GetConnID())
	fmt.Println("Remove connection, ConnID:", conn.GetConnID())
}

// 获取连接
func (cm *ConnectionManager) Get(connID uint32) (ziface.IConnection, error) {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()
	conn, ok := cm.connections[connID]
	if !ok {
		return nil, errors.New("connection not found")
	}
	fmt.Println("Get connection, ConnID:", connID)
	return conn, nil
}

// 获取连接数量
func (cm *ConnectionManager) Len() int {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()
	fmt.Println("Connection count:", len(cm.connections))
	return len(cm.connections)
}

// 清空所有连接
func (cm *ConnectionManager) Clear() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	for connID, conn := range cm.connections {
		// 关闭连接
		conn.Stop()
		fmt.Println("Close connection, ConnID:", connID)
	}
	cm.connections = make(map[uint32]ziface.IConnection)
	fmt.Println("Clear all connections")
}
