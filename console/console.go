package console

import (
	"github.com/wudaoluo/fyleaf/conf"
	"github.com/wudaoluo/fyleaf/network"
	"strconv"
	"github.com/wudaoluo/fyleaf/glog"
)

var cfg = conf.GetInstance()
var server *network.TCPServer

// Init() 运行 tcp server
func Init() {
	if cfg.Cfg.ConsolePort == 0 {
		return
	}

	server = new(network.TCPServer)
	server.Addr = "localhost:" + strconv.Itoa(cfg.Cfg.ConsolePort)
	server.NewAgent = newAgent
	server.Start()
}





// Destory 关闭 tcp server
func Destroy() {
	if server != nil {
		glog.Warning("关闭console 监听",cfg.Cfg.ConsolePort)
		server.Close()
	}
}
