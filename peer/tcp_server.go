package peer

import (
	"net"
	"fyleaf/glog"
	"sync"
	"sync/atomic"
	"time"
)

type TCPServer struct {
	Addr            string
	MaxConnNum      int32
	NewAgent        func(*TCPConn) Agent
	ln              net.Listener
	conns           ConnSet
	mutexConns      sync.RWMutex
	connCount       int32			//当前连接数
}



//原子性
/*
初始化 atomic.StoreInt32(&connCount, 0)
获取 atomic.LoadInt32(&connCount)
增加atomic.AddInt32(&connCount, 1)
减少 atomic.AddInt32(&connCount, -1)

*/

func (server *TCPServer) Start() {

}


func (server *TCPServer) init() {
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		glog.Fatal("%v", err)
	}

	if server.MaxConnNum <= 0 {
		server.MaxConnNum = 100
		glog.Info("invalid MaxConnNum, reset to %v", server.MaxConnNum)
	}


	if server.NewAgent == nil {
		glog.Fatal("NewAgent must not be nil")
	}


	atomic.StoreInt32(&server.connCount, 0)

	server.ln = ln
	server.conns = make(ConnSet)

	// msg parser
	msgParser := NewMsgParser()
	msgParser.SetMsgLen(server.LenMsgLen, server.MinMsgLen, server.MaxMsgLen)
	msgParser.SetByteOrder(server.LittleEndian)
	server.msgParser = msgParser

}


func (server *TCPServer) run() {
	var tempDelay time.Duration
	for {
		conn, err := server.ln.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				glog.Warning("accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return
		}
		tempDelay = 0


		go server.handle(conn)
	}

}


func (server *TCPServer) handle(conn net.Conn) {
	atomic.AddInt32(&server.connCount, 1)
	//连接超过最大值，释放连接，减少当前连接数(server.connCount)
	if atomic.LoadInt32(&server.connCount) >= server.MaxConnNum {
		conn.Close()
		atomic.AddInt32(&server.connCount, -1)
		glog.Warning("too many connections")
		return
	}

	//TODO 暂时不知道这个有啥鸟用
	server.conns[conn] = struct{}{}

	tcpConn := newTCPConn(conn, server.PendingWriteNum, server.msgParser)
	agent := server.NewAgent(tcpConn)


	agent.Run()

	// cleanup
	tcpConn.Close()
	server.mutexConns.Lock()
	delete(server.conns, conn)
	server.mutexConns.Unlock()
	agent.OnClose()

	server.wgConns.Done()
}


func (server *TCPServer) Close() {
	server.ln.Close()
}