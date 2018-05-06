package gate

import (
	"fyleaf/conf"
	"fyleaf/network"
	"fyleaf/glog"
)

var cfg = conf.GetInstance()

type Gate struct {
	wsServer *network.WSServer
	tcpServer *network.TCPServer

}


// Run 启动 wss/tcp_sever
func (gate *Gate) Run(closeSig chan bool) {
	if cfg.Cfg.WSAddr != "" {

	}


	if cfg.Cfg.TCPAddr != "" {

	}


	if gate.wsServer != nil {
		glog.Info("启动wss 监听地址:",cfg.Cfg.WSAddr)
		gate.wsServer.Start()
	}

	if gate.tcpServer != nil {
		glog.Info("启动tcp 监听地址:",cfg.Cfg.TCPAddr)
		gate.tcpServer.Start()
	}
}


// OnDestrony 停止wss 和 tcp
func (gate *Gate) OnDestroy(closeSig chan bool) {
	<-closeSig

	if gate.wsServer != nil {
		glog.Warning("关闭wss server")
		gate.wsServer.Close()
	}
	if gate.tcpServer != nil {
		glog.Warning("关闭tcp server")
		gate.tcpServer.Close()
	}

}