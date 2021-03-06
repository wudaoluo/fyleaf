package network

import (
	"net"
	"github.com/wudaoluo/fyleaf/glog"
	"sync"
	"sync/atomic"
	"time"
)

type TCPServer struct {
	Addr            string
	NewAgent        func(*TCPConn) Agent
	ln              net.Listener
	conns           ConnSet
	mutexConns      sync.RWMutex
	connCount       int32			//当前连接数

	// msg parser
	LenMsgLen    int
	MinMsgLen    uint32
	MaxMsgLen    uint32
	LittleEndian bool
	msgParser    *MsgParser
}





//原子性
/*
初始化 atomic.StoreInt32(&connCount, 0)
获取 atomic.LoadInt32(&connCount)
增加atomic.AddInt32(&connCount, 1)
减少 atomic.AddInt32(&connCount, -1)

*/

func (server *TCPServer) Start() {
	server.init()
	go server.run()
}


func (server *TCPServer) init() {
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		glog.Fatal("%v", err)
	}

	glog.Info("tcp server listening addr ",server.Addr)
	//if server.NewAgent == nil {
	//	glog.Fatal("NewAgent must not be nil")
	//}


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

		//在同一个goroutine，可以减少并发的竞争
		//连接超过最大值，释放连接
		if atomic.LoadInt32(&server.connCount) >= cfg.Cfg.MaxConnNum {
			conn.Close()
			glog.Warning("too many connections")
			continue
		}

		atomic.AddInt32(&server.connCount, 1)


		//把这个当做 session 池吧
		server.conns[conn] = struct{}{}

		go server.handle(conn)

	}

}


func (server *TCPServer) handle(conn net.Conn) {

	server.conns[conn] = struct{}{}

	tcpConn := newTCPConn(conn, server.msgParser)
	agent := server.NewAgent(tcpConn)


	agent.Run()

	// cleanup
	tcpConn.Close()

	//删除 map 中的连接
	server.mutexConns.Lock()
	delete(server.conns, conn)
	server.mutexConns.Unlock()

	agent.OnClose()

}


func (server *TCPServer) Close() {
	server.ln.Close()
	//清除连接池
	server.mutexConns.Lock()
	for conn := range server.conns {
		conn.Close()
	}
	server.conns = nil
	server.mutexConns.Unlock()
}