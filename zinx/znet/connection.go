package znet

import (
	"errors"
	"fmt"
	"net"

	"github.com/zfz-725/zinx/utils"
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
	// 告知当前连接已经退出/停止 channel，由reader告知writer
	ExitChan chan bool
	// 该连接处理的方法MsgHandler
	MsgHandler ziface.IMsgHandler
	// 无缓冲的管道，用于读、写
	MsgChan chan []byte
}

// 初始化连接模块的方法
func NewConnection(connID uint32, conn *net.TCPConn, msgHandler ziface.IMsgHandler) *Connection {
	return &Connection{
		ConnID:     connID,
		Conn:       conn,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		MsgHandler: msgHandler,
		MsgChan:    make(chan []byte),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Connection StartReader... ConnID:", c.ConnID)
	// 启动从当前连接的读数据的业务
	go func() {
		defer fmt.Println("StartReader goroutine exit... ConnID:", c.ConnID, "RemoteAddr:", c.Conn.RemoteAddr())
		defer c.Stop()
		for {

			// 创建一个拆包器
			dp := NewDataPack()

			// 读取客户端的Msg Head
			msgHead := make([]byte, dp.GetHeadLen())
			_, err := c.Conn.Read(msgHead)
			if err != nil {
				fmt.Printf("Read failed, err: %v\n", err)
				break
			}
			// 拆包，得到msgID 和 dataLen 放在msg消息中
			msg, err := dp.Unpack(msgHead)
			if err != nil {
				fmt.Printf("Unpack failed, err: %v\n", err)
				break
			}
			// 根据dataLen，再次读取data
			if msg.GetMsgLen() > 0 {
				data := make([]byte, msg.GetMsgLen())
				_, err = c.Conn.Read(data)
				if err != nil {
					fmt.Printf("Read failed, err: %v\n", err)
					break
				}
				msg.SetData(data)
			}

			// 得到当前conn数据的Request请求数据
			req := &Request{
				conn: c,
				msg:  msg,
			}

			if utils.GlobalObject.WorkerPoolSize > 0 {
				// 开启工作池处理请求
				c.MsgHandler.SendMsgToTaskQueue(req)
			} else {
				// 执行注册的路由方法
				go c.MsgHandler.DoMsgHandler(req)
			}
		}
	}()
}

// 写消息goroutine，专门给客户端发送消息
func (c *Connection) StartWriter() {
	fmt.Println("Connection StartWriter... ConnID:", c.ConnID)
	// 启动从当前连接的写数据的业务
	defer fmt.Println("StartWriter goroutine exit... ConnID:", c.ConnID, "RemoteAddr:", c.Conn.RemoteAddr())

	for {
		select {
		case data := <-c.MsgChan:
			_, err := c.Conn.Write(data)
			if err != nil {
				fmt.Printf("Write failed, err: %v\n", err)
				continue
			}
		case <-c.ExitChan:
			return
		}
	}
}

// 发送数据
func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send data")
	}
	dp := NewDataPack()
	// 先将msg进行封包
	msg, err := dp.Pack(NewMessage(msgID, data))
	if err != nil {
		fmt.Printf("Pack failed, err: %v\n", err)
		return err
	}
	// 写入数据
	c.MsgChan <- msg
	return nil
}

// 启动连接
func (c *Connection) Start() {
	fmt.Println("Connection Start... ConnID:", c.ConnID)
	// 启动从当前连接的读数据的业务
	go c.StartReader()

	// 启动从当前连接写数据的业务
	go c.StartWriter()
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
	// 通知writer goroutine退出
	c.ExitChan <- true
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
