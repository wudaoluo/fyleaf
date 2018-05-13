package internal

import (
	"github.com/name5566/leaf/module"
	"golang/base"
)

var (
	skeleton = base.NewSkeleton()
	//ChanRPC  = skeleton.ChanRPCServer
)


/*
	OnInit()
	OnDestroy(closeSig chan bool)  //关闭的操作
	Run(closeSig chan bool)         //

*/

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	//m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
