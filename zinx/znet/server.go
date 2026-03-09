package znet

import (
	"errors"
	"fmt"
	"net"

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
}

// 当前客户端连接所绑定hadle api
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallbackToClient ...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Printf("Write failed, err: %v\n", err)
		return errors.New("CallbackToClient failed:" + err.Error())
	}
	return nil
}

// 实现IServer接口的方法
// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP : %s, Port %d, is starting\n", s.IP, s.Port)

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

			dealConn := NewConnection(cid, conn, CallBackToClient)
			cid++
			go dealConn.Start()
		}
	}()
}

// 停止服务器
func (s *Server) Stop() {
	// TODO 将服务器的资源、连接信息、路由信息等进行释放
}

// 运行服务器
func (s *Server) Serve() {
	s.Start()

	// TODO 启动服务器之后其他业务

	// 阻塞，等待其他指令
	select {}
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}
