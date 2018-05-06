package console

import (
	"fyleaf/conf"
	"fyleaf/network"
)

var cfg = conf.GetInstance()
var server *network.TCPServer

// Init() 运行 tcp server
func Init() {
	if cfg.Cfg.ConsolePort == 0 {
		return
	}

	server = new(network.TCPServer)
	server.Start()
}


// Destory 关闭 tcp server
func Destroy() {
	if server != nil {
		server.Close()
	}
}
