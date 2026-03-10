package znet

import (
	"errors"
	"fmt"
	"net"

	"github.com/zfz-725/zinx/ziface"
)

/*
	连接模式
*/

type Connection struct {
	// 当前连接的ID
	ConnID uint32
	// 当前连接的socket TCP套接字
	Conn *net.TCPConn
	// 当前连接的状态
	isClosed bool
	// 告知当前连接已经退出/停止 channel
	ExitChan chan bool
	// 该连接处理的方法Router
	Router ziface.IRouter
}

// 初始化连接模块的方法
func NewConnection(connID uint32, conn *net.TCPConn, router ziface.IRouter) *Connection {
	return &Connection{
		ConnID:   connID,
		Conn:     conn,
		isClosed: false,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Connection StartReader... ConnID:", c.ConnID)
	// 启动从当前连接的读数据的业务
	go func() {
		defer fmt.Println("StartReader goroutine exit... ConnID:", c.ConnID, "RemoteAddr:", c.Conn.RemoteAddr())
		defer c.Stop()
		for {
			// 读取客户端数据到buf中，最大512字节
			buf := make([]byte, 512)
			cnt, err := c.Conn.Read(buf)
			if err != nil {
				fmt.Printf("Read failed, err: %v\n", err)
				continue
			}

			// 得到当前conn数据的Request请求数据
			req := &Request{
				conn: c,
				data: buf[:cnt],
			}

			// 执行注册的路由方法
			go func(request ziface.IRequest) {
				c.Router.PreHandle(request)
				c.Router.Handle(request)
				c.Router.PostHandle(request)
			}(req)
		}
	}()
}

// 启动连接
func (c *Connection) Start() {
	fmt.Println("Connection Start... ConnID:", c.ConnID)
	// 启动从当前连接的读数据的业务
	go c.StartReader()

	// TODO 启动从当前连接写数据的业务
}

// 停止连接
func (c *Connection) Stop() {
	fmt.Println("Connection Stop... ConnID:", c.ConnID)
	if c.isClosed {
		return
	}
	c.isClosed = true
	// 关闭socket连接
	c.Conn.Close()
	// 关闭管道
	close(c.ExitChan)
}

// 获取当前连接的绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取连接的远程节点地址
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据
func (c *Connection) Send(data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send data")
	}
	return nil
}
