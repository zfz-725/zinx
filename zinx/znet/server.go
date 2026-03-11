package znet

import (
	"fmt"
	"net"

	"github.com/zfz-725/zinx/utils"
	"github.com/zfz-725/zinx/ziface"
)

// 定义一个Server的服务器模块
type Server struct {
	// 服务器名称
	Name string
	// 服务器IP版本
	IPVersion string
	// 服务器IP
	IP string
	// 服务器端口
	Port int
	// 路由
	MsgHandler ziface.IMsgHandler
	// 连接管理
	ConnManager ziface.IConnectionManager

	// 服务器创建连接之后的钩子函数
	OnConnStart func(conn ziface.IConnection)
	// 服务器关闭连接之前的钩子函数
	OnConnStop func(conn ziface.IConnection)
}

// 实现IServer接口的方法
// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name: %s, Version: %s, Host: %s, Port: %d\n", utils.GlobalObject.Name, utils.GlobalObject.Version, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] MaxConn: %d, MaxPackageSize: %d\n", utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	fmt.Printf("[Start] Server Listener at IP : %s, Port %d, is starting\n", s.IP, s.Port)

	// 启动Worker工作池
	s.MsgHandler.StartWorkerPool()

	go func() {
		// 1 获取TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Printf("ResolveTCPAddr failed, err: %v\n", err)
			return
		}
		// 2 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("ListenTCP failed, err: %v\n", err)
			return
		}

		fmt.Println("start Zinx server succ.", s.Name, "succ, Listenning...")

		var cid uint32 = 0

		// 3 阻塞的等待客户端连接，处理客户端连接业务（读写）
		for {
			// 如果有客户端连接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Printf("AcceptTCP failed, err: %v\n", err)
				continue
			}

			// 设置最大连接
			if s.ConnManager.Len() >= utils.GlobalObject.MaxConn {
				fmt.Printf("MaxConn Reached, refuse new connection, connID: %d\n", cid)
				// TODO给客户端放回一个错误信息
				conn.Close()
				continue
			}

			dealConn := NewConnection(s, cid, conn, s.MsgHandler)
			cid++
			go dealConn.Start()
		}
	}()
}

// 停止服务器
func (s *Server) Stop() {
	// 将服务器的资源、连接信息、路由信息等进行释放
	fmt.Println("Stop Zinx server...")
	s.ConnManager.Clear()
}

// 运行服务器
func (s *Server) Serve() {
	s.Start()

	// TODO 启动服务器之后其他业务

	// 阻塞，等待其他指令
	select {}
}

// 添加路由
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("Add Router Succeed!")
}

// 获取当前Server的连接管理器
func (s *Server) GetConnManager() ziface.IConnectionManager {
	return s.ConnManager
}

func NewServer() ziface.IServer {
	return &Server{
		Name:        utils.GlobalObject.Name,
		IPVersion:   "tcp4",
		IP:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.TcpPort,
		MsgHandler:  NewMsgHandler(),
		ConnManager: NewConnectionManager(),
	}
}

func (s *Server) SetOnConnStart(hookFunc func(conn ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc func(conn ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

// 调用服务器创建连接之后的钩子函数
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Printf("CallOnConnStart, connID: %d\n", conn.GetConnID())
		s.OnConnStart(conn)
	}
}

// 调用服务器关闭连接之前的钩子函数
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Printf("CallOnConnStop, connID: %d\n", conn.GetConnID())
		s.OnConnStop(conn)
	}
}
