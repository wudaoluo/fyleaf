package gate

import (
	"github.com/wudaoluo/fyleaf/network"
	"github.com/wudaoluo/fyleaf/glog"
)

type Gate struct {
	//wsServer 	*network.WSServer
	tcpServer 	*network.TCPServer
	Processor   network.Processor

}


// Run 启动 wss/tcp_sever
func (gate *Gate) Run(closeSig chan bool) {
	//if cfg.Cfg.WSAddr != "" {
	//	gate.wsServer.NewAgent = func(conn *network.WSConn) network.Agent {
	//		a := &agent{conn: conn, gate: gate}
	//		//TODO 这行代码的作用
	//		//if gate.AgentChanRPC != nil {
	//		//	gate.AgentChanRPC.Go("NewAgent", a)
	//		//}
	//		return a
	//	}
	//}


	if cfg.Cfg.TCPAddr != "" {
		gate.tcpServer.Addr = cfg.Cfg.TCPAddr
		gate.tcpServer.NewAgent = func(conn *network.TCPConn) network.Agent {
			a := &agent{conn: conn, gate: gate}
			//TODO 这行代码的作用
			//if gate.AgentChanRPC != nil {
			//	gate.AgentChanRPC.Go("NewAgent", a)
			//}
			return a
		}

	}


	//if gate.wsServer != nil {
	//	glog.Info("启动wss 监听地址:",cfg.Cfg.WSAddr)
	//	gate.wsServer.Start()
	//}

	if gate.tcpServer != nil {
		glog.Info("启动tcp 监听地址:",cfg.Cfg.TCPAddr)
		gate.tcpServer.Start()
	}
}


// OnDestrony 停止wss 和 tcp
func (gate *Gate) OnDestroy(closeSig chan bool) {
	<-closeSig

	//if gate.wsServer != nil {
	//	glog.Warning("关闭wss server")
	//	gate.wsServer.Close()
	//}
	if gate.tcpServer != nil {
		glog.Warning("关闭tcp server")
		gate.tcpServer.Close()
	}

}