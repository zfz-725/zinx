package utils

import (
	"encoding/json"
	"os"

	"github.com/zfz-725/zinx/ziface"
)

/*
	存储zinx框架的全部参数，供其他模块使用
	一些参数是可以通过zinx.json由用户进行配置
*/

type GlobalObj struct {
	/*
		Server
	*/
	// zinx全局的Server对象
	TCPServer ziface.IServer
	// 服务器IP
	Host string
	// 服务器端口
	TcpPort int
	// 服务器名称
	Name string

	/*
		Zinx
	*/
	// Zinx版本号
	Version string
	// 服务器最大连接数
	MaxConn int
	// 每个连接的最大消息包大小
	MaxPackageSize uint32
}

// 定义一个全局的对象
var GlobalObject *GlobalObj

// 从conf/zinx.json加载配置
func (g *GlobalObj) Reload() {
	configData, err := os.ReadFile("../conf/zinx.json")
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		panic(err)
	}
	// 将json文件解析到struct
	err = json.Unmarshal(configData, g)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		Host:           "0.0.0.0",
		TcpPort:        8999,
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	// 从conf/zinx.json加载配置
	GlobalObject.Reload()
}
