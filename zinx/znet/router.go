package znet

import "github.com/zfz-725/zinx/ziface"

// 实现router时，先嵌入这个BaseRouter基类，然后根据需要对这个基类的方法进行重写
type BaseRouter struct {
}

func (b *BaseRouter) PreHandle(request ziface.IRequest) {
	// 空实现
}

func (b *BaseRouter) Handle(request ziface.IRequest) {
	// 空实现
}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {
	// 空实现
}
