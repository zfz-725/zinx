package main

import (
	"fmt"
	"net"
	"time"

	"github.com/zfz-725/zinx/znet"
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
		dp := znet.NewDataPack()
		msg, err := dp.Pack(znet.NewMessage(1, []byte("Hello Zinx V0.6")))
		if err != nil {
			fmt.Printf("Pack failed, err: %v\n", err)
			continue
		}

		_, err = conn.Write(msg)
		if err != nil {
			fmt.Printf("Write failed, err: %v\n", err)
			continue
		}

		// 读取客户端的Msg Head
		msgHead := make([]byte, dp.GetHeadLen())
		_, err = conn.Read(msgHead)
		if err != nil {
			fmt.Printf("Read failed, err: %v\n", err)
			continue
		}
		// 拆包，得到msgID 和 dataLen 放在msg消息中
		msgObj, err := dp.Unpack(msgHead)
		if err != nil {
			fmt.Printf("Unpack failed, err: %v\n", err)
			continue
		}
		// 根据dataLen，再次读取data
		if msgObj.GetMsgLen() > 0 {
			data := make([]byte, msgObj.GetMsgLen())
			_, err = conn.Read(data)
			if err != nil {
				fmt.Printf("Read failed, err: %v\n", err)
				continue
			}
			msgObj.SetData(data)
			// 打印接收到的消息
			fmt.Printf("Received message: ID=%d, Data=%s\n", msgObj.GetMsgID(), string(msgObj.GetData()))
		}

		time.Sleep(time.Second)
	}
}
