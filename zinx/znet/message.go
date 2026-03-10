package znet

type Message struct {
	ID      uint32
	DataLen uint32
	Data    []byte
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		ID:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// 获取消息ID
func (m *Message) GetMsgID() uint32 {
	return m.ID
}

// 获取消息数据长度
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

// 获取消息数据
func (m *Message) GetData() []byte {
	return m.Data
}

// 设置消息ID
func (m *Message) SetMsgID(id uint32) {
	m.ID = id
}

// 设置消息数据长度
func (m *Message) SetMsgLen(len uint32) {
	m.DataLen = len
}

// 设置消息数据
func (m *Message) SetData(data []byte) {
	m.Data = data
}
