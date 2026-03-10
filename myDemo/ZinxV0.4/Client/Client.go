package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 1 连接远程服务器，得到一个conn连接
	fmt.Println("Client Start...")

	// 2 包一个conn连接，得到一个Connection接口
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Printf("NewConnection failed, err: %v\n", err)
		return
	}

	// 3 写数据
	for {
		_, err := conn.Write([]byte("Hello Zinx V0.1"))
		if err != nil {
			fmt.Printf("Write failed, err: %v\n", err)
			continue
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Read failed, err: %v\n", err)
			continue
		}
		fmt.Printf("Server Call Back: %s\n", buf[:cnt])

		time.Sleep(time.Second)
	}
}
