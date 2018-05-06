package gate

import (
	"net"
	"fyleaf/network"
	"reflect"
	"fyleaf/glog"
)


type Agent interface {
	WriteMsg(msg interface{})
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}



type agent struct {
	conn     network.Conn
	gate     *Gate
	userData interface{}   // TODO 用户数据在这里的作用
}

// Run 读取的数据,解析数据, route 到响应的 handle 处理
func (a *agent) Run() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			glog.Warning("read message:",err)
			break
		}

		if a.gate.Processor != nil {
			msg, err := a.gate.Processor.Unmarshal(data)
			if err != nil {
				glog.Warning("unmarshal message error:",err)
				break
			}
			err = a.gate.Processor.Route(msg, a)
			if err != nil {
				glog.Warning("route message error:",err)
				break
			}
		}
	}
}

//TODO 没看出来有啥用
func (a *agent) OnClose() {
	//if a.gate.AgentChanRPC != nil {
	//	err := a.gate.AgentChanRPC.Call0("CloseAgent", a)
	//	if err != nil {
	//		log.Error("chanrpc error: %v", err)
	//	}
	//}
}


//TODO 这里为什么还要 if a.gate.Processor != nil判断一下
//TODO Processor 有问题的 readmsg 就出错了啊
func (a *agent) WriteMsg(msg interface{}) {
	if a.gate.Processor != nil {
		data, err := a.gate.Processor.Marshal(msg)
		if err != nil {
			glog.Warning("marshal message %v error: %v", reflect.TypeOf(msg), err)
			return
		}
		err = a.conn.WriteMsg(data...)
		if err != nil {
			glog.Warning("write message %v error: %v", reflect.TypeOf(msg), err)
		}
	}
}


func (a *agent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *agent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *agent) Close() {
	a.conn.Close()
}

//TODO 这个的作用是什么
func (a *agent) Destroy() {
	a.conn.Destroy()
}

//TODO 用户数据是什么数据
func (a *agent) UserData() interface{} {
	return a.userData
}

func (a *agent) SetUserData(data interface{}) {
	a.userData = data
}




