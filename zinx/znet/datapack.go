package znet

import (
	"encoding/binary"
	"errors"

	"github.com/zfz-725/zinx/utils"
	"github.com/zfz-725/zinx/ziface"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取消息头长度
func (dp *DataPack) GetHeadLen() uint32 {
	return 8
}

// 封包
// datalen|megID|data
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个字节数组，用于存储封包后的结果
	data := make([]byte, dp.GetHeadLen()+msg.GetMsgLen())

	// 写入消息头
	binary.BigEndian.PutUint32(data[:4], msg.GetMsgID())
	binary.BigEndian.PutUint32(data[4:8], msg.GetMsgLen())

	// 写入消息数据
	copy(data[dp.GetHeadLen():], msg.GetData())

	return data, nil
}

// 拆包
func (dp *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	// 创建一个消息头结构体
	msgHead := &Message{}

	// 从数据中读取消息头
	msgHead.ID = binary.BigEndian.Uint32(data[:4])
	msgHead.DataLen = binary.BigEndian.Uint32(data[4:8])

	if utils.GlobalObject.MaxPackageSize > 0 && msgHead.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("package size is too large")
	}

	// 创建一个消息体结构体
	msg := &Message{
		ID:      msgHead.ID,
		DataLen: msgHead.DataLen,
		Data:    data[dp.GetHeadLen():],
	}

	return msg, nil
}
