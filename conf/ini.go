package conf

import (
	"sync"
	"github.com/Unknwon/goconfig"
	"fyleaf/glog"
)

type iniConf struct {
	mu      sync.Mutex
	server  *Server
	fileName string
	cfg      *goconfig.ConfigFile
	goconfig.ParseError
}


func newiniConf(f string,cfg *Server) *iniConf {
	return &iniConf{
		fileName:f,
		server:cfg,
	}
}



func (i *iniConf) load() {
	var err error
	i.cfg, err = goconfig.LoadConfigFile(i.fileName)
	if err != nil {
		glog.Fatal("第一次载入配置文件失败",err)
	}

	//s.Global.Port,_ = i.cfg.Int(goconfig.DEFAULT_SECTION,"port")
	//s.Global.Debug,_ = cfg.Bool("Global","debug")
	//s.Global.LogFile,_ = cfg.GetValue("Global","log_file")


	//s.TCP.State,_ = cfg.Bool("TCP","state")
	//s.TCP.Timeout,_ = cfg.Int("TCP","timeout")

	//s.UDP.State,_ = cfg.Bool("UDP","state")
	//s.UDP.Timeout,_ = cfg.Int("UDP","timeout")

	return

}


func (i *iniConf) Reload() {
	err := i.cfg.Reload()
	if err !=nil {
		glog.Error("重新载入配置文件失败",err)
	}

}


func (i *iniConf) SaveFile() {
	err := goconfig.SaveConfigFile(i.cfg,i.fileName)
	if err != nil {
		glog.Error("保存配置到本地失败",err)
	}
}