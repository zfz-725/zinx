package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestDataPack(t *testing.T) {
	/*
		模拟的服务器
	*/
	// 1创建socketTCP
	listener, err := net.Listen("tcp", "127.0.0.1:8999")
	if err != nil {
		t.Errorf("Listen failed, err: %v\n", err)
		return
	}
	defer listener.Close()

	// 2读取客户端数据，拆包处理
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				// 监听器关闭时退出循环，不报错
				break
			}
			defer conn.Close()
			go func(conn net.Conn) {
				// 处理客户端请求
				// ----- 拆包过程 -----
				// 定义一个拆包对象
				dp := NewDataPack()
				for {
					// 1第一次从conn读，读取head
					headData := make([]byte, dp.GetHeadLen())
					_, err := conn.Read(headData)
					if err != nil {
						// 连接关闭时退出循环，不报错
						break
					}
					// 2根据head中的datalen，读取data
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						t.Errorf("Unpack head failed, err: %v\n", err)
						break
					}
					if msgHead.GetMsgLen() > 0 {
						data := make([]byte, msgHead.GetMsgLen())
						_, err = conn.Read(data)
						if err != nil {
							// 连接关闭时退出循环，不报错
							break
						}

						// 完整的消息已经读取完毕
						fmt.Println("ID:", msgHead.GetMsgID(), "DataLen:", msgHead.GetMsgLen(), "Data:", string(data))
					}
				}
			}(conn)
		}
	}()

	/*
		模拟客户端
	*/
	// 1创建socketTCP
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		t.Errorf("Dial failed, err: %v\n", err)
	}
	defer conn.Close()

	// 2发送数据
	dp := NewDataPack()

	// 模拟粘包过程，封装msg一同发送
	msg1 := &Message{
		ID:      1,
		DataLen: 5,
		Data:    []byte("hello"),
	}
	msg2 := &Message{
		ID:      2,
		DataLen: 5,
		Data:    []byte("world"),
	}
	data1, err := dp.Pack(msg1)
	if err != nil {
		t.Errorf("Pack failed, err: %v\n", err)
	}
	data2, err := dp.Pack(msg2)
	if err != nil {
		t.Errorf("Pack failed, err: %v\n", err)
	}
	sendData := append(data1, data2...)
	_, err = conn.Write(sendData)
	if err != nil {
		t.Errorf("Write failed, err: %v\n", err)
	}

	// 客户端阻塞
	<-time.After(time.Second * 3)
	fmt.Println("Client Test Pass")
}
